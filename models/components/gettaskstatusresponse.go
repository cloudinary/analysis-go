// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GetTaskStatusResponseData struct {
	TaskID *string `json:"task_id,omitempty"`
	// The status of an async task
	Status *TaskStatus `json:"status,omitempty"`
}

func (o *GetTaskStatusResponseData) GetTaskID() *string {
	if o == nil {
		return nil
	}
	return o.TaskID
}

func (o *GetTaskStatusResponseData) GetStatus() *TaskStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetTaskStatusResponse struct {
	Data      *GetTaskStatusResponseData `json:"data,omitempty"`
	RequestID *string                    `json:"request_id,omitempty"`
}

func (o *GetTaskStatusResponse) GetData() *GetTaskStatusResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetTaskStatusResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}
