# EasyDeploy

A command line tool which helps you to deploy your website/webapps in AWS servers.

## Dependencies 
- Golang 
- AWS account
- Viper(github.com/spf13/viper)
- Cobra(github.com/spf13/cobra)
- AWS-SDK(github.com/aws/aws-sdk-go/...)

```
$ export GOPATH=/home/$USER/go

```

## Build From Source

```
$ git clone https://github.com/LoneWolf38/EasyDeploy.git /home/$USER/go/github.com/LoneWolf38/.

$ cd /home/$USER/go/github.com/LoneWolf38/EasyDeploy/

$ go build -o EasyDeploy main.go

$ ./EasyDeploy

```

## Basic Usage (Only One Command is now working)
```
$ ./EasyDeploy init [--region][--key]

```

## Commands
- Init : First command to run, only for one time to create a config with user details and creation of a ssh key for the instance.
- Deploy : Creates one instance, installs necessary packages and clones your repo from github and deploys it in the server.
- Destroy : Destroys the instance along with all the resources created.