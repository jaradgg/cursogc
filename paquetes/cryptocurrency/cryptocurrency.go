package cryptocurrency

import (
	"fmt"
)

// Cryptocurrency es una estructura que representa una criptomoneda
var (
	bitcoinPrice float32
	ethereumPrice float32
	litecoinPrice float32
)

func init() {
	bitcoinPrice = 50000
	ethereumPrice = 3000
	litecoinPrice = 150
}

func ToBitcoin(dolars float32) {
	bitcoins := dolars / bitcoinPrice
	fmt.Printf("%.2f USD equivalen a ~%.6f Bitcoins\n",dolars, bitcoins)
}

// ToEthereum convierte una cantidad de dolares a Ethereum
func ToEthereum(dolars float32) {
	ethereum := dolars / ethereumPrice
	fmt.Printf("%.2f USD equivalen a ~%.6f Ethereum\n",dolars, ethereum)
}

// ToLitecoin convierte una cantidad de dolares a Litecoin
func ToLitecoin(dolars float32) {
	litecoin := dolars / litecoinPrice
	fmt.Printf("%.2f USD equivalen a ~%.6f Litecoin\n",dolars, litecoin)
}