package orm

type VehicleGift struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
}
