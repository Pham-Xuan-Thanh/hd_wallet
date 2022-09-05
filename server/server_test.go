package server

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "github.com/Pham-Xuan-Thanh/hdwallet/proto/btchdwallet"
)

const (
	passpharse      = "gentle sleep artist give drip rare sister change couple script era shove"
	prvkey          = "xprv9s21ZrQH143K3GgHrM6LSwX1d1kt5RrCb3YJzNUvfBvU7Dswk5ktYXJHxA9qRRwG3qbnLpFx56FhUPFntbcXLC21Mai1wnUzMXV9QxDCQeQ"
	pubkey          = "xpub661MyMwAqRbcFkkkxNdLp5TkB3bNUta3xGTunktYDXTSz2D6Hd596KcmoSaoaK4bGbALmGTe43xRDaWxKwDArjjwqzzRae1Wz5YtsiFH9DC"
	address         = "1Emr7QPngRH9VmJn2DdFR5PCPWkuuFz82x"
	wrongPasspharse = "power1 army2 apology3 model4 thumb foil place cactus dutch diary ahead drill science pride agree"
	wrongAddress    = "1Emr7QPngRH9VmJn2DdFR5PCPWkuuFz82xa"
)

func TestGetWallet(t *testing.T) {
	t.Run("Correct type mnemonics: ", func(t *testing.T) {
		server := Server{}

		want := &pb.Response{Address: address, PubKey: pubkey, PrivKey: prvkey}

		request := &pb.Request{Mnemonic: passpharse}
		cancleContext, cancle := context.WithCancel(context.Background())
		got, _ := server.GetWallet(cancleContext, request)
		time.AfterFunc(3*time.Second, cancle)

		if reflect.DeepEqual(got, *want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("Wrong type mnemonics: ", func(t *testing.T) {
		server := Server{}

		want := pb.Response{}

		request := &pb.Request{Mnemonic: wrongPasspharse}

		got, err := server.GetWallet(context.TODO(), request)

		if err != nil {
			t.Errorf("Get error %v while running @@", err)
		}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}

	})
}

func TestGetBalance(t *testing.T) {
	t.Run("Correct type mnemonics: ", func(t *testing.T) {
		server := Server{}

		want := &pb.Response{Address: address, Balance: 0, TotalReceived: 0, TotalSent: 0, UnconfirmedBalance: 0}

		request := &pb.Request{Address: address}
		cancleContext, cancle := context.WithCancel(context.Background())
		got, _ := server.GetBalance(cancleContext, request)
		time.AfterFunc(3*time.Second, cancle)

		if reflect.DeepEqual(got, *want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

}
