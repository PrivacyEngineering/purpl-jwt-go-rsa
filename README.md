<p align="center">
	<img src="purpl.png" width=50" />
</p>

# grpc-go-auth
This Go module generates JWTs using RSA encryption, tailored to specific services and purposes as defined in a policy file. 
It dynamically adjusts token claims based on the policy, including permissions and conditions for the token's use, then 
signs it with a private RSA key. The token's expiration is set according to the specified duration.

The function structure is:
```
GenerateToken(policyPath string, serviceName string, purpose string, keyPath string, expirationInHours time.Duration)
```

The key should be an RSA private key and the JSON structure of the policy should be:

```
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
          ...
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
You can find an example [here](https://github.com/PEngG7/jwt-go-rsa/blob/main/policy.json).

# Usage

To use this module run:
```shell
go get -u github.com/Siar-Akbayin/jwt-go-auth@v0.1.3
``` 

and add this import statement to your Go file:
```go
import ("github.com/Siar-Akbayin/jwt-go-auth")
```

# Testing
The test.go file contains a test for the GenerateToken function. It uses the policy.json file and the private key
provided in this repo. The provided test generates the following token:

```
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJwb2xpY3kiOnsiYWxsb3dlZCI6e30sImdlbmVyYWxpemVkIjp7ImNpdHkiOlsic3RyaW5nIiwiMiJdLCJjcmVkaXRfY2FyZF9jdnYiOlsiaW50IiwiMyJdLCJjcmVkaXRfY2FyZF9leHBpcmF0aW9uX3llYXIiOlsiaW50IiwiMTAiXSwiY3JlZGl0X2NhcmRfbnVtYmVyIjpbInN0cmluZyIsIjUiXSwiemlwX2NvZGUiOlsiaW50IiwiOCJdfSwibm9pc2VkIjp7ImFnZSI6WyJpbnQiLCJMYXBsYWNlIl0sImNyZWRpdF9jYXJkX2V4cGlyYXRpb25fbW9udGgiOlsiaW50IiwiTGFwbGFjZSJdLCJzdHJlZXRfbmFtZSI6WyJzdHJpbmciLCJMYXBsYWNlIl0sInN0cmVldF9udW1iZXIiOlsiaW50IiwiTGFwbGFjZSJdfSwicmVkdWNlZCI6eyJjb3VudHJ5IjpbInN0cmluZyIsIjMiXSwiZW1haWwiOlsic3RyaW5nIiwiNCJdLCJuYW1lIjpbInN0cmluZyIsIjQiXSwicGhvbmUiOlsic3RyaW5nIiwiMyJdfX0sImlzcyI6InRva2VuR2VuZXJhdG9yIiwiZXhwIjoxNzA3NDgzNzg4fQ.H-m005YL06s5ZcMeyhda9EX20tjzZv1RSpC7W32EPA-MplKGT7bmU4n8Kwntfr-yGi9Xv8vhqqDjjBUjhuHiKs9kzeBwhsDhzsB2j-W5C1V6NWsrCEZFcw0_w35jGVv1EhTC02qcPoPbfthzM2_6rWmcJX1IDeQDQu4ZwdUOWdYU3i4nw6HwDJIfUbNSdr9bPQ0RX50HT4xWuKX2KaG7OAYcn_i2tawmaJ7gcMDsNZFiO8DuyzeuoPePafMumDQbArDNL0_PdwWCqZddQbFtUl8M0auQDUBnFJlIC75afC09aH3JdjhMWU-hIoz8m26v-2T57Zr0P705thEJWoh1IA
```

The content can be decoded using a JWT decoder, such as [jwt.io](https://jwt.io/).

In this case it looks like this:
HEADER
```json
{
  "alg": "RS256",
  "typ": "JWT"
}
```
PAYLOAD
```json
{
  "policy": {
    "allowed": {},
    "generalized": {
      "city": [
        "string",
        "2"
      ],
      "credit_card_cvv": [
        "int",
        "3"
      ],
      "credit_card_expiration_year": [
        "int",
        "10"
      ],
      "credit_card_number": [
        "string",
        "5"
      ],
      "zip_code": [
        "int",
        "8"
      ]
    },
    "noised": {
      "age": [
        "int",
        "Laplace"
      ],
      "credit_card_expiration_month": [
        "int",
        "Laplace"
      ],
      "street_name": [
        "string",
        "Laplace"
      ],
      "street_number": [
        "int",
        "Laplace"
      ]
    },
    "reduced": {
      "country": [
        "string",
        "3"
      ],
      "email": [
        "string",
        "4"
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
  },
  "iss": "tokenGenerator",
  "exp": 1707483788
}
```

Furthermore, you can verify the signature by copying the public key from the public.pem file and the private key from the 
key.pem file and pasting them into the "Verify Signature" section of the jwt.io website.