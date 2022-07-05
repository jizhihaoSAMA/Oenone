package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func BuildESBody(body gin.H) (bytes.Buffer, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		log.Printf("Error encoding query: %s\n", err)
		return bytes.Buffer{}, err
	}
	return buf, nil
}
