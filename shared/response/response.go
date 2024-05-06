package response

import (
	"encoding/json"
	"fmt"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Base struct {
	Data    *any    `json:"data,omitempty"`
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// WithMessage sends a response with a simple text message
func WithMessage(w http.ResponseWriter, code int, message string) {
	respond(w, code, Base{Message: &message})
}

// WithJSON sends a response containing a JSON object
func WithJSON(w http.ResponseWriter, code int, jsonPayload any) {
	respond(w, code, Base{Data: &jsonPayload})
}

// WithError sends a response with an error message
func WithError(w http.ResponseWriter, err error) {
	code := failure.GetCode(err)
	errMsg := err.Error()
	respond(w, code, Base{Error: &errMsg})
}

func (e *Failure) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(e.Code), e.Message)
}

func respond(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed Error!")
	}
}
