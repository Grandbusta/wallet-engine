# Wallet engine API - Go🔥

A wallent engine built with Golang.
NB: This is a very basic version

## Technologies Used

- Golang
- Gorilla Mux
- GORM
- MySQL

## Features

- Create a wallet

```
route:/wallet/create
body: { "email" : "johndoe@gmail.com" }
```

- Credit a wallet
```
route: /wallet/credit
body: { "amount" : "500", wallet_address : 647077292}
```

- Debit a wallet
```
route: /wallet/debit
body: { "amount" : "500", wallet_address : 647077292}
```
- Activate a wallet
```
route: /wallet/{wallet_address}
body: nil
```

- Deactivate a wallet
```
route: /wallet/{wallet_address}
body: nil
```

## 🤓 Author(s)

Olaifa Boluwatife