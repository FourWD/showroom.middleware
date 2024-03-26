package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Insurance struct {
	midOrm.Insurance
	FundPercent int `json:"fund_percent" query:"fund_percent" gorm:"type:int"`
}
