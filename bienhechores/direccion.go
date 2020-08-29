package bienhechores

import (
	"ejemplo/bd"
	"ejemplo/estructura"
	"encoding/json"
	"net/http"
)

func GetDireccionEndPoint(w http.ResponseWriter, r *http.Request) {
	//Se codifica la informacion en lenguaje 'json' y se imprime en el navegador 'json.NewEncoder(w)'
	//la informacion que se codifica es el resultado de la funcion 'bd.GetDireccion()' el cual será un slice
	json.NewEncoder(w).Encode(bd.GetDireccion())
}

func PostDireccionEndPoint(w http.ResponseWriter, r *http.Request) {
	var direccion estructura.Direccion
	//Se decodifica del lengiaje 'json' la información del cuerpo 'r.Body' del formulario que es enviado
	//Y se almacena en la variable 'direccion'
	json.NewDecoder(r.Body).Decode(&direccion)
	//Se manda la variable 'direccion' a la funcion 'bd.PostMembresia' y el error que regrese se manda al navegador
	json.NewEncoder(w).Encode(bd.PostMembresia(direccion))

}
