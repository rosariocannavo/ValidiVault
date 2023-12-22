// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//proxy dp 

contract ProductProxy {
    address public logicContract;

    constructor(address _logicContract) {
        logicContract = _logicContract;
    }

    fallback() external payable {
        address target = logicContract;
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, 0, calldatasize())
            let result := delegatecall(gas(), target, ptr, calldatasize(), 0, 0)
            let size := returndatasize()
            returndatacopy(ptr, 0, size)

            switch result
            case 0 {
                revert(ptr, size)
            }
            default {
                return(ptr, size)
            }
        }
    }

    receive() external payable {
        // Receive function to handle incoming Ether (optional logic)
        // This function will be triggered when Ether is sent directly to the contract
    }

    function changeLogicContract(address newLogicContract) external {
        logicContract = newLogicContract;
    }
}