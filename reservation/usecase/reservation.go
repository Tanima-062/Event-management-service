package usecase

import (
	"event-management-service/event"
	eventInfra "event-management-service/event/infra"
	"event-management-service/reservation"
	reservationInfra "event-management-service/reservation/infra"
	"event-management-service/workshop"
	workshopInfra "event-management-service/workshop/infra"

	"gorm.io/gorm"
)

// ReservationUsecase event related usecase
type ReservationUsecase struct {
	ReservationRepository reservation.IReservationRepository
	EventRepository       event.IEventRepository
	WorkshopRepository    workshop.IWorkshopRepository
}

// NewReservationUsecase instantiation
func NewReservationUsecase(db *gorm.DB) reservation.IReservationUsecase {
	return &ReservationUsecase{
		ReservationRepository: reservationInfra.NewReservationRepository(db),
		EventRepository:       eventInfra.NewEventRepository(db),
		WorkshopRepository:    workshopInfra.NewWorkshopRepository(db),
	}
}

// Create Reservation
func (r *ReservationUsecase) Create(request *reservation.SaveInput, workshopID int64) (*reservation.SaveOutput, error) {
	response := &reservation.SaveOutput{}

	reservationTable := &reservation.ReservationTable{}
	reservationTable.Name = request.Name
	reservationTable.Email = request.Email
	reservationTable.WorkshopID = workshopID

	if err := r.ReservationRepository.Create(reservationTable); err != nil {
		return response, err
	}

	reservationOutput, err := r.ReservationRepository.GetLastReservation()
	if err != nil {
		return response, err
	}

	response.Reservation = reservationOutput

	workshopDetail, _ := r.WorkshopRepository.FetchWorkshopDetail(workshopID)

	workshopResult := &workshop.WorkshopOutput{}
	workshopResult.ID = workshopDetail.ID
	workshopResult.Title = workshopDetail.Title
	workshopResult.Description = workshopDetail.Description
	workshopResult.StartAt = workshopDetail.StartAt
	workshopResult.EndAt = workshopDetail.EndAt

	response.Workshop = *workshopResult

	eventDetail, _ := r.EventRepository.FetchEventByID(workshopDetail.EventID)

	response.Event = eventDetail.EventTable

	return response, err
}
