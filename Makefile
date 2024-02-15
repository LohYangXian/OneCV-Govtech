build:
	go build -o bin/OneCV-Govtech ./cmd/api

run: build
	./bin/OneCv-Govtech

test:
	#go test -v ./..