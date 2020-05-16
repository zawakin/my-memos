.PHONY: url-to-good-link remove-citation format

FORMAT_DIR = computer-architecture

url-to-good-link:
	find ${FORMAT_DIR} -name '*.md' | xargs sed -i -E "s ^https://en.wikipedia.org/wiki/(.+)$$ [\1(en.wikipedia)](\0) "

remove-citation:
	find ${FORMAT_DIR} -name '*.md' | xargs sed -i -E "s/\[[0-9]+\]//g"

format: url-to-good-link remove-citation

local-test:
	go test ./...
