package whois

import (
	"strings"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	scenarios := []struct {
		domain  string
		wantErr bool
	}{
		{
			domain:  "name.com",
			wantErr: false,
		},
		{
			domain:  "name.org",
			wantErr: false,
		},
		{
			domain:  "name.net",
			wantErr: false,
		},
		{
			domain:  "name.sh",
			wantErr: false,
		},
		{
			domain:  "name.io",
			wantErr: false,
		},
		{
			domain:  "name.dev",
			wantErr: false,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.domain+"_Query", func(t *testing.T) {
			output, err := NewClient().Query(scenario.domain)
			if scenario.wantErr && err == nil {
				t.Error("expected error, got none")
				t.FailNow()
			}
			if !scenario.wantErr && err != nil {
				t.Error("expected no error, got", err.Error())
			}
			if !strings.Contains(strings.ToLower(output), scenario.domain) {
				t.Errorf("expected %s in output, got %s", scenario.domain, output)
			}
		})
		time.Sleep(50 * time.Millisecond) // Give the WHOIS servers some breathing room
		t.Run(scenario.domain+"_QueryAndParse", func(t *testing.T) {
			response, err := NewClient().QueryAndParse(scenario.domain)
			if scenario.wantErr && err == nil {
				t.Error("expected error, got none")
				t.FailNow()
			}
			if !scenario.wantErr && err != nil {
				t.Error("expected no error, got", err.Error())
			}
			if response.ExpirationDate.Unix() == 0 {
				t.Errorf("expected to have an expiry date")
			}
			if len(response.NameServers) == 0 {
				t.Errorf("expected to have at least one name server")
			}
			if len(response.DomainStatuses) == 0 {
				t.Errorf("expected to have at least one domai status")
			}
		})
		time.Sleep(50 * time.Millisecond) // Give the WHOIS servers some breathing room
	}
}
