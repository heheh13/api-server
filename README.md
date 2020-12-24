# api-server

A simple api server as a learning purpose

## Running the server

`go run . start`

`go run . start --port=<portNumber>`

## performing api tasks

|method|url|payload|actions|
|---|---|---|---|
|GET|`https://localhost:8080/users`||returns the list of all users|
|GET|`https://localhost:8080/users/{id}`||returns a specific users|
|POST|`https://localhost:8080/users`|payload| add a new user to the list|
|PUT|`http://localhost:8080/users/{id}`|payload|update specific existing user|
|DELETE|`http://localhost:8080/users/{id}`||delete specific user|

-------

## structure

- apiServer
  - api
    - api.go `performs related task about the api processing`
  - auth
    - auth.go `works a middleware before executing any of the api functions`
  - cmd
    - startCmd `provide some features of cli`

## implementation details

to perform a http request we need to write a handler functions which accepts
`http response and request`. we decode the request and write  to response as a expected functionality.
then we wrap the handler function to the authentication middleware to perform the security check.which follows tow simple technique Basic Auth and Jwt auth.
To execute every http request has to go through the security check before it can hit the handler function.
Which can simply implemented by adding the authorization in the request header

## resources

`https://www.youtube.com/watch?v=W5b64DXeP0o`

`https://www.youtube.com/watch?v=YMQUQ6XQgz8`

`https://www.youtube.com/watch?v=-Scg9INymBs&t=1072s`

`https://sysdevbd.com/go/#http-basic`

`https://sysdevbd.com/go/#go-nethttp`
