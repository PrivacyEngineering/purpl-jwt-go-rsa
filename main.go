package grpc_go_auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateToken(filepath string, serviceName string, keypath string) (string, error) {

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

	// Retrieve the services array from the policy
	servicesArr, exists := policyMap["services"].([]interface{})
	if !exists {
		return "", fmt.Errorf("Invalid policy format: services not found")
	}

	// Get the service policy based on the service name
	servicePolicy := make(map[string]interface{})
	for _, service := range servicesArr {
		serviceObj := service.(map[string]interface{})
		if _, exists := serviceObj[serviceName]; exists {
			servicePolicy = serviceObj[serviceName].(map[string]interface{})
			break
		}
	}

	if len(servicePolicy) == 0 {
		return "", fmt.Errorf("Service %s not found in policy file", serviceName)
	}

	// Create the reduced policy based on the service policy
	reducedPolicy := map[string]interface{}{
		"allowed":     servicePolicy["allowed"],
		"generalized": servicePolicy["generalized"],
		"noised":      servicePolicy["noised"],
		"reduced":     servicePolicy["reduced"],
	}

	// Convert the reduced policy to JSON
	reducedPolicyJSON, err := json.Marshal(reducedPolicy)
	if err != nil {
		log.Fatalf("Error marshaling reduced policy: %v", err)
	}

	// Load the RSA private key from file
	keyData, err := ioutil.ReadFile(keypath)
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
		jwt.StandardClaims
	}{
		reducedPolicyJSON,
		jwt.StandardClaims{
			// Valid for 24 hrs
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
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
