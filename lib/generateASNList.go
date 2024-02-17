package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateASNList() ([]string, error) {
	var (
		sep          = "\t"
		readFileName = "./files/sample_data.tsv"
	)

	file, err := os.Open(readFileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") || strings.Contains(line, ":") {
			continue
		}

		cols := strings.Split(line, sep)
		asn := cols[3]

		data = append(data, fmt.Sprintf("%s\n", asn))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
