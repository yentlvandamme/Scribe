build:
	go build -o ./dist/main

run:
	go run main.go $(ARGS)

cleanMacDev:
	rm -r ~/Documents/Scribe