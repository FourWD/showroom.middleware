package orm

import (
	"time"

	midOrm "github.com/FourWD/middleware/orm"
)

type Booking struct {
	midOrm.Booking
	VehicleYear       string    `json:"vehicle_year" query:"vehicle_year" gorm:"type:varchar(4)"`
	IDCard            string    `json:"id_card" query:"id_card" gorm:"type:varchar(13)"`
	CompanyRegisterNo string    `json:"company_register_no" query:"company_register_no" gorm:"type:varchar(13)"`
	BirthDate         time.Time `json:"birth_date" query:"birth_date" gorm:"type:date"`
	OccupationID      string    `json:"occupation_id" query:"occupation_id" gorm:"type:varchar(255)"`
	Telephone         string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`
	CustomerSourceID  string    `json:"customer_source_id" query:"customer_source_id" gorm:"type:varchar(36)"`
	CustomerTypeID    string    `json:"customer_type_id" query:"customer_type_id" gorm:"type:varchar(2)"`
	CustomerSignature string    `json:"customer_signature" query:"customer_signature" gorm:"type:varchar(500)"`
	EmployeeID        string    `json:"employee_id" query:"employee_id" gorm:"type:varchar(36)"`
	ManagerID         string    `json:"manager_id" query:"manager_id" gorm:"type:varchar(36)"`
	// EmployeeSignature string    `json:"employee_signature" query:"employee_signature" gorm:"type:varchar(500)"`
	// ManagerSignature  string    `json:"manager_signature" query:"manager_signature" gorm:"type:varchar(500)"`
}
