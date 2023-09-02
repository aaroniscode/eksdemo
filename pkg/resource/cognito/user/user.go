package user

import (
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/resource"
)

func New() *resource.Resource {
	options, getFlags := newOptions()

	return &resource.Resource{
		Command: cmd.Command{
			Name:        "user",
			Description: "Cognito User",
			// Args:        []string{"DOMAIN"},
		},

		GetFlags: getFlags,

		Getter: &Getter{},

		Options: options,
	}
}
