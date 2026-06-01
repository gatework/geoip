package lib

import "testing"

func TestContainerLookupSkipsEntriesWithoutRequestedIPType(t *testing.T) {
	container := NewContainer()

	ipv4Entry := NewEntry("ipv4")
	if err := ipv4Entry.AddPrefix("192.0.2.0/24"); err != nil {
		t.Fatalf("AddPrefix IPv4 returned error: %v", err)
	}
	if err := container.Add(ipv4Entry); err != nil {
		t.Fatalf("Add IPv4 returned error: %v", err)
	}

	ipv6Entry := NewEntry("ipv6")
	if err := ipv6Entry.AddPrefix("2001:db8::/32"); err != nil {
		t.Fatalf("AddPrefix IPv6 returned error: %v", err)
	}
	if err := container.Add(ipv6Entry); err != nil {
		t.Fatalf("Add IPv6 returned error: %v", err)
	}

	result, found, err := container.Lookup("2001:db8::1")
	if err != nil {
		t.Fatalf("Lookup returned error: %v", err)
	}
	if !found {
		t.Fatal("Lookup did not find IPv6 entry")
	}
	if len(result) != 1 || result[0] != "IPV6" {
		t.Fatalf("Lookup result = %v, want [IPV6]", result)
	}
}
