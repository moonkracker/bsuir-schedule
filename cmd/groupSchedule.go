package cmd

import (
	"fmt"

	"github.com/dshipenok/gomorphos/russian/noun"
	"github.com/dshipenok/gomorphos/str"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"bsuir-schedule/controllers/bsuir"
	"bsuir-schedule/helpers"
	"bsuir-schedule/printers"
)

var (
	GroupWeekNumbers     []int32
	GroupWeekDays        []string
	GroupSubgroupsNumber []int32
	GroupAllWeekDays     bool
	GroupAllWeekNumbers  bool

	groupScheduleCmd = &cobra.Command{
		Use:   "group-schedule",
		Short: "Get group schedule",
		Run:   RunGroupSchedule,
	}
)

func init() {
	rootCmd.AddCommand(groupScheduleCmd)
	groupScheduleCmd.PersistentFlags().Int32SliceVar(&GroupWeekNumbers, "week-numbers", []int32{bsuir.GetCurrentWeekNumber()}, "Week numbers to get schedule")
	groupScheduleCmd.PersistentFlags().Int32SliceVarP(&GroupSubgroupsNumber, "subgroups", "s", []int32{0, 1, 2}, "Number of subgroup")
	groupScheduleCmd.PersistentFlags().StringSliceVar(&GroupWeekDays, "week-days", []string{helpers.GetCurrentWeekDay()}, "Week days to get schedule")
	groupScheduleCmd.PersistentFlags().BoolVar(&GroupAllWeekDays, "all-week-days", false, "Print schedule for all week days")
	groupScheduleCmd.PersistentFlags().BoolVar(&GroupAllWeekNumbers, "all-week-numbers", false, "Print schedule for all BSUIR weeks")
}

func getNumberFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("group number parameter is missing")
	}
	return args[0], nil
}

func RunGroupSchedule(cmd *cobra.Command, args []string) {
	groupNumber, err := getNumberFromArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, days := bsuir.GetGroupSchedule(groupNumber)

	if GroupAllWeekNumbers {
		GroupWeekNumbers = []int32{1, 2, 3, 4}
	}
	if GroupAllWeekDays {
		GroupWeekDays = []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}
	}

	GroupSubgroupsNumber = append(GroupSubgroupsNumber, 0)

	fmt.Printf("Расписание группы %s\n", groupNumber)
	fmt.Println(cases.Title(language.Und).String(data.StudentGroupDto.FacultyName))
	fmt.Println(data.StudentGroupDto.SpecialityName)

	for _, weekNumber := range GroupWeekNumbers {
		fmt.Printf("Неделя %d\n", weekNumber)
		for _, day := range GroupWeekDays {
			if !helpers.IsEmptyGroupDay(weekNumber, GroupSubgroupsNumber, days[day]) {
				sp := printers.NewGroupPrinter()
				fmt.Printf("%s\n", day)
				sp.WriteGroupDay(weekNumber, GroupSubgroupsNumber, days[day])
				sp.RenderTable()
			} else {
				if !GroupAllWeekDays {
					fmt.Println("Сегодня пар нет, отдыхай")
				} else {
					fmt.Printf("В %s пар нет, можно и отдохнуть\n", declension.GetCase(str.Word(day), "винительный", false))
				}
			}
		}
	}
}
