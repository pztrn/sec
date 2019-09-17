package sec

import (
	// stdlib
	"errors"
	"os"
	"reflect"
	"strconv"
)

var (
	errNotBool = errors.New("environment variable doesn't contain boolean")

	errNotFloat   = errors.New("environment variable doesn't contain floating point number")
	errNotFloat32 = errors.New("environment variable doesn't contain float32")
	errNotFloat64 = errors.New("environment variable doesn't contain float64")

	errNotInt   = errors.New("environment variable doesn't contain integer")
	errNotInt8  = errors.New("environment variable doesn't contain int8")
	errNotInt16 = errors.New("environment variable doesn't contain int16")
	errNotInt32 = errors.New("environment variable doesn't contain int32")
	errNotInt64 = errors.New("environment variable doesn't contain int64")

	errNotUint   = errors.New("environment variable doesn't contain unsigned integer")
	errNotUint8  = errors.New("environment variable doesn't contain uint8")
	errNotUint16 = errors.New("environment variable doesn't contain uint16")
	errNotUint32 = errors.New("environment variable doesn't contain uint32")
	errNotUint64 = errors.New("environment variable doesn't contain uint64")
)

// Parses environment for data.
func parseEnv() error {
	printDebug("Starting parsing data into tree from environment variables...")

	for _, element := range parsedTree {
		printDebug("Processing element '%s'", element.EnvVar)
		data, found := os.LookupEnv(element.EnvVar)
		if !found {
			printDebug("Value for '%s' environment variable wasn't found", element.EnvVar)
			continue
		} else {
			printDebug("Value for '%s' will be: %s", element.EnvVar, data)
		}

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
		}
	}

	return nil
}
