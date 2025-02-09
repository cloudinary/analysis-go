# Cloudinary Analysis Go SDK

<!-- Start Summary [summary] -->
## Summary

Analyze API (Beta): Use the Analyze API to analyze any external asset and return details based on the type of analysis requested.

Currently supports the following analysis options:
  * [AI Vision - Tagging](https://cloudinary.com/documentation/cloudinary_ai_vision_addon#tagging_mode)
  * [AI Vision - Moderation](https://cloudinary.com/documentation/cloudinary_ai_vision_addon#moderation_mode)
  * [AI Vision - General](https://cloudinary.com/documentation/cloudinary_ai_vision_addon#general_mode)
  * [Captioning](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning)
  * [Cld Fashion](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Cld Text](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Coco](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Google Tagging](https://cloudinary.com/documentation/google_auto_tagging_addon)
  * [Human Anatomy](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Image Quality Analysis](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#image_quality_analysis)
  * [Lvis](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Shop Classifier](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Unidet](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Watermark Detection](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#watermark_detection)

  **Notes**: 
  * The Analyze API is currently in development and is available as a Public Beta, which means we value your feedback, so please feel free to [share any thoughts with us](https://support.cloudinary.com/hc/en-us/requests/new).
  * The analysis options require an active subscription to the relevant add-on. Learn more about [registering for add-ons](https://cloudinary.com/documentation/cloudinary_add_ons#registering_for_add_ons).

  The API supports both Basic Authentication using your Cloudinary API Key and API Secret (which can be found on the Dashboard page of your [Cloudinary Console](https://console.cloudinary.com/pm)) or OAuth2 ([Contact support](https://support.cloudinary.com/hc/en-us/requests/new) for more information regarding OAuth).
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Cloudinary Analysis Go SDK](#cloudinary-analysis-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Server Selection](#server-selection)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/cloudinary/analysis-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Server Selection [server] -->
## Server Selection

### Server Variables

The default server `https://api.cloudinary.com/v2/analysis/{cloud_name}` contains variables and is set to `https://api.cloudinary.com/v2/analysis/CLOUD_NAME` by default. To override default values, the following options are available when initializing the SDK client instance:
 * `WithCloudName(cloudName string)`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		cloudinaryanalysis.WithServerURL("https://api.cloudinary.com/v2/analysis/CLOUD_NAME"),
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
<!-- End Server Selection [server] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name             | Type   | Scheme       |
| ---------------- | ------ | ------------ |
| `CloudinaryAuth` | http   | Custom HTTP  |
| `OAuth2`         | oauth2 | OAuth2 token |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Analyze](docs/sdks/analyze/README.md)

* [AiVisionGeneral](docs/sdks/analyze/README.md#aivisiongeneral) - Analyze - AI Vision General
* [AiVisionModeration](docs/sdks/analyze/README.md#aivisionmoderation) - Analyze - AI Vision Moderation
* [AiVisionTagging](docs/sdks/analyze/README.md#aivisiontagging) - Analyze - AI Vision Tagging
* [Captioning](docs/sdks/analyze/README.md#captioning) - Analyze - Captioning
* [CldFashion](docs/sdks/analyze/README.md#cldfashion) - Analyze - Cld-Fashion
* [CldText](docs/sdks/analyze/README.md#cldtext) - Analyze - Cld-Text
* [Coco](docs/sdks/analyze/README.md#coco) - Analyze - Coco
* [GoogleLogoDetection](docs/sdks/analyze/README.md#googlelogodetection) - Analyze - Google Logo Detection
* [GoogleTagging](docs/sdks/analyze/README.md#googletagging) - Analyze - Google Tagging
* [HumanAnatomy](docs/sdks/analyze/README.md#humananatomy) - Analyze - Human Anatomy
* [ImageQuality](docs/sdks/analyze/README.md#imagequality) - Analyze - Image Quality Analysis
* [Lvis](docs/sdks/analyze/README.md#lvis) - Analyze - Lvis
* [ShopClassifier](docs/sdks/analyze/README.md#shopclassifier) - Analyze - Shop Classifier
* [Unidet](docs/sdks/analyze/README.md#unidet) - Analyze - Unidet
* [WatermarkDetection](docs/sdks/analyze/README.md#watermarkdetection) - Analyze - Watermark Detection


### [Tasks](docs/sdks/tasks/README.md)

* [GetStatus](docs/sdks/tasks/README.md#getstatus) - Get analysis task status

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/components"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/retry"
	"log"
	"models/operations"
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
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AnalyzeAIVisionGeneralResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/components"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := cloudinaryanalysis.New(
		cloudinaryanalysis.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `AiVisionGeneral` function may return the following errors:

| Error Type                    | Status Code        | Content Type     |
| ----------------------------- | ------------------ | ---------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404 | application/json |
| apierrors.RateLimitedResponse | 429                | application/json |
| apierrors.ErrorResponse       | 500                | application/json |
| apierrors.APIError            | 4XX, 5XX           | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis"
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/apierrors"
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

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.RateLimitedResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/cloudinary/analysis-go/cloudinaryanalysis&utm_campaign=go)
