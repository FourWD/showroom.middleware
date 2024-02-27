package orm

import "github.com/FourWD/middleware/orm"

type TempBookingAccessory struct {
	orm.BookingAccessory
	GroupID string `json:"group_id" query:"group_id" gorm:"type:varchar(2)"`
	ShopID  string `json:"shop_id" query:"shop_id" gorm:"type:varchar(36)"`
}
