all:
	go build -o ./bin/web-native main.go certifacate.go
windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/web-windows-amd64 main.go certifacate.go
mac:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/web-mac-amd64 main.go certifacate.go
linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/web-linux-amd64 main.go certifacate.go
clean:
	rm -fr ./bin/