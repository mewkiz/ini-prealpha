package ini

import (
	"log"
)

// GetBool returns the boolean value with the specified key name in sect. The
// value of ok is true if the key is present in sect, and false otherwise.
func (sect *Sect) GetBool(name string) (v bool, ok bool) {
	log.Fatalln("Sect.GetBool: not yet implemented.")
	return false, false
}

// GetString returns the string with the specified key name in sect. The value
// of ok is true if the key is present in sect, and false otherwise.
func (sect *Sect) GetString(name string) (v string, ok bool) {
	log.Fatalln("Sect.GetString: not yet implemented.")
	return "", false
}

// GetInt returns the integer value with the specified key name in sect. The
// value of ok is true if the key is present in sect, and false otherwise.
func (sect *Sect) GetInt(name string) (v int, ok bool) {
	log.Fatalln("Sect.GetInt: not yet implemented.")
	return 0, false
}

// GetUint returns the unsigned integer value with the specified key name in
// sect. The value of ok is true if the key is present in sect, and false
// otherwise.
func (sect *Sect) GetUint(name string) (v uint, ok bool) {
	log.Fatalln("Sect.GetUint: not yet implemented.")
	return 0, false
}

// GetFloat64 returns the float64 value with the specified key name in sect. The
// value of ok is true if the key is present in sect, and false otherwise.
func (sect *Sect) GetFloat64(name string) (v float64, ok bool) {
	log.Fatalln("Sect.GetFloat64: not yet implemented.")
	return 0.0, false
}
