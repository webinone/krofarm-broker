package cfg

import (
	"log"
	"github.com/kardianos/osext"
)

func NewAppCfg() *Cfg {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("################# folderPath : " , folderPath)

	//c := NewCfg(folderPath + "/property/app-config.properties")
	c := NewCfg("D:/Project/krofarm-broker/src/property/app-config.properties")

	if err := c.Load() ; err != nil {
		//panic(err)
		log.Fatal(err)
	}
	return c
}

//func NewCfgMapper(configName string) *Cfg {
//	c := NewCfg("./resources/app-config.properties")
//	if err := c.Load() ; err != nil {
//		panic(err)
//	}
//	return c
//}


