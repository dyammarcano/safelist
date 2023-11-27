package main

import (
	"context"
	"github.com/dyammarcano/module-template-go/internal/metadata"
	"github.com/spf13/cobra"
)

var (
	Version    = "v0.0.1-manual-build"
	CommitHash string
	Date       string
)

var rootCmd = &cobra.Command{
	Use:   "template-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := rootCmd.ExecuteContext(ctx)
	cobra.CheckErr(err)
}

func init() {
	metadata.Set(Version, CommitHash, Date)
}
