# run all tests
[group('dev')]
test:
    @echo "running tests..."
    @go test -p 1 -v -count=1 ./internal/... -json | tparse -all -format plain
    @echo "tests passed."

# format, vet, and lint
[group('dev')]
lint:
    @echo "formatting..."
    @go fmt ./...
    @echo "vetting..."
    @go vet ./...
    @echo "linting..."
    @golangci-lint run ./...
    @echo "all checks passed."

# build and install tapi binary to $GOPATH/bin
[group('setup')]
install:
    @echo "installing dependencies..."
    @brew install go@1.26 jq
    @echo "intalling go packages"
    @go install github.com/mfridman/tparse@latest
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
