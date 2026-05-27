package contact

type Contact struct {
	Email Email `gorm:"embedded"`
	Phone Phone `gorm:"embedded"`
}


func NewContact(email Email, phone Phone) (Contact, []*models.ErrorMessage) {
	return Contact{
		Email: email,
		Phone: phone,
	}, nil
}