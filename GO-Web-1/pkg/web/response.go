package web

import "strconv"

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{
			Code:  strconv.FormatInt(int64(code), 10),
			Data:  data,
			Error: "",
		}
	}
	if code == 401 {
		return Response{
			Code:  strconv.FormatInt(int64(code), 10),
			Data:  nil,
			Error: "no tienes permisos para la peticion solicitada",
		}

	}
	return Response{
		Code:  strconv.FormatInt(int64(code), 10),
		Data:  nil,
		Error: err,
	}
}
