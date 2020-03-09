package handler

import (
	"encoding/json"
	"github.com/yusank/klyn"
	"log"
)

// NewRouter - register router
func NewRouter(r *klyn.RouterGroup) {
	githubRouter := r.Group("/github")
	{
		githubRouter.POST("/webhook", githubWebhookHandler)
	}
}

func githubWebhookHandler(c *klyn.Context) {
	var f = make(map[string]interface{})
	if err := bindJson(&f, c);err != nil {
		c.AbortWithJSON(klyn.K{"errcode": -1})
		return
	}

	log.Println(f)
	c.JSON(200, klyn.K{"errcode": 0})
}

func bindJson(v interface{}, c *klyn.Context) error {
	d := json.NewDecoder(c.Request.Body)
	return d.Decode(v)
}