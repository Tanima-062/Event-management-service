package event

import (
	"time"
	"event-management-service/common"
)

// Event Table
type Events struct {
	EventTable `gorm:"embedded"`
}

type EventTable struct {
	ID     		int64     		`gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Title  		string    		`json:"title" validate:"required"`
	StartAt     time.Time 		`gorm:"type:time" json:"start_at"`
	EndAt       time.Time 		`gorm:"type:time" json:"end_at"`
}

// ListInput
type ListInput struct {
	common.Paging
}

type ListOutput struct {
	Events       []EventTable    	`json:"events"`
	Pagination   common.Pagination	`json:"pagination"`
}

// DetailOutput
type DetailOutput struct {
	EventTable
	TotalWorkshops    int64     `json:"total_workshops"`
}

type IEventUsecase interface {
	FetchList(request *ListInput) (*ListOutput, error)
	Detail(eventID int64) (*DetailOutput, error)
}

// IEventRepository
type IEventRepository interface {
	common.Repository

	FetchAll(req ListInput) ([]Events, error)
	FetchEventByID(eventID int64) (Events, error)
	FetchOne(eventID int64)(DetailOutput, error)
	TotalEvents()int64
}