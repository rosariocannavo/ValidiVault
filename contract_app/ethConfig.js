const {Web3} = require('web3');

const ganacheURL = 'http://host.docker.internal:7545';  //localhost

const web3 = new Web3(new Web3.providers.HttpProvider(ganacheURL));

const contractAddress = '0x6dEbf52F3a0b8b90Ab1544C34C6CBfb9d1dd7286'; 

const contractABI = require('./contractABI');

module.exports = {
    web3,
    contractAddress,
    contractABI
};
