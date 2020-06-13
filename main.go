//go:generate go run frontend/generate.go

package main

import (
	"github.com/rs/zerolog/log"

	"github.com/clintjedwards/scheduler/cmd"
	"github.com/clintjedwards/scheduler/config"
)

func main() {
	conf, err := config.FromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("could not load env config")
	}

	setupLogging(conf.LogLevel, conf.Debug)

	cmd.RootCmd.Execute()
}

// func createSchedule(request proto.NewScheduleRequest) {

// 	startDate, _ := time.Parse("01-02-2006", request.StartDate)
// 	endDate, _ := time.Parse("01-02-2006", request.EndDate)

// 	// schedule key is day
// 	// value is another map whose key is position and value is employee id
// 	schedule := map[string]map[string]string{}

// 	currentDate := startDate

// 	for {
// 		fmt.Println(currentDate)
// 		fmt.Println(endDate)
// 		if currentDate == endDate {
// 			break
// 		}

// 		day := generateDay(currentDate)
// 		schedule[currentDate.Format("01-02-2006")] = day

// 		currentDate = currentDate.AddDate(0, 0, 1)
// 	}

// 	fmt.Println(schedule)

// }

// func generateDay(day time.Time) map[string]string {
// 	john := proto.Employee{
// 		Id:        "1",
// 		Name:      "john",
// 		Positions: map[string]bool{"baker": true},
// 	}
// 	don := proto.Employee{
// 		Id:        "2",
// 		Name:      "don",
// 		Positions: map[string]bool{"baker": true},
// 	}
// 	juan := proto.Employee{
// 		Id:        "3",
// 		Name:      "juan",
// 		Positions: map[string]bool{"cashier": true},
// 	}
// 	caroline := proto.Employee{
// 		Id:        "4",
// 		Name:      "caroline",
// 		Positions: map[string]bool{"baker": true},
// 	}
// 	destiny := proto.Employee{
// 		Id:        "5",
// 		Name:      "destiny",
// 		Positions: map[string]bool{"cashier": true},
// 	}

// 	employees := []proto.Employee{john, don, juan, caroline, destiny}
// 	positions := []string{"baker", "cashier"}
// 	finalMap := map[string]string{}

// 	for _, position := range positions {
// 		for _, employee := range employees {
// 			if _, ok := employee.Positions[position]; ok {
// 				finalMap[position] = employee.Id
// 				break
// 			}
// 			// employee does not work current position
// 			continue
// 		}
// 	}

// 	return finalMap
// }
