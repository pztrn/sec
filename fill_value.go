package sec

import (
	// stdlib
	"errors"
	"reflect"
	"strconv"
)

func fillValue(element *field, data string) error {
	switch element.Kind {
	case reflect.String:
		element.Pointer.SetString(data)
	case reflect.Bool:
		val, err := strconv.ParseBool(data)
		if err != nil {
			printDebug("Error occurred while parsing boolean: %s", err.Error())
			if options.ErrorsAreCritical {
				return errNotBool
			}
		}
		element.Pointer.SetBool(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Bitsize 64 here specified for a reason - actual ints
		// ranges checking goes below and we should expect it to
		// be 0 in case of configuration.
		val, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			printDebug("Error occurred while parsing int: %s", err.Error())
			if options.ErrorsAreCritical {
				return errNotInt
			}
		}
		switch element.Kind {
		case reflect.Int8:
			// int8 is an integer in [-128...127] range.
			if val >= -128 && val <= 127 {
				element.Pointer.SetInt(val)
			} else {
				printDebug("Data in environment variable '%s' isn't int8", element.EnvVar)
				element.Pointer.SetInt(0)
				if options.ErrorsAreCritical {
					return errNotInt8
				}
			}
		case reflect.Int16:
			// int16 is an integer in [-32768...32767] range.
			if val >= -32768 && val <= 32767 {
				element.Pointer.SetInt(val)
			} else {
				printDebug("Data in environment variable '%s' isn't int16", element.EnvVar)
				element.Pointer.SetInt(0)
				if options.ErrorsAreCritical {
					return errNotInt16
				}
			}
		case reflect.Int32:
			// int32 is an integer in [-2147483648...2147483647] range.
			if val >= -2147483648 && val <= 2147483647 {
				element.Pointer.SetInt(val)
			} else {
				printDebug("Data in environment variable '%s' isn't int32", element.EnvVar)
				element.Pointer.SetInt(0)
				if options.ErrorsAreCritical {
					return errNotInt32
				}
			}
		case reflect.Int64, reflect.Int:
			// int64 is an integer in [-9223372036854775808...9223372036854775807] range.
			// This is currently maximum allowed int values, so we'll
			// just set it.
			element.Pointer.SetInt(val)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(data, 10, 64)
		if err != nil {
			printDebug("Error occurred while parsing unsigned integer: %s", err.Error())
			if options.ErrorsAreCritical {
				return errNotUint
			}
		}
		switch element.Kind {
		case reflect.Uint8:
			// uint8 is an integer in [0...255] range.
			if val <= 255 {
				element.Pointer.SetUint(val)
			} else {
				printDebug("Data in environment variable '%s' isn't uint8", element.EnvVar)
				element.Pointer.SetUint(0)
				if options.ErrorsAreCritical {
					return errNotUint8
				}
			}
		case reflect.Uint16:
			// uint16 is an integer in [0...65535] range.
			if val <= 65535 {
				element.Pointer.SetUint(val)
			} else {
				printDebug("Data in environment variable '%s' isn't uint16", element.EnvVar)
				element.Pointer.SetUint(0)
				if options.ErrorsAreCritical {
					return errNotUint16
				}
			}
		case reflect.Uint32:
			// uint32 is an integer in [0...4294967295] range.
			if val <= 4294967295 {
				element.Pointer.SetUint(val)
			} else {
				printDebug("Data in environment variable '%s' isn't uint32", element.EnvVar)
				element.Pointer.SetUint(0)
				if options.ErrorsAreCritical {
					return errNotUint32
				}
			}
		case reflect.Uint64:
			// uint64 is an integer in [0...18446744073709551615] range.
			// This is currently maximum allowed int values, so we'll
			// just set it.
			element.Pointer.SetUint(val)
		}
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(data, 64)
		if err != nil {
			printDebug("Error occurred while parsing float: %s", err.Error())
			if options.ErrorsAreCritical {
				return errNotFloat
			}
		}
		element.Pointer.SetFloat(val)
	case reflect.Interface:
		// We should not attempt to work with data in interface{}
		// unless it is a pointer to value.
		if element.Pointer.Elem().Kind() != reflect.Ptr {
			printDebug("Element for environment variable '%s' isn't a pointer and put into interface{}. Nothing will be done with this element.", element.EnvVar)
			if options.ErrorsAreCritical {
				return errors.New("element for environment variable '" + element.EnvVar + "' isn't a pointer and put into interface")
			}
			return nil
		}

		// We should get actual value. Two Elem()'s for that.
		// It goes interface{} -> ptr -> real element.
		element.Pointer = element.Pointer.Elem().Elem()
		element.Kind = element.Pointer.Kind()
		return fillValue(element, data)
	}
	return nil
}
