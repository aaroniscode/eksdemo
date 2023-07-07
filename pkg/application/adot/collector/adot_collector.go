package collector

import (
	"github.com/awslabs/eksdemo/pkg/application"
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/installer"
	"github.com/awslabs/eksdemo/pkg/resource"
	"github.com/awslabs/eksdemo/pkg/resource/irsa"
	"github.com/awslabs/eksdemo/pkg/template"
)

// Docs:    https://opentelemetry.io/docs/collector/
// GitHub:  https://github.com/aws-observability/aws-otel-collector/
// GitHub:  https://github.com/open-telemetry/opentelemetry-collector
// Helm:    https://github.com/open-telemetry/opentelemetry-helm-charts/tree/main/charts/opentelemetry-collector
// Repo:    gallery.ecr.aws/aws-observability/aws-otel-collector
// Version: Latest is ADOT Collector v0.30.0 (as of 7/6/23)
//          Translates to OTEL Collector v0.78.2, OTEL Collector Chart 0.59.3

func NewApp() *application.Application {
	options, flags := newOptions()

	return &application.Application{
		Command: cmd.Command{
			Parent:      "adot",
			Name:        "collector",
			Description: "AWS Distro for OpenTelemetry (ADOT) Collector",
		},

		Dependencies: []*resource.Resource{
			irsa.NewResourceWithOptions(&irsa.IrsaOptions{
				CommonOptions: resource.CommonOptions{
					Name: "adot-collector-irsa",
				},
				PolicyType: irsa.PolicyARNs,
				Policy: []string{
					"arn:aws:iam::aws:policy/AmazonPrometheusRemoteWriteAccess",
					"arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess",
					"arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy",
				},
			}),
		},

		Flags: flags,

		Installer: &installer.HelmInstaller{
			ChartName:     "opentelemetry-collector",
			ReleaseName:   "adot-collector",
			RepositoryURL: "https://open-telemetry.github.io/opentelemetry-helm-charts",
			ValuesTemplate: &template.TextTemplate{
				Template: valuesTemplate,
			},
		},

		Options: options,
	}
}

// https://github.com/open-telemetry/opentelemetry-helm-charts/blob/main/charts/opentelemetry-collector/values.yaml
const valuesTemplate = `---
fullnameOverride: adot-collector
# Valid values are "daemonset", "deployment", and "statefulset".
mode: {{ .Mode }}
image:
  # repository: otel/opentelemetry-collector-contrib
  repository: public.ecr.aws/aws-observability/aws-otel-collector
  tag: {{ .Version }}
# only used with deployment mode
replicaCount: 1
`
