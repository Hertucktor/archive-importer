package importer

import (
	"net/http"
	"testing"
)

func TestRequestAllCardsFromAPI(t *testing.T) {
	testCases := []struct {
		page          int
		expectedError bool
	}{
		{1, false},
		{-1, true},
	}

	for _, tc := range testCases {
		resp, err := RequestAllCardsFromAPI(tc.page)

		if err != nil {
			if !tc.expectedError {
				t.Errorf("unexpected error: %v", err)
			}
			continue
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d; got %d", http.StatusOK, resp.StatusCode)
		}
	}
}
