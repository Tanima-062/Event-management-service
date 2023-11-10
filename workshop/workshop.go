package workshop

import (
	"event-management-service/common"
	"event-management-service/event"
	"time"
)

// Workshop Table
type WorkshopTable struct {
	ID     		int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	EventID     int64           `json:"event_id"`
	Title  		string    		`json:"title"`
	Description string    		`json:"description"`
	StartAt     time.Time 		`gorm:"type:time" json:"start_at"`
	EndAt       time.Time 		`gorm:"type:time" json:"end_at"`
}

type WorkshopOutput struct {
	ID     				int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Title  				string    		`json:"title"`
	Description 		string    		`json:"description"`
	StartAt     		time.Time 		`gorm:"type:time" json:"start_at"`
	EndAt       		time.Time 		`gorm:"type:time" json:"end_at"`
}

// DetailOutput
type DetailOutput struct {
	ID     				int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Title  				string    		`json:"title"`
	Description 		string    		`json:"description"`
	StartAt     		time.Time 		`gorm:"type:time" json:"start_at"`
	EndAt       		time.Time 		`gorm:"type:time" json:"end_at"`
	TotalReservations   int64     		`json:"total_reservations"`
}

type ListOutput struct {
	event.EventTable
	Workshops       []WorkshopTable    `json:"workshops"`
}

type IWorkshopUsecase interface {
	FetchList(eventID int64) (*ListOutput, error)
	Detail(eventID int64) (*DetailOutput, error)
}

// IWorkshopRepository
type IWorkshopRepository interface {
	common.Repository

	FetchAllByEventID(eventID int64) ([]WorkshopTable, error)
	FetchOne(workshopID int64)(DetailOutput, error)
	FetchWorkshopDetail(workshopID int64)(WorkshopTable, error)
}