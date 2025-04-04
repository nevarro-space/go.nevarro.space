package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"go.mau.fi/util/exerrors"
)

//go:embed repos.json
var reposJSON []byte

//go:embed vanity.tmpl
var vanityTemplate string

type RepoInfo struct {
	Owner      string `json:"owner"`
	Repo       string `json:"repo"`
	ImportPath string `json:"import_path"`
}

func main() {
	var repos []RepoInfo
	exerrors.PanicIfNotNil(json.NewDecoder(bytes.NewReader(reposJSON)).Decode(&repos))

	exerrors.PanicIfNotNil(os.MkdirAll("public", 0755))
	for _, info := range repos {
		templ := exerrors.Must(template.New("vanity").Parse(vanityTemplate))
		f := exerrors.Must(os.Create(fmt.Sprintf("public/%s.html", info.Repo)))
		exerrors.PanicIfNotNil(templ.Execute(f, info))
	}
}
