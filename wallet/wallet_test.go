package wallet

import (
	"reflect"
	"testing"

	pb "github.com/Pham-Xuan-Thanh/hdwallet/proto/btchdwallet"
)

const (
	passpharse      = "gentle sleep artist give drip rare sister change couple script era shove"
	prvkey          = "xprv9s21ZrQH143K3GgHrM6LSwX1d1kt5RrCb3YJzNUvfBvU7Dswk5ktYXJHxA9qRRwG3qbnLpFx56FhUPFntbcXLC21Mai1wnUzMXV9QxDCQeQ"
	pubkey          = "xpub661MyMwAqRbcFkkkxNdLp5TkB3bNUta3xGTunktYDXTSz2D6Hd596KcmoSaoaK4bGbALmGTe43xRDaWxKwDArjjwqzzRae1Wz5YtsiFH9DC"
	address         = "1Emr7QPngRH9VmJn2DdFR5PCPWkuuFz82x"
	wrongPasspharse = "power1 army2 apology3 model4 thumb foil place cactus dutch diary ahead drill science pride agree"
)

func TestDecodeWallet(t *testing.T) {
	t.Run("with legal syntax", func(t *testing.T) {
		want := &pb.Response{Address: address, PubKey: pubkey, PrivKey: prvkey}
		got := DecodeWallet(passpharse)

		if reflect.DeepEqual(*want, got) {
			t.Errorf("got %+v want %+v", got, want)

		}
	})

	t.Run("with illegal syntax", func(t *testing.T) {
		want := &pb.Response{}
		got := DecodeWallet(wrongPasspharse)
		if reflect.DeepEqual(want, got) {
			t.Errorf("got %+v want %+v", got, want)

		}
	})
}

func TestGetBalance(t *testing.T) {
	want := &pb.Response{Address: address, Balance: 0, TotalReceived: 0, TotalSent: 0, UnconfirmedBalance: 0}
	got := GetBalance(address)

	if reflect.DeepEqual(*want, got) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
