package models

import (
	"database/sql"
	"fmt"
	"time"
)

//Producto estructura de un producto
type Producto struct {
	IDProducto  int       `json:"idProducto"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Texto       string    `json:"texto"`
	Categoria   string    `json:"categoria"`
	ImageURL    string    `json:"imageUrl"`
	Fecha       time.Time `json:"fecha"`
	EsWeb       string    `json:"esweb"`
}

//Productos array de productos
type Productos []Producto

var queryPregunta = `CREATE TABLE if NOT EXISTS producto(
	idproducto int primary KEY not null AUTO_INCREMENT,
	nombre varchar(45) not null,
	descripcion varchar(150),
	texto varchar(600),
	categoria varchar(45),
	imagenurl varchar(255),
	fecha DATETIME(0),
	esweb varchar(4)
)`

//CrearTablaProducto crea una tabla de productos
func CrearTablaProducto() error {
	_, err := EjecutarExec(queryPregunta)
	if err != nil {
		return err
	}
	return nil
}

//AddProducto agrega un producto
func (p *Producto) AddProducto() (sql.Result, error) {
	query := `INSERT INTO producto(nombre,descripcion,texto,categoria,imagenurl,fecha,esweb) VALUES (?,?,?,?,?,?,?)`
	result, err := EjecutarExec(query, &p.Nombre, &p.Descripcion, &p.Texto, &p.Categoria, &p.ImageURL, FechaAString(p.Fecha), &p.EsWeb)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//UpdateProducto actualiza en la tabla de la base de datos el producto
func (p *Producto) UpdateProducto() error {
	query := `UPDATE producto SET nombre=?,descripcion=?, texto=?, categoria=?,imagenurl=?,fecha=?,esweb=? WHERE idproducto=?`
	_, err := EjecutarExec(query, &p.Nombre, &p.Descripcion, &p.Texto, &p.Categoria, &p.ImageURL, FechaAString(p.Fecha), &p.EsWeb, &p.IDProducto)
	return err
}

//GetProducto devuelve un producto que coincida con el id de la tabla
func GetProducto(id int) (*Producto, error) {
	query := `SELECT idproducto,nombre,descripcion,texto,categoria,imagenurl,fecha,esweb FROM producto WHERE idproducto=?`
	row, err := EjecutarQuery(query, id)
	defer row.Close()
	prod := &Producto{}
	if err != nil {
		return prod, err
	}

	for row.Next() {

		var f string
		row.Scan(&prod.IDProducto, &prod.Nombre, &prod.Descripcion, &prod.Texto, &prod.Categoria, &prod.ImageURL, &f, &prod.EsWeb)
		fmt.Println(f)
		//prod.Fecha = StringAFecha(f)
	}
	return prod, nil

}

//GetProductos obtienes todos los productos de la tabla
func GetProductos() (Productos, error) {
	query := `SELECT idproducto,nombre,descripcion,texto,categoria,imagenurl,fecha,esweb FROM producto`
	rows, err := EjecutarQuery(query)
	defer rows.Close()
	productos := Productos{}

	for rows.Next() {
		producto := Producto{}
		var f string
		rows.Scan(&producto.IDProducto, &producto.Nombre, &producto.Descripcion,
			&producto.Texto, &producto.Categoria, &producto.ImageURL, &f, &producto.EsWeb)
		//producto.Fecha = StringAFecha(f)
		productos = append(productos, producto)
	}
	return productos, err
}

//DeleteProducto elimina un elemento de la tabla de productos
func DeleteProducto(id int) error {
	query := `DELETE FROM producto WHERE idproducto=?`
	_, err := EjecutarExec(query, id)
	return err
}
