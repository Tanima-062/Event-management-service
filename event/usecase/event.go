package usecase

import (
	"event-management-service/common"
	"event-management-service/event"
	eventInfra "event-management-service/event/infra"
	"log"

	"gorm.io/gorm"
)

// EventUsecase event related usecase
type EventUsecase struct {
	EventRepository event.IEventRepository
}

// NewEventUsecase instantiation
func NewEventUsecase(db *gorm.DB) event.IEventUsecase {
	return &EventUsecase{
		EventRepository:  eventInfra.NewEventRepository(db),
	}
}

// FetchList Get event list
func (e *EventUsecase) FetchList(request *event.ListInput) (*event.ListOutput, error) {

	eventCh := make(chan []event.Events)
	paginationCh := make(chan common.Pagination)
    go e.fetchEvents(eventCh, request)
	go e.pagination(paginationCh, request)
	eventList, pagination := <-eventCh, <- paginationCh

	response := event.ListOutput{}
	response.Pagination = pagination
	responseData := []event.EventTable{}
	for _, eventData := range eventList {
		record := &event.EventTable{}
		record.ID = eventData.EventTable.ID;
		record.Title = eventData.EventTable.Title;
		record.StartAt = eventData.EventTable.StartAt;
		record.EndAt = eventData.EventTable.EndAt;
		responseData = append(responseData, *record)
	}

	response.Events = responseData
	return &response, nil
}

// Detail Event details
func (e *EventUsecase) Detail(eventID int64) (*event.DetailOutput, error) {
	eventDetail, err := e.EventRepository.FetchOne(eventID)
	if err != nil{
		log.Fatal(err)
	}
	return &eventDetail, nil
}

func (e *EventUsecase) fetchEvents(ch chan<- []event.Events, request *event.ListInput) {
	eventR, eventErr := e.EventRepository.FetchAll(*request)
	if eventErr != nil {
		ch <- []event.Events{}
	}
	ch <- eventR
}

func(e *EventUsecase) pagination(ch chan<- common.Pagination, request *event.ListInput){
	totalEvents := e.EventRepository.TotalEvents()

	pagination := &common.Pagination{}
	pagination.Total = totalEvents
	pagination.PerPage = request.Paging.Limit
	pagination.TotalPages = totalEvents/ int64(request.Paging.Limit)
	pagination.CurrentPage = int64(request.Paging.Page)

	ch <- *pagination

}