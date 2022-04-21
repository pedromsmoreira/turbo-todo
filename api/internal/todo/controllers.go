package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/api/internal/errors"
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
	id := c.Param("id")

	todo, err := tc.tcsvc.Get(id)

	if err != nil {
		pd := &errors.ProblemDetails{}
		c.JSON(http.StatusBadRequest, pd)
	}

	dto, err := FromModelToDto(todo)

	if err != nil {
		pd := &errors.ProblemDetails{}
		c.JSON(http.StatusInternalServerError, pd)
	}

	resp := &ApiResponse{
		Data: dto,
	}

	c.JSON(http.StatusOK, resp)
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
