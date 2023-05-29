build: ## lint check ## Generate the website from the markdown and configuration files in the directory dist/ and compile the source code of the application to a binary named awesome-api.
	hugo
	go build 

.PHONY: clean
clean: ## Cleanup the content of the directory dist/.Stop the application. Delete the binary awesome-api and the log file awesome-api.log
	rm -fr ./dist/*
	rm -f ./awesome-api ./awesome-api.log
	rm -f ./coverage-units.out ./coverage-integrations.out
	make stop

post: ## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME.
	hugo new posts/${POST_NAME}.md
	sed -i '2s/.*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md
	sed -i '4s/.*/draft: false/' content/posts/$(POST_NAME).md

run: ## Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log.
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application.
	-pgrep awesome-api && pkill awesome-api

test: ## Test it to ensure that it behaves as expected.
	make unit-tests
	make integration-tests

lint: ## Execute a static analysis to lint this code.
	golangci-lint run || exit

unit-tests: ## Make Unit tests
	go test -v -short -coverprofile=coverage-units.out

integration-tests: ## Make integration tests
	go test -v -coverprofile=coverage-integrations.out

check: ## Check the Markdown source files for linting and dead links
	markdownlint ./content/posts/* || exit
	markdown-link-check ./content/posts/* || exit

validate: ## Validate the generated HTML file
	@echo "Validating HTML..."
	./opt/W3C-Validator/w3c_validator.py ./dist/index.html || true

help: ## Show this help usage
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
