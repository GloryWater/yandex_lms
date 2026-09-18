package main

import (
	"errors"
	"strings"
	"time"
)

func currentDayOfTheWeek() string {
	switch TimeNow().Weekday() {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 0:
		return "Воскресенье"
	default:
		return ""
	}
}

func dayOrNight() string {
	hourNow := TimeNow().Hour()
	if hourNow >= 10 && hourNow <= 22 {
		return "День"
	} else {
		return "Ночь"
	}
}

func nextFriday() int {
	return int(time.Friday) - int(TimeNow().Weekday())
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	if strings.EqualFold(answer, currentDayOfTheWeek()) {
		return true
	} else {
		return false
	}
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if len(answer) == 8 {
		if strings.EqualFold(answer, dayOrNight()) {
			return true, nil
		} else {
			return false, nil
		}
	} else {

		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
}
