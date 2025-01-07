APP=chess-daily-puzzle
ENV=prod

install: build schedule
uninstall: schedule-off clean-binary

run: build
	./${APP} -e prod

run-debug: build
	./${APP} -e debug

build:
	go build -o ${APP} ./cmd

clean-binary:
	rm -f ${APP}

schedule:
	echo "11 11 * * 1,3,5 cd /usr/local/src/${APP} && ./${APP} -e ${ENV} >> ${APP}-${ENV}.log 2>&1" | crontab -

schedule-off:
	crontab -r

format: 
	goimports -w .
	gofumpt -w .

lint:
	golangci-lint run

.PHONY: schedule