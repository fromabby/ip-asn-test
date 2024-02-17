package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateTestScript() error {
	var (
		readFileName  = "./files/ip_test_list.tsv"
		writeFileName = "./test.sh"
	)

	file, err := os.Open(readFileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	f, err := os.Create(writeFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	var data []byte
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		cols := strings.Split(line, "\t")
		ip := cols[0]
		asn := cols[1]
		isp := cols[2]
		country_code := cols[3]
		data = []byte(fmt.Sprintf("curl \"http://localhost:8081/ip-info?ip=%s\" | jq -c '.ip_info | { as_number: .as_number, ip: .ip, expected: %s, isp: \"%s\", country_code: \"%s\" }' &\n", ip, asn, isp, country_code))

		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
