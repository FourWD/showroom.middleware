package orm

import "github.com/FourWD/middleware/model"

type CustomerSource struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id;primary_key"`
	model.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(200)"`

	model.GormRowOrder
}
