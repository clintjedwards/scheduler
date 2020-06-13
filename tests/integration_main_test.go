package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/clintjedwards/scheduler/app"
	"github.com/clintjedwards/scheduler/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type testHarness struct {
	client proto.SchedulerAPIClient
}

func (info *testHarness) setup() {
	os.Setenv("SCHEDULER_LOGLEVEL", "error")

	go app.StartServices()
	time.Sleep(time.Second)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "8081"), grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to server")
	}

	client := proto.NewSchedulerAPIClient(conn)
	info.client = client
}

func (info *testHarness) cleanup() {
	os.Unsetenv("SCHEDULER_LOGLEVEL")
}

func TestFullApplication(t *testing.T) {

	info := testHarness{}
	info.setup()

	info.TestSetSchedulerSettings(t)
	info.TestGetSchedulerSettings(t)

	info.TestAddEmployee(t)
	info.TestGetEmployee(t)

	info.cleanup()
}
