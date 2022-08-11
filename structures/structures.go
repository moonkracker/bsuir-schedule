package structures

type APIBsuirScheduleStruct struct {
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	StudentGroupDto struct {
		Name           string `json:"name"`
		FacultyName    string `json:"facultyName"`
		SpecialityName string `json:"specialityName"`
		Course         int    `json:"course"`
	} `json:"studentGroupDto"`
	EmployeeDto struct {
		ID         int    `json:"id"`
		FirstName  string `json:"firstName"`
		MiddleName string `json:"middleName"`
		LastName   string `json:"lastName"`
		PhotoLink  string `json:"photoLink"`
		Degree     string `json:"degree"`
		Rank       string `json:"rank"`
		URLID      string `json:"urlId"`
	} `json:"employeeDto"`
	TodayDate string                 `json:"todayDate"`
	Schedules map[string]interface{} `json:"schedules"`
}

type APIDay []struct {
	Auditories       []string    `json:"auditories"`
	EndLessonTime    string      `json:"endLessonTime"`
	LessonTypeAbbrev string      `json:"lessonTypeAbbrev"`
	Note             interface{} `json:"note"`
	NumSubgroup      int         `json:"numSubgroup"`
	StartLessonTime  string      `json:"startLessonTime"`
	Subject          string      `json:"subject"`
	SubjectFullName  string      `json:"subjectFullName"`
	WeekNumber       []int32     `json:"weekNumber"`
	Employees        `json:"employees"`
	StudentGroups    `json:"studentGroups"`
}

type Employees []struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	PhotoLink  string `json:"photoLink"`
}

type StudentGroups []struct {
	SpecialityName   string `json:"specialityName"`
	SpecialityCode   string `json:"specialityCode"`
	NumberOfStudents int    `json:"numberOfStudents"`
	Name             string `json:"name"`
	EducationDegree  int    `json:"educationDegree"`
}

type APITeachers []struct {
	FirstName          string   `json:"firstName"`
	LastName           string   `json:"lastName"`
	MiddleName         string   `json:"middleName"`
	Degree             string   `json:"degree"`
	Rank               string   `json:"rank"`
	PhotoLink          string   `json:"photoLink"`
	CalendarID         string   `json:"calendarId"`
	AcademicDepartment []string `json:"academicDepartment"`
	ID                 int      `json:"id"`
	URLID              string   `json:"urlId"`
	Fio                string   `json:"fio"`
}
