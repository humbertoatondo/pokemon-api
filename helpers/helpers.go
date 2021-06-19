package helpers

import (
	"encoding/json"
	"net/http"
)

// LanguageMap is used to reference the available languages
// for pokemon moves.
var LanguageMap = map[string]int{
	"ja-Hrkt": 0,
	"ko":      1,
	"zh-Hant": 2,
	"fr":      3,
	"de":      4,
	"es":      5,
	"it":      6,
	"en":      7,
	"ja":      8,
	"zh-Hans": 9,
}

// ParseKeyFromURL retreives the value for a parameter in the url.
func ParseKeyFromURL(key string, r *http.Request) (string, bool) {
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys) < 1 {
		return "", false
	}
	return keys[0], true
}

// RespondWithJSON sets and writes the headers and the response for the http request.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError uses RespondWithJSON to write an error.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	payload := make(map[string]string)
	payload["error"] = message
	RespondWithJSON(w, code, payload)
}
