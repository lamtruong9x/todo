package data

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Content   string    `json:"content" binding:"required,"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"-"`
}

type TaskModel struct {
	DB *gorm.DB
}

//func (t *Task) TableName() string {
//	return "task"
//}

func (t TaskModel) Insert(task *Task) error {
	tx := t.DB.Begin()
	tx.Select("content", "user_id").Create(task)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return nil
}

func (t TaskModel) Get(id int) ([]*Task, error) {
	var task []*Task
	tx := t.DB.Where(map[string]interface{}{"user_id": id}).Find(&task)
	return task, tx.Error
}

func (t TaskModel) GetByID(id int) (*Task, error) {
	var task *Task
	tx := t.DB.First(task, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return task, nil
}

func (t TaskModel) Update(task *Task) error {
	tx := t.DB.Begin()
	tx.Exec(`UPDATE tasks
				 SET content = ?, done = ?`, task.Content, task.Done)
	if tx.Error != nil {
		fmt.Println("errors")
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return nil
}

func (t TaskModel) Delete(id int) error {
	tx := t.DB.Unscoped().Delete(&Task{}, id)
	fmt.Println("pass")
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
