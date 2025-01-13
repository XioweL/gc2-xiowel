package handler

import (
	"context"
	"gc2-p3-xiowel/config"
	"gc2-p3-xiowel/internal/models"
	"gc2-p3-xiowel/pb"
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
