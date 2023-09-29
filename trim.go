package trim

import (
	"net/http"
	"strings"
)

func init() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/sca/", http.StripPrefix("/sca/", fs))
}

func Trim(v string) string {
	return strings.TrimSpace(v)
}
