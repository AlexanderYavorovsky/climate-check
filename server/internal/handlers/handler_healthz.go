package handlers

import (
	"net/http"
	"server/internal/jsonutils"
)

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	jsonutils.RespondWithJSON(w, 200, struct{}{})
}
