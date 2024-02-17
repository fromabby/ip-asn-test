#!/bin/bash -e

echo "Downloading ip2asn combined ipv4/ipv6 database"

curl -O https://iptoasn.com/data/ip2asn-combined.tsv.gz && gunzip ip2asn-combined.tsv.gz
