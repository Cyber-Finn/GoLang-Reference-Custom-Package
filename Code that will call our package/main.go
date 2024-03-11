package main

import (
	//remember to run: go env -w GOPRIVATE=github.com/your_username_or_org
	// this basically lets you skip the checksum validation, which will cause GO to never find your package (because it's not hosted in the checksum db)
	//	additionally, you can also run this, to make it so that go skips this check for any repos under that umbrella: go env -w GOPRIVATE=github.com/your_username_or_org/*
	custom_logging_package "github.com/Cyber-Finn/GoLang-Reference-Custom-Package"
	//after this line is added, run this: go get github.com/Cyber-Finn/GoLang-Reference-Custom-Package
	//	this will actually download the package, for you to use in your app
	//then run go mod init and go mod tidy
)

func main() {
	//because I've already defined the "message_object" in my package, I dont have to redefine here. Simply create a new one
	var obj custom_logging_package.Message_object

	obj.Message = "hello world"
	obj.Message_type = 1
	obj.Severity = 100

	custom_logging_package.Log_info(obj)
	custom_logging_package.Log_warning(obj)
	custom_logging_package.Log_error(obj)
	custom_logging_package.Log_fatal(obj)

}

// notes:
// basically, follow these steps:
// 1) write the code you need, following the package you want to import - and all its inputs
// 2) run the "GOPRIVATE" command
// 3) run the "go get" command
// 4) run go mod init {file name}.go
// 5) run go mod tidy -> this fetches the dependencies and links them up
// 6) happy coding!

// notes for private repos or organisational repos - where CICD pipelines need to do all of this:
// 	follow the steps above, or make sure your pipeline is following them
// 1) set up an "Access Token" (in gitlab) or "Personal Access Token" (in GitHub -> you can limit permissions, etc.) for the repo you want to reference in your project
// 2) set up a "variable" (in Gitlab) or "secret" (in GitHub) on your CICD pipeline/main repo
// 3) place the value of Step 1's Access Token into the "Variable" in Step 2
// 4) ensure that your CICD pipelines are using Step 2's variable/secret to access the repository
