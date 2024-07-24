package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(response); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

	}

	w.WriteHeader(http.StatusCreated)

}
