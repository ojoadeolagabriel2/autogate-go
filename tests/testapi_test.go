package tests

import (
	"autogate/api"
	"autogate/app"
	"autogate/data"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGivenRequestForAllUsersConfirmUsersReturned(t *testing.T) {
	application := app.App{}
	application.InitializeRouter()
	api.HandleRequests(application.Router)

	response, _ := http.Get("http://localhost:12345/users")
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	bodyStr := string(body)
	fmt.Println(bodyStr)

	var payload []data.User
	json.Unmarshal(body, &payload)

	if len(payload) != 3 {
		t.Errorf("got %d expected %d", len(payload), 3)
	}
}

func TestGivenRequestForUserByIdConfirmUserReturned(t *testing.T) {
	application := app.App{}
	application.InitializeRouter()
	api.HandleRequests(application.Router)

	response, _ := http.Get("http://localhost:12345/users/1")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	bodyStr := string(body)
	fmt.Println(bodyStr)
	var payload data.User
	json.Unmarshal(body, &payload)

	assert.NotNil(t, payload)
	assert.Equal(t, payload.Id, int64(1))
}
