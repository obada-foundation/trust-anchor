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
