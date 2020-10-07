# qna_service

# before you run: 
- Make sure your computer has docker, docker-compose installed

# Run the service: 
- To run this service, open your favorite terminal then run: make build
- If you see a line of server logs looks like below one, the server stared successfully: 
qna_server | 2020/10/07 06:05:48 stdout: {"time":"2020-10-07T06:05:48Z","level":"info","message":"starting on port 7000"}
- To stop this service: Ctrl-C
- To completely shutdown this service, run: make down

# troubleshoot:
- If your computer required root priviledge, add sudo in every commands in Makefile
- Sometime docker cannot build service successfully, try to delete mongo_volume folder and make build again

