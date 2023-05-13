package main

import "log"

type DollarValue struct {
	dollars rune
	cents   rune
}

type DollarWallet struct {
	balance DollarValue
}

func (dw *DollarWallet) ReceiveUSD(usd DollarValue) {
	dw.balance.dollars += usd.dollars
	dw.balance.cents += usd.cents

	log.Printf("Received: (%d.%d)USD", usd.dollars, usd.cents)
}
