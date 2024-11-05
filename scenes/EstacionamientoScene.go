package scenes

import (
	"223194-simulador-estacionamiento/models"
	"image/color"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

var (
	cajones = []*models.CajonEstacionamiento{
		models.NewCajonEstacionamiento(410, 210, 440, 240, 1, 1),
		models.NewCajonEstacionamiento(410, 255, 440, 285, 1, 2),
		models.NewCajonEstacionamiento(410, 300, 440, 330, 1, 3),
		models.NewCajonEstacionamiento(410, 345, 440, 375, 1, 4),

		models.NewCajonEstacionamiento(320, 210, 350, 240, 2, 5),
		models.NewCajonEstacionamiento(320, 255, 350, 285, 2, 6),
		models.NewCajonEstacionamiento(320, 300, 350, 330, 2, 7),
		models.NewCajonEstacionamiento(320, 345, 350, 375, 2, 8),

		models.NewCajonEstacionamiento(230, 210, 260, 240, 3, 9),
		models.NewCajonEstacionamiento(230, 255, 260, 285, 3, 10),
		models.NewCajonEstacionamiento(230, 300, 260, 330, 3, 11),
		models.NewCajonEstacionamiento(230, 345, 260, 375, 3, 12),

		models.NewCajonEstacionamiento(140, 210, 170, 240, 4, 13),
		models.NewCajonEstacionamiento(140, 255, 170, 285, 4, 14),
		models.NewCajonEstacionamiento(140, 300, 170, 330, 4, 15),
		models.NewCajonEstacionamiento(140, 345, 170, 375, 4, 16),

		models.NewCajonEstacionamiento(50, 210, 80, 240, 5, 17),
		models.NewCajonEstacionamiento(50, 255, 80, 285, 5, 18),
		models.NewCajonEstacionamiento(50, 300, 80, 330, 5, 19),
		models.NewCajonEstacionamiento(50, 345, 80, 375, 5, 20),
	}
	cajonesEstacionamiento = models.NewCajonEstacionamineto(cajones)
	doorMutex              sync.Mutex
	carAdmin               = models.NewCarAdministrador()
)

type EstacionamientoScene struct{}

func NewEstacionamientoScene() *EstacionamientoScene {
	return &EstacionamientoScene{}
}

func (ps *EstacionamientoScene) Start() {
	isFirstTime := true

	oak.AddScene("estacionamiento", scene.Scene{
		Start: func(ctx *scene.Context) {
			err := oak.SetTitle("C2 Simulador Estacionamiento")
			if err != nil {
				return
			}
			setUpScene(ctx)
			event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
				if !isFirstTime {
					return 0
				}
				isFirstTime = false

				for i := 0; i < 100; i++ {
					go carCycle(ctx)
					time.Sleep(time.Millisecond * time.Duration(getNumeroAleatorio(1000, 2000)))
				}
				return 0
			})
		},
	})
}

func setUpScene(ctx *scene.Context) {
	backgroundRender, err := render.LoadSprite("assets/images/background.jpg")
	if err != nil {
		log.Fatal(err)
	}

	entities.New(ctx,
		entities.WithRenderable(backgroundRender),
		entities.WithDrawLayers([]int{-1}),
	)

	puerta := floatgeom.NewRect2(440, 170, 500, 180)
	entities.New(ctx,
		entities.WithRect(puerta),
		entities.WithColor(color.RGBA{255, 255, 255, 255}),
	)
}

func carCycle(ctx *scene.Context) {
	carro := models.NewCarro(ctx)
	//println("Carro generado")
	carAdmin.AgregarCarro(carro)
	carro.Enqueue(carAdmin)

	cajonDisponible := cajonesEstacionamiento.GetCajonEstacionamientoDisponible()

	doorMutex.Lock()
	carro.PuertaEntrada(carAdmin)
	doorMutex.Unlock()

	carro.Estacionamiento(cajonDisponible, carAdmin)
	tiempoEstacionamiento := time.Millisecond * time.Duration(getNumeroAleatorio(30000, 50000))
	println("Carro estacionado por: ", tiempoEstacionamiento/time.Second, "s en el cajón", cajonDisponible.GetNumeroDeCajon())
	time.Sleep(tiempoEstacionamiento)
	carro.DejarLugar(carAdmin)
	cajonesEstacionamiento.LiberarCajon(cajonDisponible)
	carro.DejarCajon(cajonDisponible, carAdmin)

	doorMutex.Lock()
	carro.PuertaSalida(carAdmin)
	doorMutex.Unlock()

	carro.Irse(carAdmin)
	carro.Eliminar()
	carAdmin.EliminarCarroScene(carro)
}

func getNumeroAleatorio(min, max int) float64 {
	return float64(rand.Intn(max-min+1) + min)
}
