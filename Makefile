BIN=spraak

run: $(BIN)
	bin/$(BIN)

$(BIN): main.go
	go build -o bin

