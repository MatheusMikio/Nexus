package dates

import "time"

type GoalDates struct {
	startDate            time.Time `gorm:"not null"`
	FinalizationForecast *time.Time
}

type TaskDates struct {
	startDate        time.Time `gorm:"not null"`
	FinalizationDate *time.Time
	TimeSpent        *time.Time
}
