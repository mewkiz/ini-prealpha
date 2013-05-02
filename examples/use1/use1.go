package main

import (
	"fmt"
	"log"

	ini "github.com/mewkiz/ini-prealpha"
)

func main() {
	err := use1()
	if err != nil {
		log.Fatalln(err)
	}
}

func use1() (err error) {
	// Parse ini file.
	conf, err := ini.ReadFile("use1.ini")
	if err != nil {
		return err
	}

	// Keys in the default section can be implicitly accessed and are available
	// directly through conf.
	v, err := conf.GetInt("implicit")
	if err == nil {
		fmt.Println("implicit:", v)
	}

	// Keys from other sections must be explicitly accessed and are available
	// from their sect.
	sect, err := conf.Section("basic")
	if err == nil {
		a, err := sect.GetBool("bool")
		if err == nil {
			fmt.Println("basic.bool:", a)
		}

		b, err := sect.GetString("string")
		if err == nil {
			fmt.Println("basic.string:", b)
		}

		c, err := sect.GetInt("int")
		if err == nil {
			fmt.Println("basic.int:", c)
		}

		e, err := sect.GetFloat64("float")
		if err == nil {
			fmt.Println("basic.float:", e)
		}
	}

	return nil
}
