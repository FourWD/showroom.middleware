package orm

import "github.com/FourWD/middleware/orm"

type TempBooking struct {
	orm.Booking
	VehicleYear string `json:"vehicle_year" query:"vehicle_year" gorm:"type:varchar(4)"`
}
