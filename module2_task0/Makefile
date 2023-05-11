## APP_NAME: the name of the binary that will be generated
APP_NAME=awesome-api

## BUILD_DIR: the directory where the binary will be built
BUILD_DIR=

## LOG_FILE: the name of the log file that will be used for logging
LOG_FILE=

## PORT: the port on which the server will listen
PORT=9999

## PID_FILE: the name of the file that will store the process ID of the running server
PID_FILE=


build: ## compile the source code of the application to a binary named awesome-api (the name awesome-api comes from the command go mod init github.com/<your github handle>/awesome-api) with the command go build. The first build may takes some times.
	go build -o $(BUILD_DIR)/$(APP_NAME)

run: ## Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log with the command ./awesome-api >./awesome-api.log 2>&1 &.
	go run

stop: ## Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application. For instance: kill "$(pgrep awesome-api)".
	kill `cat $(PID_FILE)`

clean: ## Stop the application. Delete the binary awesome-api and the log file awesome-api.log

test: ## You want to test it to ensure that it behaves as expected. With the application started, you may want to use the command line curl (or your web browser, or the command wget or any other HTTP client):

help: ## Show this help usage
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
