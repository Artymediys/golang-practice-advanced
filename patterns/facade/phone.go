package main

import (
	"fmt"
)

// Phone - contains info about a phone
type Phone struct {
	Desc  Description
	Comps Components
}

func (p *Phone) GetInfo() string {
	return fmt.Sprintf("%s %s %dGB, camera: %dMP, processor: \"%s\")",
		p.Desc.name, p.Desc.model, p.Comps.storage, p.Comps.camera, p.Comps.processor)
}

// Description - contains info about a phone's description
type Description struct {
	name  string
	model string
}

func (d *Description) SetDescription(name, model string) {
	d.name = name
	d.model = model
}

func (d *Description) GetDescription() (string, string) {
	return d.name, d.model
}

// Components - contains info about a phone's components
type Components struct {
	camera    uint8
	storage   uint16
	processor string
}

func (c *Components) SetComponents(camera uint8, storage uint16, processor string) {
	c.camera = camera
	c.storage = storage
	c.processor = processor
}

func (c *Components) GetComponents() (uint8, uint16, string) {
	return c.camera, c.storage, c.processor
}
