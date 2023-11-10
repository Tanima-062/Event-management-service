package handler

import (
	"event-management-service/reservation"
	"event-management-service/reservation/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReservationHandler struct {
	ReservationUsecase reservation.IReservationUsecase
}

// NewReservationHandler
func NewReservationHandler(db *gorm.DB) *ReservationHandler {
	return &ReservationHandler{
		ReservationUsecase: usecase.NewReservationUsecase(db),
	}
}

func (r *ReservationHandler) Create(c *gin.Context){
	id, _ := c.Params.Get("workshopID")
	workshopID, _ := strconv.Atoi(id)

	request := &reservation.SaveInput{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	response, _ := r.ReservationUsecase.Create(request, int64(workshopID))

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
	return
}
