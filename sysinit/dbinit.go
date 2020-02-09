package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// dbinit() //初始化主库
// dbinit("w") 或 dbinit("default")
// dbinit("w","r"...)
func dbinit(alias string)  {
	dbAlias := alias //default
	if "w" == alias || "default"==alias || len(alias)<=0{
		dbAlias = "default"
		alias = "w"
	}
	// 数据库名称
	dbName := beego.AppConfig.String("db_"+alias+"_database")
	// 数据库用户名
	dbUser := beego.AppConfig.String("db_"+alias+"_username")
	// 数据库密码
	dbPwd := beego.AppConfig.String("db_"+alias+"_password")
	// 数据库IP
	dbHost := beego.AppConfig.String("db_"+alias+"_host")
	// 数据库端口
	dbPost := beego.AppConfig.String("db_"+alias+"_post")
	// user:pwd@tcp(host:port)/db?charset=utf8 30闲置连接数；最大连接数
	orm.RegisterDataBase(dbAlias,"mysql",
		dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPost+")/"+dbName+"?charset=utf8",30)

	isDev := ("dev" == beego.AppConfig.String("runmode"))
	if "w" == alias{
		orm.RunSyncdb("default",false,isDev)
	}

	if isDev {
		orm.Debug = isDev
	}
}
