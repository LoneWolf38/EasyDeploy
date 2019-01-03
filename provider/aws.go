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




func CreateSecGroup(vpcId, secName,des string) {
	if len(vpcId) == 0{
		svc := CreateEc2Session(Region)
		vpcinfo, err = svc.DescribeVpcs(nil)
		if err!= nil {
			exitErrorf("Error in describing VPCs")
		}

		vpcId = aws.StringValue(vpcinfo.Vpcs[0].VpcId)	
		
	}
	



}


//Create A IAM user

//Create A EC2 instance

//Create A security group




//Create A VPC

//Create A Subnet

//Create A Route53 entry