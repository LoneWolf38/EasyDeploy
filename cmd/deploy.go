package cmd

import (
		"fmt"
		"os"
		"time"
		"github.com/spf13/cobra"
		"github.com/LoneWolf38/EasyDeploy/provider"
		"github.com/aws/aws-sdk-go/aws"
		"github.com/aws/aws-sdk-go/service/ec2"
		"github.com/LoneWolf38/EasyDeploy/provisioner"
		"github.com/aws/aws-sdk-go/aws/awserr"
)

var HOME = os.Getenv("HOME")

var URL string


var DeployAppCmd = &cobra.Command{
	Use: "deploy",
	Short: "To Deploy the App",
	Run: deploy,
}

func deploy(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(CPath); os.IsNotExist(err) {
		fmt.Println("No Config File found...")
		fmt.Println("Creating a New one......")
		ConfigInit()
		os.Exit(1)	
	}else{
		ExecuteDeploy()
	} 
}

func init() {
	DeployAppCmd.PersistentFlags().StringVar(&secName,"firewall",secName,"Name of the security Group")
	DeployAppCmd.PersistentFlags().StringVar(&repo, "repo", repo, "Name of the project")
	DeployAppCmd.PersistentFlags().StringVar(&URL, "url","", "Github URL of the project")
}



func ExecuteDeploy() {
	readConfig.SetConfigFile(CPath)
	readConfig.ReadInConfig()
	fmt.Println("Setting up Environment variables....")
	os.Setenv("AWS_ACCESS_KEY_ID",readConfig.GetString("aws.access_key"))
	os.Setenv("AWS_SECRET_ACCESS_KEY",readConfig.GetString("aws.secret_key"))
	fmt.Println("Collecting Info...")
	svc := provider.CreateEc2Session(Region)
	vpcId := provider.VpcDetails(svc)
	subnetId := provider.SubnetDetails(svc)
	readConfig.Set("server.subnetid",subnetId)
	readConfig.Set("server.vpcid",vpcId)
	keyName := readConfig.GetString("user.keyname")
	fmt.Println("Creating a Security Group....")
	secGroup := provider.CreateSecGroup(secName,secDes,svc)
	readConfig.Set("server.secgroup",secGroup)
	fmt.Println("Creating a EC2 Instance...")
	instanceId := provider.CreateOneInstance(subnetId,secName,secGroup,instancetype,ami,keyName,svc)
	fmt.Println("Server Created")
	readConfig.Set("server.InstanceId",instanceId)
	readConfig.WriteConfig()
	publicIp := GetInstanceIP(instanceId,svc)
	PublicDnsName := GetInstanceDNS(instanceId,svc)
	readConfig.Set("server.ip",publicIp)
	readConfig.Set("server.dns",PublicDnsName)

	readConfig.WriteConfig()

	fmt.Println("Installing Necessary Software...")
	time.Sleep(5 * time.Second)
	user := readConfig.GetString("user.github")
	if len(repo)==0{
		provisioner.Deploy(URL,CPath,repo)
	}else {
	githubUrl := "https://github.com/"+user+"/"+repo+".git"
	provisioner.Deploy(githubUrl,CPath,repo)
	}

	fmt.Println("Site is Deployed at: "+publicIp)
	fmt.Println("Public DNS: "+PublicDnsName)
	
}	

func GetInstanceIP(instanceID string, svc *ec2.EC2) string{
	var ip string
    input := &ec2.DescribeInstancesInput{
        InstanceIds: []*string{
        aws.String(instanceID),
    },
    }
    result, err := svc.DescribeInstances(input)
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

    for _, reservation := range result.Reservations {
        for _, instance := range reservation.Instances {
            ip = aws.StringValue(instance.PublicIpAddress)
        }
    }
    return ip
}

func GetInstanceDNS(instanceID string, svc *ec2.EC2) string{
	var dns string
    input := &ec2.DescribeInstancesInput{
        InstanceIds: []*string{
        aws.String(instanceID),
    },
    }
    result, err := svc.DescribeInstances(input)
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

    for _, reservation := range result.Reservations {
        for _, instance := range reservation.Instances {
            dns = aws.StringValue(instance.PublicDnsName)
        }
    }
    return dns
}
