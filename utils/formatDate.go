package utils

import "time"

func FormatDate(date string) (string, error) {
	time, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		return "", err
	}

	return time.Format("Jan 2, 2006 3:04 PM"), nil
}