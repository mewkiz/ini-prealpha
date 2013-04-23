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
	v, ok := conf.GetInt("implicit")
	if ok {
		fmt.Println("implicit:", v)
	}

	// Keys from other sections must be explicitly accessed and are available
	// from their sect.
	sect, ok := conf.Section("basic")
	if ok {
		a, ok := sect.GetBool("bool")
		if ok {
			fmt.Println("basic.bool:", a)
		}

		b, ok := sect.GetString("string")
		if ok {
			fmt.Println("basic.string:", b)
		}

		c, ok := sect.GetInt("int")
		if ok {
			fmt.Println("basic.int:", c)
		}

		d, ok := sect.GetUint("uint")
		if ok {
			fmt.Println("basic.uint:", d)
		}

		e, ok := sect.GetFloat64("float")
		if ok {
			fmt.Println("basic.float:", e)
		}
	}

	return nil
}
