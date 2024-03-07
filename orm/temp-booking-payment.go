package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type TempBookingPayment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key"`
	model.GormModel

	BookingID              string    `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	TotalPrice             float64   `json:"total_price" query:"total_price" gorm:"type:decimal(14,2)"`
	DownPrice              float64   `json:"down_price" query:"down_price" gorm:"type:decimal(14,2)"`
	Downpercent            float64   `json:"down_precent" query:"down_precent" gorm:"type:decimal(14,2)"`
	FinancePrice           float64   `json:"finance_price" query:"finance_price" gorm:"type:decimal(14,2)"`
	FinanceInstallment     int       `json:"finance_installment" query:"finance_installment" gorm:"type:int(3)"`
	InterestTypeID         string    `json:"interest_type_id" query:"interest_type_id" gorm:"type:varchar(36)"`
	SumInsured             float64   `json:"sum_insured" query:"sum_insured" gorm:"type:decimal(14,2)"`
	FinanceCompanyID       string    `json:"finance_company_id" query:"finance_company_id" gorm:"type:varchar(10)"`
	FinanceInterest        float64   `json:"finance_interest" query:"finance_interest" gorm:"type:decimal(3,3)"`
	FinanceInstallmentRate float64   `json:"finance_installment_rate" query:"finance_installment_rate" gorm:"type:decimal(14,2)"`
	InsuranceCompanyID     string    `json:"insurance_company_id" query:"insurance_company_id" gorm:"type:varchar(10)"`
	PaymentTypeID          string    `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`
	DepositPrice           float64   `json:"deposit_price" query:"deposit_pricet" gorm:"type:decimal(14,2)"`
	BankID                 string    `json:"bank_id" query:"bank_id" gorm:"type:varchar(2)"`
	CreditNo               string    `json:"credit_no" query:"credit_no" gorm:"type:varchar(36)"`
	CreditExpireDate       string    `json:"credit_expire_date" query:"credit_expire_date" gorm:"type:varchar(36)"`
	TransferDate           time.Time `json:"transfer_date" query:"transfer_date"`
	FileSlip               string    `json:"file_slip" query:"file_slip" gorm:"type:varchar(500)"`
	model.GormRowOrder
}
