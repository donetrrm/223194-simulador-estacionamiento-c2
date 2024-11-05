package main

import (
	"223194-simulador-estacionamiento/scenes"
	"github.com/oakmound/oak/v4"
)

func main() {

	scenes.NewEstacionamientoScene().Start()

	if err := oak.Init("estacionamiento"); err != nil {
		panic(err)
	}

}
