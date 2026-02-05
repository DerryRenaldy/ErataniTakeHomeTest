package response

import (
	"encoding/json"
	"net/http"
)

type Base struct {
	Data     *interface{} `json:"data,omitempty"`
	Metadata *interface{} `json:"metadata,omitempty"`
	Error    *string      `json:"error,omitempty"`
	Message  *string      `json:"message,omitempty"`
}

func NoContent(w http.ResponseWriter) {
	respond(w, http.StatusNoContent, nil)
}

func WithMessage(w http.ResponseWriter, code int, message string) {
	msg := message
	respond(w, code, Base{Message: &msg})
}

func WithJSON(w http.ResponseWriter, code int, jsonPayload interface{}) {
	respond(w, code, Base{Data: &jsonPayload})
}

func WithMetadata(w http.ResponseWriter, code int, jsonPayload interface{}, metadata interface{}) {
	respond(w, code, Base{Data: &jsonPayload, Metadata: &metadata})
}

func WithError(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	errMsg := err.Error()
	respond(w, code, Base{Error: &errMsg})
}

func WithPreparingShutdown(w http.ResponseWriter) {
	code := http.StatusServiceUnavailable
	msg := "Server is preparing to shutdown, please try again later"
	respond(w, code, Base{Error: &msg})
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
