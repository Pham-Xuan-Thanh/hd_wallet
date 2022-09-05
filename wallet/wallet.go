package wallet

import (
	"fmt"

	"github.com/Pham-Xuan-Thanh/hdwallet/config"
	"github.com/Pham-Xuan-Thanh/hdwallet/crypt"
	pb "github.com/Pham-Xuan-Thanh/hdwallet/proto/btchdwallet"
	"github.com/blockcypher/gobcy"
	"github.com/brianium/mnemonic"
	"github.com/wemeetagain/go-hdwallet"
)

var conf = config.ParseConfig("../config.yml")

// CreateWallet is in charge of creating a new root wallet
func CreateWallet() *pb.Response {
	// Generate a random 256 bit seed
	seed := crypt.CreateHash()
	mnemonic, _ := mnemonic.New([]byte(seed), mnemonic.English)

	// Create a master private key
	masterprv := hdwallet.MasterKey([]byte(mnemonic.Sentence()))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	return &pb.Response{Address: address, PubKey: masterpub.String(), PrivKey: masterprv.String(), Mnemonic: mnemonic.Sentence()}
}

// DecodeWallet is in charge of decoding wallet from mnemonic
func DecodeWallet(mnemonic string) *pb.Response {
	// Get private key from mnemonic
	masterprv := hdwallet.MasterKey([]byte(mnemonic))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	return &pb.Response{Address: address, PubKey: masterpub.String(), PrivKey: masterprv.String()}
}

func GetBalance(address string) *pb.Response {
	btc := gobcy.API{Token: conf.Blockchain.Token, Coin: "btc", Chain: "main"}
	addr, err := btc.GetAddrBal(address, nil)
	if err != nil {
		fmt.Println("Wallet: Error why get balance from Blockcypher: ", err)
	}

	balance := addr.Balance
	totalReceived := addr.TotalReceived
	totalSent := addr.TotalSent
	unconfirmedBalanced := addr.UnconfirmedBalance

	return &pb.Response{Address: address, Balance: balance.Int64(), TotalReceived: totalReceived.Int64(), TotalSent: totalSent.Int64(), UnconfirmedBalance: unconfirmedBalanced.Int64()}
}
