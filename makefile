build-gtask:
	@go build -o cmd\gtask.exe .\cmd\main.go

run-gtask: build-gtask
	@ .\cmd\gtask.exe 