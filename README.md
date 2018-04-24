[![Build Status](https://travis-ci.org/rfrench/ectoken-go.svg?branch=master)](https://travis-ci.org/rfrench/ectoken-go)

![Verizon Digital Media Services](https://images.verizondigitalmedia.com/2016/03/vdms-30.png)

# Token Generator for Edgecast Token-Based Authentication

Token-Based Authentication safeguards against hotlinking by adding a token requirement to requests for content secured by it. This token, which must be defined in the request URL's query string, defines the criteria that must be met before the requested content may be served via the CDN. This repository contains the following token generation resources:
- [Linux binaries and Windows executable](https://github.com/VerizonDigital/ectoken/releases/latest)

## golang Token Generator for EdgeCast Token-Based Authentication
* **Encrypt**(key string, params string, verbose bool)
* **Decrypt**(key string, token string, verbose bool)

## Install
First use "go get" to install the latest version of the library.

    $ go get github.com/VerizonDigital/ectoken-go

## Example
```go
import "github.com/VerizonDigital/ectoken-go/v3"

// encrypt
token, err := v3.Encrypt("keyvalue", "ec_expire=1257642471&ec_clientip=11.22.33.1", false)
if err != nil {
    log.Fatal(err)
}

// decrypt
params, err := v3.Decrypt("keyvalue", token, false)
if err != nil {
    log.Fatal(err)
}
```

## Contributing
Contributions are welcome! Please review our [contribution guidelines](CONTRIBUTING.md).

## More Information
Please refer to the CDN Help Center, which is available from within the MCC, for more information (e.g., parameter names and usage).

## License
[View legal and licensing information.](LICENSE.txt)

## Other Languages & Documentation
You can find implementations of EC Token in other languages and further documentation on the [EC Token Repo](https://github.com/hattan/ectoken).