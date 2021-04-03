package dateformat

import "time"

//return the current date for the given format
func DateFormatterDB(format string) string {

	currentTime := time.Now()

	return currentTime.Format(format)
}
