package sec

import (
	// stdlib
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	// Errors.
	errNotPTR       = errors.New("passed data is not a pointer")
	errNotStructure = errors.New("passed data is not a structure")

	// Debug flag.
	debugFlagEnvName = "SEC_DEBUG"
	debug            bool

	// Parsed structure fields.
	parsedTree []*field

	// Options for current run.
	options *Options
)

// Parse parses environment variables into passed structure.
func Parse(structure interface{}, config *Options) error {
	parsedTree = []*field{}

	options = config
	if config == nil {
		options = defaultOptions
	}

	// Set debug flag if defined in environment.
	debugFlagRaw, found := os.LookupEnv(debugFlagEnvName)
	if found {
		var err error
		debug, err = strconv.ParseBool(debugFlagRaw)
		if err != nil {
			log.Println("Invalid '" + debugFlagEnvName + "' environment variable data: '" + debugFlagRaw + "'. Error: " + err.Error())
			if options.ErrorsAreCritical {
				return err
			}
		} else {
			printDebug("Debug mode activated")
		}
	}

	printDebug("Parsing started with configuration: %+v", options)

	value := reflect.ValueOf(structure)

	// Figure out passed data type. We should accept ONLY pointers
	// to structure.
	printDebug("Passed structure kind: %s, want: %s", value.Type().Kind().String(), reflect.Ptr.String())

	// If passed data isn't a pointer - return error in any case because
	// we can't support anything except pointer.
	if value.Type().Kind() != reflect.Ptr {
		return errNotPTR
	}

	printDebug("Passed data kind: %s, want: %s", value.Elem().Type().Kind().String(), reflect.Struct.String())

	value = value.Elem()

	// Passed data should be a pointer to structure. Otherwise we should
	// return error in any case.
	if value.Type().Kind() != reflect.Struct {
		return errNotStructure
	}

	// Parse structure.
	composeTree(value, strings.ToUpper(value.Type().Name()))

	return parseEnv()
}

// Produces debug output into stdout using standard log module if debug
// mode was activated by setting SEC_DEBUG environment variable to true.
func printDebug(text string, params ...interface{}) {
	if debug {
		if len(params) == 0 {
			log.Println(text)
		} else {
			log.Printf(text, params...)
		}
	}
}
