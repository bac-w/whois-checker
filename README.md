# whois-and-ssl-checker
Simple golang whois checker

# Build
``go build -o wchecker -ldflags "-X 'pkg.BuildDate=$(date -u +%d-%m-%Y/%H:%M)'" main.go``

# Usage
#### Run help

``wchecker -h``

#### *REQUIRED
#### Indicate domain name*
``Use -d* to indicate domain name``

#### Check whois information

``wchecker -d google.com -w``

#### Check domain epire (in days)

``wchecker -d google.com -e``


#### Check ssl certificate expire (in days)

``wchecker -d google.com -s``
