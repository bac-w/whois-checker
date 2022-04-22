# whois-and-ssl-checker
Simple golang whois checker

# Build
go build -o wchecker -ldflags "-X 'pkg.BuildDate=$(date -u +%d-%m-%Y/%H:%M)'" main.go

# Usage
#### Run help

```wchecker -h```


#### Check domain epire (in days)

```wchecker -d google.com -e```


#### Check ssl certificat epire (in days)

```wchecker -d google.com -s```