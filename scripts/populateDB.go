package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/icrowley/fake"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type harness struct {
	positionsList []string
	employeesList []string
	conn          *grpc.ClientConn
}

func (h *harness) setup() {

	const certPath string = "./localhost.crt"

	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		log.Fatalf("failed to get certificates: %v", err)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "8080"), opts...)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	h.conn = conn
}

func (h *harness) cleanup() {
	defer h.conn.Close()
}

func (h *harness) createEmployees(num int) {
	for i := 0; i < num; i++ {
		newEmployee := &proto.AddEmployeeRequest{
			Name:  fake.FullName(),
			Notes: fake.WordsN(30),
		}

		client := proto.NewSchedulerAPIClient(h.conn)

		response, err := client.AddEmployee(context.Background(), newEmployee)
		if err != nil {
			log.Fatalf("could not create employee: %v", err)
		}

		h.employeesList = append(h.employeesList, response.Employee.Id)
		log.Printf("created employee %s", newEmployee.Name)
	}
}

func (h *harness) createPositions() {
	positions := []proto.Position{
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
		newPosition := &proto.AddPositionRequest{
			PrimaryName:   position.PrimaryName,
			SecondaryName: position.SecondaryName,
			Description:   position.Description,
		}

		client := proto.NewSchedulerAPIClient(h.conn)
		response, err := client.AddPosition(context.Background(), newPosition)
		if err != nil {
			log.Fatalf("could not create position: %v", err)
		}
		h.positionsList = append(h.positionsList, response.Position.Id)
		log.Printf("created position %s", newPosition.PrimaryName)
	}
}

func main() {

	h := harness{}
	h.setup()

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run populateDB.go <numEmployees>")
		os.Exit(1)
	}
	numEmployees, _ := strconv.Atoi(os.Args[1])
	h.createEmployees(numEmployees)
	h.createPositions()

	h.cleanup()
}
