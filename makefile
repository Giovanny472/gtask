build-gtask:
	@go build -o bin\gtask.exe .\cmd\main.go

run-gtask: build-gtask 
	@.\bin\gtask.exe 