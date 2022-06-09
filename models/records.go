package models

import (
	"time"

	"github.com/lib/pq"
)

type Records struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	Name      string        `json:"name"`
	Marks     pq.Int64Array `json:"marks" gorm:"type:integer[]"`
	CreatedAt time.Time     `json:"createdAt"`
}
