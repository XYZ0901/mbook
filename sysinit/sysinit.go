package sysinit

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
)

func sysinit()  {
	uploads := filepath.Join("./","uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads

	// 注册前端使用函数
	registerFunctions()
}

func registerFunctions()  {
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs","")
		if strings.HasPrefix(p,"/") && strings.HasSuffix(cdn,"/") {
			return cdn + p[1:]
		}
	})
}