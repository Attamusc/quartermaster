package resolver

import (
    "fmt"
)

type Resolver struct {
    domain string
    path   string
}

func GetResolverFor(repo string) *Resolver {
    return &Resolver{}
}

func (r *Resolver) Install(location string) (bool, error) {
    if len(location) == 0 {
        return false, fmt.Errorf("Installation requires a valid location")
    }

    fmt.Println("Installing")

    return true, nil
}
