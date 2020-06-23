package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/clintjedwards/scheduler/models"
	"github.com/icrowley/fake"
)

type harness struct {
	positionsList []string
	employeesList []string
}

func (h *harness) setup() {
}

func createEmployee(employee models.Employee) string {

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

	response := models.Employee{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.ID
}

func (h *harness) createEmployees(num int) {

	positions := map[string]bool{}
	for _, id := range h.positionsList {
		positions[id] = true
	}

	for i := 0; i < num; i++ {
		newEmployee := models.Employee{
			Name:      fake.FullName(),
			Notes:     fake.WordsN(30),
			Status:    models.EmployeeActive,
			Positions: positions,
		}

		id := createEmployee(newEmployee)

		h.employeesList = append(h.employeesList, id)
		log.Printf("created employee %s", newEmployee.Name)
	}
}

func createPosition(position models.Position) string {

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

	response := models.Position{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.ID
}

func (h *harness) createPositions() {
	positions := []models.Position{
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
		newPosition := models.Position{
			PrimaryName:   position.PrimaryName,
			SecondaryName: position.SecondaryName,
			Description:   position.Description,
		}

		id := createPosition(newPosition)
		h.positionsList = append(h.positionsList, id)
		log.Printf("created position %s", newPosition.PrimaryName)
	}
}

func createSchedule(schedule models.Schedule) string {

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

	response := models.Position{}

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

	schedule := models.Schedule{
		Start:          "06-19-1990",
		End:            "06-27-1990",
		EmployeeFilter: []string{},
		Program: models.Program{
			Monday: map[string][]models.Shift{
				position1: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position2: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position3: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
			},
			Tuesday: map[string][]models.Shift{
				position1: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position2: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position3: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
			},
			Wednesday: map[string][]models.Shift{
				position1: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position2: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position3: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
			},
			Thursday: map[string][]models.Shift{
				position1: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position2: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position3: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
			},
			Friday: map[string][]models.Shift{
				position1: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position2: {
					{
						Start: "0800",
						End:   "1300",
					},
				},
				position3: {
					{
						Start: "0800",
						End:   "1300",
					},
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

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run populateDB.go <numEmployees>")
		os.Exit(1)
	}
	numEmployees, _ := strconv.Atoi(os.Args[1])
	h.createPositions()
	h.createEmployees(numEmployees)
	h.generateSchedule()
}
