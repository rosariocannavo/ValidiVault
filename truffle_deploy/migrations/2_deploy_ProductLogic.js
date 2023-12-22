const ProductLogic = artifacts.require("./ProductLogic.sol");

module.exports = function(deployer) {
  deployer.deploy(ProductLogic);
};

//truffle migrate --network development

//truffle migrate --f 2 --to 2
