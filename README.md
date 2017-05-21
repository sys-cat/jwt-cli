# jwt-cli

'jwt-cli' is generate jwt centence client on cli.

## Install

0. Before install `github.com/dgrijalva/jwt-go` and `github.com/urfave/cli` or `glide i`
1. clone this repository under the `$GOPATH` .
2. run `go install` on jwt-cli directory.

## Usage

```
jwt-cli [Global Option] argument
```

### Example

**on inline json**

```
$ jwt-cli -m "HS512" -s "1234567890" -e "{\"test\":\"testtest\"}"
Original: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdHRlc3QifQ.2l10UxO1WGGaoCTyL8aFOa564pC_HkorJwCy_gVhJ0e6NvWZQ-fMLUdt_Mxbi-8Qmdw9lo4bRAMDXBCIw1AMBQ
Header: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9
Body  : eyJ0ZXN0IjoidGVzdHRlc3QifQ
Verify: 2l10UxO1WGGaoCTyL8aFOa564pC_HkorJwCy_gVhJ0e6NvWZQ-fMLUdt_Mxbi-8Qmdw9lo4bRAMDXBCIw1AMBQ
```

**from json file**

```
$ jwt-cli -m "HS512" -s "1234567890" -f "./d.json"
Original: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdCBkYXRhIn0.C2n3frZ-L_oO2hycC1ND6SAN8HLAnrE6IJUEAJxVY8uchQXUGf7MYwMBWNbZYTqhi24nqgZ1rt5IodTpab1BvQ
Header  : eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9
Body    : eyJ0ZXN0IjoidGVzdCBkYXRhIn0
Verify  : C2n3frZ-L_oO2hycC1ND6SAN8HLAnrE6IJUEAJxVY8uchQXUGf7MYwMBWNbZYTqhi24nqgZ1rt5IodTpab1BvQ
```

### Change the format of JSON

edit `Input` struct on `main.go` when change json format.

default is 

```
Input struct {
	Test string `json:"test"`
	jwt.StandardClaims
}
```