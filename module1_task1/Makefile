build: 
	hugo

clean:
	rm -fr ./dist/*

post:
	hugo new posts/${POST_NAME}.md
	sed -i "1s/title: .*$/title: $(POST_TITLE)/" ./content/posts/$(POST_NAME).
