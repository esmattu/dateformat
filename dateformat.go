package dateformat

import "time"

//return the current date for the given format
func DateFormatterDB(format string) string {

	currentTime := time.Now()
	currentTime.Format(format)

	return currentTime.String()
}
