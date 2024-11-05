package models

type CajonEstacionamientoDireccion struct {
	Direccion string
	Ubicacion float64
}

func newDireccionCajon(direccion string, ubicacion float64) *CajonEstacionamientoDireccion {
	return &CajonEstacionamientoDireccion{
		Direccion: direccion,
		Ubicacion: ubicacion,
	}
}
