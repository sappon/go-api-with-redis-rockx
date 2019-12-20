package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"testing"
	// "fmt"
)

const (
	BadRequestCode     = 400
	SuccessRequestCode = 200
)

type TestStruct struct {
	requestBody        string
	responseBody       string
	expectedStatusCode int
	observedStatusCode int
}

func TestIndex(t *testing.T) {
	// The test passes if the response status code is success (200)

	url := "http://localhost:8080/"

	tests := []TestStruct{
		{`{}`, `{}`, SuccessRequestCode, 0},
	}


	var reader io.Reader
	reader = strings.NewReader(tests[0].requestBody)

	request, err := http.NewRequest("GET", url, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	tests[0].responseBody = strings.TrimSpace(string(body))
	tests[0].observedStatusCode = res.StatusCode

	DisplayTestCaseResults("index", tests, t)
}

func TestUpdateData(t *testing.T) {
	// The test passes if the response created time (HH:mm) is current

	url := "http://localhost:8080/updateData"

	//get response from /updateData
	tests := []TestStruct{
		{`{"pairs":"ETH,ATOM","currency":"BTC"}`, ``, SuccessRequestCode, 0},
	}
	
	var reader io.Reader
	reader = strings.NewReader(tests[0].requestBody)

	request, err := http.NewRequest("POST", url, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	s := strings.TrimSpace(string(body))
	tests[0].responseBody = s[:35] + "..."

	// Test by compare latest update time with current
	if strings.Contains(s, time.Now().Local().Format("2006-01-02T15:04:05")) {
		tests[0].observedStatusCode = res.StatusCode
	}

	DisplayTestCaseResults("updateData", tests, t)
}

func TestLatestUpdate(t *testing.T) {
	// The test passes if the response status code is success (200)

	url := "http://localhost:8080/getLatest"

	tests := []TestStruct{
		{``, ``, SuccessRequestCode, 0},
	}

	var reader io.Reader
	reader = strings.NewReader(tests[0].requestBody)

	request, err := http.NewRequest("GET", url, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	s := strings.TrimSpace(string(body))
	tests[0].responseBody = s[:35] + "..."
	tests[0].observedStatusCode = res.StatusCode

	DisplayTestCaseResults("getLatest", tests, t)
}

func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {

	for _, test := range tests {

		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}