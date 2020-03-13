package cmd

import (
	"fmt"
	"github.com/allaryin/rimworld-modlist-sync/pkg/rimworld/xml"
	"os"
	"path/filepath"
	"runtime"

	"github.com/allaryin/rimworld-modlist-sync/pkg/util"
	"github.com/spf13/cobra"
)

var (
	// command parameters
	rwDir string

	// and default values
	defaultRwDir string

	// and laziness, oh my - these are derived and exported by the pre-run validation pass
	saveDir string
	configDir string
	configFile string

	// the actual ModsConfig.xml data itself
	modsConfig *xml.ModsConfigData
)

func init() {
	home := util.GetHomeDir()

	// figure out our default data directory
	switch runtime.GOOS {
	case "windows":
		// C:\Users\%username%\AppData\LocalLow\Ludeon Studios\RimWorld by Ludeon Studios\
		defaultRwDir = filepath.Join(home, "AppData", "LocalLow", "Ludeon Studios", "RimWorld by Ludeon Studios")
	case "darwin":
		// ~/Library/Application Support/RimWorld/
		defaultRwDir = filepath.Join(home, "Library", "Application Support", "RimWorld")
	default:
		// ~/.config/unity3d/Ludeon Studios/RimWorld by Ludeon Studios/
		defaultRwDir = filepath.Join(home, ".config", "unity3d", "Ludeon Studios", "RimWorld by Ludeon Studios")
	}
}

// perform some directory structure validation before we get too carried away
func preRun(cmd *cobra.Command, args []string) (err error) {
	// make sure the directory exists
	if !util.FileExists(rwDir) {
		return fmt.Errorf("could not find %q", rwDir)
	} else if !util.IsDir(rwDir) {
		return fmt.Errorf("%q is a file, not a directory", rwDir)
	} else {
		fmt.Printf("Found directory at %s...\n", rwDir)
	}

	// make sure config dir exists before going further
	configDir = filepath.Join(rwDir, "Config")
	if !util.FileExists(configDir) || !util.IsDir(configDir){
		// there is no config dir, we are probably pointed at the wrong place, bail
		return fmt.Errorf("could not find config dir at %q", configDir)
	} else {
		fmt.Printf("Found config dir...\n")
	}

	// and verify that a config file exists and is writable
	configFile = filepath.Join(configDir, xml.ModsConfigFilename)
	if !util.FileExists(configFile) {
		// we don't have a config - they've probably never run a compatible build of rimworld with this config dir
		fmt.Printf("!! Could not find mods config, have you launched RimWorld?\n")
		return fmt.Errorf("could not find config file at %q", configFile)
	}

	// make sure we have save dir - creating if necessary
	saveDir = filepath.Join(rwDir, "Saves")
	if !util.FileExists(saveDir) {
		if err = os.Mkdir(saveDir, 0755); err != nil {
			return fmt.Errorf("unable to create missing save dir %w", err)
		} else {
			fmt.Printf("Created missing save dir...\n")
		}
	} else if !util.IsDir(saveDir) {
		return fmt.Errorf("could not use %q as save dir", saveDir)
	} else {
		fmt.Printf("Found save dir...\n")
	}

	// and now that everything else is known okay - let's read in the current config file
	modsConfig, err = xml.LoadModsConfig(configFile)
	if err == nil {
		fmt.Printf("Loaded config xml...\n")
	}

	return err
}

func runRoot(cmd *cobra.Command, args []string) (err error) {
	return cmd.Help()
}

func Execute() {
	rootCmd := &cobra.Command{
		Use: "rwms",
		Short: "Rimworld Modlist Sync updates your modlist to match your save",
		PersistentPreRunE: preRun,
		RunE: runRoot,
	}

	rootCmd.PersistentFlags().StringVar(&rwDir, "dir", defaultRwDir, "data directory")

	rootCmd.AddCommand(savesCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}