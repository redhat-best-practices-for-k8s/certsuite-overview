package pkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const DateTemplate = "01/02/2006" // Example format for dates

func TestGetDateRangeLastYear(t *testing.T) {
	// Define test cases
	tests := []struct {
		name            string
		expectedToday   string
		expectedLastYear string
	}{
		{
			name:            "Regular Case - One Year Range",
			expectedToday:   time.Now().Format(DateFormat),
			expectedLastYear: time.Now().AddDate(-1, 0, 0).Format(DateFormat),
		},
	}

	// Iterate through test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			actualLastYear, actualToday := getDateRangeLastYear()

			// Validate the results with assertions
			assert.Equal(t, tc.expectedToday, actualToday, "Mismatch for today in test: %s", tc.name)
			assert.Equal(t, tc.expectedLastYear, actualLastYear, "Mismatch for last year in test: %s", tc.name)
		})
	}
}
