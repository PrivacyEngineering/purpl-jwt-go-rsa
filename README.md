To use this module run:
```shell
go get -u github.com/Siar-Akbayin/jwt-go-auth@v0.1.2
``` 

and add this import statement to your Go file:
import ("github.com/Siar-Akbayin/jwt-go-auth")

# grpc-go-auth
This module contains a function that generates and returns a JWT based on a policy file. 

The function structure is:
```
GenerateToken(policyPath string, serviceName string, purpose string, keyPath string, expirationInHours time.Duration)
```

The key should be a RSA private key and the JSON structure of the policy should be:

```json
{
  "services": [
    {
      "service1": {
        "purpose1": {
          "allowed":
          {
            "field1": [
              "string"
            ],
            "field2": [
              "string"
            ],
            "field3": [
              "string"
            ]
          },
          "generalized":
          {
            "field1": [
              "string",
              "parameter if necessary"
            ],
            "field2": [
              "string",
              "parameter if necessary"
            ],
            "field3": [
              "string",
              "parameter if necessary"
            ]
          },
          "noised":
          {
            "field1": [
              "string",
              "parameter if necessary"
            ],
            "field2": [
              "string",
              "parameter if necessary"

            ],
            "field3": [
              "string",
              "parameter if necessary"
            ]
          },
          "reduced":
          {
            "field1": [
              "string",
              "parameter if necessary"
            ],
            "field2": [
              "string",
              "parameter if necessary"
            ],
            "field3": [
              "string",
              "parameter if necessary"
            ]
          }
        },
        "purpose2": {
          "..."
        }
      }
    },
    "service2"
    :
    {
      "..."
    },
    "..."
  ]
}
```

Example:
You can find an example [here](https://github.com/Siar-Akbayin/jwt-go-auth/blob/main/policy.json).