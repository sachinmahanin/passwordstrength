# passwordstrength service 
returns the strength of the password.

## Developer Setup
  Add launch.json to provide the environment variable
```
  {
    "version": "0.2.0",
    "configurations": [{
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "HOSTNAME": "127.0.0.1",
            "program": "${workspaceRoot}",
            "env": {
                "HOSTNAME": "127.0.0.1",
                "APP_PORT": 18605,
                "APP_VERSION": "0.15.25" ,              
            },
            "args": [],
            "showLog": true
        }
    ]
  }
```
## API Request
How to send API Request to the code running locally through tool like Postman etc.?

Method : POST

URL : http://localhost:18605/PassportStrength

Body :

```
{
    "password": "JamesBond#007forYou"
}
```
## Code structure

### passwordstrength

  .github        -CI workflow
  
  config         -Application level configuration eg. AppName,AppVersion etc
  
  customization  -Customization related to bootstrapping,logging,hosting,handlers and web requests
  
  deployment     -k8 yaml files like deployment.yaml,service.yaml
  
  handler        -hanlers for API request
  
  middleware     -middleware related to logging or other
  
  sessionutil    -Used for Session related operations
  
  timeutil       -Used for time related operations
  
  vendor         -stored dependent libraries
  
  web            - RegisteredStatics Static,Business & Utility routes

## CI
As soon as commit is pushed to the master branch, CI workflow is triggered. Which include
1. Build
2. Running Test cases with code coverage
3. Building docker image
4. push the image to the docker hub account https://hub.docker.com/r/sachinmahanin/passwordstrength

## Running the service in minikube
1. Run the minikube service on your local box - minkube start
2. Run following command from the root passwordstrength will create deployment+service in the minikube
```
kubectl apply -f deployment
```

