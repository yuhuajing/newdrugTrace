// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract RockPaperScissor {
    mapping(string => MintStageInfo) public tracedata;

    struct MintStageInfo {
        string traceid;
        string prodinfo;
        string storeinfo;
        string logisinfo;
        string salestring;
        string batchid;
    }
    address producer;
    address storer;
    address logostic;
    address sales;

    constructor(
        address _producer,
        address _storer,
        address _logostic,
        address _sales
    ) {
        producer = _producer;
        storer = _storer;
        logostic = _logostic;
        sales = _sales;
    }

    function addOrupdateProdData(string memory id, string memory data)
        external
    {
        tracedata[id].prodinfo = data;
    }

    function addOrupdateStoreData(string memory id, string memory data)
        external
    {
        tracedata[id].storeinfo = data;
    }

    function addOrupdateLogisData(string memory id, string memory data)
        external
    {
        tracedata[id].logisinfo = data;
    }

    function addOrupdateSaleData(string memory id, string memory data)
        external
    {
        tracedata[id].salestring = data;
    }

    function tracedataById(string memory id)
        external
        view
        returns (MintStageInfo memory)
    {
        return tracedata[id];
    }
}