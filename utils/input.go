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

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	res, _ := http.DefaultClient.Do(req)

	resBody, _ := ioutil.ReadAll(res.Body)

	return string(resBody)
}