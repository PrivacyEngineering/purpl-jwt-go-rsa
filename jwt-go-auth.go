package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(filepath string, serviceName string, purpose string, keyPath string) (string, error) {

	// Load policy from file
	policyData, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading policy.json: %v", err)
	}

	// Parse the policy JSON into a map
	var policyMap map[string]interface{}
	err = json.Unmarshal(policyData, &policyMap)
	if err != nil {
		log.Fatalf("Error parsing policy.json: %v", err)
	}

	// Retrieve the services object from the policy
	servicesObj, exists := policyMap["services"].(map[string]interface{})
	if !exists {
		return "", fmt.Errorf("Invalid policy format: services not found")
	}

	// Get the service policy based on the service name
	servicePolicy, exists := servicesObj[serviceName].(map[string]interface{})
	if !exists {
		return "", fmt.Errorf("Service %s not found in policy file", serviceName)
	}

	// Get the policy for the specified purpose
	purposePolicy, exists := servicePolicy[purpose].(map[string]interface{})
	if !exists {
		return "", fmt.Errorf("Purpose %s not found in service policy", purpose)
	}

	// Create the reduced policy based on the purpose policy
	reducedPolicy := map[string]interface{}{
		"allowed":     purposePolicy["allowed"],
		"generalized": purposePolicy["generalized"],
		"noised":      purposePolicy["noised"],
		"reduced":     purposePolicy["reduced"],
	}

	// Convert the reduced policy to JSON
	reducedPolicyJSON, err := json.Marshal(reducedPolicy)
	if err != nil {
		log.Fatalf("Error marshaling reduced policy: %v", err)
	}

	// Load the RSA private key from file
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}

	// Parse the RSA private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	// Create the Claims
	claims := struct {
		Policy json.RawMessage `json:"policy"`
		jwt.RegisteredClaims
	}{
		reducedPolicyJSON,
		jwt.RegisteredClaims{
			// Valid for 2 hrs
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "tokenGenerator",
		},
	}

	// Sign the token using RSA-SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
		return "", err
	}

	return tokenString, nil
}
