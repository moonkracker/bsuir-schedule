package printers

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	TypeField       = table.ColumnConfig{Name: "Тип", Hidden: false}
	TimeField       = table.ColumnConfig{Name: "Время", Hidden: false}
	SubjectField    = table.ColumnConfig{Name: "Предмет", Hidden: false}
	AuditorieField  = table.ColumnConfig{Name: "Аудитория", Hidden: false}
	TeacherField    = table.ColumnConfig{Name: "Преподаватель", Hidden: false}
	SubgroupField   = table.ColumnConfig{Name: "Подгруппа", Hidden: false}
	GroupsField     = table.ColumnConfig{Name: "Группы", Hidden: false, AutoMerge: true}
	SpecialityField = table.ColumnConfig{Name: "Специальность", Hidden: false, AutoMerge: true}
)

type SchedulePrinter struct {
	table table.Writer
}

func (sp *SchedulePrinter) GetTable() table.Writer {
	return sp.table
}

func (sp *SchedulePrinter) RenderTable() {
	sp.table.Render()
}
