# mygo-api
Basic api using go language, docker and github actions CI workflow

## Run locally using docker 
```bash
docker build -t mygo-api .   
docker run -it -p 8080:8080 mygo-api
```
## Available endpoints
- /hello    
- /health
- /metadata

## Deployment secrets
- Create a personal access token by going to Settings -> Developer settings -> Personal access tokens and name it: CI_GITHUB_TOKEN
- Copy the generated code and paste it on repo's Settings -> Secrets -> Value text area after clicking on "New repository secret" button
- Create additional secrets for docker hub details: DOCKER_USERNAME, DOCKER_PASSWORD

## Additional challenges
- Adding more tests coverage
- Trigger dev deployment to actual cloud provider



