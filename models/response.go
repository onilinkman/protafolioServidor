package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Response estructura del response
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

//CreateDefaultResponse crea el valor por defecto del response
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}
}

//SendNotFound envia cuando no haya contenido
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

//NotFound Envia cuando no hay datos
func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource Not Found"
}

//SendUnprocessableEntity cuado una entidad no se puede procesar
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

//UnprocessableEntity prepara para cuando una netidad no se puede procesar
func (r *Response) UnprocessableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Unprocessable Entity"
}

//SendNoContent envia la respuesta cuando no hay contenido
func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

//NoContent lo prepara cuando no haya contenido
func (r *Response) NoContent() {
	r.Status = http.StatusNoContent
	r.Message = "No Content"
}

//SendData prepara para enviar los datos
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

//Send envia los datos
func (r *Response) Send() {
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.Header().Set("Access-Control-Allow-Origin", "*")
	r.writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.writer.WriteHeader(r.Status)
	output, _ := json.Marshal(&r)
	fmt.Fprintf(r.writer, "%s", string(output))
}
