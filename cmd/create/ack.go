package create

import (
	"eksdemo/pkg/resource"
	"eksdemo/pkg/resource/ack/ec2"
	"eksdemo/pkg/resource/ack/ecr"
	"eksdemo/pkg/resource/ack/eks"
	"eksdemo/pkg/resource/ack/s3"

	"github.com/spf13/cobra"
)

var ack []func() *resource.Resource

func NewAckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ack",
		Short: "AWS Controllers for Kubernetes (ACK)",
	}

	// Don't show flag errors for `create ack`` without a subcommand
	cmd.DisableFlagParsing = true

	for _, r := range ack {
		cmd.AddCommand(r().NewCreateCmd())
	}

	return cmd
}

func init() {
	ack = []func() *resource.Resource{
		ec2.NewSecurityGroupResource,
		ec2.NewSubnetResource,
		ec2.NewVpcResource,
		ecr.NewResource,
		eks.NewFargateProfileResource,
		s3.NewResource,
	}
}
