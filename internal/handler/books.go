package handler

import (
	"errors"
	"gc2-p3-xiowel/config"
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
	if err := config.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return &pb.CreateBookResponse{Message: "Book created successfully"}, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var book models.Book
	result := config.DB.First(&book, req.BookId)
	if result.Error != nil {
		return nil, result.Error
	}
	book.Title = req.Title
	book.Author = req.Author
	book.PublishedDate = req.PublishedDate
	book.Status = req.Status

	config.DB.Save(&book)
	return &pb.UpdateBookResponse{Message: "Book updated successfully"}, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	result := config.DB.Delete(&models.Book{}, req.BookId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pb.DeleteBookResponse{Message: "Book deleted successfully"}, nil
}

func (s *BookServiceServer) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	// Log sebelum mengambil user_id
	log.Println("Available context keys:", ctx)
	userID := ctx.Value("user_id")
	log.Println("Extracted user_id from context:", userID)

	log.Printf("User ID extracted from context in gRPC handler: %v", userID)

	log.Println("Extracted user_id from gRPC context:", userID)
	if userID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "User ID not found in context")
	}

	// Pastikan user_id dalam format yang benar (int)
	userIDInt, ok := userID.(int)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid User ID format")
	}

	// Cari buku berdasarkan ID
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

	// Simpan data peminjaman buku ke dalam tabel borrowedbooks
	borrowedBook := models.BorrowedBook{
		BookID:       book.BookID,
		UserID:       userIDInt,
		BorrowedDate: in.BorrowedDate,
		ReturnDate:   in.ReturnDate,
	}

	if err := config.DB.Create(&borrowedBook).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to borrow book: %v", err)
	}

	// Return response sukses
	return &pb.BorrowBookResponse{
		Status:  "success",
		Message: "Book borrowed successfully",
	}, nil
}
