package orm

import "github.com/FourWD/middleware/orm"

type BookingAccessory struct {
	orm.BookingAccessory
	GroupID    string  `json:"group_id" query:"group_id" gorm:"type:varchar(2)"`
	ShopID     string  `json:"shop_id" query:"shop_id" gorm:"type:varchar(36)"`
	IsTrade    bool    `json:"is_trade" query:"is_trade" gorm:"type:bool"`
	TradePrice float64 `json:"trade_price" query:"trade_price" gorm:"type:decimal(14,2)"`
}
