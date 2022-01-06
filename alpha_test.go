package go_demo_module

import (
	"testing"
	"github.com/shekhar-jha/go-demo-module/mod1"
	"github.com/shekhar-jha/go-demo-module/mod2"
)

 func TestMod1_V2(t *testing.T) {
	if mod1.Mod1() != "Mod1-V1" {
		t.Error("Mod1_V1 failed. Expected: Mod1-V1 Actual: >",mod1.Mod1(),"<")
 	} 
 }

func TestMod2_V1(t *testing.T) {
        if mod2.Mod2() != "Mod2-V1" {
                t.Error("Mod2_V1 failed. Expected: Mod2-V1 Actual: >",mod2.Mod2(),"<")
        }
}


