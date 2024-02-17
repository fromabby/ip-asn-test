package main

import "github.com/fromabby/ip-asn-test/lib"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// get random sample data from ip2asn-combined.tsv
	err := lib.GenerateSampleDataList()
	check(err)

	{
		// create list of asn in sample data list
		asnList, err := lib.GenerateASNList()
		check(err)

		// create random ip for each line sample data ip range
		err = lib.GenerateSampleTestIPList(asnList)
		check(err)
	}

	// create test script
	// run './test.sh > result.txt' to pipe the result to file
	err = lib.GenerateTestScript()
	check(err)
}
