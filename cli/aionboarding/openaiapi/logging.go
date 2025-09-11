package openaiapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/openai/openai-go/v2/option"
)

// createLoggingMiddleware creates a middleware that logs request bodies for debugging
func createLoggingMiddleware() option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		// Read the request body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
		} else {
			// Pretty print the JSON with indentation
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
				log.Printf("OpenAI API Request Body (raw): %s", string(body))
			} else {
				log.Printf("OpenAI API Request Body:\n%s", prettyJSON.String())
			}
		}

		// Reset the request body for the next middleware/handler
		req.Body = io.NopCloser(bytes.NewReader(body))

		// Call the next middleware/handler
		return next(req)
	}
}
