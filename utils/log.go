package utils

import (
	"fmt"
	"log"
	"os"
)

// Log fonction
func Logger() error {
	// Ouverture du fichier
	f, err := os.Open("log")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Fermeture du fichier
	defer f.Close()

	logger := log.New(f, "error: ", log.Lmsgprefix)
	logger.Println(err.Error())
	return err
}
