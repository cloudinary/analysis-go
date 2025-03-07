// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ImageQualityAnalysisData struct {
	Quality      *string  `json:"quality,omitempty"`
	Score        *float64 `json:"score,omitempty"`
	Confidence   *float64 `json:"confidence,omitempty"`
	ModelVersion *int64   `json:"model_version,omitempty"`
}

func (o *ImageQualityAnalysisData) GetQuality() *string {
	if o == nil {
		return nil
	}
	return o.Quality
}

func (o *ImageQualityAnalysisData) GetScore() *float64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *ImageQualityAnalysisData) GetConfidence() *float64 {
	if o == nil {
		return nil
	}
	return o.Confidence
}

func (o *ImageQualityAnalysisData) GetModelVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.ModelVersion
}
