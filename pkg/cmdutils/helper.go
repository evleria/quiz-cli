package cmdutils

import (
	"fmt"
	"io"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ClearConsole(writer io.Writer) error {
	_, err := fmt.Fprint(writer, "\033[H\033[2J")
	return err
}
