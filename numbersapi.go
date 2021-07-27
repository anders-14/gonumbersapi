package numbersapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const BASE_URL = "http://numbersapi.com/"

var reDate = regexp.MustCompile(`\d\d\/\d\d`)

// fetch the given url and parse and return the body
func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("err (http.Get): %+v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("error fetching, returned status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("err (io.ReadAll): %+v", err)
	}

	return string(body)
}

// MathFact return a mathematical fact about the number passed to it
func MathFact(number int) string {
	url := fmt.Sprintf("%s%d/math", BASE_URL, number)

	fact := fetch(url)

	return fact
}

// FunFact returns some random trivia about the number passed to it
func FunFact(number int) string {
	url := fmt.Sprintf("%s%d/trivia", BASE_URL, number)

	fact := fetch(url)

	return fact
}

// DateFact returns a fact about the date passed on the format dd/mm
// No freedom formatting allowed
func DateFact(date string) string {
	if !reDate.MatchString(date) {
		return "wrong"
	}

	dateNums := strings.Split(date, "/")
	newDate := fmt.Sprintf("%s/%s", dateNums[1], dateNums[0])

	url := fmt.Sprintf("%s%s/%s", BASE_URL, newDate, "date")

	fact := fetch(url)

	return fact
}

// YearFact returns a fact about the year passed to it
func YearFact(year int) string {
	url := fmt.Sprintf("%s%d/year", BASE_URL, year)

	fact := fetch(url)

	return fact
}
