package gateway

import (
	"errors"
	"net"
	"os/exec"
)

func Default(os string) (net.IP, error) {
	switch os {
	case "darwin":
		return darwinGatewayDiscover()
	case "freebsd":
		return freebsdGatewayDiscover()
	case "linux":
		return linuxGatewayDiscover()
	case "windows":
		return windowsGatewayDiscover()
	}
	return nil, errors.New("not supported OS: " + os)
}

func darwinGatewayDiscover() (net.IP, error) {
	route := exec.Command("sh", "-c", "/sbin/route -n get 0.0.0.0 | grep gateway | awk '{print $2}'")
	b, err := route.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return darwinParser(b)
}

func freebsdGatewayDiscover() (net.IP, error) {
	route := exec.Command("sh", "-c", "netstat -rn | grep default | awk '{print $2}'")
	b, err := route.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return freebsdParser(b)
}

func linuxGatewayDiscover() (net.IP, error) {
	route := exec.Command("sh", "-c", "cat /proc/net/route | awk '{print $3}'")
	b, err := route.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return linuxParser(b)
}

func windowsGatewayDiscover() (net.IP, error) {
	route := exec.Command("sh", "-c", "route print 0.0.0.0 | findstr 0.0.0.0")
	b, err := route.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return windowsParser(b)
}
