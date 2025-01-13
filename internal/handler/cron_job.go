package handler

import (
	"gc2-p3-xiowel/config"
	"gc2-p3-xiowel/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"time"
)

func RunJob(c echo.Context) error {
	cronJob := cron.New()

	_, err := cronJob.AddFunc("0 0 * * *", func() {
		if err := updateBookStatus(); err != nil {
			log.Printf("Error updating book status: %v\n", err)
		}
	})
	if err != nil {
		log.Printf("Error adding cron job: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to schedule job",
		})
	}

	cronJob.Start()
	log.Println("Cron job started for updating book status")

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Cron job started successfully",
	})
}

func updateBookStatus() error {
	db := config.DB

	var overdueBooks []models.Book

	// Find books where returnDate < today and status is Borrowed
	today := time.Now().Truncate(24 * time.Hour)
	if err := db.Where("status = ? AND return_date < ?", "Borrowed", today).Find(&overdueBooks).Error; err != nil {
		log.Printf("Error fetching books: %v\n", err)
		return err
	}

	if len(overdueBooks) == 0 {
		log.Println("No overdue books found to update")
		return nil
	}

	log.Printf("Found %d overdue books to reset status\n", len(overdueBooks))

	// Update status to "Available"
	if err := db.Model(&models.Book{}).
		Where("status = ? AND return_date < ?", "Borrowed", today).
		Update("status", "Available").Error; err != nil {
		log.Printf("Error updating books: %v\n", err)
		return err
	}

	log.Printf("Successfully updated %d books to 'Available'\n", len(overdueBooks))
	return nil
}
