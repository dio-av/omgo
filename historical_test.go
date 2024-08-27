package omgo_test

import (
	"context"
	"testing"

	"github.com/dio-av/omgo"
	"github.com/stretchr/testify/require"
)

func TestHistorical(t *testing.T) {
	c, err := omgo.NewClient()
	require.NoError(t, err)

	loc, err := omgo.NewLocation(52.3738, 4.8910) // Amsterdam
	require.NoError(t, err)

	hopts := omgo.HistoricalOptions{
		TemperatureUnit:   "fahrenheit",
		WindspeedUnit:     "mph",
		PrecipitationUnit: "inch",
		Timezone:          "US/Eastern",
		StartDate:         "2023-05-01",
		EndDate:           "2023-06-01",
		HourlyMetrics:     []string{"cloudcover", "relativehumidity_2m"},
		DailyMetrics:      []string{"temperature_2m_max"},
	}

	res, err := c.Historical(context.Background(), loc, &hopts)
	require.NoError(t, err)

	require.Greater(t, len(res.HourlyTimes), 0)
	require.Equal(t, 2, len(res.HourlyMetrics))
	require.Greater(t, len(res.DailyTimes), 0)
	require.Equal(t, 1, len(res.DailyMetrics))
}
