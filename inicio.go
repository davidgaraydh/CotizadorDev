package main

import (
	"fmt"
	"strings"
)

type Proyecto struct {
	TiempoHoraP, PrecioHora int
	NivelComplejidad        string
}

type Cliente struct {
	NivelCalidad, TipoAdquisicion, tipoCliente string
}
type Interno struct {
	porcentajeNego int
	NivelExpertise, tecnologias,
	documentacion, personalExterno,
	liderazgo string
}

var Valores [3]int
var Valoresring [9]string
var contador int = 0
var contador2 int = 0
var suma float32 = 0

func main() {
	var TiempoHora, PrecioHora, porcentajeNego int
	var NivelExpertise, Tecnologias, NivelCalidad, TipoAdquisicion,
		documentacion, tipoCliente, personalExterno, NivelComplejidad, Liderazgo string

	Proyecto := Proyecto{
		TiempoHoraP:      TiempoHora,
		PrecioHora:       PrecioHora,
		NivelComplejidad: NivelComplejidad,
	}

	Cliente := Cliente{
		NivelCalidad:    NivelCalidad,
		TipoAdquisicion: TipoAdquisicion,
		tipoCliente:     tipoCliente,
	}

	Interno := Interno{
		porcentajeNego:  porcentajeNego,
		NivelExpertise:  NivelExpertise,
		tecnologias:     Tecnologias,
		documentacion:   documentacion,
		personalExterno: personalExterno,
		liderazgo:       Liderazgo,
	}

	AllReaders(Proyecto, Cliente, Interno)

}

func AllReaders(ProyectoVar Proyecto, ClienteVar Cliente, InternoVar Interno) {

	LecturaDatosInt("¿Ingresa el tiempo del proyecto (Horas)? (I)", ProyectoVar.TiempoHoraP)
	LecturaDatosString("Seleccione el nivel de expertise del stack tech (S)", InternoVar.NivelExpertise)
	LecturaDatosString("Ingrese las tecnologias (S)", InternoVar.tecnologias)
	LecturaDatosString("Seleccione el nivel de calidad (S)", ClienteVar.NivelCalidad)
	LecturaDatosInt("Introduce el Precio por Hora (I)", ProyectoVar.PrecioHora)
	LecturaDatosString("Seleccione el tipo de adquisicion (S)", ClienteVar.TipoAdquisicion)
	LecturaDatosString("Seleccione la documentacion que se creara (S)", InternoVar.documentacion)
	LecturaDatosString("¿Que tipo de cliente es? (S)", ClienteVar.tipoCliente)
	LecturaDatosString("¿Habra personal externo? (S)", InternoVar.personalExterno)
	LecturaDatosInt("Ingrese el porcentaje de Negociacion (I)", InternoVar.porcentajeNego)
	LecturaDatosString("Ingrese el nivel de complejidad (S)", ProyectoVar.NivelComplejidad)
	LecturaDatosString("El proyecto requiere liderazgo (S)", InternoVar.liderazgo)

	CalculoProyectoInicial()
	CalculoProyectoIntermedio()
	CalculoAvanzado()

	fmt.Println(suma)

}

func LecturaDatosInt(impresion string, varLeidaI int) {
	fmt.Println(impresion)
	fmt.Scan(&varLeidaI)
	Valores[contador] = varLeidaI
	contador++
}

func LecturaDatosString(impresion, varLeidaS string) {
	fmt.Println(impresion)
	fmt.Scan(&varLeidaS)
	Valoresring[contador2] = varLeidaS
	contador2++
}

func CalculoProyectoInicial() {

	var indiceExpertise float32

	if Valoresring[3] == "Empresarial" {

		switch Valoresring[0] {
		case "Beg":
			indiceExpertise = 1.5
		case "JR":
			indiceExpertise = 1.4
		case "Semi":
			indiceExpertise = 1.1
		case "Senior":
			indiceExpertise = 1.05
		default:
			break
		}

	} else if Valoresring[3] == "Freelance" {

		switch Valoresring[0] {
		case "Beg":
			indiceExpertise = 0.7
		case "JR":
			indiceExpertise = 0.8
		case "Semi":
			indiceExpertise = 1.05
		case "Senior":
			indiceExpertise = 1.1
		default:
			break
		}

	}

	suma = (indiceExpertise * float32(Valores[1])) * float32(Valores[0])

}

func CalculoProyectoIntermedio() {

	var indiceExpertiseIntermedio float32 = 0
	_ = indiceExpertiseIntermedio

	if Valoresring[2] == "detalle" {
		suma = float32(suma) * 1.150

	}

	split := strings.Split(Valoresring[4], ",")

	for i := 0; i < len(split); i++ {
		switch split[i] {
		case "CU":
			indiceExpertiseIntermedio = 1.050 + indiceExpertiseIntermedio
		case "CP":
			indiceExpertiseIntermedio = 1.050 + indiceExpertiseIntermedio
		case "DF":
			indiceExpertiseIntermedio = 1.3 + indiceExpertiseIntermedio
		case "DC":
			indiceExpertiseIntermedio = 1.5 + indiceExpertiseIntermedio
		case "DD":
			indiceExpertiseIntermedio = 0.7 + indiceExpertiseIntermedio
		case "MT":
			indiceExpertiseIntermedio = 1.6 + indiceExpertiseIntermedio
		case "MU":
			indiceExpertiseIntermedio = 1.2 + indiceExpertiseIntermedio
		default:
		}
	}

	if Valoresring[7] == "JR" {
		suma = suma * 1.050
	} else if Valoresring[7] == "SEMI" {
		suma = suma * 1.2
	} else if Valoresring[7] == "SENIOR" {
		suma = suma * 1.4
	}

	suma = suma * indiceExpertiseIntermedio

}

func CalculoAvanzado() {

	var porcentaje float32
	_ = porcentaje

	if Valoresring[5] == "CH" {
		suma = suma * 1.050
	} else if Valoresring[5] == "MD" {
		suma = suma * 1.2
	} else if Valoresring[5] == "GD" {
		suma = suma * 1.3
	}

	if Valoresring[8] == "SI" {
		suma = suma * 1.1
	}

	porcentaje = (float32(Valores[2]) * suma)
	suma = suma + porcentaje

}
