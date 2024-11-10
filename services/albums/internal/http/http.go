package http

type HttpError struct {
    Reason string `json:"reason"`
}

func NewHttpError(reason string) HttpError {
    return HttpError{Reason: reason}
}
