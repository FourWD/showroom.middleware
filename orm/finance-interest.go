package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type FinanceInterest struct {
	midOrm.Insurance
	VehicleModelID     string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	IsIncludeInsurance bool   `json:"is_include_insurance" query:"is_include_insurance" gorm:"type:bool"`
}
