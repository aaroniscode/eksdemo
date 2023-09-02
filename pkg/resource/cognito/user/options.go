package user

import (
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/eksdemo/pkg/aws"
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/resource"
	"github.com/awslabs/eksdemo/pkg/resource/cognito/userpool"
	"github.com/spf13/cobra"
)

type Options struct {
	resource.CommonOptions

	UserPoolID   string
	UserPoolName string
}

func newOptions() (options *Options, getFlags cmd.Flags) {
	options = &Options{
		CommonOptions: resource.CommonOptions{
			ClusterFlagDisabled: true,
		},
	}

	getFlags = cmd.Flags{
		&cmd.StringFlag{
			CommandFlag: cmd.CommandFlag{
				Name:        "user-pool-id",
				Description: "id of the user pool",
				Shorthand:   "I",
				Validate: func(_ *cobra.Command, _ []string) error {
					if options.UserPoolID == "" && options.UserPoolName == "" {
						return &cmd.MustIncludeEitherOrFlagError{Flag1: "--user-pool-id", Flag2: "--user-pool-name"}
					}
					return nil
				},
			},
			Option: &options.UserPoolID,
		},
		&cmd.StringFlag{
			CommandFlag: cmd.CommandFlag{
				Name:        "user-pool-name",
				Description: "name of the user pool",
				Shorthand:   "U",
				Validate: func(cmd *cobra.Command, args []string) error {
					if options.UserPoolName == "" {
						return nil
					}

					up, err := userpool.NewGetter(aws.NewCognitoUserPoolClient()).GetUserPoolByName(options.UserPoolName)
					if err != nil {
						return err
					}
					options.UserPoolID = awssdk.ToString(up.Id)
					return nil
				},
			},
			Option: &options.UserPoolName,
		},
	}

	return
}
