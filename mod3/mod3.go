package mod3

import "github.com/shekhar-jha/go-demo-module/mod1"

func Mod3() string {
        return "Mod3-V1-with-Mod1-V2("+ mod1.Mod1() + ")"
}

