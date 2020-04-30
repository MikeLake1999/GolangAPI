"# GolangAPI" 
Run App
"go run main.go"
Use Curl in Git Bash
Example: 
- curl -X GET http://127.0.0.1:3000/v1/post?accountId=1
- curl -X POST -H "Content-Type: application/json" -d '{"email":"123", "password":"123"}' http://127.0.0.1:3000/v1/authentication  
- curl -X GET -H "Authorization: Bearer (token)" http://127.0.0.1:3000/v1/account
