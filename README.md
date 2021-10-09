# vanity - A tiny server for golang vanity redirects

## Who is this for?

`vanity` is for anyone who wants to self host their own 
[go vanity imports](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths) without 
having to maintain a full-scale webserver.

`vanity` deploys as a single binary with no configuration files and runs well 
in a docker container.


## Usage

Run `vanity` on a host with DNS that matches your desired namespace. Set the 
destination to the actual location of your git repositories.


```
Usage: vanity -d <destination URI>

  -d, --dest string        Destination URI. Ex: https://git.markhuge.com
  -p, --port int           Port to listen on
      --bind string        Optional bind address (default "0.0.0.0")
  -n, --namespace string   Optional namespace Ex: markhuge.com. Default is the host in the request header
      --verbose            Verbose logging
      --ssl-cert string    Path to fully concatinated SSL certificate. Used optionally to enable SSL and serve HTTPS. (--ssl-key is also required with this option)
      --ssl-key string     Path to SSL Keyfile (ex: key.pem). Used in conjunction with --ssl-cert
      --v                  Print version
  -h, --help               Print this help
      --donate string      Display QR code to donate to the project. possible values: btc, xmr, eth
```

### Example

Let's say you have a package `foo` hosted on github at github.com/yourgithub/foo,
and you want to import it (or any other packages on your account) from 
`yourname.biz/<package name>`

On a host with DNS for yourname.biz: `vanity --dest https://github.com/yourgithub --port 80`

## Install

`go install markhuge.com/vanity`

## Contribute

- Bugs:		bugs@markhuge.com 
- Patches:	patches@markhuge.com 

Submitted patches should include tests that cover the change.

## Donate

Donate to this project:

- Monero (XMR): 88vd4Fxy3AdcUpZp3FChgu5RGBBoEANdpXaB5Bm47JRGKqYbxwQZo1MMwkguQAUDioEPyf4rFK1yMUCgrE7ojVpAVEEzXVD
- Bitcoin (BTC): bc1qk22yx0gfce54gx9csy6dp6kl629wm0m9kscwl8
- Ethereum (ETH): 0x10517dcb7f3357aB6888cD6067b12D1ce2727B26
