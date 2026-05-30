package dates

import "time"

func DateOnlyIn(value time.Time, location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}

	value = value.In(location)
	year, month, day := value.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

func CompareDate(a, b time.Time, location *time.Location) int {
	aDate := DateOnlyIn(a, location)
	bDate := DateOnlyIn(b, location)

	switch {
	case aDate.Before(bDate):
		return -1
	case aDate.After(bDate):
		return 1
	default:
		return 0
	}
}
