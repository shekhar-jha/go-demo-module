package mod1

import "github.com/shekhar-jha/go-demo-module/mod2"

type Mod1Struct struct {
	Astring string
}

func Mod1() string {
        return "Mod1-V1-with-Mod2-V1(" + mod2.Mod2()+")"
}
