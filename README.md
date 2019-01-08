# EasyDeploy [![Build Status](https://travis-ci.org/LoneWolf38/EasyDeploy.svg?branch=dev)](https://travis-ci.org/LoneWolf38/EasyDeploy)



![](https://img.shields.io/badge/Golang-1.11-blue.svg?style=for-the-badge&logo=go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=for-the-badge)](https://opensource.org/licenses/Apache-2.0)


A command line tool which helps you to deploy your website/webapps in AWS servers.

## Install

- No extra dependencies required to be installed
- Just download the binary from Releases and copy the binary to /usr/bin

```
$ wget https://github.com/LoneWolf38/EasyDeploy/releases/download/1.0/EasyDeploy-v1.0-PreRelease.zip
$ unzip EasyDeploy-v1.0-PreRelease.zip
$ sudo cp EasyDeploy /usr/local/bin/.
$ EasyDeploy --version
```
## Dependencies 
- Golang 
- AWS account (Access key and Secret Key needed)
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

## Basic Usage 
```
$ ./EasyDeploy init [--region][--key]

$ ./EasyDeploy deploy [--repo][--firewall]

```

## Commands
- Init : First command to run, only for one time to create a config with user details and creation of a ssh key for the instance.
- Deploy : Creates one instance, installs necessary packages and clones your repo from github and deploys it in the server.
- Destroy : Destroys the instance along with all the resources created.

## TODO
- Complete all commands and flags
- Add more resources of AWS
- More configuration of the resources
- Direct upload of files from user 
- Support of GitLab and other private repositories
- Available for all Operating Systems
- Single executable with no dependencies
- Colored output for interactive user experience
- Automatic updates of the repo in the server
- Building of software(other than webapps) support to be added
- More expalined help command
- MageFile integration for manual building of EasyDeploy