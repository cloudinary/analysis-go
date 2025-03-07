// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AnalyzeShopClassifierResponseData struct {
	Entity   *string                     `json:"entity,omitempty"`
	Analysis *ShopClassifierAnalysisData `json:"analysis,omitempty"`
}

func (o *AnalyzeShopClassifierResponseData) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *AnalyzeShopClassifierResponseData) GetAnalysis() *ShopClassifierAnalysisData {
	if o == nil {
		return nil
	}
	return o.Analysis
}

type AnalyzeShopClassifierResponse struct {
	Limits    *LimitsObject                      `json:"limits,omitempty"`
	RequestID *string                            `json:"request_id,omitempty"`
	Data      *AnalyzeShopClassifierResponseData `json:"data,omitempty"`
}

func (o *AnalyzeShopClassifierResponse) GetLimits() *LimitsObject {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *AnalyzeShopClassifierResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *AnalyzeShopClassifierResponse) GetData() *AnalyzeShopClassifierResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
