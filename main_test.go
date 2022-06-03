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
    }
    for _, record := range dnsRecords {
        recordType := valdateRecordType(record)
        if recordType != record {
            t.Errorf("Failed to validate record string")
        }
    }

}