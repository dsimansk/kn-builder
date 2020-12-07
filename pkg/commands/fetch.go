package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "os/exec"
)

// TODO: it'd nice to have release list cmd, or interactive fetch to choose version
func NewFetchCmd() *cobra.Command {
    var version string
    var fetchCmd = &cobra.Command{
        Use:   "fetch",
        Short: "Fetch kn repository.",
        RunE: func(cmd *cobra.Command, args []string) error {
            out, err := exec.Command("git", "clone", "--branch", version, "https://github.com/knative/client.git").Output()
            if err != nil {
                return fmt.Errorf("git clone failed: %s", err)
            }
            fmt.Printf("%s\n", out)
            return nil
        },
    }
    fetchCmd.Flags().StringVar(&version, "version", "master", "Specify kn release version to fetch, e.g branch or tag.")
    fetchCmd.SilenceUsage = true
    return fetchCmd
}
