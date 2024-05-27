package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	midOrm.VehicleModel
	StartPrice int `json:"start_price" query:"start_price" gorm:"type:decimal(14,2)"`
}
