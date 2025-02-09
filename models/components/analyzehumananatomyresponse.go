// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AnalyzeHumanAnatomyResponseData struct {
	Entity   *string                   `json:"entity,omitempty"`
	Analysis *HumanAnatomyAnalysisData `json:"analysis,omitempty"`
}

func (o *AnalyzeHumanAnatomyResponseData) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *AnalyzeHumanAnatomyResponseData) GetAnalysis() *HumanAnatomyAnalysisData {
	if o == nil {
		return nil
	}
	return o.Analysis
}

type AnalyzeHumanAnatomyResponse struct {
	Limits    *LimitsObject                    `json:"limits,omitempty"`
	RequestID *string                          `json:"request_id,omitempty"`
	Data      *AnalyzeHumanAnatomyResponseData `json:"data,omitempty"`
}

func (o *AnalyzeHumanAnatomyResponse) GetLimits() *LimitsObject {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *AnalyzeHumanAnatomyResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *AnalyzeHumanAnatomyResponse) GetData() *AnalyzeHumanAnatomyResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
