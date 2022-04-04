package dynamicgetter

import (
	"errors"
	"fmt"
	"reflect"
)

func GetField(v interface{}, name string) (interface{}, error) {
	// v debe ser un puntero a una estructura
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return nil, errors.New("v debe ser un puntero a una estructura")
	}

	// Obtengo el valor subyacente al puntero
	// Textual de la documentacion de reflect:
	// Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	rv = rv.Elem()

	// Obtengo el campo de la estruct a partir de su nombre
	fv := rv.FieldByName(name)
	// Verifico que el campo exista dentro de mi estructura
	if !fv.IsValid() {
		return nil, fmt.Errorf("%s no existe en la estructura", name)
	}

	// Si el campo no está exportado, no deberiamos poder acceder a el
	if !fv.CanSet() {
		return nil, fmt.Errorf("no es posible acceder al campo %s", name)
	}

	/*
		//EXTRA:
		// Si quisiesemos ademas verificar que el valor del campo sea de un tipo determinado
		// podemos usar un codigo similar al siguiente
		if fv.Kind() != reflect.String {
			return nil, fmt.Errorf("%s is not a string field", name)
		}

	*/
	// Si el valor es el zero value de su tipo, devolvemos un error
	if fv.IsZero() {
		return nil, fmt.Errorf("el campo %s esta vacio", name)
	}

	// retornamos el valor del campo, y un error nulo
	return fv, nil
}
