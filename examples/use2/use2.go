package main

import (
	"errors"
	"log"

	ini "github.com/mewkiz/ini-prealpha"
)

func main() {
	err := use2()
	if err != nil {
		log.Fatalln(err)
	}
}

func use2() (err error) {
	// Parse ini file.
	conf, err := ini.ReadFile("use2.ini")
	if err != nil {
		return err
	}

	// Get user id.
	sect, ok := conf.Section("user")
	if !ok {
		return errors.New(`unable to locate "user" section.`)
	}
	id, ok := sect.GetInt("id")
	if !ok {
		return errors.New(`unable to locate "id" key in "user" section.`)
	}

	// Get email based on user id.
	email, ok := GetEmail(id)

	if ok {
		// Add a "contact" section to the ini configuration.
		sect, err := conf.AddSection("contact")
		if err != nil {
			return err
		}

		// Add the email to the "email" key of the "contact" section.
		err = sect.SetKey("email", email)
		if err != nil {
			return err
		}

		// Write the updated ini configuration to the "use2.ini" file.
		err = conf.WriteFile("use2.ini")
		if err != nil {
			return err
		}
	}

	return nil
}

func GetEmail(id int) (email string, ok bool) {
	// do stuff...
	return "user@example.org", true
}
