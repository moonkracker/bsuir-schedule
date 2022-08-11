package helpers

import (
	"reflect"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"

	"bsuir-schedule/structures"
)

const (
	LabColor       = text.FgRed
	LectureColor   = text.FgGreen
	PracticalColor = text.FgYellow
)

func GetLessonColor(lessonType string) text.Color {
	switch lessonType {
	case "ЛК":
		return LectureColor
	case "ПЗ":
		return PracticalColor
	case "ЛР":
		return LabColor
	default:
		return text.FgWhite
	}
}

func GetCurrentWeekDay() string {
	weekday := time.Now().Weekday()
	switch weekday {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	default:
		return ""
	}
}

func Int32SliceContains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckIfMapIsEmpty[T any](m map[string]T) bool {
	for _, v := range m {
		if !reflect.ValueOf(v).IsZero() {
			return false
		}
	}
	return true
}

func IsEmptyGroupDay(weeknumber int32, SubgroupsNumber []int32, data structures.APIDay) bool {
	cnt := 0
	for _, v := range data {
		if !Int32SliceContains(v.WeekNumber, weeknumber) {
			continue
		}
		if !Int32SliceContains(SubgroupsNumber, int32(v.NumSubgroup)) {
			continue
		}
		cnt++
	}
	return cnt == 0
}

func IsEmptyTeacherDay(weeknumber int32, data structures.APIDay) bool {
	cnt := 0
	for _, v := range data {
		if !Int32SliceContains(v.WeekNumber, weeknumber) {
			continue
		}
		cnt++
	}
	return cnt == 0
}

func IsEmptyTeacherWeek(weeknumber int32, data map[string]structures.APIDay) bool {
	cnt := 0
	for _, weekday := range []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"} {
		for _, v := range data[weekday] {
			if !Int32SliceContains(v.WeekNumber, weeknumber) {
				continue
			}
			cnt++
		}
	}
	return cnt == 0
}
