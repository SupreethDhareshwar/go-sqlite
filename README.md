# Go Sqlite Setup

## Pre-requisites

- Install Git `sudo apt-get install git`
- Configure user.name and user.email
  `git config --global user.email "you@example.com"`
  `git config --global user.name "Your Name"`
- Setup Github Desktop
  Go to https://github.com/shiftkey/desktop/releases and download latest release
  To install `sudo dpkg -i <file>.deb`
- Install Docker and Docker-Compose
- Download latest Go Version from https://golang.org/dl/
- Setup Go (We are using 1.14) at /usr/local/go
- Setup Env variables like GOPATH and GOROOT in ~/.profile
  `export PATH=$PATH:/usr/local/go/bin`
- Install Go Plugin for VSCode
- Setup Project anywhere and initialize with `go mod init <projecturl>` Example: `go mod init github.com/SupreethDhareshwar/go-sqlite`
- Refer this article for managing project dependecies - `https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31`
- Project is all setup, time to start installing packages and start coding

## Setup Project

- Create a new file main.go. This will be a starting point of application 
- Create a folder model. This will store all your application model and schema files
- Create a folder config. This will contain logic related to configuration
- Create a Makefile. This will hold all your project commands
- 
