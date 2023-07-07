package collector

import (
	"github.com/awslabs/eksdemo/pkg/application"
	"github.com/awslabs/eksdemo/pkg/cmd"
)

type Options struct {
	application.ApplicationOptions

	Mode string
}

func newOptions() (options *Options, flags cmd.Flags) {
	options = &Options{
		ApplicationOptions: application.ApplicationOptions{
			DefaultVersion: &application.LatestPrevious{
				LatestChart:   "0.59.3",
				Latest:        "v0.30.0",
				PreviousChart: "0.59.3",
				Previous:      "v0.30.0",
			},
			Namespace:      "adot-system",
			ServiceAccount: "adot-collector",
		},
		Mode: "deployment",
	}

	return
}
