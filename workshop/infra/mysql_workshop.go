package infra

import (
	"event-management-service/workshop"
    "time"
	"gorm.io/gorm"
)

// WorkshopRepository workshop related repository
type WorkshopRepository struct {
	db *gorm.DB
}

// NewWorkshopRepository instantiation
func NewWorkshopRepository(db *gorm.DB) workshop.IWorkshopRepository {
	return &WorkshopRepository{
		db: db,
	}
}

// TxStart transaction start
func (w *WorkshopRepository) TxStart() (*gorm.DB, error) {
	tx := w.db.Begin()
	return tx, tx.Error
}

// TxCommit transaction commit
func (w *WorkshopRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

// TxRollback transaction rollback
func (w *WorkshopRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

// FetchAll Acquire multiple workshops
func (w *WorkshopRepository) FetchAllByEventID(eventID int64) ([]workshop.WorkshopTable, error) {
	result := []workshop.WorkshopTable{}
	err := w.db.
		Table("workshops").
		Where("event_id = ?", eventID).
		Where("start_at > ?", time.Now().Format("2006-01-02")).
	    Find(&result).Error
	return result, err
}

// FetchOne Get one event with event_id
func (w *WorkshopRepository) FetchOne(workshopID int64) (workshop.DetailOutput, error) {
	result := workshop.DetailOutput{}
	err := w.db.
	    Select("count(res.`id`) as total_reservations, w.id, w.title, w.description, w.start_at, w.end_at").
		Table("workshops as w").
		Joins("INNER JOIN reservations as res ON res.workshop_id = w.id").
		Where("w.id = ?", workshopID).
		Group("res.workshop_id").
		First(&result).Error
	return result, err
}

func (w *WorkshopRepository) FetchWorkshopDetail(workshopID int64)(workshop.WorkshopTable, error){
	result := workshop.WorkshopTable{}
	err := 	w.db.
			Table("workshops").
			Where("id = ?", workshopID).
			First(&result).Error
	return result, err
}
