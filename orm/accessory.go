package orm

import midOrm "github.com/FourWD/middleware/orm"

type Accessory struct {
	midOrm.Accessory
	ShopID    string `json:"shop_id" query:"shop_id" gorm:"type:varchar(36)"`
	ImagePath string `json:"image_path" query:"image_path" gorm:"type:varchar(36)"`
}
