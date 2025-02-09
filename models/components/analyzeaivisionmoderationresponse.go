// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AnalyzeAIVisionModerationResponseData struct {
	Entity   *string                         `json:"entity,omitempty"`
	Analysis *AIVisionModerationAnalysisData `json:"analysis,omitempty"`
}

func (o *AnalyzeAIVisionModerationResponseData) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *AnalyzeAIVisionModerationResponseData) GetAnalysis() *AIVisionModerationAnalysisData {
	if o == nil {
		return nil
	}
	return o.Analysis
}

type AnalyzeAIVisionModerationResponse struct {
	Limits    *LimitsObject                          `json:"limits,omitempty"`
	RequestID *string                                `json:"request_id,omitempty"`
	Data      *AnalyzeAIVisionModerationResponseData `json:"data,omitempty"`
}

func (o *AnalyzeAIVisionModerationResponse) GetLimits() *LimitsObject {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *AnalyzeAIVisionModerationResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *AnalyzeAIVisionModerationResponse) GetData() *AnalyzeAIVisionModerationResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
