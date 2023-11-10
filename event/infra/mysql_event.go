package infra

import (
	"event-management-service/event"
    "time"
	"gorm.io/gorm"
)

// EventRepository event related repository
type EventRepository struct {
	db *gorm.DB
}

// NewEventRepository instantiation
func NewEventRepository(db *gorm.DB) event.IEventRepository {
	return &EventRepository{
		db: db,
	}
}

// TxStart transaction start
func (e *EventRepository) TxStart() (*gorm.DB, error) {
	tx := e.db.Begin()
	return tx, tx.Error
}

// TxCommit transaction commit
func (e *EventRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

// TxRollback transaction rollback
func (e *EventRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

// FetchAll Acquire multiple events with pagination
func (e *EventRepository) FetchAll(req event.ListInput) ([]event.Events, error) {
	result := []event.Events{}
	query := e.db.
		Table("events").
		Where("start_at > ?", time.Now().Format("2006-01-02"))
	if req.Paging.Limit > 0 {
		query = query.Limit(req.Paging.Limit).Offset(req.Paging.Offset)
	}
	err := query.Find(&result).Error
	return result, err
}

func (e *EventRepository) FetchEventByID(eventID int64) (event.Events, error) {
	result := event.Events{}
	err := e.db.
		Table("events").
		Where("id = ?", eventID).
		First(&result).Error
	return result, err
}

// FetchOne Get one event with event_id
func (e *EventRepository) FetchOne(eventID int64) (event.DetailOutput, error) {
	result := event.DetailOutput{}
	err := e.db.
	    Select("count(w.`id`) as total_workshops, ev.id, ev.title, ev.start_at, ev.end_at").
		Table("events as ev").
		Joins("INNER JOIN workshops as w ON w.event_id = ev.id").
		Where("ev.id = ?", eventID).
		Group("w.event_id").
		First(&result).Error
	return result, err
}

func(e *EventRepository) TotalEvents()int64{
	var result int64
	e.db.Table("events").Count(&result)
	return result
}
