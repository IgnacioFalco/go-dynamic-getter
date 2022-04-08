package dynamicgetter

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrInvalidField = fmt.Errorf("invalid field")
var ErrInvalidObject = errors.New("v must be a pointer to struct")
var ErrUnaccesableValue = fmt.Errorf("cannot read value of unexported field")
var ErrZeroValue = fmt.Errorf("value of field is zero")

func GetField(v interface{}, name string, ignoreZero bool) (interface{}, error) {
	// v debe ser un puntero a una estructura
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return nil, ErrInvalidObject
	}

	// Obtengo el valor subyacente al puntero
	// Textual de la documentacion de reflect:
	// Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	rv = rv.Elem()

	// Obtengo el campo de la estruct a partir de su nombre
	fv := rv.FieldByName(name)
	// Verifico que el campo exista dentro de mi estructura
	if !fv.IsValid() {
		return nil, ErrInvalidField
	}

	// Si el campo no está exportado, no deberiamos poder acceder a el
	if !fv.CanSet() {
		return nil, ErrUnaccesableValue
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
	if !ignoreZero {
		if fv.IsZero() {
			return nil, ErrZeroValue
		}

	}
	// retornamos el valor del campo, y un error nulo
	return fv.Interface(), nil
}
