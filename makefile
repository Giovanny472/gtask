build-gtask-terminal:
	@go build -o bin\gtask-terminal.exe .\cmd\cli\main.go

run-gtask-terminal: build-gtask-terminal
	@cd .\bin
	@gtask-terminal.exe 

build-gtask-web:
	@go build -o bin\gtask-web.exe .\cmd\web\main.go

run-gtask-web: build-gtask-web
	@.\bin\gtask-web.exe 	

build-gtask-desktop:
	@go build -o bin\gtask-desktop.exe .\cmd\desktop\main.go

run-gtask-desktop: build-gtask-desktop
	@cd .\bin
	@gtask.exe 
