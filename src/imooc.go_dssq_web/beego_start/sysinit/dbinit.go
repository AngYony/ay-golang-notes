package sysinit

// 数据库初始化
func dbinit(alias string) {
	dbAlias := alias
	if "w" == alias || "default" == alias || len(alias) <= 0 {
		dbAlias = "default"
		alias = "w"
	}

	// 数据库名称
	dbName := beego

}
