package data

import (
	"errors"
	"gorm.io/gorm"
)

type Models struct {
	Tasks interface {
		Insert(movie *Task) error
		Get(id int) ([]*Task, error)
		GetByID(id int) (*Task, error)
		Update(movie *Task) error
		Delete(id int) error
	}
}

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when // looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

func NewModel(db *gorm.DB) Models {
	return Models{
		Tasks: TaskModel{DB: db},
	}
}
