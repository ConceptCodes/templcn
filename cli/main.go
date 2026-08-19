package main

import (
	"os"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

var cliVersion string

func version() string {
	if cliVersion != "" {
		return strings.TrimPrefix(cliVersion, "v")
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return strings.TrimPrefix(info.Main.Version, "v")
	}
	return "devel"
}

func main() {
	if err := newRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "templcn",
		Short:         "templcn CLI for Go/templ components",
		Version:       version(),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(newInitCommand())
	cmd.AddCommand(newAddCommand())
	cmd.AddCommand(newApplyCommand())
	cmd.AddCommand(newViewCommand())
	cmd.AddCommand(newDiffCommand())
	cmd.AddCommand(newSearchCommand())
	cmd.AddCommand(newBuildCommand())
	cmd.AddCommand(newDocsCommand())
	cmd.AddCommand(newInfoCommand())
	cmd.AddCommand(newParityCommand())

	return cmd
}

func newParityCommand() *cobra.Command {
	opts := ParityOptions{}
	cmd := &cobra.Command{
		Use:   "parity",
		Short: "compare local components with upstream shadcn/ui",
		RunE: func(cmd *cobra.Command, args []string) error {
			return CheckParityCommand(opts)
		},
	}

	cmd.Flags().StringVar(&opts.Source, "source", "", "upstream registry index URL or file")
	cmd.Flags().BoolVar(&opts.JSON, "json", false, "output as JSON")
	return cmd
}

func newDiffCommand() *cobra.Command {
	opts := DiffOptions{}
	cmd := &cobra.Command{
		Use:   "diff <items...>",
		Short: "show component source differences",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Items = args
			return DiffItems(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	return cmd
}

func newInitCommand() *cobra.Command {
	opts := InitOptions{}
	cmd := &cobra.Command{
		Use:     "init [components...]",
		Aliases: []string{"create"},
		Short:   "scaffold a Go/templ project",
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Items = args
			return InitProject(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	cmd.Flags().StringVarP(&opts.Template, "template", "t", "next", "template to scaffold")
	cmd.Flags().StringVarP(&opts.Base, "base", "b", "radix", "component base")
	cmd.Flags().StringVarP(&opts.Preset, "preset", "p", "nova", "preset to apply")
	cmd.Flags().BoolVarP(&opts.Yes, "yes", "y", true, "skip confirmation prompts")
	cmd.Flags().BoolVarP(&opts.Defaults, "defaults", "d", false, "use default configuration")
	cmd.Flags().BoolVarP(&opts.Force, "force", "f", false, "overwrite existing files")
	cmd.Flags().BoolVarP(&opts.Silent, "silent", "s", false, "mute output")
	cmd.Flags().BoolVar(&opts.Monorepo, "monorepo", false, "scaffold a monorepo")
	cmd.Flags().BoolVar(&opts.RTL, "rtl", false, "enable RTL support")
	cmd.Flags().BoolVar(&opts.Reinstall, "reinstall", false, "reinstall existing components")
	return cmd
}

func newAddCommand() *cobra.Command {
	opts := AddOptions{}
	cmd := &cobra.Command{
		Use:   "add [components...]",
		Short: "sync templcn/ui source into your project",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Items = args
			return AddComponents(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().StringVarP(&opts.Path, "path", "p", "", "target ui package path")
	cmd.Flags().BoolVarP(&opts.Yes, "yes", "y", false, "skip confirmation prompts")
	cmd.Flags().BoolVarP(&opts.Overwrite, "overwrite", "o", false, "overwrite existing files")
	cmd.Flags().BoolVarP(&opts.All, "all", "a", false, "sync every available component")
	cmd.Flags().BoolVarP(&opts.Silent, "silent", "s", false, "mute output")
	cmd.Flags().BoolVar(&opts.DryRun, "dry-run", false, "preview changes without writing")
	cmd.Flags().StringVar(&opts.Diff, "diff", "", "show the diff for a file")
	cmd.Flags().StringVar(&opts.View, "view", "", "show file contents")
	return cmd
}

func newApplyCommand() *cobra.Command {
	opts := ApplyOptions{}
	cmd := &cobra.Command{
		Use:   "apply [preset]",
		Short: "apply a preset to an existing project",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.Preset = args[0]
			}
			return ApplyPreset(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().StringVarP(&opts.Preset, "preset", "p", "", "preset to apply")
	cmd.Flags().StringSliceVar(&opts.Only, "only", nil, "apply only parts of a preset: theme, font")
	cmd.Flags().BoolVarP(&opts.Yes, "yes", "y", false, "skip confirmation prompts")
	cmd.Flags().BoolVarP(&opts.Silent, "silent", "s", false, "mute output")
	return cmd
}

func newViewCommand() *cobra.Command {
	opts := ViewOptions{}
	cmd := &cobra.Command{
		Use:   "view <items...>",
		Short: "view a component source file",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Items = args
			return ViewItems(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	return cmd
}

func newSearchCommand() *cobra.Command {
	opts := SearchOptions{}
	cmd := &cobra.Command{
		Use:     "search <registries...>",
		Aliases: []string{"list"},
		Short:   "search available components",
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Registries = args
			return SearchComponents(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().StringVarP(&opts.Query, "query", "q", "", "search query")
	cmd.Flags().IntVarP(&opts.Limit, "limit", "l", 100, "maximum results")
	cmd.Flags().IntVarP(&opts.Offset, "offset", "o", 0, "items to skip")
	return cmd
}

func newBuildCommand() *cobra.Command {
	opts := BuildOptions{}
	cmd := &cobra.Command{
		Use:   "build [registry]",
		Short: "generate a registry manifest",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.Registry = args[0]
			}
			return BuildRegistry(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().StringVarP(&opts.OutputDir, "output", "o", "./public/r", "output directory")
	return cmd
}

func newDocsCommand() *cobra.Command {
	opts := DocsOptions{}
	cmd := &cobra.Command{
		Use:   "docs [component]",
		Short: "show component docs metadata",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.Component = args[0]
			}
			return ShowDocs(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().StringVarP(&opts.Base, "base", "b", "project", "base to use")
	cmd.Flags().BoolVar(&opts.JSON, "json", false, "output as JSON")
	return cmd
}

func newInfoCommand() *cobra.Command {
	opts := InfoOptions{}
	cmd := &cobra.Command{
		Use:   "info",
		Short: "inspect the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return PrintInfo(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CWD, "cwd", "c", ".", "working directory")
	cmd.Flags().BoolVar(&opts.JSON, "json", false, "output as JSON")
	return cmd
}
