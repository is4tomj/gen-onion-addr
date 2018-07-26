Generate Onion Addresses
============

Generating [v2](https://gitweb.torproject.org/torspec.git/tree/rend-spec-v2.txt) and [v3 (part of proposal 224)](https://gitweb.torproject.org/torspec.git/tree/rend-spec-v3.txt) onion addresses are discussed below. The online descriptions I found had various degrees of ambiguity. Any perceived ambiguity was probably caused by my lack of experience with Tor. Hopefully, this project will help clarify the process for others.

[This page on the tor wiki](https://trac.torproject.org/projects/tor/wiki/doc/NextGenOnions#Howtosetupyourownprop224service) describes how to use your newly generated onion address.

## Usage

```bash
$ go build && go install
$ gen-onion-addr --help
```

## Generating a v2 Onion Address

1. Generate 1024-bit RSA key-pair comprising private/secret key, referred to as *SK*, and public key, referred to as *PK*.
2. Encode *SK* into PKCS#1, [ASN.1 DER](https://cryptologie.net/article/262/what-are-x509-certificates-rfc-asn1-der/) format.
3. Encode *PK* into PKCS#1, ASN.1 DER format.
4. Generate the SHA1 hash of the *PK* in the PKCS#1, ASN.1 DER format.
5. Encode the first 10 bytes of the SHA1 hash into a base32 string. The result will be 16 bytes in all caps.
6. Lowercase the base32 string and append ".onion" to generate the v2 onion address.
7. Store the v2 onion address in a file named `hostname`.
8. Encode the PKCS#1 ASN.1 DER form of *SK* in PEM form.
9. Store the PEM encoded *SK* in a file named `private_key`.

**Note**, the *PK* need not be persistently stored.

## Generating a v3 Onion Address

