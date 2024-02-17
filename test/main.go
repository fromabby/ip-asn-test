package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type TestCase struct {
	Got         int    `json:"as_number"`
	Want        int    `json:"expected"`
	Ip          string `json:"ip"`
	Isp         string `json:"isp"`
	CountryCode string `json:"country_code"`
}

func main() {
	var (
		readFileName = "../testcases.txt"
	)

	file, err := os.Open(readFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var failed int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		var testCase TestCase
		if err := json.Unmarshal([]byte(line), &testCase); err != nil {
			panic(err)
		}

		if testCase.Got != testCase.Want {
			fmt.Printf("FAILED ip: \"%s\" | isp: \"%s\" | country_code: \"%s\"\n\t-got(+) %d, want(-) %d\n\n", testCase.Ip, testCase.Isp, testCase.CountryCode, testCase.Got, testCase.Want)
			failed++
		}
		// else {
		// 	fmt.Printf(\"\t\tPASSED %s | %s: got %d, want %d\n", testCase.Ip, testCase.Isp, testCase.Got, testCase.Want)
		// }
	}

	fmt.Printf("Total failed: %d\n", failed)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
