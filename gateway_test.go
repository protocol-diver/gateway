package gateway

import (
	"net"
	"os/exec"
	"testing"
)

func TestDefault(t *testing.T) {
	if _, err := Default("some"); err == nil {
		t.Fatal("not supported OS but returns success")
	}
}

func TestDarwin(t *testing.T) {
	route := exec.Command("sh", "-c", "cat ./terminal/darwin.txt | grep gateway | awk '{print $2}'")
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
	route := exec.Command("sh", "-c", "cat ./terminal/freebsd.txt| grep default | awk '{print $2}'")
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
	route := exec.Command("sh", "-c", "cat ./terminal/linux.txt | awk '{print $3}'")
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
	route := exec.Command("sh", "-c", "cat ./terminal/windows.txt | grep 0.0.0.0")
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
