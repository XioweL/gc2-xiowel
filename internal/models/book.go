package models

type Book struct {
	BookID        int    `json:"book_id" gorm:"primaryKey;column:book_id"`
	Title         string `json:"title" gorm:"column:title"`
	Author        string `json:"author" gorm:"column:author"`
	PublishedDate string `json:"published_date" gorm:"column:published_date"`
	Status        string `json:"status" gorm:"column:status"`
	//UserID        int    `json:"user_id" gorm:"column:user_id;foreignKey:UserID;references:user_id"`
	//User          User   `gorm:"foreignKey:UserID;references:user_id"`
}

type BorrowedBook struct {
	ID           int    `json:"id" gorm:"primaryKey;column:id"`
	BookID       int    `json:"book_id" gorm:"column:book_id;foreignKey:BookID;references:book_id"`
	UserID       int    `json:"user_id" gorm:"column:user_id;foreignKey:UserID;references:user_id"`
	BorrowedDate string `json:"borrowed_date" gorm:"column:borrowed_date"`
	ReturnDate   string `json:"return_date" gorm:"column:return_date"`
}

func (BorrowedBook) TableName() string {
	return "borrowedbooks"
}
