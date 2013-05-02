package ini

import (
	"fmt"
	"log"
)

// GetBool returns the boolean value, with the specified key name in sect. The
// value of err is nil if the key is present in sect, and not nil otherwise.
func (sect *Sect) GetBool(name string) (v bool, err error) {
	log.Fatalln("Sect.GetBool: not yet implemented.")
	return false, fmt.Errorf("Sect.GetBool: unable to locate key %q in section %q.", name, sect)
}

// GetString returns the string, with the specified key name in sect. The value
// of err is nil if the key is present in sect, and not nil otherwise.
func (sect *Sect) GetString(name string) (v string, err error) {
	log.Fatalln("Sect.GetString: not yet implemented.")
	return "", fmt.Errorf("Sect.GetString: unable to locate key %q in section %q.", name, sect)
}

// GetInt returns the integer value, with the specified key name in sect. The
// value of err is nil if the key is present in sect, and not nil otherwise.
func (sect *Sect) GetInt(name string) (v int, err error) {
	log.Fatalln("Sect.GetInt: not yet implemented.")
	return 0, fmt.Errorf("Sect.GetInt: unable to locate key %q in section %q.", name, sect)
}

// GetFloat64 returns the float64 value, with the specified key name in sect. The
// value of err is nil if the key is present in sect, and not nil otherwise.
func (sect *Sect) GetFloat64(name string) (v float64, err error) {
	log.Fatalln("Sect.GetFloat64: not yet implemented.")
	return 0.0, fmt.Errorf("Sect.GetFloat64: unable to locate key %q in section %q.", name, sect)
}
