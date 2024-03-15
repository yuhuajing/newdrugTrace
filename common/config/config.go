package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/ethconn"
)

var (
	EthServer = "https://eth-sepolia.g.alchemy.com/v2/t39LUOfkEMNMz_uxQk2gkwK3kJ1HwDZF"
	Client    = ethconn.ConnBlockchain(EthServer)
	TraceSC   = "0x9f07b5fE8D3b0Fd020144Bbd70E3481223fD547c"
)
var (
	Mongodburl           = "mongodb://clay:password@127.0.0.1:27017" // "mongodb://clay:password@127.0.0.1:27017"
	Dbname               = "drugtra"
	DbcollectionUserInfo = "userinfo"
	DbcollectionEcoInfo  = "prodinfo"
	DbcollectionTousu    = "tousu"
	Mongoclient          *mongo.Client
)
