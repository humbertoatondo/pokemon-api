package helpers // import "github.com/humbertoatondo/pokemon-api/helpers"

VARIABLES

var LanguageMap = map[string]int{
"ja-Hrkt": 0,
"ko": 1,
"zh-Hant": 2,
"fr": 3,
"de": 4,
"es": 5,
"it": 6,
"en": 7,
"ja": 8,
"zh-Hans": 9,
}
LanguageMap is used to reference the available languages for pokemon moves.

FUNCTIONS

func ParseKeyFromURL(key string, r \*http.Request) (string, bool)
ParseKeyFromURL retreives the value for a parameter in the url.

func RespondWithError(w http.ResponseWriter, code int, message string)
RespondWithError uses RespondWithJSON to write an error.

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{})
RespondWithJSON sets and writes the headers and the response for the http
request.

TYPES

type HTTPGet func(string) (\*http.Response, error)
HTTPGet handles the http get requests.
