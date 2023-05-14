build: ## Generate the website from the markdown and configuration files in the directory dist/ and compile the source code of the application to a binary named awesome-api.
	hugo
	make lint
	go build 

.PHONY: clean
clean: ## Cleanup the content of the directory dist/.Stop the application. Delete the binary awesome-api and the log file awesome-api.log
	rm -fr ./dist/*
	rm ./awesome-api ./awesome-api.log
	pkill awesome-api

post: ## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME.
	hugo new posts/${POST_NAME}.md
	sed -i '2s/.*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md
	sed -i '4s/.*/draft: false/' content/posts/$(POST_NAME).md

run: ## Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log.
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application.
	pkill awesome-api

test: ## Test it to ensure that it behaves as expected.
	curl http://localhost:9999
	curl http://localhost:9999/health

lint: ## Execute a static analysis to lint this code.
	golangci-lint run || exit

help: ## Show this help usage
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
