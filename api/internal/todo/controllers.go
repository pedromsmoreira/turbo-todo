package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	tcsvc *TodoService
}

func NewTodoController(tcsvc *TodoService) *TodoController {
	return &TodoController{
		tcsvc: tcsvc,
	}
}
func (tc *TodoController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List",
	})
}

func (tc *TodoController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get by id",
	})
}

func (tc *TodoController) Create(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Create",
	})
}

func (tc *TodoController) Update(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Update",
	})
}

func (tc *TodoController) Delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Delete",
	})
}
