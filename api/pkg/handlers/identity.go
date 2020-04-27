package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/replicatedhq/kots-validation/api/pkg/identity"
	"github.com/replicatedhq/kots-validation/api/pkg/logger"
)

func GenerateIdentity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	if err := identity.Generate(n); err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func GetIdentity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	ids, err := identity.List()
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	data, err := json.MarshalIndent(ids, "", "    ")
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}
