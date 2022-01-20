# trust-anchor

## Running OBADA TrustAnchor software

### Generating EdDSA cryptographical keys

Running software requires two pairs of keys, one is used for private API access and other one is for signing and verifing tokens for OBADA blockchain.

#### Creating a private API key pair

TODO: add this as a subcommand binary executable

```bash
openssl genpkey -algorithm ED25519 | tee ~/.ssh/ta_api.pem
openssl pkey -in ~/.ssh/ta_api.pem  -pubout | tee ~/.ssh/ta_api.pub.pem
```

#### Creating OBADA blockchain key pair

```bash
openssl genpkey -algorithm ED25519 | tee ~/.ssh/ta_obada.pem
openssl pkey -in ~/.ssh/ta_obada.pem  -pubout | tee ~/.ssh/ta_obada.pub.pem
```

### Running binary

```bash
trust-anchor server --auth.public-key-path=~/.ssh/ta_api.pub.pem --auth.private-key-path=~/.ssh/ta_api.pub.pem
```

## API endpoints

### Get auth token (not finalized)

#### Request

```bash
curl -X GET http://localhost/api/v1/token
```

#### Response

```javascript
{"token":"eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiODE1YmMwN2EtNDUwMi00ZjY4LTkyNTQtNDZkN2ZhNjk4ZDEyIn0.BAkg48vWAuFiZP5_KhC-yRMCMtYXrvkS3qxlW4WQQFYcZCsUTl6AI4IFcJGuFPzlTYeXp7Or6vBpeJKK7NmUCg"}
```

#### Decoded body

```javascript
{

  "alg": "EdDSA",
  "typ": "JWT"
}
{
	"user_id": "815bc07a-4502-4f68-9254-46d7fa698d12"
}
```

### Issue OBADA TrustAnchor token

#### Request

```bash
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiODE1YmMwN2EtNDUwMi00ZjY4LTkyNTQtNDZkN2ZhNjk4ZDEyIn0.BAkg48vWAuFiZP5_KhC-yRMCMtYXrvkS3qxlW4WQQFYcZCsUTl6AI4IFcJGuFPzlTYeXp7Or6vBpeJKK7NmUCg' http://localhost/api/v1/issue-token
```

#### Response

```javascript
{"token":"eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDI3MjA4MDMsImlzcyI6Im9iYWRhLXRydXN0LWFuY2hvci1vcmciLCJzdWIiOiI4MTViYzA3YS00NTAyLTRmNjgtOTI1NC00NmQ3ZmE2OThkMTIiLCJ2ZXJpZnlVcmwiOiJodHRwOi8vbG9jYWxob3N0L2FwaS92MS92ZXJpZnkifQ.vXq5_SGghS_HGq4Ms8qmDXNaIa0KEzF15m-9X0B9yrqNn6DjNML4oVJ1-uQiE7Qg00Yq__BuKE0qNHliSAbvDA"}
```

#### Decoded body

```javascript
{
  "iat": 1642720803,
  "iss": "obada-trust-anchor-org",
  "sub": "815bc07a-4502-4f68-9254-46d7fa698d12",
  "verifyUrl": "http://localhost/api/v1/verify"
}
```

### Verify that blockchain token owner is compliant

#### Request

```bash
curl -X POST http://localhost/api/v1/verify -d '{"token":"eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDI3MjA4MDMsImlzcyI6Im9iYWRhLXRydXN0LWFuY2hvci1vcmciLCJzdWIiOiI4MTViYzA3YS00NTAyLTRmNjgtOTI1NC00NmQ3ZmE2OThkMTIiLCJ2ZXJpZnlVcmwiOiJodHRwOi8vbG9jYWxob3N0L2FwaS92MS92ZXJpZnkifQ.vXq5_SGghS_HGq4Ms8qmDXNaIa0KEzF15m-9X0B9yrqNn6DjNML4oVJ1-uQiE7Qg00Yq__BuKE0qNHliSAbvDA"}'
```

#### Response

```javascript
{"is_compliant":true}
```


