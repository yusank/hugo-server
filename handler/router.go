package handler

import (
	"bufio"
	"encoding/json"
	"github.com/yusank/klyn"
	"io"
	"log"
	"net/http"
	"os/exec"
)

// NewRouter - register router
func NewRouter(r *klyn.RouterGroup) {
	githubRouter := r.Group("/github")
	{
		githubRouter.POST("/webhook", githubWebhookHandler)
	}

	r.POST("/admin/web/restart",restartHandler)
}

func githubWebhookHandler(c *klyn.Context) {
	var f = make(map[string]interface{})
	if err := bindJson(&f, c);err != nil {
		c.AbortWithJSON(klyn.K{"errcode": -1})
		return
	}

	if err := restartHugoWeb();err != nil {
		log.Println("restart err:",err)
	}

	log.Println(f)
	c.JSON(200, klyn.K{"errcode": 0})
}

func bindJson(v interface{}, c *klyn.Context) error {
	d := json.NewDecoder(c.Request.Body)
	return d.Decode(v)
}

func restartHandler(c *klyn.Context) {
	if err := restartHugoWeb();err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK,"ok")
}

func restartHugoWeb() error {
	cmd :=  exec.Command("/bin/sh", "-c","/home/dev/hugo.yusank.space/deploy.sh")
	log.Println(2,cmd.String())
	out,err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	eOut,err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start();err != nil {
		return err
	}

	go asyncScanner(out)
	go asyncScanner(eOut)

	return cmd.Wait()
}

func asyncScanner(r io.Reader) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		log.Println(sc.Text())
	}
}