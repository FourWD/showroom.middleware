package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BookingPayment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key"`
	model.GormModel

	IsAccessory            bool      `json:"is_accessory" query:"is_accessory" gorm:"type:bool"`
	BookingID              string    `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	VehiclePrice           float64   `json:"vehicle_price" query:"vehicle_price" gorm:"type:decimal(14,2)"`
	VehicleDiscountPrice   float64   `json:"vehicle_discount_price" query:"vehicle_discount_price" gorm:"type:decimal(14,2)"`
	VehicleTotalPrice      float64   `json:"vehicle_total_price" query:"vehicle_total_price" gorm:"type:decimal(14,2)"`
	AccessoryPrice         float64   `json:"accessory_price" query:"accessory_price" gorm:"type:decimal(14,2)"`
	AccessoryDiscountPrice float64   `json:"accessory_discount_price" query:"accessory_discount_price" gorm:"type:decimal(14,2)"`
	AccessoryTotalPrice    float64   `json:"accessory_total_price" query:"accessory_total_pric" gorm:"type:decimal(14,2)"`
	TotalPrice             float64   `json:"total_price" query:"total_price" gorm:"type:decimal(14,2)"`
	DownPrice              float64   `json:"down_price" query:"down_price" gorm:"type:decimal(14,2)"`
	DownPercent            float64   `json:"down_precent" query:"down_precent" gorm:"type:decimal(14,2)"`
	Discount               float64   `json:"discount" query:"discount" gorm:"type:decimal(14,2)"`
	FinancePrice           float64   `json:"finance_price" query:"finance_price" gorm:"type:decimal(14,2)"`
	FinanceMonth           int       `json:"finance_month" query:"finance_month" gorm:"type:int(3)"`
	FinancePricePerMonth   float64   `json:"finance_price_per_month" query:"finance_price_per_month" gorm:"type:decimal(14,2)"`
	InterestTypeID         string    `json:"interest_type_id" query:"interest_type_id" gorm:"type:varchar(36)"`
	InsuranceFund          float64   `json:"insurance_fund" query:"insurance_fund" gorm:"type:decimal(14,2)"` //
	FinanceID              string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	FinanceInterest        float64   `json:"finance_interest" query:"finance_interest" gorm:"type:decimal(5,3)"`
	InsuranceID            string    `json:"insurance_id" query:"insurance_id" gorm:"type:varchar(10)"`
	PaymentTypeID          string    `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`
	DepositPrice           float64   `json:"deposit_price" query:"deposit_pricet" gorm:"type:decimal(14,2)"`
	BankID                 string    `json:"bank_id" query:"bank_id" gorm:"type:varchar(2)"`
	CreditCardNo           string    `json:"credit_card_no" query:"credit_card_no" gorm:"type:varchar(36)"` //mark-up
	CreditCardExpireDate   string    `json:"credit_card_expire_date" query:"credit_card_expire_date" gorm:"type:varchar(36)"`
	TransferDate           time.Time `json:"transfer_date" query:"transfer_date"`
	// PickupDate             time.Time `json:"pickup_date" query:"pickup_date"`
	// ExpectPickupDate       time.Time `json:"expect_pickup_date" query:"expect_pickup_date"`
	FileSlip     string    `json:"file_slip" query:"file_slip" gorm:"type:varchar(500)"`
	DeliveryDate time.Time `json:"delivery_date" query:"delivery_date" gorm:"type:date"`
	model.GormRowOrder
}
