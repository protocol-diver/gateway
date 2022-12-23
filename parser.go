package gateway

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"net/netip"
	"strconv"
	"strings"
)

func darwinParser(b []byte) (net.IP, error) {
	// '192.168.1.1\n'
	addr, err := netip.ParseAddr(string(b[:len(b)-1]))
	if err != nil {
		return nil, err
	}

	return addr.AsSlice(), nil
}

func freebsdParser(b []byte) (net.IP, error) {
	// '192.168.1.1\n'
	return darwinParser(b)
}

func linuxParser(b []byte) (net.IP, error) {
	// 'Gateway\nC900A8C0\n00000000\n'
	arr := bytes.Split(b, []byte("\n"))

	var gstr string

	for _, elem := range arr[1:] {
		if !bytes.Equal(elem, []byte("00000000")) {
			gstr = string(elem)
			break
		}
	}

	if gstr == "" {
		return nil, errors.New("not exist gateway")
	}

	gstr = "0x" + gstr

	i, err := strconv.ParseInt(gstr, 0, 64)
	if err != nil {
		return nil, err
	}

	ip := make(net.IP, 4)
	binary.LittleEndian.PutUint32(ip, uint32(i))

	return ip, nil
}

func windowsParser(b []byte) (net.IP, error) {
	// '                        0.0.0.0          0.0.0.0      192.168.1.1    192.168.1.100     20'
	s := string(b)

	var rawAddr string
	arr := strings.Fields(s)
	switch len(arr) {
	case 5:
		rawAddr = arr[2]
	case 6:
		rawAddr = arr[3]
	}

	addr, err := netip.ParseAddr(rawAddr)
	if err != nil {
		return nil, err
	}

	return addr.AsSlice(), nil
}
