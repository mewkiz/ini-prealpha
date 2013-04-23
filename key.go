package ini

import (
	"errors"
	"log"
)

// SetKey sets the value of the key with the specified key name in sect to v.
// The key is created if it isn't already present in sect.
func (sect *Sect) SetKey(name string, v interface{}) (err error) {
	return errors.New("Sect.SetKey: not yet implemented.")
}

// DelKey deletes the key with the specified key name from sect.
func (sect *Sect) DelKey(name string) {
	log.Fatalln("Sect.DelKey: not yet implemented.")
}

// HasKey returns true if the key with the specified key name is present in
// sect, and false otherwise.
func (sect *Sect) HasKey(name string) (ok bool) {
	log.Fatalln("Sect.HasKey: not yet implemented.")
	return false
}
