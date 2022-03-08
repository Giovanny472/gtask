build-gtask:
	@go build -o bin\gtask.exe .\cmd\main.go

run-gtask: build-gtask 
	@cd .\bin
	@gtask.exe 