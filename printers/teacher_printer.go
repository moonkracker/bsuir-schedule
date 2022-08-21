package printers

import (
	"os"

	"bsuir-schedule/helpers"
	"bsuir-schedule/structures"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (sp *SchedulePrinter) WriteTeacherDay(weeknumber int32, data structures.APIDay) {
	pairNumber := 0
	var previousLessonTime, currentLessonTime string
	for _, v := range data {
		if !helpers.Int32SliceContains(v.WeekNumber, weeknumber) {
			continue
		}
		for group, speciality := range getGroupsAndSpecialityMap(v.StudentGroups) {
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
				group,
				speciality,
			})
		}
	}
}

func getGroupsAndSpecialityMap(groups structures.StudentGroups) map[string]string {
	var result = make(map[string]string)
	for _, v := range groups {
		result[v.Name] = v.SpecialityName
	}
	return result
}

func NewTeacherPrinter() *SchedulePrinter {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"№", "Тип", "Время", "Предмет", "Аудитория", "Группы", "Специальность"})
	t.SetStyle(table.StyleLight)
	t.SetColumnConfigs([]table.ColumnConfig{
		NumberField,
		TypeField,
		TimeField,
		SubjectField,
		AuditorieField,
		GroupsField,
		SpecialityField,
	})
	t.Style().Options.SeparateRows = true
	t.SuppressEmptyColumns()
	t.SetOutputMirror(os.Stdout)
	return &SchedulePrinter{
		table: t,
	}
}
