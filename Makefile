build:
	./set_version.sh
	go mod tidy
	go build -o ./bin/pgcabletester.$(uname_p) ./cmd/pgquartz

# Use the following on m1:
# alias make='/usr/bin/arch -arch arm64 /usr/bin/make'
debug:
	go build -gcflags "all=-N -l" -o ./bin/pgcabletester.debug.$(uname_p) ./cmd/pgquartz
	~/go/bin/dlv --headless --listen=:2345 --api-version=2 --accept-multiclient exec ./bin/pgcabletester.debug.$(uname_p) -- -c '$(JOB)'

debug_test:
	~/go/bin/dlv --headless --listen=:2345 --api-version=2 --accept-multiclient test ./pkg/git/

run:
	./bin/pgcabletester.$(uname_p) -c testconfig.yml

fmt:
	gofmt -w .
	goimports -w .
	gci write .

compose:
	./docker-compose-tests.sh

test: sec lint

sec:
	gosec ./...

lint:
	golangci-lint run
