package picoclient

import (
	"encoding/json"
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

func (s *srv) TakePhoto() (string, error) {
	resp, err := http.Get(s.BaceApi + "/take_photo")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r PicoResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}
	return r.Path, nil
}
