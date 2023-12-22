// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//Emergency Stop / Circuit Breaker dp

contract ProductLogic {
    address public owner;
    bool public stopped;

    struct Product {
        uint productId;
        string productName;
        address manufacturer;
        bool isRegistered;
    }

    mapping(uint => Product) public products;
    uint public productCount;

    event ProductRegistered(uint productId, string productName, address indexed manufacturer);
    event EmergencyStopSet(bool stopped);

    modifier onlyOwner {
        require(msg.sender == owner, "Only contract owner can perform this action");
        _;
    }

    modifier whenNotStopped {
        require(!stopped, "Contract is stopped");
        _;
    }

    constructor() {
        owner = msg.sender;
        stopped = false;
    }

    function toggleContractStopped() external onlyOwner {
        stopped = !stopped;
        emit EmergencyStopSet(stopped);
    }

    function registerProduct(uint _productId, string memory _productName) external whenNotStopped {
        require(!products[_productId].isRegistered, "Product already registered");
        
        products[_productId] = Product(_productId, _productName, msg.sender, true);
        productCount++;

        emit ProductRegistered(_productId, _productName, msg.sender);
    }

    function getProduct(uint _productId) external view returns (
        uint productId,
        string memory productName,
        address manufacturer,
        bool isRegistered
    ) {
        require(products[_productId].isRegistered, "Product not registered");

        Product memory product = products[_productId];
        return (
            product.productId,
            product.productName,
            product.manufacturer,
            product.isRegistered
        );
    }

    function getRegisterProductHash() external pure returns (bytes4) {
        return bytes4(keccak256("registerProduct(uint256,string)"));
    }

    function getGetProductHash() external pure returns (bytes4) {
        return bytes4(keccak256("getProduct(uint)"));
    }

    function test() external pure returns (bytes memory) {
        return bytes("ciao");
    }   

    function getTestHash() external pure returns (bytes4) {
        return bytes4(keccak256("test()"));
    }

    


}