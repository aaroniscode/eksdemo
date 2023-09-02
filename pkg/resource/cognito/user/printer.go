package user

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/eksdemo/pkg/printer"
	"github.com/hako/durafmt"
)

type Printer struct {
	users []types.UserType
}

func NewPrinter(users []types.UserType) *Printer {
	return &Printer{users}
}

func (p *Printer) PrintTable(writer io.Writer) error {
	table := printer.NewTablePrinter()
	table.SetHeader([]string{"Age", "Name", "Email"})

	for _, u := range p.users {
		age := durafmt.ParseShort(time.Since(aws.ToTime(u.UserCreateDate)))

		table.AppendRow([]string{
			age.String(),
			aws.ToString(u.Username),
			p.getAttribute(u, "email"),
		})
	}

	table.Print(writer)

	return nil
}

func (p *Printer) PrintJSON(writer io.Writer) error {
	return printer.EncodeJSON(writer, p.users)
}

func (p *Printer) PrintYAML(writer io.Writer) error {
	return printer.EncodeYAML(writer, p.users)
}

func (p *Printer) getAttribute(user types.UserType, attribute string) string {
	for _, a := range user.Attributes {
		if aws.ToString(a.Name) == attribute {
			return aws.ToString(a.Value)
		}
	}
	return "-"
}
