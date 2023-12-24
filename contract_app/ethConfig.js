const {Web3} = require('web3');

// within the container "localhost" refers to the container itself, as ganache is local you have to enter the address of the machine
const ganacheURL = "http://192.168.1.15:7545"; 

const web3 = new Web3(new Web3.providers.HttpProvider(ganacheURL));

const contractAddress = '0x6dEbf52F3a0b8b90Ab1544C34C6CBfb9d1dd7286'; 

const contractABI = require('./contractABI');

module.exports = {
    web3,
    contractAddress,
    contractABI
};
