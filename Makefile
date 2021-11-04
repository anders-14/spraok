BIN=spraok

run: $(BIN)
	@printf "[CMD] "
	bin/$(BIN)

$(BIN): main.go
	@mkdir -p bin
	@printf "[CMD] "
	go build -o bin

all: fmt lint

fmt:
	@printf "[CMD] "
	goimports -w -l .

lint:
	@printf "[CMD] "
	golint ./...

test:
	@printf "[CMD] "
	go test ./...

.PHONY: run fmt lint all
