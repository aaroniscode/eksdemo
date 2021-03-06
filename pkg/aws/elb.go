package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func ELBDescribeLoadBalancersv1(name string) ([]*elb.LoadBalancerDescription, error) {
	sess := GetSession()
	svc := elb.New(sess)

	elbs := []*elb.LoadBalancerDescription{}
	input := &elb.DescribeLoadBalancersInput{}
	pageNum := 0

	if name != "" {
		input.LoadBalancerNames = aws.StringSlice([]string{name})
	}

	err := svc.DescribeLoadBalancersPages(input,
		func(page *elb.DescribeLoadBalancersOutput, lastPage bool) bool {
			pageNum++
			elbs = append(elbs, page.LoadBalancerDescriptions...)
			return pageNum <= maxPages
		},
	)

	return elbs, err
}

func ELBDescribeLoadBalancersv2(name string) ([]*elbv2.LoadBalancer, error) {
	sess := GetSession()
	svc := elbv2.New(sess)

	elbs := []*elbv2.LoadBalancer{}
	input := &elbv2.DescribeLoadBalancersInput{}
	pageNum := 0

	if name != "" {
		input.Names = aws.StringSlice([]string{name})
	}

	err := svc.DescribeLoadBalancersPages(input,
		func(page *elbv2.DescribeLoadBalancersOutput, lastPage bool) bool {
			pageNum++
			elbs = append(elbs, page.LoadBalancers...)
			return pageNum <= maxPages
		},
	)

	return elbs, err
}
