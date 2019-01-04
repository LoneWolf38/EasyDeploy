package provider

import (
		"fmt"
		"github.com/aws/aws-sdk-go/aws"
		"github.com/aws/aws-sdk-go/aws/session"
		"github.com/aws/aws-sdk-go/service/ec2"
		"github.com/aws/aws-sdk-go/aws/awserr"
		"github.com/aws/aws-sdk-go/aws/awsutil"
		"github.com/LoneWolf38/EasyDeploy/cmd"
)

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




func CreateSecGroup(GvpcId, secName,des string) {
	if len(GvpcId) == 0{
		svc := CreateEc2Session(Region)
		vpcinfo, err = svc.DescribeVpcs(nil)
		if err!= nil {
			exitErrorf("Error in describing VPCs")
		}

		vpcId = aws.StringValue(vpcinfo.Vpcs[0].VpcId)	
		secgr, err := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
			GroupName: aws.String(secName),
			Description: aws.String(des),
			VpcId: aws.String(vpcId),
			})
		if err != nil {
    		if aerr, ok := err.(awserr.Error); ok {
        	switch aerr.Code() {
        		case "InvalidVpcID.NotFound":
            		exitErrorf("Unable to find VPC with ID %q.", vpcID)
        		case "InvalidGroup.Duplicate":
            		exitErrorf("Security group %q already exists.", name)
        		}
    		}
    	exitErrorf("Unable to create security group %q, %v", name, err)
		}
	}else{
		secgr, err := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
			GroupName: aws.String(secName),
			Description: aws.String(des),
			VpcId: aws.String(GvpcId),
			})
		if err != nil {
    		if aerr, ok := err.(awserr.Error); ok {
        	switch aerr.Code() {
        		case "InvalidVpcID.NotFound":
            		exitErrorf("Unable to find VPC with ID %q.", vpcID)
        		case "InvalidGroup.Duplicate":
            		exitErrorf("Security group %q already exists.", name)
        		}
    		}
    	exitErrorf("Unable to create security group %q, %v", name, err)
		}
	}
	secGrId := aws.StringValue(secgr.GroupId)
	 _, err := svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
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
    if err != nil {
        exitErrorf("Unable to set security group %q ingress, %v", name, err)
    }
    _,rerr := svc.AuthorizeSecurityGroupEgress(&ec2.AuthorizeSecurityGroupEgressInput{
    	GroupId: aws.String(secGrId),
    	IpPermission: []*ec2.IpPermission{
    		(&ec2.IpPermission{}).
    		SetIpProtocol("tcp").
    		SetFromPort(0).
    		SetToPort(65535).
    		SetIpRanges([]*ec2.IpRange{
    			{CidrIp: aws.String("0.0.0.0/0")},
    			}),
    	},
    	})
    if rerr != nil {
        exitErrorf("Unable to set security group %q ingress, %v", name, err)
    }

}



//Create A IAM user

//Create A EC2 instance

//Create A security group




//Create A VPC

//Create A Subnet

//Create A Route53 entry