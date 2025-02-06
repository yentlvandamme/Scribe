build:
	go build -o ./dist/main $(ARGS)

run:
	go run main.go $(ARGS)

testAll:
	go test ./...

cleanMacDev:
	rm -r ~/Documents/Scribe