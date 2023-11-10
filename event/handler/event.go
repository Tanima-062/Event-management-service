package handler

import (
	"event-management-service/event"
	"event-management-service/event/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventHandler struct {
	EventUsecase   event.IEventUsecase
}

// NewEventHandler 
func NewEventHandler(db *gorm.DB) *EventHandler {
	return &EventHandler{
		EventUsecase:   usecase.NewEventUsecase(db),
	}
}

// List 
func (e *EventHandler) List(c *gin.Context) error {
	request := &event.ListInput{}
	if err := c.Bind(request); err != nil {
		return err
	}
    pg := c.Query("page")
	page,_ := strconv.Atoi(pg)
	lt := c.Query("limit")
	limit,_ := strconv.Atoi(lt)

	offset := (page-1)*limit

	request.Paging.Page  = page
	request.Paging.Limit = limit
	request.Paging.Offset = offset

	list, _ := e.EventUsecase.FetchList(request)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, list)

	return nil
}

// Detail
func (e *EventHandler) Detail(c *gin.Context) error {
	id, _ := c.Params.Get("id")
	eventID,_ := strconv.Atoi(id)

	event, _ := e.EventUsecase.Detail(int64(eventID))

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, event)

	return nil
}