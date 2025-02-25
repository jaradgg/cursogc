package main

import (
	"crypto"

	"github.com/jaradgg/go-course/003-packages/"
)

func main() {
	cryptocurrency.ToBitcoin(10000)
	cryptocurrency.ToEthereum(10000)
	cryptocurrency.ToLitecoin(10000)
}