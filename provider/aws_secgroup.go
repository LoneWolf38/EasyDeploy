package provider

import (
		"github.com/aws/aws-sdk-go/aws"
		"github.com/aws/aws-sdk-go/aws/session"
		"github.com/aws/aws-sdk-go/service/ec2"
		"github.com/aws/aws-sdk-go/aws/awserr"
	    "github.com/LoneWolf38/EasyDeploy/src"
)



func (*src.EasyDeploy)CreateSecGroup(secName,secDes string, svc *ec2.EC2) string  {
	vpcInfo, err := svc.DescribeVpcs(nil)
	if err != nil{
		ExitErrorf("Error in describing Vpcs")
	}
	vpcID := aws.StringValue(vpcInfo.Vpcs[0].VpcId)
	secGroup, secerr := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName: aws.String(secName),
		Description: aws.String(des),
		VpcId: aws.String(vpcID)
	})
	if secerr != nil{
		if aerr, ok := err.(awserr.Error); ok{
			switch aerr.code() {
			case "InvalidVpcID.NotFound":
				ExitErrorf("Unable to find VPC with ID %q", vpcID)
			case "InvalidGroup.Duplicate":
				ExitErrorf("Security Group %q already exists", secName)
			}
		}
		ExitErrorf("Unable to create secruity group %q. %v",secName, secerr)
	}
	secGroupId := aws.String(secGroup.GroupId)
	_, gerr := svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
        GroupId: aws.String(secGrId),
        IpPermissions: []*ec2.IpPermission{
            (&ec2.IpPermission{}).
                SetIpProtocol("tcp").
                SetFromPort(80).
                SetToPort(80).
                SetIpRanges([]*ec2.IpRange{
                    {CidrIp: aws.String("0.0.0.0/0")},
                }),
            (&ec2.IpPermission{}).
                SetIpProtocol("tcp").
                SetFromPort(22).
                SetToPort(22).
                SetIpRanges([]*ec2.IpRange{
                    (&ec2.IpRange{}).
                        SetCidrIp("0.0.0.0/0"),
                }),
            (&ec2.IpPermission{}).
                SetIpProtocol("tcp").
                SetFromPort(443).
                SetToPort(443).
                SetIpRanges([]*ec2.IpRange{
                    (&ec2.IpRange{}).
                        SetCidrIp("0.0.0.0/0"),
                }),

        },
    })
    if gerr != nil {
        ExitErrorf("Unable to set security group %q ingress, %v", secName, err)
	}
	return secGroupId
}