package boot

import "mermaid/global"

func bootDB() (err error) {
	dbs := global.Conf.Db

	return global.InitDB(dbs)
}
