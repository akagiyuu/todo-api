package util

type ApiError struct {
	Inner   error          `json:"-"`
	Code    int            `json:"-"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}

func (e ApiError) Error() string { return e.Message }

func (e ApiError) StatusCode() int { return e.Code }

// TODO: convert Details to string
func (e ApiError) DetailMsg() string { return e.Message }
