package mod3

import "testing"

func TestMod3_V1(t *testing.T) {
        if Mod3() != "Mod3-V1-with-Mod1-V2(Mod1-V2-with-Mod2-V1(Mod2-V1))" {
                t.Error("Mod3_V1 failed. Expected: Mod3-V1-with-Mod1-V2(Mod1-V2-with-Mod2-V1(Mod2-V1)) Actual: >",Mod3(),"<")
        }
}

