package sec

import (
	"reflect"
)

// This structure represents every parsable field that was found while
// reading passed structure.
type field struct {
	// Name is a field name. Mostly for debugging purpose.
	Name string
	// Pointer is a pointer to field wrapped in reflect.Value.
	Pointer reflect.Value
	// EnvVar is a name of environment variable we will try to read.
	EnvVar string
	// Kind is a reflect.Kind value.
	Kind reflect.Kind

	// Next variables are tag-related.
	// optional bool
}
