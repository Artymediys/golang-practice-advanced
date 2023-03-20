package main

import "log"

func main() {
	var ps *PhoneStorage = PhoneStorageInit()

	ps.NewPhone("iPhone")

	log.Println(ps.GetTotalSupply())
	log.Println(ps.GetPhone(0))
}
