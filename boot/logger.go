package boot

import (
	"mermaid/global"
)

func bootLogger() error {
	logInfo  := global.Conf.Log
	if err := global.InitLog(logInfo); err != nil {
		return err
	}
	return nil
}