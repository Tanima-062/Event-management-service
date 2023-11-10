package reservation

import (
	"event-management-service/event"
	"event-management-service/common"
	"event-management-service/workshop"
)

// Reservation Table
type Reservations struct {
	ReservationTable `gorm:"embedded"`
}

// Reservation Table
type ReservationTable struct {
	ID     			int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name  		 	string    		`json:"name"`
	Email  		 	string    		`json:"email"`
	WorkshopID      int64           
}

type ReservationOutput struct {
	ID     			int64     		`json:"id"`
	Name  		 	string    		`json:"name"`
	Email  		 	string    		`json:"email"`
}

type SaveInput struct{
	Name  		 	string    		`json:"name"  binding:"required"`
	Email  		 	string    		`json:"email" binding:"required"`
}

type SaveOutput struct{
	Reservation  ReservationOutput   			`json:"reservation"`
	Event        event.EventTable    			`json:"event"`
	Workshop     workshop.WorkshopOutput      	`json:"workshop"`
}
type IReservationUsecase interface {
	Create(request *SaveInput, workshop int64) (*SaveOutput, error)
}

// IReservationRepository
type IReservationRepository interface {
	common.Repository
	Create(reservationTable *Reservations) error
	GetLastReservation() (ReservationOutput, error)
}