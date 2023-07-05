# grpc-go-auth
This module contains a function that generates and returns a JWT based on a policy file. 

The function structure is:
```
GenerateToken(filepath string, serviceName string, keypath string)
```

To use this module run:
```shell
go get -u github.com/Siar-Akbayin/jwt-go-auth@v0.1.1
``` 

and add this import statement to your Go file:
import ("github.com/Siar-Akbayin/jwt-go-auth")