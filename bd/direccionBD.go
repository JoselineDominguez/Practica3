package bd

import "ejemplo/estructura"

func GetDireccion() (direcciones []estructura.Direccion) {
	//se crean dos variables en las que se almacenan los valores de forma temporal antes de guardarlos en el slice
	var (
		idDireccion int
		calle string
		colonia string
		alcaldia string
		codigoPostal string
		numeroInterior string
		numeroExterior string
		idUsuario int
		idTipoDireccion byte
		direccions []Direccion
	  )
	//Se ejecuta el query, como se están solicitando 'n' cantidad de datos se usa 'Query'
	//Si se necesitara obtener un solo dato (como en el caso de requirir un dato mediante el id) se usaria 'QueryRow'
	direccion, _ := db.Query("SELECT IdDireccion, calle, colonia, alcaldia, codigoPostal,
								numeroInterior, numeroExterior, idUsuario, idTipoDireccion FROM Direccion")
	revisarError(err)
	for direccion.Next() {
		//se aconsiguen los datos de la BD
		err = direccion.Scan(
			&idDireccion,
			&calle,
			&colonia,
			&alcaldia,
			&codigoPostal,
			&numeroInterior,
			&numeroExterior,
			&idUsuario,
			&idTipoDireccion,
		)
		revisarError(err)
		//Se almacenan en el slice para ser regresados
		membresias = append(membresias, estructura.Membresia{
			IdDireccion:   		idDireccion,
			Calle:				calle,
			Colonia:			colonia,
			Alcaldia:			alcaldia,
			CodigoPostal:		codigoPostal,
			NumeroInterior:		numeroInterior,
			NumeroExterior:		numeroExterior,
			IdUsuario:			idUsuario,
			IdidTipoDireccion:	TipoDireccion,
		})
	}

	//Se retorna el slice 'direcciones'
	return
}