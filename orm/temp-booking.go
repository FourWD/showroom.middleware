package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type TempBooking struct {
	orm.Booking
	VehicleYear      string    `json:"vehicle_year" query:"vehicle_year" gorm:"type:varchar(4)"`
	IDCard           string    `json:"id_card" query:"id_card" gorm:"type:varchar(13)"`
	BirthDate        time.Time `json:"birth_date" query:"birth_date"`
	Occupation       string    `json:"occupation" query:"occupation" gorm:"type:varchar(255)"`
	Telephone        string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`
	CustomerSourceID string    `json:"customer_source_id" query:"customer_source_id" gorm:"type:varchar(36)"`
}
