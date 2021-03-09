package apiserver

import (
	"net/http"
	"strconv"

	"github.com/artem-shestakov/pin-generator/internal/app/answer"
)

func (s *APIServer) handlePin() http.Handler {

	answer := answer.NewAnswer()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params from request
		query := r.URL.Query()
		strong := query.Get("strong")
		pinLen, err := strconv.ParseInt(query.Get("pin_len"), 10, 8)
		if err != nil {
			http.Error(w, "Missing or incorrect parameter \"pin_len\"", http.StatusBadRequest)
			s.logger.Error(r.Method, " ", r.RemoteAddr, " ", r.RequestURI, " [ERROR] Missing or incorrect parameter \"pin_len\" ", http.StatusBadRequest)
			return
		}
		saltLen, err := strconv.ParseInt(query.Get("salt_len"), 10, 8)
		if err != nil || saltLen == 0 {
			saltLen = 10
			s.logger.Debug(r.Method, " ", r.RemoteAddr, " ", r.URL, " [DEBUG] Can't get parameter \"salt_len\". Parameter is ", saltLen)
		}

		if strong == "true" || strong == "True" {
			answer.Generate(int(pinLen), int(saltLen), true, 2)
		} else {
			answer.Generate(int(pinLen), int(saltLen), false)
		}

		// Decode to JSON and send response
		err = answer.ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
			s.logger.Error(r.Method, " ", r.RemoteAddr, " ", r.RequestURI, " [ERROR] Unable to marshal JSON ", http.StatusInternalServerError)
		}
		s.logger.Info(r.Method, " ", r.RemoteAddr, " ", r.RequestURI, " [INFO] Request successfully processed ", http.StatusOK)
	})
}
