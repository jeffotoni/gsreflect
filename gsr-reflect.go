/*
* Golang reflect update struct to tag
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package gsr

import (
	"errors"
	"reflect"
)

// const tag
const tagName = "default"

func Default(s interface{}) (err error) {

	st := reflect.TypeOf(s)

	if st.Kind() != reflect.Ptr {

		err = errors.New("Not a pointer")
		return
	}

	refField := st.Elem()
	if refField.Kind() != reflect.Struct {
		err = errors.New("Not a struct")
		return
	}

	//refValue := reflect.ValueOf(s).Elem()
	for i := 0; i < refField.NumField(); i++ {

		field := refField.Field(i)
		// value := refValue.Field(i)
		// kind := field.Type.Kind()
		tagVal := field.Tag.Get(tagName)

		reflect.ValueOf(s).Elem().Field(i).SetString(tagVal)
	}

	return
}
