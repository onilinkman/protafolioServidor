package models

import "fmt"

const dominioLocal = "158.124.30.108"
const puerto = 8000

func init() {
	Configuracion.Dominio = fmt.Sprintf("%s:%d", dominioLocal, puerto)
	Configuracion.Puerto = puerto
}

//Config estructura de la configuracion
type Config struct {
	Dominio string
	Puerto  int
}

//Configuracion donde se almacena las configuraciones
var Configuracion Config
