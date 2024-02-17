build:
	go build -o bin/OneCV-Govtech

run: build
	./bin/OneCV-Govtech

test:
	go test -v ./tests

docker-build:
	docker build -t my-go-app .

docker-run:
	docker run -p 8080:8080 my-go-app
