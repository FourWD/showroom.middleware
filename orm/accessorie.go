package orm

import "github.com/FourWD/middleware/model"

// midOrm "github.com/FourWD/middleware/orm"

type Accessorie struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	Name             string  `json:"name" query:"name" gorm:"type:varchar(255)"`
	AccessorieTypeID string  `json:"accessorie_type_id" query:"accessorie_type_id" gorm:"type:varchar(2)"`
	PricePreVat      float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat              float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price            float64 `json:"price" query:"price" gorm:"type:decimal(14,2)"`
	model.GormRowOrder
}
