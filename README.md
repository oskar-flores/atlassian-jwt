## Simple Atlassian JWT Generator

This is a simple console application written in Go that generates a JWT (JSON Web Token) compatible with Atlassian cloud
apps.

## Prerequisites:

- Go (version 1.14 or later)

## Usage:

- Clone this repository: `git clone https://github.com/oskar-flores/atlassian-jwt`
- Navigate to the cloned repository: cd jwt_generator
- Build the Go file: `go build jwt_generator.go`

- Run the program with the issuer and signing key as command-line arguments: ./jwt_generator -issuer= -key=
  Replace and with your actual issuer and signing key.

## Output:

The program will output a JWT token that expires in 72 hours.


