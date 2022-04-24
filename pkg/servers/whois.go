package servers

import (
	_ "embed"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"strings"
	"whois-checker/pkg/log"
)

func GetWhois(domain string) (whoisparser.WhoisInfo, error) {
	var resWho whoisparser.WhoisInfo
	var err error
	checkWhoisData := 0
checkWhois:
	whoisDataTMP, err := whois.Whois(domain)
	if err != nil {
		return resWho, err
	}
	resWho, err = whoisparser.Parse(whoisDataTMP)
	log.Debug("[WHOIS]: Check whois server for domain")
	if err != nil {
		if strings.Contains(err.Error(), "domain is not found") {
			if strings.Count(domain, ".") >= 2 {
				domainP := strings.Split(domain, ".")
				mainDomain, rootDomain := domainP[len(domainP)-2], domainP[len(domainP)-1]
				log.Debugf("[WHOIS]: Subdomain not found %s. Check root domain: %s.", domain, mainDomain+"."+rootDomain)
				domain = mainDomain + "." + rootDomain
				whoisDataTMP, err = whois.Whois(mainDomain + "." + rootDomain)
				if err != nil {
					return resWho, err
				}
				resWho, err = whoisparser.Parse(whoisDataTMP)
				if err != nil {
					return resWho, err
				}
			}
		} else {
			log.Debug(whoisDataTMP)
			log.Fatalf("[WHOIS]: ERROR: %s", err.Error())
		}
	}
	whoisData, err := whois.Whois(domain, resWho.Domain.WhoisServer)
	resWho, err = whoisparser.Parse(whoisData)
	if err != nil {
		if strings.Contains(err.Error(), "domain whois data is invalid") {
			if checkWhoisData >= 0 && checkWhoisData < 2 {
				checkWhoisData++
				goto checkWhois
			} else if checkWhoisData <= 3 {
				log.Info("Use -v for verbose and check")
				log.Debug("[WHOIS]: Rate limit exceeded")
			}
		}
		return resWho, err
	}
	log.Debugf("[WHOIS]: Set whois server: %s", resWho.Domain.WhoisServer)
	return resWho, nil
}
