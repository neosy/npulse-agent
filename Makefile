# Makefile
# fastHTTP server "npulse-agent"

include .make.env
export

VERSION_FILE := "VERSION"
VERSION_START := "0.1.0"

VERSION := $(shell cat $(VERSION_FILE))
VERSION_NEW := $(shell echo $(VERSION) | awk -F. '{print $$1"."$$2"."$$3+1}')

.DEFAULT_GOAL := help

help: ## Список команд
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: make <commands> \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: ## Билд исполняемого файла
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o $(APP_NAME) ./cmd/main.go

run: ## Запуск
	@go run ./cmd/main.go

git-push-tag-version: ## Создание тега в git для актуальной версии
	-git tag v$(VERSION)
	git push --tags

version-create: ## Создание файла с номер версии программы
	echo -n $(VERSION_START) > $(VERSION_FILE)
	
version-inc: ## Увеличение номера версии программы и сохранение в файл
	echo -n $(VERSION_NEW) > $(VERSION_FILE)

install: ## Установка приложения
	@./install.sh