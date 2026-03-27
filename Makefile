APP_NAME = dhh-material-tool

.PHONY: build-mac build-windows build-all clean

build-mac:
	go build -ldflags="-s -w" -o $(APP_NAME) main.go
	@echo "Mac 构建完成: ./$(APP_NAME)"

build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(APP_NAME).exe main.go
	@echo "Windows 构建完成: ./$(APP_NAME).exe"

build-all: build-mac build-windows

clean:
	rm -f $(APP_NAME) $(APP_NAME).exe
