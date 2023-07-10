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
      "service2": {
        "..."
      },
      "..."
  ]
}
```

Example:
```json
{
  "services": [
    {
      "service1": {
        "allowed": {
        },
        "generalized": {
          "credit_card_cvv": [
            "int",
            "3"
          ],
          "zip_code": [
            "int",
            "8"
          ],
          "city": [
            "string",
            "2"
          ],
          "credit_card_expiration_year": [
            "int",
            "10"
          ],
          "credit_card_number": [
            "string",
            "5"
          ]
        },
        "noised": {
          "age": [
            "int",
            "Laplace"
          ],
          "street_number": [
            "int",
            "Laplace"
          ],
          "street_name": [
            "string",
            "Laplace"
          ],
          "credit_card_expiration_month": [
            "int",
            "Laplace"
          ]
        },
        "reduced": {
          "email": [
            "string",
            "4"
          ],
          "country": [
            "string",
            "3"
          ],
          "name": [
            "string",
            "4"
          ],
          "phone": [
            "string",
            "3"
          ]
        }
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