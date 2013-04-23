// Package ini implements access to ini configurations.
//
// NOTE: This project is in a *pre-alpha* stage. It is mainly intended to
// speculate about API design. The API may change significantly!
package ini

import (
	"errors"
	"io"
	"log"
)

// Conf is a configuration containing keys: name and value pairs. Key are
// grouped into sections, either explicitly with a section name or implicitly to
// the default section.
type Conf struct {
	// Sect is the default section.  Keys in the default section can be
	// implicitly accessed and are available directly through conf.
	*Sect
}

// New returns a blank ini configuration.
func New() (conf *Conf) {
	log.Fatalln("ini.New: not yet implemented.")
	return nil
}

// Parse reads from r and returns a parsed ini configuration.
func Parse(r io.Reader) (conf *Conf, err error) {
	return nil, errors.New("ini.Parse: not yet implemented.")
}

// ReadFile reads from the file specified by filePath and returns a parsed ini
// configuration.
func ReadFile(filePath string) (conf *Conf, err error) {
	return nil, errors.New("ini.ReadFile: not yet implemented.")
}

// Write writes the ini configuration to w.
func (conf *Conf) Write(w io.Writer) (err error) {
	return errors.New("Conf.Write: not yet implemented.")
}

// WriteFile writes the ini configuration to a file specified by filePath. If
// the file doesn't exist, WriteFile creates it, otherwise WriteFile truncates
// it before writing.
func (conf *Conf) WriteFile(filePath string) (err error) {
	return errors.New("Conf.WriteFile: not yet implemented.")
}
