# PrintNumbers
Print the numbers from 1 to 10 in random order 


If binary will target MacOS, please make sure that given env (GOOS GOARCH) should be same as result.
 
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
For build you need execute given command.
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
 
