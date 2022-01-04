package mod1

import "testing"

func TestMod1_V2(t *testing.T) {
	if Mod1() != "Mod1-V1-with-Mod2-V1(Mod2-V1)" {
		t.Error("Mod1_V2 failed. Expected: Mod1-V1-with-Mod2-V1(Mod2-V1) Actual: >",Mod1(),"<")
	} 
}

