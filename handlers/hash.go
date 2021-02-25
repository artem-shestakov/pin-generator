package handlers

import (
	"log"
	"net/http"
	"pin-salt-hash/models"
	"strconv"
)

// HashMux handler for Mux Server
type HashMux struct {
	l *log.Logger
}

func (h *HashMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		// Getting param from request
		query := r.URL.Query()
		strong := query.Get("strong")
		pinLen, err := strconv.ParseInt(query.Get("pin_len"), 10, 8)
		if err != nil {
			http.Error(w, "Missing or incorrect parameter \"pin_len\"", http.StatusBadRequest)
			return
		}
		saltLen, err := strconv.ParseInt(query.Get("salt_len"), 10, 8)
		if saltLen == 0 || err != nil {
			saltLen = 10
		}

		answer := models.Answer{}

		if strong == "true" || strong == "True" {
			answer.Generate(int(pinLen), int(saltLen), true, 2)
		} else {
			answer.Generate(int(pinLen), int(saltLen), false)
		}

		err = answer.ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// NewHashMux create handler for Mux Server
func NewHashMux(l *log.Logger) *HashMux {
	return &HashMux{l}
}
