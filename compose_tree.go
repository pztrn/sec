package sec

import (
	"reflect"
	"strings"
)

// Composes full tree for every structure member.
func composeTree(value reflect.Value, prefix string) {
	typeOf := value.Type()

	// Compose prefix for everything below current field.
	var curPrefix string
	if prefix != "" {
		curPrefix = prefix
		if !strings.HasSuffix(curPrefix, "_") {
			curPrefix += "_"
		}
	}

	// Resolve Ptrs if any.
	for {
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
			typeOf = value.Type()
		} else {
			break
		}
	}

	if value.Kind() != reflect.Struct {
		if value.Kind() == reflect.Map {
			newElementPrefix := curPrefix

			mapIter := value.MapRange()
			for mapIter.Next() {
				composeTree(mapIter.Value().Elem(), newElementPrefix+"_"+strings.ToUpper(mapIter.Key().String()))
			}
		} else {
			f := &field{
				Name:    typeOf.Name(),
				EnvVar:  curPrefix + strings.ToUpper(typeOf.Name()),
				Pointer: value,
				Kind:    value.Kind(),
			}

			parsedTree = append(parsedTree, f)

			printDebug("Field data constructed (start): %+v", f)
		}

		return
	}

	for i := 0; i < value.NumField(); i++ {
		fieldToProcess := value.Field(i)
		fieldToProcessType := typeOf.Field(i)

		// If currently processed field - interface, then we should
		// get underlying value.
		if fieldToProcess.Kind() == reflect.Interface {
			if fieldToProcess.Elem().Kind() == reflect.Ptr {
				fieldToProcess = fieldToProcess.Elem().Elem()
			} else if fieldToProcess.Elem().Kind() == reflect.Struct {
				fieldToProcess = fieldToProcess.Elem()
			}
		}

		// In 99% of cases we will get uninitialized things we should
		// initialize.
		switch fieldToProcess.Kind() {
		case reflect.Ptr, reflect.Map, reflect.Slice:
			if fieldToProcess.IsNil() {
				printDebug("Field '%s' is nil, initializing new one", fieldToProcessType.Name)

				// We should use only exported fields as unexported aren't
				// settable using 'reflect' package. Can be possibly solved
				// using unsafe pointers?
				if fieldToProcess.CanSet() {
					fieldToProcess.Set(reflect.New(fieldToProcess.Type().Elem()))
				} else {
					printDebug("Field '%s' is unexported and will be ignored", fieldToProcessType.Name)

					continue
				}

				fieldToProcess = fieldToProcess.Elem()
			}
		}

		printDebug("Field: '%s', type: %s (anonymous or embedded: %t)",
			fieldToProcessType.Name,
			fieldToProcess.Type().Kind().String(),
			fieldToProcessType.Anonymous,
		)

		// Dealing with embedded things.
		if fieldToProcessType.Anonymous {
			// We should not allow anything other than struct.
			if fieldToProcess.Kind() != reflect.Struct {
				printDebug("Field is embedded, but not a struct (%s), which cannot be used", fieldToProcess.Kind().String())

				continue
			}
		}

		if fieldToProcess.Kind() != reflect.Struct && !fieldToProcess.CanSet() {
			printDebug("Field '%s' of type '%s' can't be set, skipping",
				fieldToProcessType.Name,
				fieldToProcess.Type().Kind().String())

			continue
		}

		printDebug("All underlying elements will have prefix '%s'", curPrefix)

		// Hello, I'm recursion and I'm here to make you happy.
		// I'll be launched only for structures to get their fields.
		switch fieldToProcess.Kind() {
		case reflect.Struct:
			newElementPrefix := curPrefix
			if !fieldToProcessType.Anonymous {
				newElementPrefix = strings.ToUpper(newElementPrefix + typeOf.Field(i).Name)
			}

			composeTree(fieldToProcess, newElementPrefix)
		case reflect.Map:
			newElementPrefix := curPrefix
			if !fieldToProcessType.Anonymous {
				newElementPrefix = strings.ToUpper(newElementPrefix + typeOf.Field(i).Name)
			}

			mapIter := fieldToProcess.MapRange()
			for mapIter.Next() {
				composeTree(mapIter.Value().Elem(), newElementPrefix+"_"+strings.ToUpper(mapIter.Key().String()))
			}
		default:
			f := &field{
				Name:    typeOf.Field(i).Name,
				EnvVar:  curPrefix + strings.ToUpper(typeOf.Field(i).Name),
				Pointer: fieldToProcess,
				Kind:    fieldToProcess.Kind(),
			}

			parsedTree = append(parsedTree, f)

			printDebug("Field data constructed (end): %+v", f)
		}
	}
}
