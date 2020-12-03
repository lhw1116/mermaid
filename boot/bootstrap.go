package boot

import (
	"github.com/pkg/errors"
	"mermaid/global"
)

func Boot() error {
	if err := bootConfig(); err != nil {
		return errors.WithMessage(err, "初始化配置文件出错")
	}

	if err := bootLogger(); err != nil {
		return errors.WithMessage(err, "初始化日志出错")
	}

	if err := bootDB(); err != nil {
		return errors.WithMessage(err, "初始化数据库出错")
	}

	if err := bootCache(); err != nil {
		return errors.WithMessage(err, "初始化redis缓存出错")
	}

	if err := bootHTTP(); err != nil {
		return errors.WithMessage(err, "初始化http路由出错")
	}

	global.Logger.Info("[程序初始化已完成]")

	return nil
}
