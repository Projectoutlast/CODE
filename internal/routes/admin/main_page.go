package admin

import (
	"fmt"
	"net/http"
)

func (h *AdminHandlers) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index page")
}
