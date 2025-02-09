package hooks

import (
	"encoding/base64"
	"fmt"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/internal/cldconfig"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/components"
	"net/http"
	"strings"
)

type CloudinaryHook struct {
	cfg *cldconfig.Configuration
}

func (h *CloudinaryHook) SDKInit(baseURL string, client HTTPClient) (string, HTTPClient) {
	cfg, err := cldconfig.New()
	if err != nil {
		panic(err)
	}
	h.cfg = cfg

	return baseURL, client
}

func (h *CloudinaryHook) BeforeRequest(ctx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	cloudName := h.cfg.Cloud.CloudName
	apiKey := h.cfg.Cloud.APIKey
	apiSecret := h.cfg.Cloud.APISecret

	// If SecuritySource is nil, we skip calling it and just use env-based config.
	if ctx.SecuritySource != nil {
		val, err := ctx.SecuritySource(req.Context())
		if err != nil {
			return nil, err
		}

		// If val is non-nil and has override credentials, apply them.
		switch s := val.(type) {
		case components.Security:
			if s.CloudinaryAuth != nil {
				if s.CloudinaryAuth.APIKey != "" {
					apiKey = s.CloudinaryAuth.APIKey
				}
				if s.CloudinaryAuth.APISecret != "" {
					apiSecret = s.CloudinaryAuth.APISecret
				}
			}
		case nil:
			// val is literally nil -> fallback to env-based creds
		default:
			// Some other type -> fallback to env-based
			fmt.Printf("SecuritySource returned unexpected type: %T\n", s)
		}
	}

	// Replace /CLOUD_NAME/ in path if we have a cloudName
	if cloudName != "" && strings.Contains(req.URL.Path, "/CLOUD_NAME/") {
		req.URL.Path = strings.ReplaceAll(req.URL.Path, "/CLOUD_NAME/", "/"+cloudName+"/")
	}

	// If we have final creds, set Basic Auth
	if apiKey != "" && apiSecret != "" {
		authVal := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", apiKey, apiSecret)))
		req.Header.Set("Authorization", "Basic "+authVal)
	}

	return req, nil
}
