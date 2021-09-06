package apiHelpers

import "time"

func ParseTime(date string) (time.Time, error) {

	if date == "" {
		return time.Time{}, nil
	}

	return time.Parse("02-Jan-2006-15-04-05", date)
}

//ParseTime2 returns date as a pointer
func ParseTime2(date string) (*time.Time, error) {

	if date == "" {
		return nil, nil
	}

	t, err := time.Parse("02-Jan-2006-15-04-05", date)
	if err != nil {
		return nil, err
	}

	if t.IsZero() {
		return nil, nil
	}

	return &t, nil
}

//FormatTime converts a time.Time value to a neocore timestamp string
func FormatTime(t time.Time) string {
	return t.Format("02-Jan-2006-15-04-05")
}
