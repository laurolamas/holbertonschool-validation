# Prerequisites

To build this website, you will need to have the following installed:

    Go-Hugo
    GNU Make

# Lifecycle

There are three main steps involved in building and maintaining this website:

    Build: This step generates the website from the markdown and configuration files in the directory dist/. To run this step, use the command `make build`.

    Clean: This step cleans up the content of the directory dist/. To run this step, use the command `make clean`.

    Post: This step creates a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME. To run this step, use the command `make POST_NAME=your-post-name POST_TITLE="Your Post Title" post`.

    Help: Show this help usage 
