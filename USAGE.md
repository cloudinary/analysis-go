<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := cloudinaryanalysis.New(
		cloudinaryanalysis.WithSecurity(components.Security{
			CloudinaryAuth: &components.SchemeCloudinaryAuth{
				APIKey:    "CLOUDINARY_API_KEY",
				APISecret: "CLOUDINARY_API_SECRET",
			},
		}),
	)

	res, err := s.Analyze.AiVisionGeneral(ctx, components.AnalyzeAIVisionGeneralRequest{
		Source: components.CreateSourceURI(
			components.URI{
				URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
			},
		),
		Prompts: []string{
			"Describe this image in detail",
			"Does this image contain an insect?",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AnalyzeAIVisionGeneralResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->