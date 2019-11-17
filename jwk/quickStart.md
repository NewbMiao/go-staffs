#  JWKS (supported by auth0)
## what is JWKS?
|Item|	Description|
|---|---|
|JSON Web Key (JWK)|	A JSON object that represents a cryptographic key. The members of the object represent properties of the key, including its value.|
|JSON Web Key Set (JWKS)|	A JSON object that represents a set of JWKs. The JSON object MUST have a keys member, which is an array of JWKs.|

Auth0 exposes a JWKS endpoint for each tenant, which is found at https://YOUR_DOMAIN/.well-known/jwks.json. This endpoint will contain the JWK used to sign all Auth0-issued JWTs for this tenant.

[jwks](https://auth0.com/docs/jwks)
## quickStart
```
# **use auth0 free trier account for test**
# **if it expired, u can change it to yours (see env in file: `jwk/utils/token.go`)**
go run jwk/server.go
go run jwk/client.go
```
[auth0-go-tutorial](https://auth0.com/docs/quickstart/backend/golang/01-authorization)
[auth0-golang-api-samples](https://github.com/auth0-samples/auth0-golang-api-samples)

[auth0-go](https://github.com/auth0-community/auth0-go)
```
//output would like:
get token: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlFqWTRNekpCTXpjNU5rUXhPRFUwT1RKR09FUkZSRU5CTWpFeU5qQTJSakZFUmpOR1FUUTFSQSJ9.eyJpc3MiOiJodHRwczovL25ld2JtaWFvLmF1dGgwLmNvbS8iLCJzdWIiOiIyUU1zQlU2YnJ1N2d0U1IzazNTMzlMVlZaTExoMTJvMkBjbGllbnRzIiwiYXVkIjoiaHR0cHM6Ly9uZXdibWlhby5hdXRoMC5jb20vYXBpL3YyLyIsImlhdCI6MTU3Mzk2NjExMSwiZXhwIjoxNTc0MDUyNTExLCJhenAiOiIyUU1zQlU2YnJ1N2d0U1IzazNTMzlMVlZaTExoMTJvMiIsInNjb3BlIjoicmVhZDp1c2VycyByZWFkOmNsaWVudHMiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMifQ.UvhSlQ_iaKZLNTFu0L28iNGa0_ZocvdQmBQwAJncZB4pXlHVwo7f315Nzw_RIr204ygNE2q2gRdIZumz5ysHhf4FDu6aS290TmquKObilebGA0xu3Bke9k7goKHDj0ItPDGe-0o_6l8jP42juDlRyXvyenQiZuJx9_QE1bAhCsPWcpkIjDk6M55LdNzODqEMd_c2DV0ZgvKvm9pwGFdHD_OIEDPkl8tGIGhC3Cqzlu_kOMKRoFKOPPhqVKmNmkOlZc3Q_IIvt3njoi4nbjPcvonThRt-YvALiqdYrco-NdvZJEWoOpjG_unmSDQ-1kr02RO9Igtzp-nNOPEvfAtAew
request  http://localhost:3010/api/private
{"message":"Hello from a private endpoint! You need to be authenticated to see this."}
request  http://localhost:3010/api/private-scoped
{"message":"Insufficient scope."}
get public key from JWK_URI
key: {0xc0001e8170 [0xc0001b0000] QjY4MzJBMzc5NkQxODU0OTJGOERFRENBMjEyNjA2RjFERjNGQTQ1RA RS256 sig}
```

## grpc with jwt
```
go run jwk/grpc/server.go
go run jwk/grpc/client.go
```