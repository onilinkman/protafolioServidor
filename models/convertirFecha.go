package models

import (
	"time"

	"fmt"
)

//insert into horaEvaluacion (horainicio,horafinalizacion) values ('2020-02-28 17:30:00','2021-08-15 18:00:00');

//FechaAString convierte una fecha a un string para llevarlo a la base de datos
func FechaAString(t time.Time) string {
	hora := t.Hour()
	minutos := t.Minute()
	segundos := t.Second()
	anio := t.Year()
	mes := int(t.Month())
	dia := t.Day()

	return fmt.Sprintf("%d-%d-%d %d:%d:%d", anio, mes, dia, hora, minutos, segundos)
}

//StringAFecha convierte de cadena a Time
func StringAFecha(f string) time.Time {
	t1, err := time.Parse("2006-1-2 15:04:05", f)
	if err != nil {
		fmt.Println(err)
	}
	return t1
}
