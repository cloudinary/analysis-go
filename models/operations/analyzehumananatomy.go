// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/cloudinary/analysis-go/cloudinaryanalysis/models/components"
)

type AnalyzeHumanAnatomyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeHumanAnatomyResponse *components.AnalyzeHumanAnatomyResponse
	// Analysis accepted
	AsyncOperationAcceptedResponse *components.AsyncOperationAcceptedResponse
}

func (o *AnalyzeHumanAnatomyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeHumanAnatomyResponse) GetAnalyzeHumanAnatomyResponse() *components.AnalyzeHumanAnatomyResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeHumanAnatomyResponse
}

func (o *AnalyzeHumanAnatomyResponse) GetAsyncOperationAcceptedResponse() *components.AsyncOperationAcceptedResponse {
	if o == nil {
		return nil
	}
	return o.AsyncOperationAcceptedResponse
}
