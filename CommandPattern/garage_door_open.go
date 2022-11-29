package main

import "fmt"

type garageDoor struct{}

func (g *garageDoor) label() string {
	return "garage_door"
}

func (g *garageDoor) up() {
	fmt.Println("garage door is open")
}

func (g *garageDoor) down() {
	fmt.Println("garage door is closed")
}

func (g *garageDoor) stop() {
	fmt.Println("garage door is stopped")
}

func (g *garageDoor) lightOn() {
	fmt.Println("garage door is light on")
}

func (g *garageDoor) lightOff() {
	fmt.Println("garage door is light off")
}

type garageDoorOpenCommand struct {
	gd *garageDoor
}

func (gCommand *garageDoorOpenCommand) execute() {
	gCommand.gd.up()
}

func (gCommand *garageDoorOpenCommand) undo() {
	gCommand.gd.down()
}

type garageDoorCloseCommand struct {
	gd *garageDoor
}

func (gCommand *garageDoorCloseCommand) execute() {
	gCommand.gd.down()
}

func (gCommand *garageDoorCloseCommand) undo() {
	gCommand.gd.up()
}
