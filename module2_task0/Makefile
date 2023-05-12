build: ## Generate the website from the markdown and configuration files in the directory dist/
## compile the source code of the application to a binary named awesome-api (the name awesome-api comes from the command go mod init github.com/<your github handle>/awesome-api) with the command go build. The first build may takes some times.
	hugo
	go build 

clean: ## Cleanup the content of the directory dist/
## Stop the application. Delete the binary awesome-api and the log file awesome-api.log
	rm -fr ./dist/*
	make stop
	rm ./awesome-api ./awesome-api/log

post: ## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME.
	hugo new posts/${POST_NAME}.md
	sed -i '2s/.*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md
	sed -i '4s/.*/draft: false/' content/posts/$(POST_NAME).md

run: ## Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log with the command ./awesome-api >./awesome-api.log 2>&1 &.
	./awesome-api & >> awesome-api.log

stop: ## Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application. For instance: kill "$(pgrep awesome-api)".
	pkill awesome-api

test: ## You want to test it to ensure that it behaves as expected. With the application started, you may want to use the command line curl (or your web browser, or the command wget or any other HTTP client):
	curl http://localhost:9999/health
	curl http://localhost:9999/health

help: ## Show this help usage
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
