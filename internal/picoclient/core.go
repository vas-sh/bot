package picoclient

import (
	"net/http"
)

type srv struct {
	BaceApi string
}

func New(baseApi string) *srv {
	return &srv{
		BaceApi: baseApi,
	}
}

type PicoResp struct {
	Path string `json:"path"`
}

func (s *srv) TakePhoto() error {
	_, err := http.Get(s.BaceApi + "/take_photo")
	if err != nil {
		return err
	}
	return nil
}
