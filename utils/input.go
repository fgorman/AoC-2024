package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const Year = 2024

func GetDaysInput(day int) string {
	sessionCookie := os.Getenv("SESSION_COOKIE")

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", Year, day)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic("Unable to create request for day's input")
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("Unable to send request for day's input")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Unable to read bytes from request body")
	}

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("Issue with request: %d\n%s", res.StatusCode, string(resBody)))
	}

	return string(resBody)
}