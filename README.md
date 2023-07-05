# grpc-go-auth
This module contains a function that generates and returns a JWT based on a policy file. 

The function structure is:
```
GenerateToken(filepath string, serviceName string, keypath string)
```
The key should be a RSA private key and the JSON structure of the policy should be:
```json
{
  "services": [
    {
      "service1": {
        "allowed":
        {
          "field1": "type",
          "field2": "type",
          "field3": "type"
        },
        "generalized":
        {
          "field1": "type",
          "field2": "type",
          "field3": "type"
        },
        "noised": 
        {
          "field1": "type",
          "field2": "type",
          "field3": "type"
        },
        "reduced":
        {
          "field1": "type",
          "field2": "type",
          "field3": "type"
        }
      },
      "service2": {
        ...
      },
      ...
  ]
}
```

Example:
```json
{
  "services": [
    {
      "registration": {
        "allowed":
        {},
        "generalized":
        {
          "street": "string",
          "name": "string",
          "age": "int",
          "sex": "string",
          "phoneNumber": "string"
        },
        "noised": {},
        "reduced":
        {}
      },
      "advertisement": {
        "allowed":
        {
          "name": "string",
          "street": "string",
          "sex": "string"
        },
        "generalized":
        {
          "phoneNumber": "string"
        },
        "noised":
        {
          "age": "int"
        },
        "reduced":
        {}
      },
      "tracking": {
        "allowed":
        {
          "street": "string",
          "name": "string",
          "age": "int",
          "sex": "string",
          "phoneNumber": "string"
        },
        "generalized": {},
        "noised": {},
        "reduced":
        {}
      }
    }
  ]
}
```

To use this module run:
```shell
go get -u github.com/Siar-Akbayin/jwt-go-auth@v0.1.1
``` 

and add this import statement to your Go file:
import ("github.com/Siar-Akbayin/jwt-go-auth")