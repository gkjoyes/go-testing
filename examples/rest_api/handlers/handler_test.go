package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g-kutty/go-testing/tests/coverage/handlers"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func init() {
	handlers.Routes()
}

// TestSendJSON testing the GET /user internal endpoint.
func TestSendJSON(t *testing.T) {
	url := "/user"
	statusCode := http.StatusOK

	t.Log("Given the need to test the SendJson endpoint.")
	{

		// request
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d]", failed, statusCode, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", succeed, statusCode)
		}

		// user struct
		var u struct {
			Name  string
			Email string
		}

		// decode response body
		if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
			t.Fatalf("\t%s\tShould be able to decode the response.", succeed)
		}
		t.Logf("\t%s\tShould be able to decode the response.", succeed)

		// check response fields
		if u.Name == "george" {
			t.Logf("\t%s\tShould have \"george\" for the Name in the response.", succeed)
		} else {
			t.Errorf("\t%s\tShould have \"george\" for Name in the response: %q", failed, u.Name)
		}

		if u.Email == "george@gmail.com" {
			t.Logf("\t%s\tShould have \"george@gmail.com\" for the Email in the response.", succeed)
		} else {
			t.Errorf("\t%s\tShould have \"george@gmail.com\" for Email in the response: %q", failed, u.Email)
		}
	}
}
