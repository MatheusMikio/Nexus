package contact

type Phone struct {
	Value string `gorm:"column:phone;not null;size:11"`
}

func NewPhone(value string) (Phone, []*models.ErrorMessage) {
	return Phone{
		Value: value,
	}, nil
}