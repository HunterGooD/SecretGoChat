build: build_front \
	build_goland

build_wasm:
	GOARCH=wasm GOOS=js go build -o web/public/assets/js/lib/utils.wasm internal/wasm/wasm.go

build_front:
	cd web/ && yarn build


build_goland:
	go build -o dist/secret_chat/main cmd/secret_chat/main.go

install:
	go get ./...
	cd web
	npm i
	yarn build
	cd ../cmd/secret_chat
	go build main.go -o ../../main

clean:
	rm -rf web/dist
	rm -rf dist

testing:
	go test ./...

help:
	@echo "                                                             "
	@echo "                  Существующие команды                       "
	@echo "-------------------------------------------------------------"
	@echo "install       : Установка всех модулей и сборка прокта	    "
	@echo "build         : Сборка проекта                          	    "
	@echo "clean         : Очистить проект                              "
	@echo "build_wasm    : Сборка WASM модулей               	        "
	@echo "testing       : тестирование компонентов             	    "
	@echo "-------------------------------------------------------------"
	@echo "                                                             "
