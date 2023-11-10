package handler

import (
	"event-management-service/reservation"
	"event-management-service/reservation/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)
type ReservationHandler struct {
	ReservationUsecase   reservation.IReservationUsecase
}

// NewReservationHandler 
func NewReservationHandler(db *gorm.DB) *ReservationHandler {
	return &ReservationHandler{
		ReservationUsecase:   usecase.NewReservationUsecase(db),
	}
}

func (r *ReservationHandler) Create(c *gin.Context) error {
	id, _ := c.Params.Get("workshopID")
	workshopID,_ := strconv.Atoi(id)

	request := &reservation.SaveInput{}
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	response, _ := r.ReservationUsecase.Create(request, int64(workshopID))

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

}