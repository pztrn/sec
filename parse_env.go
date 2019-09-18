package sec

import (
	// stdlib
	"errors"
	"os"
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

		err := fillValue(element, data)
		if err != nil {
			return err
		}
	}

	return nil
}
