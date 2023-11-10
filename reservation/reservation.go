package reservation

import (
	"event-management-service/event"
	"event-management-service/common"
	"event-management-service/workshop"
)

// Reservation Table
type ReservationTable struct {
	ID     			int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name  		 	string    		`json:"name" validate:"required"`
	Email  		 	string    		`json:"email" validate:"required"`
	WorkshopID      int64           
}

type ReservationOutput struct {
	ID     			int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name  		 	string    		`json:"name" validate:"required"`
	Email  		 	string    		`json:"email" validate:"required"`
}

type SaveInput struct{
	Name  		 	string    		`json:"name" validate:"required"`
	Email  		 	string    		`json:"email" validate:"required"`
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
	Create(reservationTable *ReservationTable) error
	GetLastReservation() (ReservationOutput, error)
}