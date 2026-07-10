CC = go

SRC = ./src/*.go
OUT = ./build/gitcloud

build:
	$(CC) build -o $(OUT) $(SRC)

run:
	$(CC) run $(SRC)

remove:
	rm -f $(OUT)

.PHONY: build run remove

