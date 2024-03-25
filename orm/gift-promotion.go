package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type GiftPromotion struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	Name           string    `json:"name" query:"name" gorm:"type:varchar(256)"`
	VehicleModelID string    `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	StartDate      time.Time `json:"start_date" query:"start_date"`
	EndDate        time.Time `json:"end_date" query:"end_date"`
	GiftID         string    `json:"gift_id" query:"gift_id" gorm:"type:varchar(36)"`
	model.GormRowOrder
}
