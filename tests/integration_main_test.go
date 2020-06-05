package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/clintjedwards/scheduler/app"
	"github.com/clintjedwards/scheduler/proto"
	"github.com/clintjedwards/toolkit/random"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type testHarness struct {
	client       proto.SchedulerAPIClient
	databasePath string
}

func (info *testHarness) setup() {
	databasePath := fmt.Sprintf("/tmp/scheduler%s.db", random.GenerateRandString(4))
	os.Setenv("SCHEDULER_DATABASE_PATH", databasePath)
	os.Setenv("SCHEDULER_LOGLEVEL", "error")

	go app.StartServices()
	time.Sleep(time.Second)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "8081"), grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to server")
	}

	client := proto.NewSchedulerAPIClient(conn)
	info.client = client
	info.databasePath = databasePath
}

func (info *testHarness) cleanup() {
	os.Unsetenv("TLS_CERT_PATH")
	os.Unsetenv("TLS_KEY_PATH")
	os.Unsetenv("SCHEDULER_DATABASE_PATH")
	os.Unsetenv("SCHEDULER_LOGLEVEL")
	os.Remove(info.databasePath)
}

func TestFullApplication(t *testing.T) {
	info := testHarness{}
	info.setup()

	info.TestSetSchedulerSettings(t)
	info.TestGetSchedulerSettings(t)

	info.cleanup()
}
