package orm

type Users struct {
	ID       string `json:"id" query:"id" gorm:"type:varchar(36)"`
	Username string `json:"username" query:"username" gorm:"type:varchar(20)"`
	Password string `json:"password" query:"password" gorm:"type:varchar(20)"`
}
