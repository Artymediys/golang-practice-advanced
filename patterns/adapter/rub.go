package main

import "log"

type IRubReceiver interface {
	ReceiveRUB(rub RubleValue)
}

type RubleValue struct {
	rubles  uint16
	kopecks uint8
}

type RubleWallet struct {
	balance RubleValue
}

func (rw *RubleWallet) ReceiveRUB(rub RubleValue) {
	rw.balance.rubles += rub.rubles
	rw.balance.kopecks += rub.kopecks

	log.Printf("Received: (%d.%d)RUB", rub.rubles, rub.kopecks)
}
