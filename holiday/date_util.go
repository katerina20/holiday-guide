package holiday

import "time"

const RFC3339FullDate = "2006-01-02"

func StringToDate(str string) time.Time {
	var date, _ = time.Parse(RFC3339FullDate, str)
	return date
}
