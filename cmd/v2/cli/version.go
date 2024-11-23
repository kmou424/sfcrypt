package cli

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/kmou424/ero"
	"github.com/kmou424/sfcrypt/app/buildinfo"
	"github.com/kmou424/sfcrypt/app/version"
	"github.com/spf13/cobra"
	"text/template"
)

//go:embed version.go.tmpl
var embedVersion embed.FS

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and build information",
	Long:  `Show version and build information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(genVersion())
	},
}

var genVersion = func() string {
	tmplInfo := map[string]any{
		"version":     version.GetVersion(),
		"buildDate":   buildinfo.BuildDate,
		"vcsRevision": buildinfo.VCSRevision,
		"goVersion":   buildinfo.GoVersion,
		"debug":       buildinfo.Debug,
	}
	tmplBytes, _ := embedVersion.ReadFile("version.go.tmpl")
	tmpl, err := template.New("version").Parse(string(tmplBytes))
	if err != nil {
		panic(ero.Wrap(err, "failed to parse version template"))
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, tmplInfo)
	if err != nil {
		panic(ero.Wrap(err, "failed to execute version template"))
	}
	return buf.String()
}
