package config

import (
	"fmt"
	"github.com/hanyxp/mytools/util"
)

func Run(s string)  {

	ip := util.GetIp()
	fmt.Printf("out:%+v", ip)

}
