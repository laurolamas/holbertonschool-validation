build: ## Generate the website from the markdown and configuration files in the directory dist/
	hugo

clean: ## Cleanup the content of the directory dist/
	rm -fr ./dist/*

post: ## Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME.
	hugo new posts/${POST_NAME}.md
	sed -i '2s/.*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md
	sed -i '4s/.*/draft: false/' content/posts/$(POST_NAME).md

help: ## Show this help usage
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
