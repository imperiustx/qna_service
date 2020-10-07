# qna_service

# Before you run: 
- Make sure your computer has docker, docker-compose installed
- Service is using port 7000, or you change to any port you want in .env, Dockerfile, docker-compose files
- Database is using port 27017

# Run the service: 
- To run this service, open your favorite terminal then run: make build
- If you see a line of server logs looks like below one, the server stared successfully: 
qna_server | 2020/10/07 06:05:48 stdout: {"time":"2020-10-07T06:05:48Z","level":"info","message":"starting on port 7000"}
- Or you can use Postman/any web browser then enter localhost:7000 If you see response like this: {"data":"welcome to qna_backend at 2020-10-07T08:20:05Z","error":"","success":true}. You are good to go
- To stop this service: Ctrl-C
- To completely shutdown this service, run: make down

# Troubleshoot:
- If your computer required root privilege, add sudo in every commands in Makefile
- Sometime docker cannot build service successfully, try to delete mongo_volume folder and make build again

