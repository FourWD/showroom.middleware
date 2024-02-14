package orm

type Users struct {
	Username string `json:"username" query:"username" gorm:"type:varchar(20)"`
	Password string `json:"password" query:"password" gorm:"type:varchar(20)"`
}
