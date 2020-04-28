url-to-good-link:
	find . -name '*.md' | xargs sed -i -E "s ^https://en.wikipedia.org/wiki/(.+)$$ [\1(en.wikipedia)](\0) "

local-test:
	go test ./...