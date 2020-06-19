all: test

test: check
	bash ./scripts/test.sh

check:format
	bash ./scripts/check.sh

format:
	bash ./scripts/format.sh

