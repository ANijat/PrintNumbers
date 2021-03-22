# PrintNumbers
Print the numbers from 1 to 10 in random order 


### How to run the program 
Based on the target OS, you must ensure that the env (GOOS, GOARCH) should be as follow:
 
for MacOS
```shell                            
$ go env GOOS GOARCH              
darwin
amd64
```
for Linux
```shell
$ go env GOOS GOARCH
linux
amd64
```
To build an executable you need to run the given command
```shell
$ go build -o path-to-binary
```
To run the binary on machine you need to make sure that binary is executable:
```shell
$ chmod +x path-to-binary
```
and then run it in terminal:
  
```shell
$ ./path-to-binary
```
 
