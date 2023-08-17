package utils

import (
	"fmt"
	"log"
	"os"
)

// Log fonction
func Logger(e error) error {
	// Ouverture du fichier
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Fermeture du fichier
	defer f.Close()

	logger := log.New(f, "error: ", log.Lmsgprefix)

	logger.Println(e.Error())
	return e
}