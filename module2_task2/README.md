# Prerequisites

To build this website, you will need to have the following installed:

    Go-Hugo
    GNU Make

# Lifecycle

There are three main steps involved in building and maintaining this website:

Help: Show this help usage 

build: this target should compile the source code and generate the binary in the BUILD_DIR directory. You can use the "go build" command for this.

run: this target should run the binary in the background and write the logs to the LOG_FILE. You can use the "nohup" command to run the binary in the background and redirect the logs to the LOG_FILE.

stop: this target should stop the running server. You can use the "kill" command to send a SIGINT signal to the process ID stored in the PID_FILE.

clean: this target should stop the running server (if it is running), delete the binary in the BUILD_DIR directory, and delete the LOG_FILE.

test: this target should run some tests to ensure that the server is working as expected. You can use the "curl" command to send HTTP requests to the server and check the responses.

lint: lint the code

unit-tests: Make unit tests
