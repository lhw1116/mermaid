package global

import "github.com/gin-gonic/gin"

func InitHttp(addr string) error {
	r := gin.New()
	err := r.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
