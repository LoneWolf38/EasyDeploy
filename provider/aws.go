package provider

import (
		"fmt"
		"github.com/aws/aws-sdk-go/aws"
		"github.com/aws/aws-sdk-go/aws/session"
		"github.com/aws/aws-sdk-go/service/ec2"
		"github.com/aws/aws-sdk-go/aws/awserr"
		"github.com/aws/aws-sdk-go/aws/awsutil"
)

var SECID string

func CreateSession(region string) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{Region: aws.String(region)},
		}))
}

func CreateEc2Session(region string) *ec2.EC2 {
	return ec2.New(session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{Region: aws.String(region)},
		})))
}

func CreateKey(keyname string, svc *ec2.EC2) string{
	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyname),
	}
	keypair, err := svc.CreateKeyPair(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok{
			switch aerr.Code(){
				default: fmt.Println(aerr.Error())
			}
		}else{
			fmt.Println(err.Error())
		}
	}
	return string(*keypair.KeyMaterial)
}

func CreateSecGroup(secName,des string, svc *ec2.EC2) {

        fmt.Println("Creating a Security Group for Website Development")
        vpcinfo, eerr := svc.DescribeVpcs(nil)
        if eerr!= nil {
            exitErrorf("Error in describing VPCs")
        }

        vpcId := aws.StringValue(vpcinfo.Vpcs[0].VpcId) 
        secgr, err := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
            GroupName: aws.String(secName),
            Description: aws.String(des),
            VpcId: aws.String(vpcId),
            })
        if err != nil {
            if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
                case "InvalidVpcID.NotFound":
                    exitErrorf("Unable to find VPC with ID %q.", vpcId)
                case "InvalidGroup.Duplicate":
                    exitErrorf("Security group %q already exists.", secName)
                }
            }
        exitErrorf("Unable to create security group %q, %v", secName, err)
        }
    secGrId := aws.StringValue(secgr.GroupId)
    SECID = secGrId
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
        },
    })
    if gerr != nil {
        exitErrorf("Unable to set security group %q ingress, %v", secName, err)
    }
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

//Create A IAM user

//Create A EC2 instance

//Create A security group




//Create A VPC

//Create A Subnet

//Create A Route53 entry