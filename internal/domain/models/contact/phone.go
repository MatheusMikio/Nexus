package contact

type Phone struct {
	Value string `gorm:"column:phone;not null;size:11"`
}
