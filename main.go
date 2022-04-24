package main

import (
	"crypto/tls"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/araddon/dateparse"
	"golang.org/x/exp/slices"
	"net"
	"strings"
	"time"
	"whois-checker/pkg/log"
	"whois-checker/pkg/servers"
)

type args struct {
	Domain  string `arg:"-d,--domain,env:CHECKDOMAIN" help:"Domain name for check"`
	Verbose bool   `arg:"-v" help:"Verbose [DEBUG]"`
	Expire  bool   `arg:"-e,--expire" help:"Find out when domain it expires (days)"`
	CertExp bool   `arg:"-s,--ssl" help:"Find out when certificate it expires (days)"`
	DFull   bool   `arg:"-w, --whois" help:"Get all whois information by domain"`
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
	pa := arg.MustParse(&Args)
	log.Debug("I'm starting")
	if level, err := log.GetLevel(VerboseLevel()); err != nil {
		log.Fatalf("[LOG]: Verbose level is incorrect %s", err)
	} else {
		log.SetLevel(level)
	}
	tNow := time.Now()
	if Args.Domain != "" {
		if Args.Expire {
			resWho, err := servers.GetWhois(Args.Domain)
			if err != nil {
				log.Fatalf("[WHOIS]: %s", err)
			}
			log.Debugf("[WHOIS]: Whois information received complete")
			tExpr, _ := dateparse.ParseAny(resWho.Domain.ExpirationDate)
			tExprDays := tExpr.Sub(tNow)
			if Args.CertExp || Args.DFull {
				log.InfoIp("Domain_expire_through: ")
			}
			if resWho.Domain.ExpirationDate == "" {
				log.Debugf("[WHOIS]: Expiration date not found")
				fmt.Println(-1)
				log.Info("Use -v for verbose and check")
				if strings.Contains(Args.Domain, ".kz") || strings.Contains(Args.Domain, ".md") {
					log.Debug("Zones KZ and MD not supported. Information received " +
						"from whois servers in zone does not have an end date. View only if use website")
				}
			} else {
				fmt.Println(int(tExprDays.Hours() / 24))
			}
		}
		if Args.DFull {
			resWho, err := servers.GetWhois(Args.Domain)
			if err != nil {
				log.Fatalf("[WHOIS]: %s", err)
			}
			log.InfoI("Administrative: ")
			fmt.Println("Status: ", resWho.Domain.Status)
			fmt.Println("Created date: ", resWho.Domain.CreatedDate)
			fmt.Println("Updated date: ", resWho.Domain.UpdatedDate)
			fmt.Println("Expiration date: ", resWho.Domain.ExpirationDate)
			fmt.Println("Registrant info: ", resWho.Registrant)
			log.InfoI("Technical: ")
			fmt.Println("Name servers: ", resWho.Domain.NameServers)
			fmt.Println("Whois servers: ", resWho.Domain.WhoisServer)
		}
		if Args.CertExp {
			conf := &tls.Config{
				InsecureSkipVerify: true,
			}
			dialerTLS := &net.Dialer{Timeout: 10 * time.Second}
			conn, err := tls.DialWithDialer(dialerTLS, "tcp", Args.Domain+":443", conf)
			if err != nil {
				log.Fatalf("[SSL]: %s", err)
			}
			defer conn.Close()
			certs := conn.ConnectionState().PeerCertificates
			for _, cert := range certs {
				if slices.Contains(cert.DNSNames, Args.Domain) {
					tExpr := cert.NotAfter
					tExprDays := tExpr.Sub(tNow)
					if cert.DNSNames == nil {
						log.Fatalf("[SSL]: Certificate on domain %s not found", Args.Domain)
					}
					if Args.Expire || Args.DFull {
						log.InfoIp("Cert_expire_through: ")
					}
					fmt.Println(int(tExprDays.Hours() / 24))
				}
			}
		}
	} else {
		pa.Fail("You must provide domain name via -d or --domain")
	}
}
