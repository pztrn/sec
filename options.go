package sec

// Options represents configuration for SEC. Note that this is parser
// configuration, per-field configuration should be defined in tags.
type Options struct {
	// ErrorsAreCritical indicates that on every parsing error we should
	// stop processing things. By default we will proceed even if errors
	// will occur except critical errors like passing invalid value to
	// SEC_DEBUG environment variable and passing not a pointer to
	// structure to Parse() function.
	ErrorsAreCritical bool
}

var (
	defaultOptions = &Options{
		ErrorsAreCritical: false,
	}
)
