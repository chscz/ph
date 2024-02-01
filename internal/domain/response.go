package domain

import (
	"encoding/json"
	"fmt"
)

func MakeJSONResponse(httpCode int, message string, products []*Product) string {
	resp := map[string]any{
		"meta": map[string]interface{}{
			"code":    httpCode,
			"message": message,
		},
		"products": products,
	}
	j, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal error")
	}
	return string(j)
}
