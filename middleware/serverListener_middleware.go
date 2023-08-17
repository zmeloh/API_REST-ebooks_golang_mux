package middleware

import (
	"example/api/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Créez un ResponseRecorder pour capturer le statut de la réponse
		recorder := NewStatusRecorder(w)

		// Appelez la prochaine étape de la chaîne
		next.ServeHTTP(recorder, r)

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		log.Printf(
			"Request: %s %s %s | Status: %d | Latency: %v\n",
			r.Method, r.URL.Path, r.RemoteAddr,
			recorder.Status, latency,
		)
		msg := fmt.Errorf(
			"Request: %s %s %s | Status: %d | Latency: %v\n",
			r.Method, r.URL.Path, r.RemoteAddr,
			recorder.Status, latency,
		)
		utils.Logger(msg)
	})
}

// Définissez un ResponseRecorder personnalisé pour capturer le statut de la réponse
type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func NewStatusRecorder(w http.ResponseWriter) *statusRecorder {
	return &statusRecorder{w, http.StatusOK}
}

func (recorder *statusRecorder) WriteHeader(code int) {
	recorder.Status = code
	recorder.ResponseWriter.WriteHeader(code)
}
