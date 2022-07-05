package entity

//User struct represents User table in DB
type User struct {
	ID       uint64  `gorm:"type:int;primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(100)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Books    *[]Book `json:"books,omitempty"`
}
