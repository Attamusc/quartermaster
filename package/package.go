package quartermaster

import (
    "github.com/attamusc/quartermaster/resolver"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Package struct {
    Name    string `json:"name"`
    Repo    string `json:"repo"`
    Version string `json:"version"`
}

func (p *Package) Install() {
    r := resolver.GetResolverFor(p.Repo)
    success, err := r.Install(p.Repo)

    if !success || err != nil {
    }
}
