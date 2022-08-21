package printers

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	NumberField     = table.ColumnConfig{Name: "№", Hidden: false, AutoMerge: true}
	TypeField       = table.ColumnConfig{Name: "Тип", Hidden: false, AutoMerge: true}
	TimeField       = table.ColumnConfig{Name: "Время", Hidden: false, AutoMerge: true}
	SubjectField    = table.ColumnConfig{Name: "Предмет", Hidden: false, AutoMerge: true}
	AuditorieField  = table.ColumnConfig{Name: "Аудитория", Hidden: false, AutoMerge: true}
	TeacherField    = table.ColumnConfig{Name: "Преподаватель", Hidden: false, AutoMerge: true}
	SubgroupField   = table.ColumnConfig{Name: "Подгруппа", Hidden: false, AutoMerge: true}
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
