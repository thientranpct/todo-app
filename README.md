# todo-app

- To start the API Server, you should run this docker-compose command:
`docker-compose -f docker-compose.yaml up --build -d`
- To generate swagger, run this command:
`swag init --parseDependency  --parseInternal --parseDepth 1  -g main.go`
- To format swagger config, run this command:
`swag fmt`
- If you get this error **command not found: swag** you should run this command:
`export PATH=$(go env GOPATH)/bin:$PATH`
- By default:
  - Database will be running on port: **5432:5432**
  - Api server will be running on port: **1323:1323**
  - Swagger URL: **<http://localhost:1323/swagger/index.html>**
