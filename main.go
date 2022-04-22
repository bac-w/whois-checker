package main

import (
	"crypto/tls"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"golang.org/x/exp/slices"
	"net"
	"os"
	"time"
	"whois-checker/pkg/log"
)

type args struct {
	Domain  string `arg:"-d,--domain,env:CHECKDOMAIN" help:"Domain name for check"`
	Verbose bool   `arg:"-v" help:"Verbose [DEBUG]"`
	Expire  bool   `arg:"-e,--expire" help:"Find out when domain it expires (days)"`
	CertExp bool   `arg:"-s,--ssl" help:"Find out when certificate it expires (days)"`
	DStatus bool   `arg:"--status" help:"Get Domain status"`
}

var Args args

var pkgBuildDate string
var pkgMaintainer = "Leonov Vitaliy"
var pkgVersion = "0.1"
var pkgUrl = "https://github.com/bac-w/whois-checker"

func (args) Version() string {
	if pkgBuildDate == "" {
		pkgBuildDate = time.Now().Format("01-02-2006 15:04")
	}
	return fmt.Sprintf("Program: Whois and SSL Simple Checker\nVersion: %s\nMaintainer: %s\nBuildDate: %s\nGithub URL: %s", pkgVersion, pkgMaintainer, pkgBuildDate, pkgUrl)
}

func VerboseLevel() int {
	if Args.Verbose {
		return 3
	}
	return 0
}

func main() {
	if level, err := log.GetLevel(VerboseLevel()); err != nil {
		log.Fatalf("Verbose level is incorrect %s", err)
		os.Exit(1)
	} else {
		log.SetLevel(level)
	}
	pa := arg.MustParse(&Args)
	if Args.Domain == "" {
		pa.Fail("You must provide domain name via -d or --domain")
	}
	whoisData, err := whois.Whois(Args.Domain)
	if err != nil {
		log.Fatalf("Failed get %s", err)
	}
	tNow := time.Now()
	result, err := whoisparser.Parse(whoisData)
	if err == nil {
		if Args.Expire {
			tExpr, _ := time.Parse("2006-01-02T15:04:05Z", result.Domain.ExpirationDate)
			tExprDays := tExpr.Sub(tNow)
			fmt.Println(int(tExprDays.Hours() / 24))
		}
		if Args.CertExp {
			conf := &tls.Config{
				InsecureSkipVerify: true,
			}
			dialer := &net.Dialer{Timeout: 5 * time.Second}
			conn, err := tls.DialWithDialer(dialer, "tcp", Args.Domain+":443", conf)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			certs := conn.ConnectionState().PeerCertificates
			for _, cert := range certs {
				if slices.Contains(cert.DNSNames, Args.Domain) {
					tExpr := cert.NotAfter
					tExprDays := tExpr.Sub(tNow)
					fmt.Println(int(tExprDays.Hours() / 24))
				}
			}
		}
		if Args.DStatus {
			fmt.Println(result.Domain.Status)
		}
	} else {
		log.Fatal(err)
	}
}
