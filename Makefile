.PHONY: dev
dev:
	# use air to watch file changes and auto reload app
	air -- -dP "1qaz!QAZ" -dn taylorzh -di "127.0.0.1" -dp 3306

.PHONY: run
run:
	go build -o ./tmp/main .
	./tmp/main -dP "1qaz!QAZ" -dn taylorzh -di "127.0.0.1" -dp 3306