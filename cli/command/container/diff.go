package container

import (
	"context"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/completion"
	"github.com/docker/cli/cli/command/formatter"
	"github.com/spf13/cobra"
)

// NewDiffCommand creates a new cobra.Command for `docker diff`
func NewDiffCommand(dockerCli command.Cli) *cobra.Command {
	return &cobra.Command{
		Use:   "diff CONTAINER",
		Short: "Inspect changes to files or directories on a container's filesystem",
		Args:  cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDiff(cmd.Context(), dockerCli, args[0])
		},
		Annotations: map[string]string{
			"aliases": "docker container diff, docker diff",
		},
		ValidArgsFunction: completion.ContainerNames(dockerCli, false),
	}
}

func runDiff(ctx context.Context, dockerCLI command.Cli, containerID string) error {
	changes, err := dockerCLI.Client().ContainerDiff(ctx, containerID)
	if err != nil {
		return err
	}
	diffCtx := formatter.Context{
		Output: dockerCLI.Out(),
		Format: newDiffFormat("{{.Type}} {{.Path}}"),
	}
	return diffFormatWrite(diffCtx, changes)
}
