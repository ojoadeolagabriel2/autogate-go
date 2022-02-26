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

const (
	BaseServerUrl = "http://localhost:12345"
)

func TestGivenRequestForAllUsersConfirmUsersReturned(t *testing.T) {
	application := app.App{}
	application.InitializeRouter()
	api.RegisterRoutes(application.Router)

	response, _ := http.Get(fmt.Sprintf("%s/users", BaseServerUrl))

	//defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var payload []data.User
	_ = json.Unmarshal(body, &payload)

	if len(payload) != 3 {
		t.Errorf("got %d expected %d", len(payload), 3)
	}
}

func TestGivenRequestForUserByIdConfirmUserReturnedX(t *testing.T) {
	// given request handler
	application := app.App{}
	application.InitializeRouter()
	api.RegisterRoutes(application.Router)

	// when getting user by id
	response, _ := http.Get(fmt.Sprintf("%s/users/1", BaseServerUrl))

	//defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var payload data.User
	_ = json.Unmarshal(body, &payload)

	// assert payload returned
	assert.NotNil(t, payload)
	assert.Equal(t, payload.Id, int64(1))
}