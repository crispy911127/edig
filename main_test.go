package main

import "testing"

func TestValidateRecord(t *testing.T) {
    dnsRecords := []string{
        "SOA",
        "NS",
        "A",
        "AAAA",
        "CNAME",
        "MX",
        "SRV",
        "TXT",
        "PTR",
        "HINFO",
        "RP",
        "AFSDB",
        "SIG",
        "KEY",
        "LOC",
        "NAPTR",
        "KX",
        "CERT",
        "DNAME",
        "APL",
        "DS",
        "NSEC3",
        "NSEC3PARAM",
        "TLSA",
        "SMIMEA",
        "HIP",
        "CDS",
        "CDNSKEY",
        "OPENPGPKEY",
        "CSYNC",
        "ZONEMD",
        "SVCB",
        "HTTPS",
        "EUI48",
        "EUI64",
        "TKEY",
        "TSIG",
        "URI",
        "CAA",
        "TA",
        "DLV",
    }
    for _, record := range dnsRecords {
        recordType := valdateRecordType(record)
        if recordType != record {
            t.Errorf("Failed to validate record string")
        }
    }

}