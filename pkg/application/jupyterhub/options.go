package jupyterhub

import (
	"github.com/awslabs/eksdemo/pkg/application"
	"github.com/awslabs/eksdemo/pkg/cmd"
)

type Options struct {
	application.ApplicationOptions

	AdminPassword string
	AllowSudo     bool
}

func newOptions() (options *Options, flags cmd.Flags) {
	options = &Options{
		ApplicationOptions: application.ApplicationOptions{
			ExposeIngressOnly: true,
			Namespace:         "jupyterhub",
			ServiceAccount:    "hub",
			DefaultVersion: &application.LatestPrevious{
				// JupyterHub image tag must chart version
				LatestChart:   "3.1.0",
				Latest:        "3.1.0",
				PreviousChart: "3.1.0",
				Previous:      "3.1.0",
			},
		},
	}

	flags = cmd.Flags{
		&cmd.StringFlag{
			CommandFlag: cmd.CommandFlag{
				Name:        "admin-pass",
				Description: "Admin password for JupyterHub",
				Required:    true,
				Shorthand:   "P",
			},
			Option: &options.AdminPassword,
		},
		&cmd.BoolFlag{
			CommandFlag: cmd.CommandFlag{
				Name:        "allow-sudo",
				Description: "allow sudo in the Juypter Notebook container",
			},
			Option: &options.AllowSudo,
		},
	}

	return
}
