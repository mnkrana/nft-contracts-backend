build:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./cmd
	zip myFunction.zip bootstrap  