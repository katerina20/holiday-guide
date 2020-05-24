package holiday

import (
	"../http"
	"fmt"
	"time"
)

const countryCode = "UA"
const apiURL = "https://date.nager.at/api/v2/publicholidays/%v/%v"

var currentTime = time.Now()
var currYear = currentTime.Year

func getData() []Holiday {
	return parseHolidayFromJson(http.RequestApi(
		fmt.Sprintf(apiURL, currYear(), countryCode)))
}

func InfoHoliday() {
	var holidays = getData()
	fmt.Println(getNextHoliday(holidays))

}

func getNextHoliday(holidays []Holiday) Holiday {
	var holidayDate time.Time
	for _, holiday := range holidays {
		holidayDate = StringToDate(holiday.Date)
		if currentTime.Equal(holidayDate) || currentTime.Before(holidayDate) {
			return holiday
		}
	}
	return Holiday{}
}
