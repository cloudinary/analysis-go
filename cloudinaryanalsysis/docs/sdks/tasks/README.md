# Tasks
(*Tasks*)

## Overview

Query the status of analysis tasks.

### Available Operations

* [GetStatus](#getstatus) - Get analysis task status

## GetStatus

Get the status of an analysis task.

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

    res, err := s.Tasks.GetStatus(ctx, "053f4bde4b933c8ecef23724ecde63b667c1ea21816d56c161c7ec1df6297da4b43109625650e9edf0f42152cc4cc32c8ad57824ac75ba8e05020f827c415559ac1248076a2d72c0a73af0479cca77eb")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTaskStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      | Example                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |                                                                                                                                                                  |
| `taskID`                                                                                                                                                         | *string*                                                                                                                                                         | :heavy_check_mark:                                                                                                                                               | The ID of the analysis task                                                                                                                                      | 053f4bde4b933c8ecef23724ecde63b667c1ea21816d56c161c7ec1df6297da4b43109625650e9edf0f42152cc4cc32c8ad57824ac75ba8e05020f827c415559ac1248076a2d72c0a73af0479cca77eb |
| `opts`                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |                                                                                                                                                                  |

### Response

**[*operations.GetAnalysisTaskStatusResponse](../../models/operations/getanalysistaskstatusresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ErrorResponse       | 400, 401, 403, 404            | application/json              |
| apierrors.RateLimitedResponse | 429                           | application/json              |
| apierrors.ErrorResponse       | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |