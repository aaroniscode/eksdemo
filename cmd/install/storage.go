package install

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/application/storage/ebs_csi"
	"eksdemo/pkg/application/storage/efs_csi"
	"eksdemo/pkg/application/storage/fsx_lustre_csi"
	"eksdemo/pkg/application/storage/openebs"

	"github.com/spf13/cobra"
)

var storageApps []func() *application.Application

func NewInstallStorageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "storage",
		Short: "Kubernetes Storage Solutions",
	}

	// Don't show flag errors for `install storage` without a subcommand
	cmd.DisableFlagParsing = true

	for _, a := range storageApps {
		cmd.AddCommand(a().NewInstallCmd())
	}

	return cmd
}

func NewUninstallStorageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "storage",
		Short: "Kubernetes Storage Solutions",
	}

	// Don't show flag errors for `uninstall storage` without a subcommand
	cmd.DisableFlagParsing = true

	for _, a := range storageApps {
		cmd.AddCommand(a().NewUninstallCmd())
	}

	return cmd
}

func init() {
	storageApps = []func() *application.Application{
		ebs_csi.NewApp,
		efs_csi.NewApp,
		fsx_lustre_csi.NewApp,
		openebs.NewApp,
	}
}
