package mid

import (
	"github.com/yusank/klyn"
	klog "github.com/yusank/klyn-log"
)

var (
	logger klog.Logger
)

func init() {
	logger = klog.DefaultLogger()
}
func LogMid(c *klyn.Context) {
	logger.Debug(klyn.K{
		"url":c.Request.URL.String(),
		"code":c.Writer.Status(),
	})
}
