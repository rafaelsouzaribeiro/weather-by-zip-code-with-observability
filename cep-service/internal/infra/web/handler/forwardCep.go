package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
)

var errorInvalidCep = "invalid zipcode"
var cepRegex = regexp.MustCompile(`^\d{8}$`)

func (h *Handler) ForwardCep(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")

	if len(cep) != 8 {
		http.Error(w, errorInvalidCep, http.StatusUnprocessableEntity)
		return
	}

	if !cepRegex.MatchString(cep) {
		http.Error(w, errorInvalidCep, http.StatusUnprocessableEntity)
		return
	}

	localeoutput, err := h.usecase.GetInfo(r.Context(), cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(localeoutput)

}
