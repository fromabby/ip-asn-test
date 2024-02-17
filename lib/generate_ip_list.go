package lib

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"

	"github.com/praserx/ipconv"
)

func GenerateIPList(asnList []string) error {
	var (
		sep           = "\t"
		readFileName  = "./files/sample_data.tsv"
		writeFileName = "./files/ip_test_list.tsv"
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

	var i = 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		cols := strings.Split(line, sep)
		range_start := cols[1]
		range_end := cols[2]
		country_code := cols[4]
		isp := cols[5]

		ip, err := GetRandomIP(range_start, range_end)
		if err != nil {
			return err
		}

		str := fmt.Sprintf("%s\t%s\t%s\t%s\n", ip.String(), strings.TrimSpace(asnList[i]), isp, country_code)

		writeDate := []byte(str)
		_, err = f.Write(writeDate)
		if err != nil {
			return err
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func GetRandomIP(start string, end string) (net.IP, error) {
	if strings.Contains(start, ":") {
		return randomIPv6(start, end)
	}

	return randomIPv4(start, end)
}

func randomIPv4(start string, end string) (net.IP, error) {
	ipStart, err := ipconv.IPv4ToInt(net.ParseIP(start))
	if err != nil {
		return nil, err
	}

	ipEnd, err := ipconv.IPv4ToInt(net.ParseIP(end))
	if err != nil {
		return nil, err
	}

	randomIP := rand.Intn(int(ipEnd)-int(ipStart)) + int(ipStart)

	return ipconv.IntToIPv4(uint32(randomIP)), nil
}

func randomIPv6(start string, end string) (net.IP, error) {
	ipStart, err := ipconv.IPv6ToInt(net.ParseIP(start))
	if err != nil {
		return nil, err
	}

	ipEnd, err := ipconv.IPv6ToInt(net.ParseIP(end))
	if err != nil {
		return nil, err
	}

	ipStartHigh, ipStartLow := ipStart[0], ipStart[1]
	ipEndHigh, ipEndLow := ipEnd[0], ipEnd[1]

	randomIPStart := rand.Int63n(int64(ipEndHigh)-int64(ipStartHigh)) + int64(ipStartHigh)
	randomIPLow := rand.Int63n(int64(ipEndLow)-int64(ipStartLow)) + int64(ipStartLow)

	return ipconv.IntToIPv6(uint64(randomIPStart), uint64(randomIPLow)), nil
}
