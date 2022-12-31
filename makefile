i: install
install:
	go mod tidy

# Build the Go program
build: main.go
    go build -o main main.go

# Run the Go program
run: main
    ./main

# Clean up the build artifacts
clean:
    rm -f main
