package utils

import (
	"log"
)

func Background(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		fn()
	}()
}
