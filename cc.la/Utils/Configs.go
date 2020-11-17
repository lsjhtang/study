package Utils

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var ProxyConfig map[string]string

type EnvConfig *os.File

func init()  {
	ProxyConfig = make(map[string]string)
	EnvConfig,err := ini.Load(".env")
	if err!=nil {
		log.Print(err)
		return
	}
	sec,_ := EnvConfig.GetSection("proxy")
	chi := sec.ChildSections()
	for _,sec := range chi{
		path,_ := sec.GetKey("path")
		pass,_ := sec.GetKey("pass")
		if path != nil && pass != nil {
			ProxyConfig[path.Value()] = pass.Value()
		}
	}
}
