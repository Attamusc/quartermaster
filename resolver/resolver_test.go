package resolver

import (
    "testing"
)

func Test_ResolverInstallWithoutLocation(t *testing.T) {
    success, err := Install("")

    if success || err == nil {
        t.Error("Resolver shouldn't be able to install without an origin location")
    }
}
