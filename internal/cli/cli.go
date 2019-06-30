package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dictyBase-docker/slack-notify/internal/logger"
	"github.com/dictyBase-docker/slack-notify/internal/registry"
	msg "github.com/dictyBase-docker/slack-notify/internal/slack"
	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "slack-notify",
	Short: "cli to send argo workflow notification in a slack channel",
	Long: `A command line application that is used as an onExit handler
	to receive information from a stepwise argo(https://github.com/argoproj/argo)
	workflow. It then send a message to a given slack channel including the 
	workflow status and link to the argo web ui.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		l, err := logger.NewLogger(cmd)
		if err != nil {
			return err
		}
		registry.SetLogger(l)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		d, _ := cmd.Flags().GetBool("doc")
		if d {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			docDir := filepath.Join(dir, "docs")
			if err := os.MkdirAll(docDir, 0700); err != nil {
				return err
			}
			if err := doc.GenMarkdownTree(cmd, docDir); err != nil {
				return err
			}
			fmt.Printf("created markdown docs in %s\n", docDir)
			return nil
		}
		l := registry.GetLogger()
		t, _ := cmd.Flags().GetString("api-token")
		e, _ := cmd.Flags().GetString("endpoint")
		id, _ := cmd.Flags().GetString("workflow-id")
		st, _ := cmd.Flags().GetString("status")
		ch, _ := cmd.Flags().GetString("channel")
		client := slack.New(t)
		emoji := ":heavy_check_mark:"
		if st != "Succeeded" {
			emoji = ":X:"
		}
		b := msg.MsgBlock(&msg.BlockData{
			Host:   e,
			Id:     id,
			Status: st,
			Emoji:  emoji,
		})
		_, _, _, err := client.SendMessage(
			ch,
			slack.MsgOptionBlocks(b),
		)
		if err != nil {
			return err
		}
		l.Infof("posted message status %s", st)
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().Bool("doc", false, "generate markdown documentation")
	RootCmd.Flags().StringP(
		"endpoint",
		"e",
		"",
		"argo-ui http endpoint that will be used for visualizing workflow output (required)!!!!",
	)
	RootCmd.Flags().StringP(
		"workflow-id",
		"w",
		"",
		"a workflow identifier that has triggered this slack notification",
	)
	RootCmd.Flags().StringP(
		"status",
		"s",
		"",
		"the status of workflow, will be used to pick up an emoji for displaying in slack",
	)
	RootCmd.Flags().StringP(
		"api-token",
		"t",
		"",
		"slack api token",
	)
	RootCmd.Flags().StringP(
		"channel",
		"c",
		"",
		"slack channel where the message will be posted",
	)
	RootCmd.MarkFlagRequired("endpoint")
	RootCmd.MarkFlagRequired("workflow-id")
	RootCmd.MarkFlagRequired("status")
	RootCmd.MarkFlagRequired("api-token")
	RootCmd.MarkFlagRequired("channel")
	RootCmd.Flags().StringP(
		"log-level",
		"",
		"error",
		"log level for the application",
	)
	RootCmd.Flags().StringP(
		"log-format",
		"",
		"json",
		"format of the logging out, either of json or text",
	)
	RootCmd.Flags().String(
		"log-file",
		"",
		"file for log output other than standard output, written to a temp folder by default",
	)
}
