package core

import "time"

type Response struct {
	startDuration time.Time
	status        int
	success       bool
	duration      float64
	response      interface{}
	errorMessage  []string
}

func NewResponse() *Response {
	return &Response{
		startDuration: time.Now(),
		status:        200,
		success:       true,
		duration:      0.0,
		response:      nil,
		errorMessage:  make([]string, 0),
	}
}

func (r *Response) Status(statusNumber int) {
	r.status = statusNumber
}

func (r *Response) Success(value bool) {
	r.success = value
}

func (r *Response) SetResponse(res interface{}) {
	r.response = res
}

func (r *Response) Error(msg string) {
	r.errorMessage = append(r.errorMessage, msg)
	r.success = false
	r.status = 400
}

func (r *Response) Stack() map[string]interface{} {
	r.duration = time.Since(r.startDuration).Seconds()

	resForm := map[string]interface{}{
		"status":   r.status,
		"success":  r.success,
		"duration": r.duration,
		"response": r.response,
	}

	if r.status != 200 {
		resForm["error_message"] = r.errorMessage
	}

	return resForm
}
