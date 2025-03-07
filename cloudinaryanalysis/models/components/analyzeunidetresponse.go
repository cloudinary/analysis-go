// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AnalyzeUnidetResponseData struct {
	Entity   *string             `json:"entity,omitempty"`
	Analysis *UnidetAnalysisData `json:"analysis,omitempty"`
}

func (o *AnalyzeUnidetResponseData) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *AnalyzeUnidetResponseData) GetAnalysis() *UnidetAnalysisData {
	if o == nil {
		return nil
	}
	return o.Analysis
}

type AnalyzeUnidetResponse struct {
	Limits    *LimitsObject              `json:"limits,omitempty"`
	RequestID *string                    `json:"request_id,omitempty"`
	Data      *AnalyzeUnidetResponseData `json:"data,omitempty"`
}

func (o *AnalyzeUnidetResponse) GetLimits() *LimitsObject {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *AnalyzeUnidetResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *AnalyzeUnidetResponse) GetData() *AnalyzeUnidetResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
