
# Table of Contents

1.  [JWT signatures method](#org1bf3fb4)
2.  [JWT token structure and how proofs work](#org3ae7213)
    1.  [JWT token encoded by Base64](#orgdae157a)
    2.  [Decoded token body represented as a JSON (JWT claims)](#org90871ba)
        1.  ["iat" identifies the time at which the JWT was issued. The value must be a NumericDate.](#orgbaea8bf)
        2.  ["iss" identifies TrustAnchor organization that issued the JWT.](#org5143548)
        3.  ["nft" identifies obit DID address for which trust anchor will generate a token.](#org114e06a)
        4.  ["sub" (subject) identifies a user that exists in internal TrustAnchor database and against which API clients will check compliance status.](#org0efb865)
        5.  ["url" identifies TrustAnchor API url that should be used for checking compliance status of OBADA actor.](#org7b63e09)
    3.  [Verification with jwt.io](#org7b67cd2)
3.  [Verifiable credentils and DIDs](#org3afb885)



<a id="org1bf3fb4"></a>

# JWT signatures method

For signing and verification TrustAnchor tokens we are using EdDSA (Ed25519), see  [benefits](https://ed25519.cr.yp.to/)


<a id="org3ae7213"></a>

# JWT token structure and how proofs work


<a id="orgdae157a"></a>

## JWT token encoded by Base64

    eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDI2ODAxMTksImlzcyI6Im9iYWRhLXRydXN0LWFuY2hvci1vcmciLCJuZnQiOiJkaWQ6b2JhZGE6MTIzNDU2Nzg5Iiwic3ViIjoiODE1YmMwN2EtNDUwMi00ZjY4LTkyNTQtNDZkN2ZhNjk4ZDEyIiwidXJsIjoiaHR0cDovL2xvY2FsaG9zdC9hcGkvdjEvdmVyaWZ5In0.V1U1mF2qMcvDpmPPNoqlGv--_7et10U5I5uLXVMYZagzry1gR8QXg5Z8sJ8_sIW92GnuTl8n-oPUtu6Lf9LTAA


<a id="org90871ba"></a>

## Decoded token body represented as a JSON ([JWT claims](https://datatracker.ietf.org/doc/html/rfc7519#page-8))

    "iat": 1642680119,
    "iss": "obada-trust-anchor-org",
    "nft": "did:obada:123456789",
    "sub": "815bc07a-4502-4f68-9254-46d7fa698d12",
    "url": "http://localhost/api/v1/verify"


<a id="orgbaea8bf"></a>

### "[iat](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6)" identifies the time at which the JWT was issued. The value must be a NumericDate.


<a id="org5143548"></a>

### "[iss](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1)" identifies TrustAnchor organization that issued the JWT.


<a id="org114e06a"></a>

### "nft" identifies obit DID address for which trust anchor will generate a token.


<a id="org0efb865"></a>

### "sub" (subject) identifies a user that exists in internal TrustAnchor database and against which API clients will check compliance status.


<a id="org7b63e09"></a>

### "url" identifies TrustAnchor API url that should be used for checking compliance status of OBADA actor.

The origin idea was that this token can be converted back to HTTP request to the origin trust anchor API that issued it. Only API is able to tell us if user complied or not at the moment of request. 

    curl -X POST http://localhost/api/v1/verify -H 'content-type: application/json' -d '{"token":"eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDI2ODAxMTksImlzcyI6Im9iYWRhLXRydXN0LWFuY2hvci1vcmciLCJuZnQiOiJkaWQ6b2JhZGE6MTIzNDU2Nzg5Iiwic3ViIjoiODE1YmMwN2EtNDUwMi00ZjY4LTkyNTQtNDZkN2ZhNjk4ZDEyIiwidXJsIjoiaHR0cDovL2xvY2FsaG9zdC9hcGkvdjEvdmVyaWZ5In0.V1U1mF2qMcvDpmPPNoqlGv--_7et10U5I5uLXVMYZagzry1gR8QXg5Z8sJ8_sIW92GnuTl8n-oPUtu6Lf9LTAA"}'

After executing such request we can see if user compliant or not:

    {
      "is_compliant": true
    }


<a id="org7b67cd2"></a>

## Verification with [jwt.io](https://jwt.io/#debugger-io?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJ0cmFkZWxvb3AiLCJ2ZXJpZmljYXRpb25fdXJsIjoiaHR0cHM6XC9cL3RydXN0LWFuY2hvci50cmFkZWxvb3AuY29tIiwidmVyaWZpY2F0aW9uX2h0dHBfbWV0aG9kIjoiUE9TVCIsInZlcmlmaWNhdGlvbl9odHRwX3Zhcl9uYW1lIjoidG9rZW4iLCJpYXQiOjE2NDAwOTI4MDcsImF1ZCI6ImI3M2UyZDIyLTYyNjAtMTFlYy1hMTZhLTAyNDJhYzE0MDAwMyIsInJlZ2lzdHJhcl9pZCI6ImI1ZGIyZTg4YTE1ODI5Y2UzYTFiNTU2NjE3NmUxMTBjIn0.LDh0tYx_ZlS01i5hwJhuQLGrucmVtCpN5s_k0qwiWA3wXDPV31saaJKqv-RAA_h3lnSKbx6LTiTVIjcgZz_xmXyo32xG96zhCpV-QIwEGw5yV-U3IpVjvxKJD6dbrEldZxcyJalmXoQppfE7hM1kWUrrHPsLKq4UJDkN0DJnCslTkgnXsAm1JCJC9U0L9I4IEas1q2N-MsJ8iaioPc03pttllUmarVXgia3PgRK_P4cAQy_XW22WgazyxxG9v2Eo5wzKFmL90_gOjdI-N8x-3swJ0TawFt-AASoQwVnUgUmSGNUxSAXEV1tVLrzCnz4cIKFCXj5AuqZEbqvAq-m9GNv7jARrM2n9b0z9lw39EkheeYVMfRtZUVpL6CJNQUGvHPZsLCGv_mCwFH6ZUGdthmPbcP6mWfadQpZtXzOjVgl7jXMwsss-8NxusWgqhRO8YhU10yR6_S3X9shy9s6h3JCeMfIXUnFT6E-l2ntEzXWYt0HBBkSkacqpNSHSbqeRhOKZE2jprwZfKB4SRIHqAAoOTAfoLDGLdweWaNZ9nqJtjcsd1wKcDjZexpv8sN1qy6_9Td4MM7gJIRUeS4nZlVj4_OQtSMbWDMWnZku6CA7RGd7e9KDUbeWGLXJ5Smx8Z-vFCT9Is_KF5zFJhEOPvD_kbGYE4vKCUzHvdcTg9kU)

![img](./img/jwt-pub-key.png)

Copy public key below  into "VERIFY SIGNATURE" section on page:

    -----BEGIN PUBLIC KEY-----
    MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA1PdF9B367xmCATpkof5D
    GDP1ZmXc9u/AYN3Ojb7rEkhrBwai6F4SvnIz56rmtD6vKT/ZDIuslcZ3+MmUFB2z
    nD6Dyo2kfZ+P7S4tmcB0ZyGIGlOOmMrTvPME2MNmNIzDXbof3ex6IP+s0gkDnBX9
    TndZz10CwjAemP06rd9lslNUdUY+dnT7yfsmgsfjzQs1xCncRhzzAdI/ZI2Xi6a5
    xA0pcZTWOjEPk2WJLO32xUhnDg7wEgxnXKTbs7du51FpbWVshgKEhEoMTYX8nzTA
    7KxJS8WIduWllKMDB+Yt6+iSPr2E7hK8739X1a5c8gamsShdieuO9njFIwmArzIT
    cA/a2porVcKS0UIv2cXaN/KhYZEkvm/ke1ZbflNZQ3d5qj8z7yNxL4oLCVPbjGB0
    TI2y89rXLngTirHtrq5ikIHcQveeLfGTTGiKwMNjdiGLOWRe/Xd/QLFXBUWIHT3c
    Tg6nvIqjrM2x718VcpHXcDwA+ZbwQ7PTfq/5Tk2v+PIxHKs2+Cni7tHuqdFV9FM5
    jRBf1M+rlCWopf1g5+2q2RIhBaTPUez0VY+PZV3bRbj7mwLlf6kRvlHlqQQPDh/p
    zU7+Z35PcMoc8C+5yflQLMHs9DVWgO5cPOAQdgZQNO8Z8JzR277lsKtR4LZnu4Nq
    1Pygp6EGj0iDBuxxHF1O/tkCAwEAAQ==
    -----END PUBLIC KEY-----

You should see a change from "Invalid Signature"  to "Signature Verified"

![img](./img/jwt-public-invalid.png)

![img](./img/jwt-public-verified.png)

Copy private key:

    -----BEGIN PRIVATE KEY-----
    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
    22zZfpR2sI++YW7KNFPbCXbiKpgJE4SE
    -----END PRIVATE KEY-----

You should keep see "Signature Verified" same as before. Try to change token body or public or private key to see how it change signature check status.


<a id="org3afb885"></a>

# Verifiable credentils and DIDs

