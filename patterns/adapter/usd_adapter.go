package main

type DollarAdapter struct {
	dw *DollarWallet
}

func (adapter *DollarAdapter) ReceiveRUB(rub RubleValue) {
	adapter.dw.ReceiveUSD(DollarValue{
		dollars: rune(rub.rubles / 72),
		cents:   rune(rub.kopecks),
	})
}
