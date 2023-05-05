build: 
	hugo

clean:
	rm -fr ./dist/*

post:
	hugo new posts/${POST_NAME}.md
	sed -i '2s/.*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md
	sed -i '4s/.*/draft: false/' content/posts/$(POST_NAME).md
