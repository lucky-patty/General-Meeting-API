# General Meeting API 
Summarise your minute meeting with ease with our API.

## Goal
This API aims to transcript the audio meeting then use AI to summarise the meeting

## Architecture
Instead of using Javascript or Typescript we will use Golang since it excels in concurrency
### Database
#### Elasticsearch
For store the transcript and summarise of the meeting
#### PostgreSQL
Use for authentication and store transaction 
#### MongoDB
Store log

## Docker & Docker Compose
This project won't use `docker-compose` since I believe the best way to deal with the scaling is Kubernetes

## Encryption data
There won't be **encryption** script since eventually we can use it for searching and it will slow us down. The encryption layer will be implemented on **authentication**

### Installation
```
docker build -t <IMAGE_NAME> .
```
