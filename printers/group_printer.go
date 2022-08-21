package printers

import (
	"fmt"
	"os"

	"bsuir-schedule/helpers"
	"bsuir-schedule/structures"

	"github.com/jedib0t/go-pretty/v6/table"
)

func getLessonAuditorie(auditories []string) string {
	if len(auditories) == 0 {
		return "Нет аудитории"
	}
	return auditories[0]
}

func getTeacherName(teacher structures.Employees) string {
	for _, v := range teacher {
		return v.LastName + " " + v.FirstName + " " + v.MiddleName
	}
	return ""
}

func getSubgroup(subgroup int) string {
	if subgroup == 0 {
		return "Общая"
	}
	return fmt.Sprintf("Подгруппа %d", subgroup)
}

func (sp *SchedulePrinter) WriteGroupDay(weeknumber int32, SubgroupsNumber []int32, data structures.APIDay) {
	pairNumber := 0
	var previousLessonTime, currentLessonTime string
	for _, v := range data {
		if !helpers.Int32SliceContains(v.WeekNumber, weeknumber) {
			continue
		}
		if !helpers.Int32SliceContains(SubgroupsNumber, int32(v.NumSubgroup)) {
			continue
		}
		currentLessonTime = v.StartLessonTime + "-" + v.EndLessonTime
		if currentLessonTime != previousLessonTime {
			previousLessonTime = currentLessonTime
			pairNumber += 1
		}
		sp.table.AppendRow(table.Row{
			pairNumber,
			helpers.GetLessonColor(v.LessonTypeAbbrev).Sprintf(v.LessonTypeAbbrev),
			currentLessonTime,
			v.Subject,
			getLessonAuditorie(v.Auditories),
			getTeacherName(v.Employees),
			getSubgroup(v.NumSubgroup),
		})
	}
}

func NewGroupPrinter() *SchedulePrinter {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"№", "Тип", "Время", "Предмет", "Аудитория", "Преподаватель", "Подгруппа"})
	t.SetStyle(table.StyleLight)
	t.SetColumnConfigs([]table.ColumnConfig{
		NumberField,
		TypeField,
		TimeField,
		SubjectField,
		AuditorieField,
		TeacherField,
		SubgroupField,
	})
	t.Style().Options.SeparateRows = true
	t.SuppressEmptyColumns()
	t.SetOutputMirror(os.Stdout)
	return &SchedulePrinter{
		table: t,
	}
}
