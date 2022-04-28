package todos

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/internal/api/apierrors"
)

type Controller struct {
	svc *TodoService
}

func NewTodoController(tcSvc *TodoService) *Controller {
	return &Controller{
		svc: tcSvc,
	}
}
func (tc *Controller) list(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "list",
	})
}

func (tc *Controller) get(c *gin.Context) {
	id := c.Param("id")

	todo, err := tc.svc.Get(id)

	if err != nil {
		pd := &apierrors.ProblemDetails{}
		c.JSON(http.StatusBadRequest, pd)
	}

	if todo == nil {
		pd := &apierrors.ProblemDetails{
			Type:   "not_found",
			Status: http.StatusNotFound,
			Detail: fmt.Sprintf("resource with id: %s not found", id),
		}
		c.JSON(http.StatusNotFound, pd)
	}

	dto, err := FromModelToDto(todo)

	if err != nil {
		pd := &apierrors.ProblemDetails{
			Type:   "conversion",
			Status: http.StatusInternalServerError,
			Detail: "error converting model",
		}
		c.JSON(http.StatusInternalServerError, pd)
	}

	resp := &ApiResponse{
		Data: dto,
	}

	c.JSON(http.StatusOK, resp)
}

func (tc *Controller) create(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "create",
	})
}

func (tc *Controller) update(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "update",
	})
}

func (tc *Controller) delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "delete",
	})
}