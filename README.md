# General Meeting API 
Summarise your minute meeting with ease with our API.

## Goal
This API aims to transcript the audio meeting and use AI to summarise the meeting using WhisperAI / GPT

## Architecture
Instead of using Javascript or Typescript we will use Golang since it excels in concurrency
### Database
#### Elasticsearch
For store the transcript and summarise of the meeting
#### PostgreSQL
Use for authentication and store transaction 

## Docker & Docker Compose
This project won't use `docker-compose` since I believe the best way to deal with the scaling is Kubernetes

## Encryption data
There won't be **encryption** script since eventually we can use it for searching and it will slow us down. The encryption layer will be implemented on **authentication**

## Requirement
- Go version `1.23.8`
- Elasticsearch version `8.17.4`
- PostgreSQL version `15`

## Setup 
Please make sure you set `.env` to be the correct one

## Deploy (locally)
I use `docker compose` to showcase the application. However, please note that this is not for the **production environment**.
Make sure that `.env` is the same as your `docker-compose.yml`
```
docker compose up --build
```
### Note
- Elasticsearch can take a long time to launch so I wrote the script to make sure that backend will wait

## Deploy (Production)
```
kubectl apply -f deploy/deployment.yaml
```
