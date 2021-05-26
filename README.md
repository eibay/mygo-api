# mygo-api
Basic api using go language, docker and github ci workflow

## Run locally using docker 
docker build -t mygo-api .   
docker run -it -p 8080:8080 mygo-api

## Deployment
- Create a personal access token by going to Settings -> Developer settings -> Personal access tokens and name it: CI_GITHUB_TOKEN
- Copy the generated code and paste it on repo's Settings -> Secrets -> Value text area after clicking on "New repository secret" button

## Challenges
- Adding tests
- Enhance github workflow to build and deploy docker images to ECR
