package mod2

import "testing"

func TestMod2_V2(t *testing.T) {
        if Mod2() != "Mod2-V2" {
                t.Error("Mod2_V2 failed. Expected: Mod2-V2 Actual: >",Mod2(),"<")
        }
}
