## Steps to get started

### Setup project

```go
go mod init github.com/mnkrana/nft-contracts-backend
```

- add solidity bindings
- add rpc ports and adapters
- add nft ports and adapters
- add mint command
- use echo for api server
- add routes for mint command
- add transfer command
- add routes for transfer command

### To run locally

- add .env file
- add all configuration as per utils/literals
- run

```go
go run ./cmd
```

### To mint

```sh
curl -X POST -d '{"address":"0xB9160721D278482F153ae7eE9DFb037471228810","ids":[1],"amounts":[1],"uris":["ipfs://Qmb269DT2JWVq6AyibidEkDQ99CwMHsZTvwYy3AEBrfa11/1.json"]}' http://localhost:8080/mintbatch -H "Content-Type: application/json"
```

### To transfer

```sh
curl -X POST -d '{"address":"0x6f7c25e46E719d7AEcd76a2a53752353C5e19cE5","id":37,"amount":1}' http://localhost:8080/transfer -H "Content-Type: application/json"
```

### Mints

- Check all minted tokens here

  - https://puppyscan.shib.io/token/0xDB929853F31f9cfccF753A2Cec27c6A37c9D8bFa

- Faucet
  - https://shibarium.shib.io/faucet
