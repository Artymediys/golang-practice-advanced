package main

func main() {
	rub := &RubleValue{
		rubles:  13,
		kopecks: 37,
	}

	usd := &DollarValue{
		dollars: 32,
		cents:   28,
	}

	dollarAdapter := &DollarAdapter{
		&DollarWallet{
			balance: *usd,
		},
	}

	dollarAdapter.ReceiveRUB(*rub)
}
