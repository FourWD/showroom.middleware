package orm

import (
	"github.com/FourWD/middleware/model"
)

type Booking struct { ///
	// midOrm.Booking //
	ID string `json:"id" query:"id" gorm:"type:varchar(36); uniqueIndex:idx_id "` //
	model.GormModel
	BookingNo       string `json:"booking_no" query:"booking_no" gorm:"type:varchar(20)"` //หมายเลขรายการ
	BookingStatusID string `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`
}
