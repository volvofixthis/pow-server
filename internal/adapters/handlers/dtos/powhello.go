package dtos

const (
	RequestState = iota
	ResponseState
)

type PowHelloReq struct {
	State int `json:"state"`
}
