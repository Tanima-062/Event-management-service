package handler

import (
	"event-management-service/workshop"
	"event-management-service/workshop/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkshopHandler struct {
	WorkshopUsecase   workshop.IWorkshopUsecase
}

// NewWorkshopHandler 
func NewWorkshopHandler(db *gorm.DB) *WorkshopHandler {
	return &WorkshopHandler{
		WorkshopUsecase:   usecase.NewWorkshopUsecase(db),
	}
}

// List 
func (w *WorkshopHandler) List(c *gin.Context) error {

	id, _ := c.Params.Get("eventID")
	eventID,_ := strconv.Atoi(id)

	list, _ := w.WorkshopUsecase.FetchList(int64(eventID))

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, list)

	return nil
}

// Detail
func (w *WorkshopHandler) Detail(c *gin.Context) error {
	id, _ := c.Params.Get("id")
	workshopID,_ := strconv.Atoi(id)

	workshop, _ := w.WorkshopUsecase.Detail(int64(workshopID))

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, workshop)

	return nil
}