package fixture

import "testing"

func TestIdentity(t *testing.T) { if Identity("fixture") != "fixture" { t.Fatal("fixture failed") } }
