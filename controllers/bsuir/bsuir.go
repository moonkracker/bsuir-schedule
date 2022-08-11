package bsuir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"

	"bsuir-schedule/structures"
)

const (
	bsuirGroupScheduleLink   = "https://iis.bsuir.by/api/v1/schedule?studentGroup=%s"
	bsuirCurrentWeekLink     = "https://iis.bsuir.by/api/v1/schedule/current-week"
	bsuirAllTeachersLink     = "https://iis.bsuir.by/api/v1/employees/all"
	bsuirTeacherScheduleLink = "https://iis.bsuir.by/api/v1/employees/schedule/%s"
)

func queryAPI(method string, apiLink string, specs string) ([]byte, error) {
	req, err := http.NewRequest(method, apiLink+specs, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func getDataFromRepositoryJSON[T any](dataHolder T, addr string) (*T, error) {
	dataFromRepo := dataHolder
	response, err := queryAPI("GET", addr, "")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &dataFromRepo)
	if err != nil {
		return nil, err
	}
	return &dataFromRepo, nil
}

func GetCurrentWeekNumber() int32 {
	data, _ := queryAPI("GET", bsuirCurrentWeekLink, "")
	weekNum, _ := strconv.ParseInt(string(data), 10, 32)
	return int32(weekNum)
}

func GetGroupSchedule(groupNumber string) (structures.APIBsuirScheduleStruct, map[string]structures.APIDay) {
	data, _ := getDataFromRepositoryJSON(structures.APIBsuirScheduleStruct{}, fmt.Sprintf(bsuirGroupScheduleLink, groupNumber))

	days := map[string]structures.APIDay{}
	for k, v := range data.Schedules {
		day := structures.APIDay{}
		mapstructure.Decode(v, &day)
		days[k] = day
	}
	return *data, days
}

func GetTeacherUrlIds(lastName string, firstName string, middleName string) []string {
	urlIds := []string{}
	data, _ := getDataFromRepositoryJSON(structures.APITeachers{}, bsuirAllTeachersLink)
	for _, v := range *data {
		switch {
		case v.FirstName == firstName && v.LastName == lastName && v.MiddleName == middleName:
			urlIds = append(urlIds, v.URLID)
		case v.FirstName == firstName && v.LastName == lastName && middleName == "":
			urlIds = append(urlIds, v.URLID)
		case v.LastName == lastName && firstName == "" && middleName == "":
			urlIds = append(urlIds, v.URLID)
		case v.FirstName == firstName && lastName == "" && middleName == "":
			urlIds = append(urlIds, v.URLID)
		case v.MiddleName == middleName && firstName == "" && lastName == "":
			urlIds = append(urlIds, v.URLID)
		default:
			continue
		}
	}
	return urlIds
}

func GetTeacherSchedule(urlId string) (structures.APIBsuirScheduleStruct, map[string]structures.APIDay) {
	data, _ := getDataFromRepositoryJSON(structures.APIBsuirScheduleStruct{}, fmt.Sprintf(bsuirTeacherScheduleLink, urlId))

	days := map[string]structures.APIDay{}
	for k, v := range data.Schedules {
		day := structures.APIDay{}
		mapstructure.Decode(v, &day)
		days[k] = day
	}
	return *data, days
}
