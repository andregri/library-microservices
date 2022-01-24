package bookssvc

import (
	"context"

	"gorm.io/gorm"
)

// BookServiceInstance implements the BookService interface
type BookServiceInstance struct {
	Db *gorm.DB
}

// PostBook adds the book to the database
func (svc BookServiceInstance) PostBook(ctx context.Context, b Book) (int, error) {
	res := svc.Db.Create(&b)
	if res.Error != nil {
		return -1, res.Error
	}
	return b.ID, nil
}

// GetBook
func (svc BookServiceInstance) GetBook(ctx context.Context, id int) (Book, error) {
	var book Book
	res := svc.Db.First(&book, id)
	return book, res.Error
}

// PutBook
func (svc BookServiceInstance) PutBook(ctx context.Context, id int, b Book) error {
	b.ID = id
	res := svc.Db.Save(&b)
	return res.Error
}

// DeleteBook
func (svc BookServiceInstance) DeleteBook(ctx context.Context, id int) error {
	res := svc.Db.Delete(&Book{}, id)
	return res.Error
}
