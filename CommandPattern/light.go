package main

import "fmt"

type light struct{}

func (l *light) label() string {
	return "light"
}

func (l *light) on() {
	fmt.Println("light is on")
}

func (l *light) off() {
	fmt.Println("light is off")
}

type lightOnCommand struct {
	l *light
}

func (c *lightOnCommand) execute() {
	c.l.on()
}

func (c *lightOnCommand) undo() {
	c.l.off()
}

type lightOffCommand struct {
	l *light
}

func (c *lightOffCommand) execute() {
	c.l.off()
}

func (c *lightOffCommand) undo() {
	c.l.on()
}
