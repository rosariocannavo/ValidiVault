
# ValidiVault: ethereum blockchain-based property validation

ValidiVault is a distributed ecosystem designed to ensure the integrity and
traceability of products through the use of blockchain. 

The system, built according to state-of-the-art design principles, boasts a microservice-oriented architecture, which ensures exceptional scalability and resilience. This is reinforced by the implementation of rate-limiting, circuit-breaking and other mechanisms, as well as an API gateway.

# Technologies Used
* **Docker**: Docker is a powerful platform that simplifies the process of building, deploying, and managing applications within containers. These containers encapsulate applications and their dependencies, ensuring they run consistently across different environments, from development to production.

* **Gin** : Go Gin is a lightweight web framework for the Go programming language. It's designed for building web applications and microservices with a focus on simplicity, performance, and minimalism. It leverages Go's concurrency features and provides a fast HTTP router, middleware support, and a robust set of functionalities to create efficient and scalable web applications.

* **Go Breaker**: Go Breaker is a library or pattern used in Go programming that implements the Circuit Breaker design pattern. This library is designed to enhance the stability and resilience of distributed systems by handling failures gracefully. The Circuit Breaker pattern works by wrapping a function call or service invocation. It monitors for failures, such as timeouts or errors, and when the failures surpass a certain threshold, it "trips" the circuit, preventing further calls to that function or service. This prevents the system from becoming overwhelmed and gives it time to recover or fallback to alternative behavior.

* **Go Reverse Proxy**: The NewSingleHostReverseProxy function in Go's net/http/httputil package is a convenient tool for creating a reverse proxy that directs incoming HTTP requests to a single backend server. This function simplifies the creation of a reverse proxy by setting up a reverse proxy instance configured to forward all incoming requests to a specific target host (backend server). It handles the necessary HTTP request and response plumbing, including modifying headers and paths as needed for the backend server.

* **Geth**: The Go Ethereum library, often referred to as Geth, is a Go language implementation of the Ethereum protocol. It's a comprehensive toolkit providing a wide range of functionalities for building applications on the Ethereum blockchain. 

* **NATS**: NATS, or the "NATS messaging system," is an open-source, lightweight, and high-performance messaging system designed for building distributed and scalable applications. It follows a publish-subscribe (pub/sub) messaging pattern, allowing different parts of an application or various services to communicate with each other efficiently.

* **WebSocket**: WebSockets are a communication protocol that provides full-duplex communication channels over a single, long-lived connection between a client and a server. Unlike traditional HTTP connections, WebSockets allow real-time, bidirectional communication, enabling both the client and server to send messages to each other at any time without the overhead of repeatedly establishing new connections. This technology is commonly used in applications that require instant data exchange, such as chat applications, online gaming, financial trading platforms, and live data streaming services.

* **Express.js**: Express.js, often referred to as Express, is a popular and minimalist web application framework for Node.js. It simplifies the process of building web applications and APIs by providing a robust set of features and middleware.

* **Web3.js**: Web3.js is a JavaScript library that serves as a bridge between web applications and the Ethereum blockchain. It provides tools and functionalities for developers to interact with Ethereum networks, manage accounts, and interact with smart contracts. With Web3.js, developers can send transactions, query blockchain data, deploy and interact with smart contracts, and handle Ethereum accounts and cryptographic keys. This library enables the creation of decentralized applications (dApps) by facilitating seamless communication and integration with the Ethereum blockchain.

* **Truffle**: Truffle is a comprehensive development framework specifically designed for Ethereum-based projects. It simplifies the process of creating, testing, and deploying smart contracts by offering tools for contract compilation, deployment, and testing. Additionally, Truffle provides a development pipeline for managing project assets and comes with a built-in console for interacting with contracts and the Ethereum network. This framework streamlines the development of decentralized applications (dApps) by providing a standardized and efficient environment for Ethereum development.

* **Mocha/Chai**: Mocha and Chai are a powerful duo in the JavaScript testing landscape. **Mocha** is a flexible and feature-rich testing framework for JavaScript that runs on both Node.js and in the browser. It provides a versatile and easy-to-use testing environment, allowing developers to write test suites and test cases using familiar syntax like describe() and it(). Mocha supports asynchronous testing, various reporting options, and hooks for setup and teardown. **Chai**, on the other hand, is an assertion library that can be paired with Mocha (or other testing frameworks). It provides a wide range of assertion styles, such as expect, should, and assert, allowing developers to write clear and expressive tests. Chai's flexibility enables different assertion styles based on individual preferences and readability.

