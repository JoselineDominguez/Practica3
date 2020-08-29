package estructura

type Direccion struct {
	//el nombre que se utilice en 'json' debe ser el mismo que se utilice en el front para enviar la informacion
	IdDireccion int32  `json:"idDireccion"`
	Calle string  `json:"calle"`
  	Colonia string  `json:"colonia"`
  	Alcaldia string  `json:"alcaldia"`
  	CodigoPostal string  `json:"codigoPostal"`
  	NumeroInterior string  `json:"numeroInterior"`
  	NumeroExterior string  `json:"numeroExterior"`
  	IdUsuario int32  `json:"idUsuario"`
  	IdTipoDireccion byte  `json:"idTipoDireccion"`
}
