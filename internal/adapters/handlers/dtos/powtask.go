package dtos

import "github.com/volvofixthis/pow-server/internal/core/models"

type PowTaskResp struct {
	Text      string `json:"text"`
	Salt      []byte `json:"salt"`
	Iteration uint32 `json:"iteration"`
	Memory    uint32 `json:"memory"`
}

func (ptr *PowTaskResp) FromDomain(powTask *models.PowTask) {
	ptr.Text = powTask.Text
	ptr.Salt = powTask.Salt
	ptr.Iteration = powTask.Iteration
	ptr.Memory = powTask.Memory
}
