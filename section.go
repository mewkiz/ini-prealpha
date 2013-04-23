package ini

import (
	"errors"
	"fmt"
	"log"
)

// Sect is a section which serves as a logic unit to group keys together. There
// exist two categories of sections, explicit and implicit ones.
//
// Sections with section names are explicit sections. The only implicit section
// is the default section and it doesn't have a section name.
type Sect struct {
}

func (sect *Sect) String() string {
	log.Fatalln("Sect.String: not yet implemented.")
	return ""
}

// Section returns the section, with the specified section name in conf. The
// value of err is nil if the section is present in conf, and not nil otherwise.
func (conf *Conf) Section(name string) (sect *Sect, err error) {
	log.Fatalln("Conf.Section: not yet implemented.")
	return nil, fmt.Errorf("Conf.Section: unable to locate section %q.", name)
}

// AddSection adds a new section, with the specified section name to conf.
func (conf *Conf) AddSection(name string) (sect *Sect, err error) {
	return nil, errors.New("Conf.AddSection: not yet implemented.")
}

// DelSection deletes the section, with the specified section name from conf.
func (conf *Conf) DelSection(name string) {
	log.Fatalln("Conf.DelSection: not yet implemented.")
}

// HasSection returns true if the section, with the specified section name, is
// present in conf, and false otherwise.
func (conf *Conf) HasSection(name string) (ok bool) {
	log.Fatalln("Conf.HasSection: not yet implemented.")
	return false
}

// Sections returns a slice containing the sections in conf.
func (conf *Conf) Sections() (sections []*Sect) {
	log.Fatalln("Conf.Sections: not yet implemented.")
	return nil
}

// SectionNames returns a slice containing the section names in conf.
func (conf *Conf) SectionNames() (sectionNames []string) {
	log.Fatalln("Conf.SectionNames: not yet implemented.")
	return nil
}
