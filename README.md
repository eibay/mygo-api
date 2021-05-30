# mygo-api
Basic api using go language, docker and github actions CI workflow. Every commit and push to main branch, triggers a continuous integration (CI) pipeline via github actions, which in turn automatically build a docker image and pushed to docker hub.

## Run locally
  ### using docker:  
```bash
  docker build -t mygo-api .   
  docker run -it -p 8080:8080 mygo-api
```
  ### using docker-compose:
```bash
  docker-compose up --build
```

## Available endpoints
- /hello    
- /health
- /metadata

## Running test
```bash
  go test ./...
```
## Deployment secrets
- Create a personal access token by going to Settings -> Developer settings -> Personal access tokens and name it: CI_GITHUB_TOKEN
- Copy the generated code and paste it on repo's Settings -> Secrets -> Value text area after clicking on "New repository secret" button
- Create additional secrets for docker hub details: DOCKER_USERNAME, DOCKER_PASSWORD
## API docker image output
  [eibayan/mygo](https://hub.docker.com/r/eibayan/mygo)

## Optional: Kubernetes setup using kind on mac
```bash
  brew install kind
  kind create cluster
  kubectl cluster-info --context kind-kind
  kubectl get pods
  kubectl get services 
```


## Additional challenges
- Adding more tests coverage
- Add build matrix for different environment (dev, stage, uat, prod)
- Trigger dev deployment to actual cloud provider
- Setup infrastructure as code (IAC) using Cloudformation Template, CDK or Terraform
- Setting up [ArgoCD](https://argoproj.github.io/projects/argo-cd)



