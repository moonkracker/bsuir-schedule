package cmd

import (
	"fmt"
	"runtime"

	"github.com/dshipenok/gomorphos/russian/noun"
	"github.com/dshipenok/gomorphos/str"
	"github.com/spf13/cobra"

	"bsuir-schedule/controllers/bsuir"
	"bsuir-schedule/helpers"
	"bsuir-schedule/image"
	"bsuir-schedule/printers"
)

var (
	TeacherWeekNumbers    []int32
	TeacherWeekDays       []string
	TeacherAllWeekDays    bool
	TeacherAllWeekNumbers bool
	TeacherName           string
	TeacherSurname        string
	TeacherPatronymic     string

	teacherScheduleCmd = &cobra.Command{
		Use:   "teacher-schedule",
		Short: "Get teacher schedule",
		Run:   RunTeachersSchedule,
	}
)

func init() {
	rootCmd.AddCommand(teacherScheduleCmd)

	teacherScheduleCmd.PersistentFlags().Int32SliceVar(&TeacherWeekNumbers, "week-numbers", []int32{bsuir.GetCurrentWeekNumber()}, "Week numbers to get schedule")
	teacherScheduleCmd.PersistentFlags().StringSliceVar(&TeacherWeekDays, "week-days", []string{helpers.GetCurrentWeekDay()}, "Week days to get schedule")
	teacherScheduleCmd.PersistentFlags().BoolVar(&TeacherAllWeekDays, "all-week-days", false, "Print schedule for all week days")
	teacherScheduleCmd.PersistentFlags().BoolVar(&TeacherAllWeekNumbers, "all-week-numbers", false, "Print schedule for all BSUIR weeks")
	teacherScheduleCmd.PersistentFlags().StringVarP(&TeacherName, "name", "n", "", "Name of teacher")
	teacherScheduleCmd.PersistentFlags().StringVarP(&TeacherSurname, "surname", "s", "", "Surname of teacher")
	teacherScheduleCmd.PersistentFlags().StringVarP(&TeacherPatronymic, "patronymic", "p", "", "Patronymic of teacher")
	teacherScheduleCmd.MarkPersistentFlagRequired("surname")
}

func RunTeachersSchedule(cmd *cobra.Command, args []string) {
	urlIds := bsuir.GetTeacherUrlIds(TeacherSurname, TeacherName, TeacherPatronymic)
	if len(urlIds) == 0 {
		fmt.Println("Teacher not found")
		return
	}
	if TeacherAllWeekNumbers {
		TeacherWeekNumbers = []int32{1, 2, 3, 4}
	}
	if TeacherAllWeekDays {
		TeacherWeekDays = []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}
	}
	for _, urlId := range urlIds {
		data, days := bsuir.GetTeacherSchedule(urlId)
		fmt.Printf("Расписание преподавателя %s\n", data.EmployeeDto.LastName+" "+data.EmployeeDto.FirstName+" "+data.EmployeeDto.MiddleName)
		if runtime.GOOS == "darwin" {
			image.DisplayNetPicture(data.EmployeeDto.PhotoLink)
		}
		for _, weekNumber := range TeacherWeekNumbers {
			if !helpers.IsEmptyTeacherWeek(weekNumber, days) {
				fmt.Printf("Неделя %d\n", weekNumber)
				for _, day := range TeacherWeekDays {
					if !helpers.IsEmptyTeacherDay(weekNumber, days[day]) {
						sp := printers.NewTeacherPrinter()
						fmt.Printf("%s\n", day)
						sp.WriteTeacherDay(weekNumber, days[day])
						sp.RenderTable()
					} else {
						if !TeacherAllWeekDays {
							fmt.Println("Сегодня пар нет, можно и отдохнуть")
						} else {
							fmt.Printf("В %s пар нет, можно и отдохнуть\n", declension.GetCase(str.Word(day), "винительный", false))
						}
					}
				}
			} else {
				fmt.Printf("На %d неделе не найдено пар, свобода\n", weekNumber)
			}
		}
	}
}
