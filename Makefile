.PHONY: build run clean

build:
	go build -o salesforce_mcp ./cmd

run: build
	SF_MCP_PORT=8989 ./salesforce_mcp

clean:
	rm -f salesforce_mcp
