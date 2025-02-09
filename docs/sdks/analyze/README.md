# Analyze
(*Analyze*)

## Overview

Use the Analyze API to analyze any external asset and return details based on the type of analysis requested.

### Available Operations

* [AiVisionGeneral](#aivisiongeneral) - Analyze - AI Vision General
* [AiVisionModeration](#aivisionmoderation) - Analyze - AI Vision Moderation
* [AiVisionTagging](#aivisiontagging) - Analyze - AI Vision Tagging
* [Captioning](#captioning) - Analyze - Captioning
* [CldFashion](#cldfashion) - Analyze - Cld-Fashion
* [CldText](#cldtext) - Analyze - Cld-Text
* [Coco](#coco) - Analyze - Coco
* [GoogleLogoDetection](#googlelogodetection) - Analyze - Google Logo Detection
* [GoogleTagging](#googletagging) - Analyze - Google Tagging
* [HumanAnatomy](#humananatomy) - Analyze - Human Anatomy
* [ImageQuality](#imagequality) - Analyze - Image Quality Analysis
* [Lvis](#lvis) - Analyze - Lvis
* [ShopClassifier](#shopclassifier) - Analyze - Shop Classifier
* [Unidet](#unidet) - Analyze - Unidet
* [WatermarkDetection](#watermarkdetection) - Analyze - Watermark Detection

## AiVisionGeneral

The General mode serves a wide array of applications by providing detailed answers to diverse questions about an image. Users can inquire about any aspect of an image, such as identifying objects, understanding scenes, or interpreting text within the image.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.AnalyzeAIVisionGeneralRequest](../../models/components/analyzeaivisiongeneralrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AnalyzeAiVisionGeneralResponse](../../models/operations/analyzeaivisiongeneralresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## AiVisionModeration

The Moderation mode accepts multiple questions about an image, to which the response provides concise answers of "yes," "no," or "unknown." This functionality allows for a nuanced evaluation of whether the image adheres to specific content policies, creative specs, or aesthetic criteria.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.AiVisionModeration(ctx, components.AnalyzeAIVisionModerationRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
        RejectionQuestions: []string{
            "Does this image contain any violent activity?",
            "Is there any nudity in the image?",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeAIVisionModerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [components.AnalyzeAIVisionModerationRequest](../../models/components/analyzeaivisionmoderationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AnalyzeAiVisionModerationResponse](../../models/operations/analyzeaivisionmoderationresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## AiVisionTagging

The Tagging mode accepts a list of tag names along with their corresponding descriptions. If the image matches the description, which may encompass various elements, it will be appropriately tagged. This approach enables customers to align with their own brand taxonomy, offering a dynamic, flexible, and open method for image classification.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.AiVisionTagging(ctx, components.AnalyzeAIVisionTaggingRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
        TagDefinitions: []components.TagDefinitions{
            components.TagDefinitions{
                Name: "cigarettes",
                Description: "Does the image contain a pack of cigarettes?",
            },
            components.TagDefinitions{
                Name: "furniture",
                Description: "Does the image contain any tables, chairs, couches or sofas?",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeAIVisionTaggingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.AnalyzeAIVisionTaggingRequest](../../models/components/analyzeaivisiontaggingrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AnalyzeAiVisionTaggingResponse](../../models/operations/analyzeaivisiontaggingresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Captioning

Provides a caption for an image.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.Captioning(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCaptioningResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeCaptioningResponse](../../models/operations/analyzecaptioningresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CldFashion

Analyze an image using the [Cld-Fashion](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's fashion model is specifically dedicated to items of clothing. The response includes attributes of the clothing identified, for example whether the garment contains pockets, its material and the fastenings used.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.CldFashion(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCldFashionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeCldFashionResponse](../../models/operations/analyzecldfashionresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CldText

Analyze an image using the [Cld-Text](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's text model tells you if your image includes text, and where it's located. Used with image tagging, you can then search for images that contain blocks of text. Used with object-aware cropping, you can choose to keep only the text part, or specify a crop that avoids the text.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.CldText(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCldTextResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeCldTextResponse](../../models/operations/analyzecldtextresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Coco

Analyze an image using the [Coco](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. The [Common Objects in Context](https://cocodataset.org/) model contains just 80 common objects.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.Coco(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCocoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeCocoResponse](../../models/operations/analyzecocoresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GoogleLogoDetection

Detects popular product logos within an image.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.GoogleLogoDetection(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeGoogleLogoDetectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeGoogleLogoDetectionResponse](../../models/operations/analyzegooglelogodetectionresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GoogleTagging

Provides tags for an image using Google's tagging service.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.GoogleTagging(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeGoogleTaggingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeGoogleTaggingResponse](../../models/operations/analyzegoogletaggingresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## HumanAnatomy

Analyze an image using the [Human Anatomy](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's human anatomy model identifies parts of the human body in an image. It works best when the majority of a human body is detected in the image.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.HumanAnatomy(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeHumanAnatomyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeHumanAnatomyResponse](../../models/operations/analyzehumananatomyresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ImageQuality

Analyze an image using the [Image Quality Analysis](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#image_quality_analysis) model.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.ImageQuality(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeImageQualityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeImageQualityResponse](../../models/operations/analyzeimagequalityresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Lvis

Analyze an image using the [Lvis](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. The [Large Vocabulary Instance Segmentation](https://www.lvisdataset.org/) model contains thousands of general objects.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.Lvis(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeLvisResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeLvisResponse](../../models/operations/analyzelvisresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ShopClassifier

Analyze an image using the [Shop Classifier](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's shop classifier model detects if the image is a product image taken in a studio, or if it's a natural image.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.ShopClassifier(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeShopClassifierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeShopClassifierResponse](../../models/operations/analyzeshopclassifierresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Unidet

Analyze an image using the [Unidet](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. The [UniDet](https://github.com/xingyizhou/UniDet) model is a unified model, combining a number of object models, including [Objects365](https://www.objects365.org/overview.html), which focuses on diverse objects in the wild.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.Unidet(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeUnidetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeUnidetResponse](../../models/operations/analyzeunidetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## WatermarkDetection

Analyze an image using the [Watermark Detection](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#watermark_detection) detection model.

### Example Usage

```go
package main

import(
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
                APIKey: "CLOUDINARY_API_KEY",
                APISecret: "CLOUDINARY_API_SECRET",
            },
        }),
    )

    res, err := s.Analyze.WatermarkDetection(ctx, components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeWatermarkDetectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.AnalyzeWatermarkDetectionResponse](../../models/operations/analyzewatermarkdetectionresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |