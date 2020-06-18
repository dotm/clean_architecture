## Running

**Run backend before frontend.**

Backend:
- install latest Go
- run from command line: `go run backend-golang/main.go`

Frontend:
- just open `frontend-javascript/index.html` in a web browser
- to login:
    - check the backend-golang/db/users.json
    - input the existing username (`testuser`)
    - input the existing password (`1234`)
    - click the `Login` link (pressing `Enter` won't trigger login)

## Warning

**Don't use this in production**

Backend side:

- doesn't have proper error handling (to make logic clearer)
- allows CORS (to enable connecting in localhost)
- doesn't encrypt or validate password
- is very unoptimized
- doesn't guard type assertion (type casting)

Frontend side:

- is not asynchronous (request to backend may cause hang because synchronous request blocks the UI)