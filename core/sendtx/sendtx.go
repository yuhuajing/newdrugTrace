package sendtx

import "C"
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/common/config"
	"main/trace"
	"math/big"
)

func gentx(gaslimit uint64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("d6def01217bfe78db0dae301766ffc3bc42b76f4c68b2be53ffdbc3a28a95364")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chianid, _ := config.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chianid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gaslimit   // in units
	auth.GasPrice = gasPrice
	return auth
}
func AddOrUpdateProdData(id, prodData string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(800000)
	tx, err := instance.AddOrupdateProdData(auth, id, prodData)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
func AddOrUpdateStoreData(id, prodData string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(800000)
	tx, err := instance.AddOrupdateStoreData(auth, id, prodData)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
func AddOrUpdateLogisData(id, prodData string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(800000)
	tx, err := instance.AddOrupdateLogisData(auth, id, prodData)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
func AddOrUpdateSaleData(id, prodData string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(800000)
	tx, err := instance.AddOrupdateSaleData(auth, id, prodData)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func ReadDataByID(ids string) (error, string, string, string, string, string, string) {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		return err, "", "", "", "", "", ""
	}
	tx, err := instance.Tracedata(nil, ids[1:])
	if err != nil {
		return err, "", "", "", "", "", ""
	}
	return nil, tx.Traceid, tx.Batchid, tx.Prodinfo, tx.Logisinfo, tx.Storeinfo, tx.Salestring
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
