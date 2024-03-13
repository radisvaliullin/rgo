# Simple TLS HTTPS Server V2

## use curl client with cert
```
curl -Lv  --cacert ./tmp/cert.pem https://localhost:4000
```

## apps
* client/server - simple TLS client and server.
* client-mtls/server-mtls - client and server with mTLS.
* server-http - simple server without TLS for testing.
* gencert-tls - generates TLS self-signed certificate (with pubkey) and private key file
* gencert-mtls-cln - generates mTLS cert and key (for client side) with CA self-signed cert and key (CA - certificate authority)
* gencert-mtls-srv - generates mTLS self-signed cert and key (for server side)
