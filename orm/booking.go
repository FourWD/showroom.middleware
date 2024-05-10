package orm

import (
	"time"

	midOrm "github.com/FourWD/middleware/orm"
)

type Booking struct {
	midOrm.Booking
	VehicleYear           string    `json:"vehicle_year" query:"vehicle_year" gorm:"type:varchar(4)"`
	VehicleImg            string    `json:"vehicle_img" query:"vehicle_img" gorm:"type:text"`
	IDCard                string    `json:"id_card" query:"id_card" gorm:"type:varchar(13)"`
	CompanyRegisterNo     string    `json:"company_register_no" query:"company_register_no" gorm:"type:varchar(13)"`
	BirthDate             time.Time `json:"birth_date" query:"birth_date" gorm:"type:date"`
	OccupationID          string    `json:"occupation_id" query:"occupation_id" gorm:"type:varchar(255)"`
	Telephone             string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`
	CustomerSourceID      string    `json:"customer_source_id" query:"customer_source_id" gorm:"type:varchar(36)"`
	CustomerTypeID        string    `json:"customer_type_id" query:"customer_type_id" gorm:"type:varchar(2)"`
	CustomerSignature     string    `json:"customer_signature" query:"customer_signature" gorm:"type:varchar(500)"`
	CustomerSignatureDate time.Time `json:"customer_signature_date" query:"customer_signature_date"`
	EmployeeSignature     string    `json:"employee_signature" query:"employee_signature" gorm:"type:varchar(500)"`
	EmployeeSignatureDate time.Time `json:"employee_signature_date" query:"employee_signature_date"`
	ManagerSignature      string    `json:"manager_signature" query:"manager_signature" gorm:"type:varchar(500)"`
	ManagerSignatureDate  time.Time `json:"manager_signature_date" query:"manager_signature_date"`
	Moo                   string    `json:"moo" query:"moo" gorm:"type:varchar(4)"`
	Soi                   string    `json:"soi" query:"soi" gorm:"type:varchar(100)"`
	Floor                 string    `json:"floor" query:"floor" gorm:"type:varchar(3)"`

	CompanyRegisterTypeID           string    `json:"company_register_type_id" query:"company_register_type_id" gorm:"type:varchar(2)"`
	EmployeeID                      string    `json:"employee_id" query:"employee_id" gorm:"type:varchar(36)"`
	ManagerID                       string    `json:"manager_id" query:"manager_id" gorm:"type:varchar(36)"`
	TelephoneCompany                string    `json:"telephone_company" query:"telephone_company" gorm:"type:varchar(10)"`
	IsConsent                       bool      `json:"is_consent" query:"is_consent" gorm:"type:bool"`
	CustomerSignatureLink           string    `json:"customer_signature_link" query:"customer_signature_link" gorm:"type:varchar(500)"`
	CustomerPDFLink                 string    `json:"customer_pdf_link" query:"customer_pdf_link" gorm:"type:varchar(500)"`
	PDFLink                         string    `json:"pdf_link" query:"pdf_link" gorm:"type:varchar(500)"`
	CustomerSignatureLinkSendDate   time.Time `json:"customer_signature_link_send_date" query:"customer_signature_link_send_date" gorm:"type:date"`
	CustomerSignatureLinkExpireDate time.Time `json:"customer_signature_link_expire_date" query:"customer_signature_link_expire_date" gorm:"type:date"`
	IsSendBookingPDF                bool      `json:"is_send_booking_pdf" query:"is_send_booking_pdf" gorm:"type:bool"`
	SendBookingPDFDate              time.Time `json:"send_booking_pdf_date" query:"send_booking_pdf_date" gorm:"type:date"`
	PDFLinkExpireDate               time.Time `json:"pdf_link_expire_date" query:"pdf_link_expire_date" gorm:"type:date"`
}
