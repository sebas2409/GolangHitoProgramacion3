package entities

type User struct {
	Id     int `gorm:"primary_key;auto_increment"`
	Nombre string
}
