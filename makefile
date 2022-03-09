build-gtask-terminal:
	@go build -o bin\gtask-terminal.exe .\cmdterminal\main.go

run-gtask-terminal: build-gtask-terminal
	@cd .\bin
	@gtask-terminal.exe 

build-gtask-web:
	@go build -o bin\gtask-web.exe .\cmdweb\main.go

run-gtask-web: build-gtask-web
	@.\bin\gtask-web.exe 	

build-gtask-desktop:
	@go build -o bin\gtask-desktop.exe .\cmdterminal\main.go

run-gtask-desktop: build-gtask-desktop
	@cd .\bin
	@gtask.exe 
