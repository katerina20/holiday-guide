package holiday

import (
	"../http"
	"fmt"
	"time"
)

const countryCode = "UA"
const apiURL = "https://date.nager.at/api/v2/publicholidays/%v/%v"
const stringInfo = "\n%v holiday is %v, %v %v"
const weekendString = "The weekend will last 3 days: %v %v - %v %v"

const stringNext = "The next"
const stringToday = "Today's"

const fridayStr = "Friday"
const saturdayStr = "Saturday"
const sundayStr = "Sunday"
const mondayStr = "Monday"

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
	var holidayDate = StringToDate(holiday.Date)
	var weekDay = holidayDate.Weekday().String()

	var startDate time.Time
	var endDate time.Time

	switch weekDay {
	case fridayStr, saturdayStr:
		startDate = holidayDate
		endDate = holidayDate.AddDate(0, 0, 2)
	case sundayStr:
		startDate = holidayDate.AddDate(0, 0, -1)
		endDate = holidayDate.AddDate(0, 0, 1)
	case mondayStr:
		startDate = holidayDate.AddDate(0, 0, -2)
		endDate = holidayDate
	default:
		return
	}

	fmt.Printf(weekendString, startDate.Month(), startDate.Day(), endDate.Month(), endDate.Day())
}
