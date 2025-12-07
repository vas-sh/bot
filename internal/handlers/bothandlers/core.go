package bothandlers

import "github.com/gin-gonic/gin"

type servicer interface {
	SendPhotoToBot(photo []byte, fileName string) error
}

type handler struct {
	botSrv servicer
}

func New(botSrv servicer) *handler {
	return &handler{
		botSrv: botSrv,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	botRouter := router.Group("bot")
	botRouter.POST("/activity", h.getActivity)
}
