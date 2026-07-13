package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func NewSwaggerHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.Open("docs/swagger.json")
		if err != nil {
			http.Error(w, fmt.Sprintf("swagger docs not found. Run 'swag init' first: %v", err), http.StatusNotFound)
			return
		}
		defer data.Close()

		var spec map[string]interface{}
		if err := json.NewDecoder(data).Decode(&spec); err != nil {
			http.Error(w, fmt.Sprintf("failed to parse swagger docs: %v", err), http.StatusInternalServerError)
			return
		}

		if spec["swagger"] != "2.0" {
			spec["swagger"] = "2.0"
		}

		if _, ok := spec["info"]; !ok {
			spec["info"] = map[string]interface{}{}
		}
		info := spec["info"].(map[string]interface{})
		if info["title"] == nil {
			info["title"] = "Banking API"
		}
		if info["version"] == nil {
			info["version"] = "1.0"
		}
		if info["description"] == nil {
			info["description"] = "A Banking REST API built with Go"
		}
		if info["contact"] == nil {
			info["contact"] = map[string]interface{}{}
		}

		if _, ok := spec["securityDefinitions"]; !ok {
			spec["securityDefinitions"] = map[string]interface{}{
				"BearerAuth": map[string]interface{}{
					"type": "apiKey",
					"in":   "header",
					"name": "Authorization",
				},
			}
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		_ = enc.Encode(spec)
	})

	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
  <title>API Docs</title>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function() {
      SwaggerUIBundle({
        url: "/swagger.json",
        dom_id: '#swagger-ui',
        requestInterceptor: function(request) {
          if (request.headers && request.headers.Authorization) {
            var token = request.headers.Authorization;
            if (!token.startsWith("Bearer ")) {
              request.headers.Authorization = "Bearer " + token;
            }
          }
          return request;
        },
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIBundle.SwaggerUIStandalonePreset],
        layout: "BaseLayout"
      })
    }
  </script>
</body>
</html>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	return mux
}
