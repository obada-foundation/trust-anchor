bin/compile:
	cd src && go build -o trust-anchor

bin/run:
	cd src && go run main.go

docker/build:

docker/publish:
