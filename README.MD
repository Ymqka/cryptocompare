start
```
docker-compose up -d
go build . && ./cryptocompare
```

test
```
curl http://localhost:8080/price?fsyms=BTC&tsyms=USD
```

