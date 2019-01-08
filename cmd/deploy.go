package cmd

import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
		"github.com/LoneWolf38/EasyDeploy/provider"
		"github.com/spf13/viper"
		   "github.com/aws/aws-sdk-go/aws"
		  // "github.com/aws/aws-sdk-go/aws/session"
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

	} 
}

func init() {
	DeployAppCmd.PersistentFlags().StringVar(&secName,"firewall",secName,"Name of the security Group")
	DeployAppCmd.PersistentFlags().StringVar(&repo, "repo", repo, "Name of the project")
	DeployAppCmd.PersistentFlags().StringVar(&URL, "url","", "Github URL of the project")
}



func ExecuteDeploy() {
	updateConfig := viper.New()
	fmt.Println("Collecting Info...")
	svc := provider.CreateEc2Session(Region)
	vpcId := provider.VpcDetails(svc)
	subnetId := provider.SubnetDetails(svc)
	updateConfig.SetConfigFile(CPath)
	updateConfig.SetConfigType("json")
	updateConfig.Set("server.SubnetId",subnetId)
	updateConfig.Set("server.VpcId",vpcId)
	keyName := updateConfig.GetString("user.keyname")
	fmt.Println("Creating a Security Group....")
	secGroup := provider.CreateSecGroup(secName,secDes,svc)
	updateConfig.Set("server.secGroup",secGroup)
	fmt.Println("Creating a EC2 Instance...")
	instanceId := provider.CreateOneInstance(subnetId,secName,secGroup,instancetype,ami,keyName,svc)
	fmt.Println("Server Created")
	updateConfig.Set("server.InstanceId",instanceId)
	updateConfig.WriteConfig()

	publicIp := GetInstanceIP(instanceId,svc)
	PublicDnsName := GetInstanceDNS(instanceId,svc)
	updateConfig.Set("server.ip",publicIp)
	updateConfig.Set("server.dns",PublicDnsName)

	updateConfig.WriteConfig()

	fmt.Println("Installing Necessary Software...")

	user := updateConfig.GetString("user.github")
	if len(repo)==0{
		provisioner.StaticDeploy(URL,CPath)
	}else {
	githubUrl := "https://github.com/"+user+"/"+repo+".git"
	provisioner.StaticDeploy(githubUrl,CPath)
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
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
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
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
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
