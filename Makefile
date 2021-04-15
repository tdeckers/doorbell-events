build:
	go build

run:
	./doorbell-function

test:
	go test -v

deploy:
	gcloud functions deploy doorbell-function --region europe-west1 --entry-point EventHandler --runtime go113 --trigger-http

.PHONY: build test run deploy