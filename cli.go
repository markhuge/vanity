package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"

	"markhuge.com/donate"
)

type options struct {
	BindAddr, NameSpace, Dest string
	Port                      int
	Debug                     bool
	SSLCert                   string
	SSLKeyFile                string
}

func Init() options {
	var opts options

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flags.StringVarP(&opts.Dest, "dest", "d", "", "Destination URI. Ex: https://git.markhuge.com")
	flags.IntVarP(&opts.Port, "port", "p", 0, "Port to listen on")
	flags.StringVar(&opts.BindAddr, "bind", "0.0.0.0", "Optional bind address")
	flags.StringVarP(&opts.NameSpace, "namespace", "n", "", "Optional namespace Ex: markhuge.com. Default is the host in the request header")
	flags.BoolVar(&opts.Debug, "verbose", false, "Verbose logging")

	flags.StringVar(&opts.SSLCert, "ssl-cert", "", "Path to fully concatinated SSL certificate. Used optionally to enable SSL and serve HTTPS. (--ssl-key is also required with this option)")
	flags.StringVar(&opts.SSLKeyFile, "ssl-key", "", "Path to SSL Keyfile (ex: key.pem). Used in conjunction with --ssl-cert")

	askedForVersion := flags.Bool("v", false, "Print version")
	askedForHelp := flags.BoolP("help", "h", false, "Print this help")
	donationType := flags.String("donate", "", "Display QR code to donate to the project. possible values: btc, xmr, eth")

	flags.SortFlags = false
	flags.Usage = usage(os.Stderr, flags)

	err := flags.Parse(os.Args[1:])

	if err != nil {
		flags.Usage()
		os.Exit(1)
	}

	if flags.Lookup("donate").Changed {
		switch *donationType {
		case "xmr":
			fmt.Println(donate.XMR.QRCode)
			os.Exit(0)
		case "eth":
			fmt.Println(donate.ETH.QRCode)
			os.Exit(0)
		case "btc":
			fmt.Println(donate.BTC.QRCode)
			os.Exit(0)
		default:
			flags.Usage()
			os.Exit(1)

		}

	}

	if *askedForVersion {
		fmt.Printf("vanity v%s\n", VERSION)
		os.Exit(0)
	}

	if *askedForHelp {
		usage(os.Stdout, flags)()
		os.Exit(0)
	}

	// Ensure both ssl keys have the same state
	if flags.Lookup("ssl-cert").Changed != flags.Lookup("ssl-key").Changed {
		flags.Usage()
		os.Exit(1)
	}

	if flags.Lookup("ssl-cert").Changed && flags.Lookup("ssl-key").Changed {
		if len(opts.SSLCert) == 0 || len(opts.SSLKeyFile) == 0 {
			flags.Usage()
			os.Exit(1)
		}
	}

	if len(opts.Dest) == 0 || opts.Port == 0 {
		flags.Usage()
		os.Exit(1)
	}
	return opts
}

// This is me being a pedant about help output going to Stdout but
// incorrect syntax going to Stderr
func usage(w io.Writer, f *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(w, versionFmt, VERSION)
		fmt.Fprint(w, f.FlagUsages())
		fmt.Fprintf(w, footer, donate.XMR, donate.BTC, donate.ETH)
	}
}

const versionFmt = "vanity v%s - A tiny server for golang vanity redirects\n\nUsage: vanity -d <destination URI>\n\n"

const footer = `

Copyright 2021 Mark Wilkerson <mark@markhuge.com>
This is free software licensed under the GPLv3 <https://www.gnu.org/licenses/> 

Source:		https://git.markhuge.com/vanity
Bugs:		bugs@markhuge.com 
Patches:	patches@markhuge.com 

Donate to this project:
  Monero (XMR): %s
  Bitcoin (BTC): %s
  Ethereum (ETH): %s

`
