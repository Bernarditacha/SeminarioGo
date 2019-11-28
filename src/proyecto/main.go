package main

import (
	"fmt"
	"proyecto/otra"
	//"encoding/json"
	//"errors"
)

func main(){
	fmt.Println("Hello world")
	fmt.Println()

	c := otra.SumarValores(4,5)
	fmt.Println(c)

	autoX := NewAuto("Ford","Nuevo", "a3")
	autoZ := NewAuto("Chevrolet","Viejo", "b6")
	autoY := NewAuto("lalala","Antiguo", "b4")
	agenciaX := NewAgencia()
	agenciaX.AgregarAuto(autoX)
	agenciaX.AgregarAuto(autoZ)
	agenciaX.AgregarAuto(autoY)
	
	/*autoA := agenciaX.DevolverAuto(1)
	
	fmt.Println(autoA.GetModelo())
	agenciaX.EliminarAuto(1)
	autoB := agenciaX.DevolverAuto(1)*/
	autoH := agenciaX.DevolverAutoByPatente("b4")
	fmt.Println(autoH.GetModelo())



	//PUNTEROS
	/*		
	a := 5 
	b := 10

	var d*int 
	d  = &b

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(d)
	fmt.Println(&d)*/


	/*//MARSHAL
	myString := "hola como andas"

	fmt.Println(myString)

	var a []byte

	a, _ = json.Marshal(myString) //Convertir un string a un arreglo de bytes

	fmt.Println(a)

	//UNMARSHAL
	var otroString string

	_= json.Unmarshal(a, &otroString) // Convertir un  arr de bytes a string 

	fmt.Println("otroString", otroString)*/

/* 	EJEMPLO DE MARSHAL CON UN OBJETO
	miAuto := Auto{
		marca : "Renault",
		modelo : "Nuevo",
	}

	fmt.Println(miAuto)

	miAutoMarshalled, _ := json.Marshal(miAuto)
	fmt.Println(miAutoMarshalled)
	fmt.Println(string(miAutoMarshalled))

	var reqAuto Auto
	_ = json.Unmarshal(miAutoMarshalled,&reqAuto)
	fmt.Println("reqAuto", reqAuto)
*/







}

type Agencia struct{
	 autos [] Auto
}


type Auto struct{
	marca string
	modelo string
	patente string
}

func (a *Auto) GetModelo() string{
	return a.modelo
}

func (a *Auto) GetPatente() string{
	return a.patente
}


func (a *Auto) SetModelo(modelo string) {
	a.modelo = modelo
}

func (a *Agencia) setModeloAuto(auto Auto, modelo string){
	for i:= 0 ; i<len(a.autos); i++{
		if a.autos[i]  == auto {
			auto.SetModelo(modelo)
		}
	}
}

func (a *Agencia) DevolverAutoByPatente (patente string) Auto{
	for i:= 0 ; i<len(a.autos); i++{
		if a.autos[i].GetPatente() == patente{
			return a.autos[i]
		}
	}
	return a.autos[0]
}


func (a *Agencia) AgregarAuto( auto Auto){
a.autos = append(a.autos, auto)
}

func NewAuto(marca string, modelo string, patente string) Auto{
	return Auto{
		marca: marca,
		modelo: modelo,
		patente: patente,
	}
}

func (a *Agencia) EliminarAuto(pos int){

	/*if pos == 0 {
		if len(a.autos)>1{
		a.autos = a.autos[1:len(a.autos)]
	}else{
		var autos [] Auto
		a.autos = autos
	}
	}
	else if pos == len(a.autos)-1{
		a.autos = a.autos[0:len(a.autos)-1]
	}
	else{*/
	aux := a.autos[0:pos]
	aux2 := a.autos[pos+1:len(a.autos)]
	
	for i :=0; i < len(aux2); i++{
		aux = append(aux, aux2[i])
	}
	a.autos = aux
}






func (a *Agencia) DevolverAuto (pos int) Auto{
	return a.autos[pos]
}

func NewAgencia() Agencia{
	var autos [] Auto
	return Agencia	{
		autos: autos,
	}
}


