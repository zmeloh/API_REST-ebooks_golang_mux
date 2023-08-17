package rest

import (
	"example/api/utils"
	"net/http"
)

// HandleError gère les erreurs et renvoie une réponse HTTP appropriée.
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		utils.Logger(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
