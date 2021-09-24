BIN=spraak

run: $(BIN)
	bin/$(BIN)

$(BIN): main.go
	go build -o bin

fmt:
	goimports -w -l .

.PHONY: run fmt
