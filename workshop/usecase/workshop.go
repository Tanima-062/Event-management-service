package usecase

import (
	"event-management-service/event"
	eventInfra "event-management-service/event/infra"
	"event-management-service/workshop"
	workshopInfra "event-management-service/workshop/infra"
	"log"

	"gorm.io/gorm"
)

// WorkshopUsecase event related usecase
type WorkshopUsecase struct {
	WorkshopRepository workshop.IWorkshopRepository
	EventRepository    event.IEventRepository
}

// NewWorkshopUsecase instantiation
func NewWorkshopUsecase(db *gorm.DB) workshop.IWorkshopUsecase {
	return &WorkshopUsecase{
		WorkshopRepository: workshopInfra.NewWorkshopRepository(db),
		EventRepository:    eventInfra.NewEventRepository(db),
	}
}

// FetchList Get workshop list
func (w *WorkshopUsecase) FetchList(eventID int64) (*workshop.ListOutput, error) {
	response := &workshop.ListOutput{}
	eventCh := make(chan event.Events)
	workshopCh := make(chan []workshop.WorkshopTable)
	go w.fetchEvent(eventCh, eventID)
	go w.fetchWorkshops(workshopCh, eventID)
	event, workshops := <-eventCh, <-workshopCh
	response.EventTable.ID = event.EventTable.ID
	response.EventTable.Title = event.EventTable.Title
	response.EventTable.StartAt = event.EventTable.StartAt
	response.EventTable.EndAt = event.EventTable.EndAt
	response.Workshops = workshops

	return response, nil
}

// Detail Workshop details
func (w *WorkshopUsecase) Detail(eventID int64) (*workshop.DetailOutput, error) {
	workshopDetail, err := w.WorkshopRepository.FetchOne(eventID)
	if err != nil {
		log.Fatal(err)
	}
	return &workshopDetail, nil
}

func (w *WorkshopUsecase) fetchEvent(ch chan<- event.Events, eventID int64) {
	eventR, eventErr := w.EventRepository.FetchEventByID(eventID)
	if eventErr != nil {
		ch <- event.Events{}
	}
	ch <- eventR
}

func (w *WorkshopUsecase) fetchWorkshops(ch chan<- []workshop.WorkshopTable, eventID int64) {
	workshops, workshopErr := w.WorkshopRepository.FetchAllByEventID(eventID)
	if workshopErr != nil {
		ch <- []workshop.WorkshopTable{}
	}
	ch <- workshops
}
