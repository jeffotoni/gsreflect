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
	"strconv"
)

// parse in the struct that was declared, it
// will get the tag name that it finds
// in the declaration of its struct
// ex: Name string `tagname: my default content here`
func Default(str interface{}) (err error) {

	// getting the
	// memory location
	// of your struct
	rt := reflect.TypeOf(str)

	// Check if it's a pointer
	if rt.Kind() != reflect.Ptr {

		err = errors.New("It's not a pointer!")
		return
	}

	elField := rt.Elem()

	// Check if it's a struct
	if elField.Kind() != reflect.Struct {
		err = errors.New("it's not a struct!")
		return
	}

	// ref value in memory
	// refValue := reflect.ValueOf(str).Elem()

	//
	// Spanning the entire struct
	// Iterate over all available fields
	// and read the tag value
	for i := 0; i < elField.NumField(); i++ {

		// Get the type and
		// kind of our struct variable
		field := elField.Field(i)
		// value := refValue.Field(i)
		kind := field.Type.Kind()
		tagVal := field.Tag.Get(tagName)

		// converting
		// some types
		switch kind {

		case reflect.String:
			reflect.ValueOf(str).Elem().Field(i).SetString(tagVal)
			break

		case reflect.Int:
			tagValint, _ := strconv.ParseInt(tagVal, 10, 0)
			reflect.ValueOf(str).Elem().Field(i).SetInt(tagValint)
			break
		case reflect.Int32:
			tagValint, _ := strconv.ParseInt(tagVal, 10, 0)
			reflect.ValueOf(str).Elem().Field(i).SetInt(tagValint)
			break
		case reflect.Int64:
			tagValint, _ := strconv.ParseInt(tagVal, 10, 0)
			reflect.ValueOf(str).Elem().Field(i).SetInt(int64(tagValint))
			break

		case reflect.Float32:
			tagValf, _ := strconv.ParseFloat(tagVal, 32)
			reflect.ValueOf(str).Elem().Field(i).SetFloat(tagValf)
			break

		case reflect.Float64:
			tagValf, _ := strconv.ParseFloat(tagVal, 64)
			reflect.ValueOf(str).Elem().Field(i).SetFloat(tagValf)
			break

		case reflect.Bool:

			tagBool, _ := strconv.ParseBool(tagVal)
			reflect.ValueOf(str).Elem().Field(i).SetBool(tagBool)
			break

		default:
			err = errors.New("Unknown type")
		}
	}

	return
}
