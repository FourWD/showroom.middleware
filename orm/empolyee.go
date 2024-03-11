package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Employee struct {
	midOrm.Employee
	Signature string `json:"signature" query:"signature" gorm:"type:varchar(500)"`
}
