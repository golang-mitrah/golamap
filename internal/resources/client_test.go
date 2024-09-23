package resources

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccessTokens(t *testing.T) {
	mockInvalidJsonConversion := make(chan int)
	tests := []struct {
		name           string
		clientID       string
		clientSecret   string
		mockStatusCode int
		mockResponse   string
		expectedToken  string
		expectedError  error
	}{
		{
			name:           "Successful Token Retrieval",
			clientID:       "validClientID",
			clientSecret:   "validClientSecret",
			mockStatusCode: http.StatusOK,
			mockResponse:   `{"access_token": "validAccessToken"}`,
			expectedToken:  "validAccessToken",
			expectedError:  nil,
		},
		{
			name:           "HTTP Status Not OK",
			clientID:       "validClientID",
			clientSecret:   "validClientSecret",
			mockStatusCode: http.StatusUnauthorized,
			mockResponse:   `{"error": "invalid_client"}`,
			expectedToken:  "",
			expectedError:  errors.New("Failed to get token: 401 Unauthorized"),
		},
		{
			name:           "JSON Decode Error",
			clientID:       "validClientID",
			clientSecret:   "validClientSecret",
			mockStatusCode: http.StatusOK,
			mockResponse:   fmt.Sprint(mockInvalidJsonConversion),
			expectedToken:  "",
			expectedError:  errors.New("json: cannot unmarshal number into Go value of type resources.TokenResponse"),
		},
	}

	for _, testData := range tests {
		// Run each test case in parallel
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			// Create a mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Fatalf("expected POST request, got %s", r.Method)
				}
				if r.URL.Path != "/token" {
					t.Fatalf("expected request to /token, got %s", r.URL.Path)
				}
				w.WriteHeader(testData.mockStatusCode)
				fmt.Fprintln(w, testData.mockResponse)
			}))
			defer server.Close()

			OLAMap := &OLAMap{}
			// Call the function with the test case parameters
			err := OLAMap.ConfigureAccessToken(testData.clientID, testData.clientSecret)

			// Check results
			if OLAMap.Token != testData.expectedToken {
				t.Errorf("expected token %s, got %s", testData.expectedToken, OLAMap.Token)
			}
			if err != nil && err.Error() != testData.expectedError.Error() {
				t.Errorf("expected error %v, got %v", testData.expectedError, err)
			} else if err == nil && testData.expectedError != nil {
				t.Errorf("expected error %v, got nil", testData.expectedError)
			}
		})
	}
}
