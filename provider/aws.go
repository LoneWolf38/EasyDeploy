package provider

import (
		"fmt"
        "os"
        "time"
		"github.com/aws/aws-sdk-go/aws"
		"github.com/aws/aws-sdk-go/aws/session"
		"github.com/aws/aws-sdk-go/service/ec2"
		"github.com/aws/aws-sdk-go/aws/awserr"
		//"github.com/aws/aws-sdk-go/aws/awsutil"
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

func CreateSecGroup(secName,des string, svc *ec2.EC2) string{

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
    return secGrId
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

func SubnetDetails(svc *ec2.EC2) string{
    result, err := svc.DescribeSubnets(nil)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
                default:
                    fmt.Println(aerr.Error())
                }
        } else {
            fmt.Println(err.Error())
           }
        return "error in finding subnets"
    }
     subnets := aws.StringValue(result.Subnets[0].SubnetId)
     return subnets
}

func VpcDetails(svc *ec2.EC2) string{
    vpcinfo, eerr := svc.DescribeVpcs(nil)
        if eerr!= nil {
            exitErrorf("Error in describing VPCs")
        }
        vpcId := aws.StringValue(vpcinfo.Vpcs[0].VpcId) 
      return vpcId  
}

func CreateOneInstance(subnetid, tags, secgroup, instancetype, ami, keyname string, svc *ec2.EC2) string{
    runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
        ImageId:      aws.String(ami),
        InstanceType: aws.String(instancetype),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
        KeyName:      aws.String(keyname),
        SecurityGroupIds: []*string{
        aws.String(secgroup),
        },
        SubnetId: aws.String(subnetid),
    })

    if err != nil {
        fmt.Println(err)
    }
    _, errtag := svc.CreateTags(&ec2.CreateTagsInput{
        Resources: []*string{runResult.Instances[0].InstanceId},
        Tags: []*ec2.Tag{
            {
                Key:   aws.String("Name"),
                Value: aws.String(tags),
            },
        },
    })
    if errtag != nil {
        fmt.Println(errtag)
    }
    instanceId = aws.StringValue(runResult.Instances[0].InstanceId) 
    fmt.Println(aws.StringValue(runResult.Instances[0].InstanceId))

    input := &ec2.DescribeInstanceStatusInput{
    InstanceIds: []*string{
        aws.String(instanceId),
    },
}

 for{
        result, err := svc.DescribeInstanceStatus(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            default:
                fmt.Println(aerr.Error())
            }
        } else {
         fmt.Println(err.Error())
        }
        return "error"
     }
        if aws.Int64Value(result.InstanceStatuses[0].InstanceState.Code) == 16{
            break
         }else{
            fmt.Println("Creating Instance....")
            time.Sleep(1000 * time.Second)
            continue
          }
        }
    }



//Create A IAM user

//Create A EC2 instance

//Create A security group




//Create A VPC

//Create A Subnet

//Create A Route53 entry