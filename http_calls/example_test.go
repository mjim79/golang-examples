package http_calls

import (
	"errors"
	"fmt"
	"github.com/mjim79/go-httpclient/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gohttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndpointsErrorGettingFromApi(t *testing.T) {

	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method: http.MethodGet,
		Url:    "https://api.github.com",
		Error:  errors.New("timeout getting the endpoints"),
	})

	endpoints, err := GetEndPoints()

	fmt.Println(err)
	fmt.Println(endpoints)

	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "timeout getting the endpoints", err.Error())
}

func TestGetEndpointsEndpointsNotFound(t *testing.T) {

	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusNotFound,
		ResponseBody:       `{"message"; "endpoint not found"}`,
	})

	endpoints, err := GetEndPoints()

	fmt.Println(err)
	fmt.Println(endpoints)

	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error when trying to fetch github endpoints", err.Error())
}

func TestGetEndpointsInvalidJsonResponse(t *testing.T) {

	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"events_url": }`,
	})

	endpoints, err := GetEndPoints()

	fmt.Println(err)
	fmt.Println(endpoints)

	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid character '}' looking for beginning of value", err.Error())
}

func TestGetEndpointsNoError(t *testing.T) {

	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"events_url": "https://api.github.com/events"}`,
	})

	endpoints, err := GetEndPoints()

	fmt.Println(err)
	fmt.Println(endpoints)

	assert.Nil(t, err)
	assert.NotNil(t, endpoints)
	assert.EqualValues(t, "https://api.github.com/events", endpoints.EventsUrl)
}
