package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/bulletRush/grizzly/pkg/grafana"
	"github.com/bulletRush/grizzly/pkg/grizzly"
	"github.com/go-clix/cli"
)

// Version is the current version of the grr command.
// To be overwritten at build time
var Version = "dev"

func Main() {

	rootCmd := &cli.Command{
		Use:     "grr",
		Short:   "Grizzly",
		Version: Version,
	}

	grizzly.ConfigureProviderRegistry(
		[]grizzly.Provider{
			&grafana.Provider{},
		})

	// workflow commands
	rootCmd.AddCommand(
		getCmd(),
		listCmd(),
		pullCmd(),
		showCmd(),
		diffCmd(),
		applyCmd(),
		watchCmd(),
		exportCmd(),
		previewCmd(),
		providersCmd(),
	)

	// Run!
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}