package cmd

import (
		"fmt"
		"os"
		"github.com/spf13/cobra"
		"github.com/LoneWolf38/EasyDeploy/provider"
		"github.com/spf13/viper"
		  "github.com/aws/aws-sdk-go/aws"
		  "github.com/aws/aws-sdk-go/aws/session"
		  "github.com/aws/aws-sdk-go/service/ec2"
		"github.com/LoneWolf38/EasyDeploy/provisioner"
)

var HOME = os.Getenv("HOME")


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



	user := updateConfig.GetString("user.github")
	githubUrl := "https://github.com/"+user+"/"+repo+".git"
	provisioner.StaticDeploy(githubUrl,CPath)
}	

// func GetInstanceIP(svc *ec2.EC2) {
// 	input := &ec2.DescribeInstancesInput{
//     Filters: []*ec2.Filter{
//         {
//             Name: aws.String("ip-address"),
//             Values: []*string{
//                 aws.String("t2.micro"),
//             },
//         },
//     },
// }
// }

// func GetInstanceDNS(svc *ec2.EC2) {
// 	input := &ec2.DescribeInstancesInput{
//     Filters: []*ec2.Filter{
//         {
//             Name: aws.String("dns-name"),
//             Values: []*string{
//                 aws.String("t2.micro"),
//             },
//         },
//     },
// }
// }

