package dates

import "time"

type GoalDates struct {
	startDate            time.Time `gorm:"not null"`
	FinalizationForecast *time.Time
}

type TaskDates struct {
	startDate        time.Time `gorm:"not null"`
	FinalizationDate *time.Time
	TimeSpent        *int64
}


func NewGoalDates(startDate time.Time, finalizationForecast *time.Time) (GoalDates, []*models.ErrorMessage) {
	return GoalDates{
		startDate:            startDate,
		FinalizationForecast: finalizationForecast,
	}, nil
}

func NewTaskDates(startDate time.Time, finalizationDate *time.Time) (TaskDates, []*models.ErrorMessage) {
	var timeSpent *int64
	
	if finalizationDate != nil {
		timeSpent := int64(finalizationDate.Sub(startDate).Minutes())
	}

	return TaskDates{
		startDate:        startDate,
		FinalizationDate: finalizationDate,
		TimeSpent:        timeSpent,
	}, nil
}