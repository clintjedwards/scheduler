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

func main() {

	h := harness{}
	h.setup()

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run populateDB.go <numEmployees> <numPositions>")
		os.Exit(1)
	}
	numEmployees, _ := strconv.Atoi(os.Args[1])
	h.createEmployees(numEmployees)

	h.cleanup()
}
