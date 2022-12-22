package gateway

import (
	"net"
	"os/exec"
	"testing"
)

func TestDarwin(t *testing.T) {
	route := exec.Command("sh", "-c", "cat ./darwin.txt | grep gateway | awk '{print $2}'")
	b, err := route.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	ip, err := darwinParser(b)
	if err != nil {
		t.Fatal(err)
	}
	if !ip.Equal(net.IPv4(192, 168, 1, 1)) {
		t.Fatalf("invalid IP: want %v, got: %v", net.IPv4(192, 168, 1, 1), ip)
	}
}

func TestFreeBSD(t *testing.T) {
	route := exec.Command("sh", "-c", "cat ./freebsd.txt| grep default | awk '{print $2}'")
	b, err := route.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	ip, err := freebsdParser(b)
	if err != nil {
		t.Fatal(err)
	}
	if !ip.Equal(net.IPv4(10, 88, 88, 2)) {
		t.Fatalf("invalid IP: want %v, got: %v", net.IPv4(10, 88, 88, 2), ip)
	}
}

func TestLinux(t *testing.T) {
	route := exec.Command("sh", "-c", "cat ./linux.txt | awk '{print $3}'")
	b, err := route.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	ip, err := linuxParser(b)
	if err != nil {
		t.Fatal(err)
	}
	if !ip.Equal(net.IPv4(192, 168, 0, 201)) {
		t.Fatalf("invalid IP: want %v, got: %v", net.IPv4(192, 168, 0, 201), ip)
	}
}

func TestWindows(t *testing.T) {
	// findstr -> grep
	route := exec.Command("sh", "-c", "cat ./windows.txt | grep 0.0.0.0")
	b, err := route.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	ip, err := windowsParser(b)
	if err != nil {
		t.Fatal(err)
	}
	if !ip.Equal(net.IPv4(192, 168, 1, 1)) {
		t.Fatalf("invalid IP: want %v, got: %v", net.IPv4(192, 168, 1, 1), ip)
	}
}
