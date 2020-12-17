test:
	cd pokemon && go run main.go

build:
	docker build -t leo101/gomon . && docker login && docker push leo101/gomon
