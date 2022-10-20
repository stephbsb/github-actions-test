# github-actions-test
Just testing how github actions work

This projects automates the upload of a golang simple application to dockerhub using github actions.
Uses sonarqube (sonarcloud) to evaluate code quality and security.

## Run from dockerhub image

````bash
docker run --rm stephanybr/github-actions-test:latest
````