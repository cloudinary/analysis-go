// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/components"
)

type AnalyzeAiVisionModerationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeAIVisionModerationResponse *components.AnalyzeAIVisionModerationResponse
	// Analysis accepted
	AsyncOperationAcceptedResponse *components.AsyncOperationAcceptedResponse
}

func (o *AnalyzeAiVisionModerationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeAiVisionModerationResponse) GetAnalyzeAIVisionModerationResponse() *components.AnalyzeAIVisionModerationResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeAIVisionModerationResponse
}

func (o *AnalyzeAiVisionModerationResponse) GetAsyncOperationAcceptedResponse() *components.AsyncOperationAcceptedResponse {
	if o == nil {
		return nil
	}
	return o.AsyncOperationAcceptedResponse
}
