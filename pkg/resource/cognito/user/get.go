package user

import (
	"fmt"
	"os"

	"github.com/awslabs/eksdemo/pkg/aws"
	"github.com/awslabs/eksdemo/pkg/printer"
	"github.com/awslabs/eksdemo/pkg/resource"
)

type Getter struct {
	cognitoClient *aws.CognitoUserPoolClient
}

func NewGetter(cognitoClient *aws.CognitoUserPoolClient) *Getter {
	return &Getter{cognitoClient}
}

func (g *Getter) Init() {
	if g.cognitoClient == nil {
		g.cognitoClient = aws.NewCognitoUserPoolClient()
	}
}

func (g *Getter) Get(domain string, output printer.Output, o resource.Options) error {
	options, ok := o.(*Options)
	if !ok {
		return fmt.Errorf("internal error, unable to cast options to client.Options")
	}

	users, err := g.cognitoClient.ListUsers(options.UserPoolID)

	if err != nil {
		return err
	}

	return output.Print(os.Stdout, NewPrinter(users))
}
