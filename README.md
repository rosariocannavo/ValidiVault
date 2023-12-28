
# ValidiVault: ethereum blockchain-based property validation

ValidiVault è un ecosistema distribuito progettato per garantire l’integrità e
la tracciabilità dei prodotti attraverso l’impiego della blockchain.

# Tecnologie Utilizzate
* Docker
* Go Gin 
* Go Breaker
* Go Proxy
* Geth 
* NATS 
* Express.js
* Web3.js
* Truffle
* Mocha 
* Chai
* Ganache 



## Run Locally
Install and Run Ganache client and create a local ethereum Blockchain
```bash
npm install -g ganache-cli

ganache-cli
```

Create a new Ganache test network
```bash
ganache-cli --networkId 5777 --port 7545 --gasLimit 8000000 --mnemonic "mnemonic phrase"

```

Clone the project

```bash
git clone https://github.com/rosariocannavo/ValidiVault.git
```

Go to the project directory

```bash
cd ValidiVault
```

Go to the Truffle directory

```bash
cd truffle_deploy
```

Deploy ProductProxy smart contract on Ganache local instance

```bash
truffle migrate --f 2 --to 2
```

Deploy ProductProxy smart contract on Ganache local instance

```bash
truffle migrate --f 3 --to 3
```

Sets ganache to the local port so that it is reachable from within containers

```bash
ganache-cli --port 7545 --host 192.168.1.15
```

Go to deploy direcytory

```bash
cd deploy 
```

Start the project containers (api gateway, server, logger, database, redis and NATS)

```bash
docker-compose up
```

Go to login page
```bash 
localhost:8080/
```


## API Reference

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |
| `api_key` | `string` | **Required**. Your API key |


#### add(num1, num2)

Takes two numbers and returns the sum.

