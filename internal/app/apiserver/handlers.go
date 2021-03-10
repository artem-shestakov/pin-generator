package apiserver

import (
	"net/http"
	"strconv"

	"github.com/artem-shestakov/pin-generator/internal/app/answer"
	"github.com/go-openapi/runtime/middleware"
)

// swagger:parameters pin_code
type pinLen struct {
	// The length of pin code
	//
	// in: query
	// required: true
	// example: 10
	PinLen string `json:"pin_len"`
}

// swagger:parameters pin_code
type saltLen struct {
	// The length of salt.
	// default: 10
	//
	// in: query
	// required: false
	// example: 15
	SaltLen string `json:"salt_len"`
}

// swagger:parameters pin_code
type strong struct {
	// Add letter in pin code.
	// default: false
	//
	// in: query
	// required: false
	// example: true
	Stong bool `json:"strong"`
}

// swagger:route GET /pin pin_code pin_code
// Returns a pin code, salt and SHA-1 hash
// Responses:
// 	200: answerResponse
// 	400: badRequest
// 	405: notAllowed
// 	500: intError

// handlePin returns a pin code, salt and SHA-1 hash
func (s *APIServer) handlePin() http.Handler {

	answer := answer.NewAnswer()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params from request
		query := r.URL.Query()
		strong := query.Get("strong")
		pinLen, err := strconv.ParseInt(query.Get("pin_len"), 10, 8)
		if err != nil {
			http.Error(w, "Missing or incorrect parameter \"pin_len\"", http.StatusBadRequest)
			return
		}
		saltLen, err := strconv.ParseInt(query.Get("salt_len"), 10, 8)
		if err != nil || saltLen == 0 {
			saltLen = 10
			s.logger.Debugln(r.Method, r.RemoteAddr, " ", r.URL, " [DEBUG] Can't get parameter \"salt_len\". Parameter is ", saltLen)
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
		}
	})
}

// Redoc htt.Handler for API documentation
func (s *APIServer) Redoc() http.Handler {
	handlerOptions := middleware.RedocOpts{SpecURL: "./api/swagger.yml"}
	return middleware.Redoc(handlerOptions, nil)
}

// Middleware handler for logging
type (
	// struct for holding response details
	responseData struct {
		status int
		size   int
	}

	// http.ResponseWriter implementation
	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

// write response and capture size
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

// write status code and capture it
func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

// WithLogging wrapper to logging requests
func (s *APIServer) WithLogging(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, r *http.Request) {
		responseData := &responseData{
			status: 200,
			size:   0,
		}
		loginRW := loggingResponseWriter{
			ResponseWriter: rw,
			responseData:   responseData,
		}
		// serve original handler with new http.ResponseWriter
		h.ServeHTTP(&loginRW, r)

		if responseData.status < 400 {
			s.logger.Infoln(r.RequestURI, r.Method, r.RemoteAddr, loginRW.responseData.status, responseData.size)
		} else {
			s.logger.Errorln(r.RequestURI, r.Method, r.RemoteAddr, loginRW.responseData.status, responseData.size)
		}
	}
	return http.HandlerFunc(loggingFn)
}
