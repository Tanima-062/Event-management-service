package common

import (
	"gorm.io/gorm"
)

// Paging
type Paging struct {
	Page   int `json:"current_page"`
	Limit  int `json:"limit" query:"limit"`
	Offset int `json:"offset" query:"offset"`
}

// Pagination
type Pagination struct {
	Total  		 int64 		   `json:"total"`
	PerPage 	 int 		   `json:"per_page"`
	TotalPages   int64         `json:"total_pages"`
	CurrentPage  int64         `json:"current_page"`
}


// Repository
type Repository interface {
	// TxStart
	TxStart() (*gorm.DB, error)
	// TxCommit
	TxCommit(tx *gorm.DB) error
	// TxRollback
	TxRollback(tx *gorm.DB)
}
