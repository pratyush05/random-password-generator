# Random Password Generator
RESTful service which returns a random password

### Prerequisite
 - Ubuntu >= 18.04
 - Docker >= 19.03
 - Go >= 1.14 

### Getting started
To start server run: `./bld/run.sh`

### For example
```
curl --insecure --location --request GET 'https://localhost:9999/password?length=32&alphaNum=false'
{"random_password":"UThPTPHn#h#x3fqn5PyFwFOPxOessrsj","length":32}

curl --insecure --location --request GET 'https://localhost:9999/password?length=16&alphaNum=true'
{"random_password":"BpLnfgDsc2WD8F2q","length":16}

curl --insecure --location --request POST 'https://localhost:9999/password?length=32&alphaNum=false'
{"error_code":501,"error_message":"Method not implemented"}
```

