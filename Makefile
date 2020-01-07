#!/usr/bin/env make -f
GCLOUD=gcloud

appid=portfolio
version=beta

vet:
	go vet ./...

packages:
	go get ./...

build: vet
	goapp build -v ./...

hot-reload: upgrade-node
	npm install --prefix client && \
	npm start --prefix client

hot-server:
	go run main.go

development-server:
	npm install --prefix client && \
	npm start --prefix client && \
	go run main.go

staging-server:
	npm install --prefix client && \
	npm run-script build --prefix client --mode staging && \
	cd ../ && \
	go run main.go

deploy-staging:
	npm install --prefix client && \
	npm run-script build --prefix client --mode staging && \
	cd ../ && \
	gcloud app deploy && \
	gcloud app logs tail -s portfolio

deploy-production:
	npm install --prefix client && \
	npm run-script build --prefix client --mode production && \
	cd ../ && \
	gcloud app deploy && \
	gcloud app logs tail -s portfolio

livereload: build
	go run main.go  &&  npm start --prefix client

# Build web application
build: build-server build-client

build-server: vet
	go build -v

build-client:
	npm install --prefix client && \
	npm run-script build --prefix client --mode production

dev:
	cd ./client/build/ && npm run-script build  && cd ../../ && go run main.go

# Install stable version of nodejs (10.16.x)
upgrade-node:
	sudo npm cache clean -f && \
	sudo npm install -g n && \
	sudo n stable