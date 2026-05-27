package contact

type Contact struct {
	Email Email `gorm:"embedded"`
	Phone Phone `gorm:"embedded"`
}
