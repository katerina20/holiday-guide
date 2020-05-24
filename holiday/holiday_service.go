package holiday

import (
	"../http"
	"fmt"
	"time"
)

const countryCode = "UA"
const apiURL = "https://date.nager.at/api/v2/publicholidays/%v/%v"
const weekendApiURL = "https://date.nager.at/Api/v2/LongWeekend/%v/%v"
const stringInfo = "\n%v holiday is %v, %v %v"
const weekendString = "The weekend will last %v days: %v %v - %v %v"

const stringNext = "The next"
const stringToday = "Today's"

//var currentTime = StringToDate("2020-05-01")
var currentTime = time.Now()
var currYear = currentTime.Year

func InfoHoliday() {
	var holidays = getData()
	var holiday, isToday = getNextHoliday(holidays)
	printInfo(holiday, isToday)
	checkWeekend(holiday)
}

func getData() []Holiday {
	return parseHolidayFromJson(http.RequestApi(
		fmt.Sprintf(apiURL, currYear(), countryCode)))
}

func getNextHoliday(holidays []Holiday) (Holiday, bool) {
	var holidayDate time.Time
	for _, holiday := range holidays {
		holidayDate = StringToDate(holiday.Date)
		if currentTime.Equal(holidayDate) {
			return holiday, true
		} else if currentTime.Before(holidayDate) {
			return holiday, false
		}
	}
	return Holiday{}, false
}

func printInfo(holiday Holiday, isToday bool) {
	var whenStr string
	if isToday {
		whenStr = stringToday
	} else {
		whenStr = stringNext
	}

	fmt.Println(fmt.Sprintf(stringInfo, whenStr,
		holiday.Name,
		StringToDate(holiday.Date).Month(),
		StringToDate(holiday.Date).Day()))
}

func checkWeekend(holiday Holiday) {
	var isHolidayAfterStart, isHolidayBeforeEnd bool
	var holidayDate = StringToDate(holiday.Date)

	var weekends = parseWeekendFromJson(http.RequestApi(
		fmt.Sprintf(weekendApiURL, currYear(), countryCode)))

	for _, weekend := range weekends {
		isHolidayAfterStart = holidayDate.After(StringToDate(weekend.StartDate)) ||
			holidayDate.Equal(StringToDate(weekend.StartDate))
		isHolidayBeforeEnd = holidayDate.Before(StringToDate(weekend.EndDate)) ||
			holidayDate.Equal(StringToDate(weekend.EndDate))

		if isHolidayAfterStart && isHolidayBeforeEnd {
			fmt.Printf(weekendString,
				weekend.DayCount,
				StringToDate(weekend.StartDate).Month(),
				StringToDate(weekend.StartDate).Day(),
				StringToDate(weekend.EndDate).Month(),
				StringToDate(weekend.EndDate).Day())
			return
		}
	}
}
