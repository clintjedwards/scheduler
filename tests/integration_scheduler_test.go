package tests

import (
	"context"
	"testing"

	"github.com/clintjedwards/scheduler/proto"
	"github.com/stretchr/testify/require"
)

func (info *testHarness) TestSetSchedulerSettings(t *testing.T) {
	t.Run("SetSchedulerSettings", func(t *testing.T) {

		setSettingsRequest := &proto.SetSchedulerSettingsRequest{
			Positions: []*proto.Position{
				{
					PrimaryName:   "baker 1",
					SecondaryName: "shaper",
				},
				{
					PrimaryName:   "baker 2",
					SecondaryName: "mixer",
				},
				{
					PrimaryName:   "baker 3",
					SecondaryName: "cookies",
				},
				{
					PrimaryName:   "line",
					SecondaryName: "greeter",
				},
				{
					PrimaryName:   "delivery app duty",
					SecondaryName: "retail",
				},
				{
					PrimaryName:   "retail",
					SecondaryName: "",
				},
				{
					PrimaryName:   "barista",
					SecondaryName: "retail",
				},
				{
					PrimaryName:   "porter",
					SecondaryName: "",
				},
			},
		}

		response, err := info.client.SetSchedulerSettings(context.Background(), setSettingsRequest)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotEmpty(t, response)
		require.NotEmpty(t, response.Settings.Positions)
	})
}

func (info *testHarness) TestGetSchedulerSettings(t *testing.T) {
	t.Run("GetSchedulerSettings", func(t *testing.T) {

		response, err := info.client.GetSchedulerSettings(context.Background(), &proto.GetSchedulerSettingsRequest{})
		require.NoError(t, err)
		require.NotNil(t, response)
		require.NotEmpty(t, response)
		require.NotEmpty(t, response.Settings.Positions)
	})
}
