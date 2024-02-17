package main

import "github.com/fromabby/ip-asn-test/lib"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	asnList, err := lib.GenerateDataPool()
	check(err)

	err = lib.GenerateIPList(asnList)
	check(err)

	err = lib.GenerateTestScript()
	check(err)
}
