// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

contract Registery {

    struct Event {
        string avsName;
        string operatorName;
        address avsAddress;
        address operatorAddress;
    }

    Event[] public events;

    function registerEvent(string calldata avsName,string calldata operatorName,address avsAddress,address operatorAddress) public  {
        events.push(Event(avsName,operatorName,avsAddress,operatorAddress));
    }

}