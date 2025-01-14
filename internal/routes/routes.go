package routes

import (
	"fmt"
	"gc2-p3-xiowel/internal/handler"
	internal "gc2-p3-xiowel/internal/middleware"
	"gc2-p3-xiowel/pb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
)

func SetupRoutes(e *echo.Echo, client pb.BookServiceClient) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Public Routes
	e.POST("/users/login", handler.Login)
	e.POST("/users/register", handler.Register)
	e.GET("/start-job", handler.RunJob)

	auth := e.Group("")
	auth.Use(internal.CustomJwtMiddleware)
	auth.GET("/books/all", func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		md := metadata.New(map[string]string{
			"authorization": authHeader,
		})
		// Tambahkan metadata ke context
		ctx := metadata.NewOutgoingContext(c.Request().Context(), md)

		resp, err := client.GetBooks(ctx, &pb.EmptyRequest{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp.Books)
	})
	auth.GET("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		bookId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("invalid id: %v", err),
			})
		}

		// Ambil Authorization header dan tambahkan sebagai metadata untuk gRPC
		authHeader := c.Request().Header.Get("Authorization")
		md := metadata.New(map[string]string{
			"authorization": authHeader,
		})
		ctx := metadata.NewOutgoingContext(c.Request().Context(), md)

		// Panggilan gRPC dengan metadata
		resp, err := client.GetBook(ctx, &pb.GetBookByIdRequest{BookId: int32(bookId)})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp.Book)
	})

	auth.POST("/books/create", func(c echo.Context) error {
		req := new(pb.CreateBookRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		// Ambil Authorization header dan tambahkan sebagai metadata untuk gRPC
		authHeader := c.Request().Header.Get("Authorization")
		md := metadata.New(map[string]string{
			"authorization": authHeader,
		})
		ctx := metadata.NewOutgoingContext(c.Request().Context(), md)

		resp, err := client.CreateBook(ctx, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp)
	})

	auth.PUT("/books/update/:id", func(c echo.Context) error {
		id := c.Param("id")
		bookId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("invalid id: %v", err),
			})
		}

		req := new(pb.UpdateBookRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		req.BookId = int32(bookId)

		// Ambil Authorization header dan tambahkan sebagai metadata untuk gRPC
		authHeader := c.Request().Header.Get("Authorization")
		md := metadata.New(map[string]string{
			"authorization": authHeader,
		})
		ctx := metadata.NewOutgoingContext(c.Request().Context(), md)

		resp, err := client.UpdateBook(ctx, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp)
	})

	auth.DELETE("/books/delete/:id", func(c echo.Context) error {
		id := c.Param("id")
		bookId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("invalid id: %v", err),
			})
		}

		// Ambil Authorization header dan tambahkan sebagai metadata untuk gRPC
		authHeader := c.Request().Header.Get("Authorization")
		md := metadata.New(map[string]string{
			"authorization": authHeader,
		})
		ctx := metadata.NewOutgoingContext(c.Request().Context(), md)

		resp, err := client.DeleteBook(ctx, &pb.DeleteBookRequest{BookId: int32(bookId)})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp)
	})
	rpc := e.Group("")
	rpc.Use(internal.CustomGrpcJwtMiddleware)
	rpc.POST("/books/borrow", func(c echo.Context) error {
		req := new(pb.BorrowBookRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		// Teruskan Authorization header ke metadata
		authHeader := c.Request().Header.Get("Authorization")
		md := metadata.New(map[string]string{
			"authorization": authHeader,
		})
		ctx := metadata.NewOutgoingContext(c.Request().Context(), md)

		// Panggil gRPC dengan konteks metadata
		resp, err := client.BorrowBook(ctx, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, resp)
	})

	return e
}
