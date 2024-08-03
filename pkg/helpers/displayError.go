package helpers

import "log"

func FatalError(mes string, err error) {
	log.Fatal(mes, err)
}

func PrintError(mes string, err error) {
	log.Println(mes, err)
}

func PanicError(mes string, err error) {
	log.Panic(mes, err)
}
