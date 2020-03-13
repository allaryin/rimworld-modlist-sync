package cmd

import (
	"fmt"
	"github.com/allaryin/rimworld-modlist-sync/pkg/rimworld"
	"github.com/spf13/cobra"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var savesCmd = &cobra.Command{
	Use:   "saves",
	Short: "Search save files",
	RunE:  runSaves,
}
var (
	savesMatch string
)

func init() {
	savesCmd.Flags().StringVar(&savesMatch, "match", "", "A short filename segment to match")
}

func runSaves(cmd *cobra.Command, args []string) (err error) {
	if saves, e := rimworld.ScanSaves(saveDir); e != nil {
		err = fmt.Errorf("unable to scan saves %w", e)
	} else {
		files := saves.ListAll()
		fmt.Printf("Found %d saves...\n", len(files))

		if savesMatch == "" {
			// no pattern to match, just dump
			for _, save := range files {
				fmt.Printf("> %s\n", save)
			}
		} else {
			// get fuzzy
			matches := fuzzy.FindNormalizedFold(savesMatch, files)
			fmt.Printf("Found %d that matched %q...\n", len(matches), savesMatch)
			for _, save := range matches {
				fmt.Printf("> %s\n", save)
			}
		}
	}

	return
}
