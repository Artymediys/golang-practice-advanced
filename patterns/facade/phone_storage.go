package main

import (
	"log"
	"strings"
	"sync"
)

// PhoneStorage - Facade for working with phones
type PhoneStorage struct {
	mutex       sync.RWMutex
	totalSupply uint64
	phones      map[uint64]*Phone
}

func PhoneStorageInit() *PhoneStorage {
	return &PhoneStorage{
		totalSupply: 0,
		phones:      make(map[uint64]*Phone),
	}
}

func (ps *PhoneStorage) NewPhone(name string) {
	var desc Description
	var comps Components

	switch strings.ToLower(name) {
	case "iphone":
		desc.SetDescription("Apple iPhone", "14 Pro")
		comps.SetComponents(48, 256, "Apple A16 Bionic")
	case "samsung":
		desc.SetDescription("Samsung", "Galaxy S23 Ultra")
		comps.SetComponents(200, 256, "Snapdragon 8 Gen 2")
	default:
		desc.SetDescription("Bruh", "RofloPhone 3228")
		comps.SetComponents(42, 128, "Intel Pentium")
	}

	// Lock Zone (RWMutex)
	ps.mutex.RLock()

	ps.phones[ps.totalSupply] = &Phone{
		Desc:  desc,
		Comps: comps,
	}
	log.Println(ps.phones[ps.totalSupply].GetInfo())

	ps.totalSupply++

	ps.mutex.RUnlock()
}

func (ps *PhoneStorage) GetTotalSupply() uint64 {
	return ps.totalSupply
}

func (ps *PhoneStorage) GetPhone(id uint64) Phone {
	return *ps.phones[id]
}
