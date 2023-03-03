package utility

import (
	"fmt"
	"net/http"
)

// Client err
func ClientErr(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// Server err
func ServerErr(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func ErrMsg(s string) error {
	return fmt.Errorf(s)
}
