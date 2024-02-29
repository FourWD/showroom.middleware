package orm

import "github.com/FourWD/middleware/model"

type VehicleSubModelColors struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id;primary_key"`
	model.GormModel

	VehicleSubModelID string `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	VehicleColorID    string `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36)"`
	YearManufacturing int    `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:varchar(4)"`
	ColorCode         string `json:"color_code" query:"color_code" gorm:"type:varchar(10)"`
	model.GormRowOrder
}
