APP=chess-daily-puzzle
APP-DEBUG=${APP}-debug

run: build
	./${APP} -e prod

run-debug: build-debug
	./${APP-DEBUG} -e debug

build:
	go build -o ${APP} ./cmd

build-debug:
	go build -o ${APP-DEBUG} ./cmd

clean-binary:
	rm -f ${APP}
	rm -f ${APP-DEBUG}