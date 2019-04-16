package goxel

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/net/proxy"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return uint(100)
	}
	return uint(ws.Col)
}

// NewClient returns a HTTP client with the requested configuration
// It supports HTTP and SOCKS proxies
func NewClient() (*http.Client, error) {
	client := &http.Client{}

	if proxyURL != "" {
		re := regexp.MustCompile(`^(http|https|socks5)://`)
		protocol := re.Find([]byte(proxyURL))

		if protocol != nil {
			var transport *http.Transport

			if string(protocol) == "http://" || string(protocol) == "https://" {
				pURL, err := url.Parse(proxyURL)
				fmt.Println(pURL)
				if err != nil {
					return client, errors.New("Invalid proxy URL")
				}

				transport = &http.Transport{
					Proxy: http.ProxyURL(pURL),
				}
			} else if string(protocol) == "socks5://" {
				dialer, err := proxy.SOCKS5("tcp", strings.Replace(proxyURL, "socks5://", "", 1), nil, proxy.Direct)
				if err != nil {
					return client, errors.New("Invalid proxy URL")
				}
				transport = &http.Transport{
					Dial: dialer.Dial,
				}
			} else {
				return client, errors.New("Invalid proxy protocol")
			}

			if transport != nil {
				client = &http.Client{
					Transport: transport,
				}
			}
		} else {
			return client, errors.New("Invalid proxy URL")
		}

	}

	return client, nil
}
