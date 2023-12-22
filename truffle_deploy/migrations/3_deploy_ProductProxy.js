const ProductProxy = artifacts.require("./ProductProxy.sol");

module.exports = function(deployer) {
  const logicContractAddress = "0x6dEbf52F3a0b8b90Ab1544C34C6CBfb9d1dd7286";
  deployer.deploy(ProductProxy, logicContractAddress);
};

//truffle migrate --network development

//truffle migrate --f 3 --to 3
