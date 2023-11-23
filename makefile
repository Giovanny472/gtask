build-gtask-cli:
	@go build -o bin\gtask-cli.exe .\cmd\cli\main.go

run-gtask-cli: build-gtask-cli
	@cd .\bin
	@gtask-cli.exe 

