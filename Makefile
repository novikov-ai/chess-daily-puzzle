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

schedule:
	echo "11 11 * * 1,3,5  cd /usr/local/src/chess-daily-puzzle && ./chess-daily-puzzle >> chess-daily-puzzle.log 2>&1" | crontab -

schedule-disable:
	crontab -r

.PHONY: schedule
