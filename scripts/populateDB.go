package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/clintjedwards/scheduler/model"
	"github.com/icrowley/fake"
)

type harness struct {
	positionsList []string
	employeesList []string
}

func (h *harness) setup() {
}

func createEmployee(employee model.AddEmployee) string {

	requestBody, err := json.Marshal(employee)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", "http://localhost:8080/api/employees", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := model.Employee{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.ID
}

func (h *harness) createEmployees(num int) {

	positions := []string{}
	for _, id := range h.positionsList {
		positions = append(positions, id)
	}

	for i := 0; i < num; i++ {
		newEmployee := model.AddEmployee{
			Name:      fake.FullName(),
			Notes:     fake.WordsN(30),
			StartDate: strconv.Itoa(fake.Year(2000, 2020)) + "-" + strconv.Itoa(fake.MonthNum()) + "-" + strconv.Itoa(fake.Day()),
			Positions: positions,
		}

		id := createEmployee(newEmployee)

		h.employeesList = append(h.employeesList, id)
		log.Printf("created employee %s", newEmployee.Name)
	}
}

func createPosition(position model.Position) string {

	requestBody, err := json.Marshal(position)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", "http://localhost:8080/api/positions", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := model.Position{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.ID
}

func (h *harness) createPositions() {
	positions := []model.Position{
		{
			PrimaryName:   "Baking",
			SecondaryName: "Shaping",
			Description:   "Shaping is harder than mixing.",
		},
		{
			PrimaryName:   "Baking",
			SecondaryName: "Mixing",
			Description:   "Nobody with any self respect wants to work in mixing.",
		},
		{
			PrimaryName:   "Retail",
			SecondaryName: "Line/Greeter",
			Description:   "Usually our most cheerful and useless employee.",
		},
		{
			PrimaryName:   "Retail",
			SecondaryName: "Deliveries",
			Description:   "Pleasing tastebuds one ubereats pickup at a time.",
		},
		{
			PrimaryName:   "Retail",
			SecondaryName: "General",
			Description:   "For the non-specialist types",
		},
		{
			PrimaryName:   "Retail",
			SecondaryName: "Barista",
			Description:   "Who unironically comes to a bakery for coffee? Just go to starbucks like a real person",
		},
		{
			PrimaryName:   "Retail",
			SecondaryName: "Cookie Baking",
			Description:   "Demoted from being an actual baker",
		},
		{
			PrimaryName: "Porter",
			Description: "Demoted from being an actual employee",
		},
	}

	for _, position := range positions {
		newPosition := model.Position{
			PrimaryName:   position.PrimaryName,
			SecondaryName: position.SecondaryName,
			Description:   position.Description,
		}

		id := createPosition(newPosition)
		h.positionsList = append(h.positionsList, id)
		log.Printf("created position %s", newPosition.PrimaryName)
	}
}

func createSchedule(schedule model.Schedule) string {

	requestBody, err := json.Marshal(schedule)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", "http://localhost:8080/api/schedules", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := model.Position{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.ID
}

func (h *harness) generateSchedule() {

	position1 := h.positionsList[0]
	position2 := h.positionsList[1]
	position3 := h.positionsList[2]
	position4 := h.positionsList[3]
	position5 := h.positionsList[4]
	position6 := h.positionsList[5]
	position7 := h.positionsList[6]
	position8 := h.positionsList[7]

	schedule := model.Schedule{
		Start:          "1990-06-19",
		End:            "1990-06-27",
		EmployeeFilter: []string{},
		Program: model.Program{
			Monday: []model.Shift{
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position1,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position2,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position3,
				},
				{
					Start:      "1300",
					End:        "1600",
					PositionID: position4,
				},
				{
					Start:      "1300",
					End:        "1600",
					PositionID: position5,
				},
				{
					Start:      "1600",
					End:        "2200",
					PositionID: position6,
				},
				{
					Start:      "1600",
					End:        "2200",
					PositionID: position7,
				},
				{
					Start:      "1600",
					End:        "2200",
					PositionID: position8,
				},
				{
					Start:      "1600",
					End:        "2200",
					PositionID: position1,
				},
			},
			Tuesday: []model.Shift{
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position1,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position2,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position3,
				},
			},
			Wednesday: []model.Shift{
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position1,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position2,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position3,
				},
			},
			Thursday: []model.Shift{
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position1,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position2,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position3,
				},
			},
			Friday: []model.Shift{
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position1,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position2,
				},
				{
					Start:      "0800",
					End:        "1300",
					PositionID: position3,
				},
			},
		},
	}

	createSchedule(schedule)
	log.Println("created schedule")
}

func main() {

	h := harness{}
	h.setup()

	h.createPositions()
	h.createEmployees(10)
	h.generateSchedule()
}
