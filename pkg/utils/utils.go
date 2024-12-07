package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, model interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	if err := json.Unmarshal(body, model); err != nil {
		fmt.Println("Error unmarshalling body:", err)
		return
	}
}
