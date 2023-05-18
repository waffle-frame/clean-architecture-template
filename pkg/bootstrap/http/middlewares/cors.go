package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/misc/response"
)

func (m *provider) CORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			data, _ := json.Marshal(response.Build(response.ErrNoContent))
			w.Write(data)
			return
		}

		next(w, r)
	})
}
