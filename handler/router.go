package handler

import (
	"github.com/yusank/klyn"
)

// NewRouter - register router
func NewRouter(r *klyn.RouterGroup) {
	githubRouter := r.Group("/github")
	{
		githubRouter.POST("/webhook", githubWebhookHandler)
	}
}

func githubWebhookHandler(c *klyn.Context) {
	c.JSON(200, klyn.K{"errcode": 0})
}
