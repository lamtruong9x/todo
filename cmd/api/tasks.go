package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todo/internal/data"
)

func (app *application) createTaskHandler(c *gin.Context) {
	var task data.Task
	var input struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	task.Content = input.Content
	// Set default only for now, update later
	task.UserID = 1
	err := app.models.Tasks.Insert(&task)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, &task)
}

func (app *application) getTaskHandler(c *gin.Context) {
	id, ok := ReadIDParam(c, "id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "not valid id",
				"message": "please check your input",
			})
		return
	}
	tasks, err := app.models.Tasks.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound,
				gin.H{
					"error":   "not found",
					"message": "please check your input",
				})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				})
		}
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (app *application) updateTaskHandler(c *gin.Context) {
	id, ok := ReadIDParam(c, "id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "not valid id",
				"message": "please check your input",
			})
		return
	}
	task, err := app.models.Tasks.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error": gorm.ErrRecordNotFound.Error(),
				})
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				})
		}
		return
	}
	var input struct {
		Content string `json:"content"`
		Done    bool   `json:"done"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	fmt.Println(input)
	//if input.Content != nil {
	//	task.Content = *input.Content
	//}
	//if input.Done != nil {
	//	task.Done = *input.Done
	//}
	fmt.Println("Run til line 100")
	err = app.models.Tasks.Update(task)
	fmt.Println("run here?")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (app *application) deleteTaskHandler(c *gin.Context) {
	id, ok := ReadIDParam(c, "id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"erorr":   "not valid id",
				"message": "please check your input",
			})
		return
	}
	err := app.models.Tasks.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.AbortWithError(http.StatusNotFound, err)
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message:": "task deleted"})
}
