package routes

import (
	"fmt"
	"gc2-p3-xiowel/internal/handler"
	internal "gc2-p3-xiowel/internal/middleware"
	"gc2-p3-xiowel/pb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		resp, err := client.GetBooks(c.Request().Context(), &pb.EmptyRequest{})
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
		resp, err := client.GetBook(c.Request().Context(), &pb.GetBookByIdRequest{BookId: int32(bookId)})
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
		resp, err := client.CreateBook(c.Request().Context(), req)
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
		resp, err := client.UpdateBook(c.Request().Context(), req)
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
		resp, err := client.DeleteBook(c.Request().Context(), &pb.DeleteBookRequest{BookId: int32(bookId)})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp)
	})
	auth.POST("/books/borrow", func(c echo.Context) error {
		req := new(pb.BorrowBookRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		resp, err := client.BorrowBook(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, resp)
	})

	return e
}
