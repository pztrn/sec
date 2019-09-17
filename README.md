# SEC. Take a break and configure your application with ease.

**SEC** stands for "Simple Environment Configuration" and provides really simple way to configure your application.

After googling around Go applications configuration management packages that able to take parameters from environment configuration I came to a conclusion that there is none packages that able to do everything I want and yet have readable and testable source code.

Key intentions to create SEC:

* Parse configuration into structure with support of infinitely nested structures.
* Works properly with interfaces.
* No goto's.
* 100% code coverage.
* No external dependencies (only testify for tests).
* Readable code and proper variables naming.
* Debug mode

This list might be updated if new key intention arrives :).

SEC was written under impression from https://github.com/vrischmann/envconfig/.

## Installation

Go modules and dep are supported. Other package managers might or might not work, MRs are welcome!

## Usage

SEC is designed to be easy to use parser, so there is only one requirement - passed data should be a pointer to structure. You cannot do something like:

```go
var Data string
sec.Parse(&Data, nil)
```
 
This will throw errors, as any type you'll pass, except for pointer to structure.

SEC is unable to parse embedded unexported things except structures due to inability to get embedded field's address. Embed only structures, please.

The very valid way to use SEC:

```go
type config struct{
    Database struct{
        URI string
        Options string
    }
    HTTPTimeout int
}

cfg := &config{}

err := Parse(cfg, nil)
if err != nil {
    log.Fatal(err)
}
```

No field tags supported yet, this in ToDo.

### Debug

To get additional debug output set ``SEC_DEBUG`` environment variable to ``true``. If invalid boolean value will be passed it'll output error about that.
 
 Debug output uses standart log package. This may change in future.
