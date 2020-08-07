package goslp

import (
	"fmt"
	"testing"

	"github.com/gcash/bchd/chaincfg"
)

func TestDecodeSlpAddress(t *testing.T) {
	addrStr := "simpleledger:qrkjty23a5yl7vcvcnyh4dpnxxzuzs4lzqvesp65yq"
	prefix, data, _ := DecodeCashAddress(addrStr)
	if prefix != "simpleledger" {
		t.Fatal("decode failed")
	}
	if len(data) != 34 {
		t.Fatal("data wrong length")
	}
}

func TestDecodeAddressMainnet(t *testing.T) {
	addrStr := "qrkjty23a5yl7vcvcnyh4dpnxxzuzs4lzqvesp65yq"
	addr, err := DecodeAddress(addrStr, &chaincfg.MainNetParams)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addr.String())
	if addr.String() != addrStr {
		t.Fatal("decode failed")
	}

}