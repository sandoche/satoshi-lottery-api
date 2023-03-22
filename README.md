# Satoshi Lottery

> An API for a lottery to try to guess bitcoin private keys

I made this project for learning how to build API with Golang

## ⚙️ Development

```sh
# install dependencies
go install github.com/codegangsta/gin@latest
go mod vendor

# run dev server
gin run server.go
```

## 👽 API endpoints

```
GET /v1/bitcoin - get a random bitcoin private key

GET /v1/bitcoin/:private-key - get the balance in satoshis for a private key and the public address
```

### 🐋 Deploy using Docker

```sh
docker build . -t satoshi-lottery
docker run -p 8080:8080 satoshi-lottery
```

## ⭐️ Show your support

Please ⭐️ this repository if this project helped you!

<a href="https://www.patreon.com/sandoche">[![patreon.png](https://c5.patreon.com/external/logo/become_a_patron_button.png)](https://www.patreon.com/sandoche)</a>

## Buy me a beer

If you like this project, feel free to donate:

- Bitcoin: 19JiNZ1LkMaz57tewqJaTg2hQWH4RgW4Yp
- Ethereum: 0xded81fa4624e05339924355fe3504ba9587d5419
- Monero: 43jqzMquW2q989UKSrB2YbeffhmJhbYb2Yxu289bv7pLRh4xVgMKj5yTd52iL6x1dvCYs9ERg5biHYxMjGkpSTs6S2jMyJn
- Paypal: https://www.paypal.me/kanbanote
