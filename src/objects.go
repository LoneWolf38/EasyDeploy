package cmd

type Project struct{
	LocalPath string
	DeployType string
	GithubRepo string
}

type AwsProviderConfig struct{
	AccessKeyId string
	SecretAccessKey string
	Region string
	KeyPath string
	ImageId string
	InstanceType string
	SecName string
	SecDescr string
}

type AwsInstance struct{
	InstancePublicIp string
	InstancePrivateIp string
	InstanceVPC string
	InstanceSecGroupID string
	InstanceSecID string
	InstanceDNS string
	InstanceID string
	InstanceSubnet string
	InstanceKeyName string
}

type AwsVPC struct{
	VpcId string
}
type AwsSecGroup struct{
	SecGroupName string
	SecGroupDes string
	SecGroupId string
	SecGroupPort []int
}
type AWS struct{
	AwsVPC
	AwsSecGroup
	AwsProviderConfig
	AwsInstance
}

type EasyDeploy struct{
	Project
	AWS 
	ConfigPath string
	DbPath string
}