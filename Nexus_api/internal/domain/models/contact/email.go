package contact

type Email struct {
	Value string `gorm:"column:email;not null;unique;size:255"`
}

func NewEmail(value string) (Email, []*models.ErrorMessage){
	return Email{
		Value: value,
	}, nil
}