package bothandlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) getActivity(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	f, err := file.Open()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.botSrv.SendPhotoToBot(bytes, file.Filename)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, nil)
}
