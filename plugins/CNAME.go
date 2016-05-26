package plugins

import (
	"net"
	"strings"

	"github.com/miekg/dns"
)

type RecordCNAMEPlugin struct {
	Addr     net.IP
	RRheader dns.RR_Header
	Conf     map[string]interface{}
}

func (this *RecordCNAMEPlugin) New(edns, remote net.IP, rr_header dns.RR_Header) {
	if edns != nil {
		this.Addr = edns
	} else {
		this.Addr = remote
	}

	this.RRheader = rr_header
}

func (this *RecordCNAMEPlugin) Filter(conf map[string]interface{}) (answer []dns.RR, err error) {
	records := getBaseRecord(this.Addr, conf)
	return this.NormalRecord(records)
}

func (this *RecordCNAMEPlugin) NormalRecord(records []interface{}) (answer []dns.RR, err error) {
	var r []interface{}
	var e error
	for _, record := range records {
		r, e = getProofRecord(record)
		if e != nil {
			err = e
			continue
		}
		for _, v := range r {
			value := strings.TrimSpace(v.(string))
			answer = append(answer, &dns.CNAME{
				Hdr:    this.RRheader,
				Target: dns.Fqdn(value)})
		}
	}
	return
}

func init() {
	RegisterPlugin("CNAME", dns.TypeCNAME, func() interface{} {
		return new(RecordCNAMEPlugin)
	})
}
