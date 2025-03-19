build:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./cmd
	zip myFunction.zip bootstrap  

deploy:	
	aws lambda update-function-code --function-name nft-marketplace --zip-file fileb://myFunction.zip

logs:
	aws logs tail /aws/lambda/nft-marketplace --follow --format short