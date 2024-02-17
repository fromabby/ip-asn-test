package lib

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/netip"
	"os"
	"slices"
	"strings"

	"go4.org/netipx"
)

func GenerateDataPool() ([]string, error) {
	var (
		sep           = "\t"
		readFileName  = "./ip2asn-combined.tsv"
		writeFileName = "./files/sample_data.tsv"
	)

	file, err := os.Open(readFileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	f, err := os.Create(writeFileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// create a slice of random line numbers out of 634062 lines
	size := 10
	fileSize := 634062
	nums := getRandomLineIndexes(size, fileSize)

	var data []byte
	var asns []string
	var i = 0
	for scanner.Scan() {
		if !slices.Contains(nums, i) {
			i++
			continue
		}

		line := strings.TrimSpace(scanner.Text())

		// ignore ipv6 for now
		isIpv6 := strings.Contains(line, ":")
		if line == "" || strings.HasPrefix(line, "#") || isIpv6 {
			continue
		}

		cols := strings.Split(line, sep)
		range_start := cols[0]
		range_end := cols[1]

		cidr, err := convertIPRangeToCIDR(range_start, range_end)
		if err != nil {
			iprangeErr := fmt.Sprintf("failed to get CIDR: %v - %v", range_start, range_end)
			fmt.Println(iprangeErr)
			continue
		}

		for _, c := range cidr {
			data = []byte(fmt.Sprintf("%s\t%s\n", c.String(), line))
			asns = append(asns, cols[2])
		}

		_, err = f.Write(data)
		if err != nil {
			return nil, err
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return asns, nil
}

func convertIPRangeToCIDR(range_start, range_end string) ([]netip.Prefix, error) {
	var b netipx.IPSetBuilder
	b.AddRange(netipx.IPRangeFrom(
		netip.MustParseAddr(range_start),
		netip.MustParseAddr(range_end),
	))

	s, err := b.IPSet()
	if err != nil {
		return nil, err
	}

	return s.Prefixes(), nil
}

func getRandomLineIndexes(n, max int) []int {
	set := make(map[int]bool)
	var result []int
	for len(set) < n {
		value := rand.Intn(max)
		if !set[value] {
			set[value] = true
			result = append(result, value)
		}
	}
	return result
}