* **Ganache**:Ganache serves as a local Ethereum blockchain emulator, offering developers a controlled environment for testing smart contracts and decentralized applications (dApps). By creating a personal blockchain on their machines, developers can simulate Ethereum network conditions, test various interactions, and debug code without relying on the actual Ethereum network. It provides a quick and user-friendly way to iterate on Ethereum-related projects, facilitating efficient development and testing processes. 

* **Infura**: Infura is a popular infrastructure-as-a-service provider specifically tailored for Ethereum and other blockchain networks. It offers developers easy access to Ethereum nodes without the need to run and maintain their own. Developers use Infura's API endpoints to interact with the Ethereum blockchain, send transactions, query data, and access various functionalities without managing the complexities of running a full Ethereum node. This allows for rapid development of decentralized applications (dApps) and other blockchain-related projects without the overhead of node maintenance or syncing. Infura provides reliable and scalable infrastructure, serving as a gateway for developers to connect their applications to the Ethereum network, making it an essential tool in the Ethereum ecosystem for seamless and efficient blockchain development.

# draft system architecture
![image](https://drive.google.com/uc?export=view&id=1KPH8hfV-L9PRpMpF_GGgatZ0i_w6_DQ3)
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

Alternatively, you can use a Sepolia infura account and set the contract address in the config file: 
```bash
contract address on Sepolia: 0x6e525aa41918B0eAc5D7278512e9e43428Cd414A
```

Go to deploy directory

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

#### Return login page

```http
GET /
```

#### Return signup page

```http
GET /signup 
```

#### Permit new user registration

```http
POST /registration 
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required**. username to register |
| `password` | `string` | **Required**. user password |
| `metamaskAddress` | `string` | **Required**. current metamask address fetched from browser plugin |

#### Returns session user cookies

```http
GET /cookie 
```

#### Login the user

```http
POST /login 
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`| `string` | **Required**. username to register |
| `password`| `string` | **Required**. user password |
| `metamaskAddress` | `string` | **Required**. current metamask address fetched from browser plugin |

#### Verifiy the Metamask user by signing a random nonce

```http
POST /verify-signature 
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nonce`| `string` | **Required**. received nonce  |
| `address`| `string` | **Required**. signer metamask address |
| `signature` | `string` | **Required**. signed nonce  |

#### Return the admin home page 

```http
GET /admin/admin_home
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token `| `string` | **Required**. jwt token   |

#### Return the user home page 

```http
GET /user/user_home
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token `| `string` | **Required**. jwt token   |


#### Call the real server by proxy to check if a product is registered in the blockchain If authenticated as users

```http
GET /user/app/product
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`| `string` | **Required**. jwt token   |
| `account`| `string` | **Required**. address of the user who wants to make the transaction |
| `productID `| `string` | **Required**. id of the product you want to retrieve from the blockchain   |


#### Call the real server by proxy to check if a product is registered in the blockchain If authenticated as admin

```http
GET /admin/app/product
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`| `string` | **Required**. jwt token   |
| `account`| `string` | **Required**. address of the user who wants to make the transaction |
| `productID `| `string` | **Required**. id of the product you want to retrieve from the blockchain   |

#### Call the real server by proxy to register a product in  the blockchain

```http
PUT /admin/app/product
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`| `string` | **Required**. jwt token   |
| `account`| `string` | **Required**. address of the user who wants to make the transaction |
| `productName `| `string` | **Required**. name of the product you want to register in the blockchain   |


## Running Truffle Tests

The smart contract testing environment is fully automated and managed by Truffle.

Go to Truffle project

```bash
cd Truffle 
```
Run the following command:

```bash
truffle test 
```

## Running Api Gateway Mock Test

you can run a test suite that checks the operation of the rate limiter, circuit breaker, and also tests the registration and login routes

Go to Api Gateway project test directory:
```bash
cd api_gateway/internal/test 
```
Launches test suite:

```bash
go test . 
```
## Useful links 

| Container     | URL                                             | Description                           |
| ------------- | ----------------------------------------------- | ------------------------------------- |
| MongoExpress  | http://localhost:8083                           | Mongo Express dashboard               |
| RedisCommander| http://localhost:8081                           | Redis Commander dashboard             |
| Api Logger    | https://localhost:8000                          | Api Gateway websocket logger          |
| NATS          | https://natsdashboard.com/                      | NATS dashboard                        |



## Made by

- [@rosariocannavo](https://github.com/rosariocannavo)

