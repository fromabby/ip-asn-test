# Check which asn lookup values are failing

Download the ip2asn masterfile

```
./download.sh
```

Generate test sample and create a shell script that will query to geoip-proxy-registry

```
go run main.go
```

Run the test script

```
./test.sh > result.txt
```

Check the result

```
cd test
go run main.go > ./test_result.txt
```
