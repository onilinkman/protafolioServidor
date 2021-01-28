package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"../../../handlers"
	"../../../models"
	"github.com/gorilla/mux"
)

//CreateUsuario crea un usuario y lo almacena en la base de datos
func CreateProducto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		producto := models.Producto{}
		fmt.Println(r.FormValue("nombre"))

		if err := json.Unmarshal([]byte(r.FormValue("cuerpo")), &producto); err != nil {
			models.SendUnprocessableEntity(w)

		} else {
			idpreg, err2 := producto.AddProducto()
			if err2 != nil {
				models.SendUnprocessableEntity(w)
				//r.fo
			} else {
				id, _ := idpreg.LastInsertId()
				handlers.RecibirArchivo(r, "./recursos/imagenesProductos/", strconv.FormatInt(id, 10))

			}
		}
	}
}

func ObtenerProducto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		idpro, err := strconv.Atoi(vars["idproducto"])
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintf(w, "%s", "Error in url")
			return
		}

		producto, err := models.GetProducto(idpro)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintf(w, "%s", "Error no content")
			return
		}
		var urld = fmt.Sprintf("http://%s/recursos/imagenesProductos/%d",
			models.Configuracion.Dominio, producto.IDProducto)

		producto.ImageURL = buscarExtension(fmt.Sprintf("./recursos/imagenesProductos/%d",
			producto.IDProducto), urld)
		models.SendData(w, producto)
	}
}

func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		productos, err := models.GetProductos()
		if err != nil {
			models.SendUnprocessableEntity(w)
			return
		}

		for i := 0; i < len(productos); i++ {
			producto := &productos[i]
			var urld = fmt.Sprintf("http://%s/recursos/imagenesProductos/%d",
				models.Configuracion.Dominio, producto.IDProducto)
			producto.ImageURL = buscarExtension(fmt.Sprintf("./recursos/imagenesProductos/%d",
				producto.IDProducto), urld)
		}

		models.SendData(w, productos)

	}
}

func buscarExtension(url, urld string) string {
	var cad = []string{".png", ".jpeg", ".gif", ".jpg"}
	for i := 0; i < len(cad); i++ {
		if _, err := os.Stat(url + cad[i]); err == nil {
			return urld + cad[i]
		}
	}
	return ""
}
