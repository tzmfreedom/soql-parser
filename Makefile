ANTLR := java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool

.PHONY: run
run: format
	go run example/main.go ./example/hoge.formula

.PHONY: build
build: format
	go build

.PHONY: format
format:
	goimports -w .
	gofmt -w .

.PHONY: generate
generate:
	cd ./parser && \
	$(ANTLR) -Dlanguage=Go -visitor SOQL.g4
