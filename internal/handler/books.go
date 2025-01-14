package handler

import (
	"errors"
	"gc2-p3-xiowel/config"
	internal "gc2-p3-xiowel/internal/middleware"
	"gc2-p3-xiowel/internal/models"
	"gc2-p3-xiowel/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
)

// Server defines the gRPC server struct
type BookServiceServer struct {
	pb.UnimplementedBookServiceServer
}

// gRPC Services
func (s *BookServiceServer) GetBooks(ctx context.Context, req *pb.EmptyRequest) (*pb.BooksResponse, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	var grpcBooks []*pb.Book
	for _, book := range books {
		grpcBooks = append(grpcBooks, &pb.Book{
			BookId:        int32(book.BookID),
			Title:         book.Title,
			Author:        book.Author,
			PublishedDate: book.PublishedDate,
			Status:        book.Status,
		})
	}

	return &pb.BooksResponse{Books: grpcBooks}, nil
}

func (s *BookServiceServer) GetBook(ctx context.Context, req *pb.GetBookByIdRequest) (*pb.GetBookByIdResponse, error) {
	var book models.Book
	result := config.DB.First(&book, req.BookId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pb.GetBookByIdResponse{
		Book: &pb.Book{
			BookId:        int32(book.BookID),
			Title:         book.Title,
			Author:        book.Author,
			PublishedDate: book.PublishedDate,
			Status:        book.Status,
		},
	}, nil
}

func (s *BookServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := models.Book{
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: req.PublishedDate,
		Status:        req.Status,
	}

	if book.Status == "" {
		book.Status = "Available"
	} else if book.Status != "Available" && book.Status != "Borrowed" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid status")
	}

	if err := config.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return &pb.CreateBookResponse{Message: "Book created successfully"}, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var book models.Book

	result := config.DB.First(&book, req.BookId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Book not found")
		}
		return nil, result.Error
	}
	book.Title = req.Title
	book.Author = req.Author
	book.PublishedDate = req.PublishedDate
	book.Status = req.Status

	if book.Status == "" {
		book.Status = "Available"
	} else if book.Status != "Available" && book.Status != "Borrowed" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid status")
	}

	config.DB.Save(&book)
	return &pb.UpdateBookResponse{Message: "Book updated successfully"}, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	var book models.Book

	// Cari buku berdasarkan ID
	result := config.DB.First(&book, req.BookId)
	if result.Error != nil {
		// Jika buku tidak ditemukan
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Book not found")
		}
		return nil, status.Errorf(codes.Internal, "Database error: %v", result.Error)
	}

	// Hapus buku
	if err := config.DB.Delete(&book).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete book: %v", err)
	}

	return &pb.DeleteBookResponse{Message: "Book deleted successfully"}, nil
}

func (s *BookServiceServer) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	userID, ok := ctx.Value(internal.UserIDKey).(int32)
	if !ok || userID == 0 {
		log.Println("User ID not found in context")
		return nil, status.Errorf(codes.Unauthenticated, "User ID not found in context")
	}

	var book models.Book
	if err := config.DB.Where("book_id = ?", in.BookId).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.BorrowBookResponse{
				Status:  "failure",
				Message: "Book not found",
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	if book.Status != "Available" {
		return nil, status.Errorf(codes.FailedPrecondition, "Book is not available for borrowing")
	}

	borrowedBook := models.BorrowedBook{
		BookID:       book.BookID,
		UserID:       int(userID),
		BorrowedDate: in.BorrowedDate,
		ReturnDate:   in.ReturnDate,
	}

	// Simpan data peminjaman buku
	if err := config.DB.Create(&borrowedBook).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to borrow book: %v", err)
	}

	// Update status buku menjadi 'Borrowed'
	book.Status = "Borrowed"
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update book status: %v", err)
	}

	return &pb.BorrowBookResponse{
		Status:  "success",
		Message: "Book borrowed successfully",
	}, nil
}
