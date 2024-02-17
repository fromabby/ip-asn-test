TEST_RESULT_FILENAME = test_result.txt

download-db:
	./download.sh

run:
	go run main.go

get-result:
	chmod 755 ./test.sh && ./test.sh > testcases.txt

test:
	go run test/main.go

test-save:
	cd test && go run main.go > $(TEST_RESULT_FILENAME)

run-test: run get-result test-save
