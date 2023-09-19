package jupyterhub

import (
	"github.com/awslabs/eksdemo/pkg/application"
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/installer"
	"github.com/awslabs/eksdemo/pkg/template"
)

// Docs:     https://jupyterhub.readthedocs.io/
// K8s Docs: https://z2jh.jupyter.org/
// GitHub:   https://github.com/jupyterhub/jupyterhub
// Helm:     https://github.com/jupyterhub/zero-to-jupyterhub-k8s/tree/main/jupyterhub
// Repo:     https://hub.docker.com/r/jupyterhub/k8s-hub
// Version: Latest is Chart 3.1.0, App v4.0.2 (as of 10/28/23)

func NewApp() *application.Application {
	options, flags := newOptions()

	return &application.Application{
		Command: cmd.Command{
			Parent:      "ai",
			Name:        "jupyterhub",
			Description: "Multi-user server for Jupyter notebooks",
			Aliases:     []string{"jupyter"},
		},

		Flags: flags,

		Installer: &installer.HelmInstaller{
			ChartName:     "jupyterhub",
			ReleaseName:   "jupyterhub",
			RepositoryURL: "https://jupyterhub.github.io/helm-chart",
			ValuesTemplate: &template.TextTemplate{
				Template: valuesTemplate,
			},
			PVCLabels: map[string]string{
				"release": "jupyterhub",
			},
		},

		Options: options,
	}
}

// https://z2jh.jupyter.org/en/latest/resources/reference.html
const valuesTemplate = `---
hub:
  config:
    Authenticator:
      admin_users:
      - admin
    DummyAuthenticator:
      password: {{ .AdminPassword }}
    JupyterHub:
      admin_access: true
      authenticator_class: dummy
  service:
    type: {{ .ServiceType }}
    annotations:
      {{- .ServiceAnnotations | nindent 6 }}
  image:
    tag: {{ .Version }}
  # Network Policy needs to be disabled or it will cause the Hub healthcheck to fail, root cause currently unknown
  networkPolicy:
    enabled: false
  serviceAccount:
    name: {{ .ServiceAccount }}
proxy:
  networkPolicy:
    enabled: false
singleuser:
{{- if .AllowSudo }}
  # allow 'sudo' in the JupyterLab containers
  allowPrivilegeEscalation: true
  cmd: start-singleuser.sh
  extraEnv:
    GRANT_SUDO: "yes"
  uid: 0
{{- end}}
  image:
    name: jupyter/scipy-notebook
    tag: python-3.10
  networkPolicy:
    enabled: false
  #extraResource:
  #  limits:
  #    aws.amazon.com/neuron: 12
  profileList:
    - display_name: "Scientific Python Stack"
      description: "Foo"
      default: true
    - display_name: "PyTorch 1.13.1 (torch-neuronx) 2.13.2"
      description: "PyTorch with support for Inf2"
      kubespawner_override:
        extra_resource_guarantees:
          aws.amazon.com/neuron: 1
        image: 763104351884.dkr.ecr.us-west-2.amazonaws.com/pytorch-inference-neuronx:1.13.1-neuronx-py310-sdk2.13.2-ubuntu20.04
    - display_name: "Learning Data Science"
      description: "Datascience Environment with Sample Notebooks"
      kubespawner_override:
        image: jupyter/datascience-notebook:2343e33dec46
        lifecycle_hooks:
          postStart:
            exec:
              command:
                - "sh"
                - "-c"
                - >
                  gitpuller https://github.com/data-8/materials-fa17 master materials-fa;
  storage:
    capacity: 100Gi
prePuller:
  hook:
    enabled: false
ingress:
  enabled: true
  annotations:
    {{- .IngressAnnotations | nindent 4 }}
  ingressClassName: {{ .IngressClass }}
  hosts:
  - {{ .IngressHost }}
  pathType: Prefix
  tls:
  - hosts:
    - {{ .IngressHost }}
`
