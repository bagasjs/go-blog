package helper

type Response struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data any `json:"data"`
}

func ResponseOK(msg string, data any) Response {
    return Response {
        Success: true,
        Message: msg,
        Data: data,
    }
}

func ResponseErr(msg string) Response {
    return Response {
        Success: false,
        Message: msg,
        Data: nil,
    }
}
