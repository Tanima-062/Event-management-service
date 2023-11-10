package infra

import (
	"event-management-service/reservation"
	"gorm.io/gorm"
)

// ReservationRepository event related repository
type ReservationRepository struct {
	db *gorm.DB
}

// NewReservationRepository instantiation
func NewReservationRepository(db *gorm.DB) reservation.IReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

// TxStart transaction start
func (r *ReservationRepository) TxStart() (*gorm.DB, error) {
	tx := r.db.Begin()
	return tx, tx.Error
}

// TxCommit transaction commit
func (r *ReservationRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

// TxRollback transaction rollback
func (r *ReservationRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

// CreateReservation
func (r *ReservationRepository) Create(reservationTable *reservation.ReservationTable) error {
	return r.db.Create(&reservationTable ).Error
}

func (r *ReservationRepository) GetLastReservation() (reservation.ReservationOutput, error) {
	result := reservation.ReservationOutput{}
	err := r.db.Last(&result).Error
	return result, err
}
