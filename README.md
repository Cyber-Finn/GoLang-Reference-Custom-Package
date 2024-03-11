# How to Reference Custom Packages from within another separate Go function/app/method
This is a code-sample detailing how to reference custom (and sometimes private/organisational) GO packages - hosted in gitLab/GitHub - within another separate project.
<br></br>
Side note: I've simply included both projects into 1, separated by a single folder called "Code that will ref our package", for easier reading,
<br>
but these projects dont need to be in the same directory. Most solutions will simply have you copy/download the files into your work folder, 
<br>but this shows you how to avoid that, and simply reference the existing remote package
<br></br>

# Notes for local building:
basically, just follow these steps: (more details are inside the actual main.go code file)
1. write the code you need, following the package you want to import - and all its inputs
2. run the "GOPRIVATE" command
3. run the "go get" command
4. run go mod init {file name}.go
5. run go mod tidy -> this fetches the dependencies and links them up
6. happy coding!

# Notes for private or organisational repos - where CICD pipelines need to do all of the steps above:
follow the steps above, or make sure your pipeline is following them
1. set up an "Access Token" (in gitlab) or "Personal Access Token" (in GitHub -> you can limit permissions, etc.) for the repo you want to reference in your project
2. set up a "variable" (in Gitlab) or "secret" (in GitHub) on your CICD pipeline/main repo
3. place the value of Step 1's "Access Token" into the "Variable" in Step 2
4. ensure that your CICD pipelines are using Step 2's variable/secret to access the repository (in addition to running the commands above)
