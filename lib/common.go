package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var remoteClient = &http.Client{
	Timeout: 30 * time.Second,
}

func GetRemoteURLContent(url string) ([]byte, error) {
	body, err := GetRemoteURLReader(url)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	return io.ReadAll(body)
}

func GetRemoteURLReader(url string) (io.ReadCloser, error) {
	resp, err := remoteClient.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get remote content -> %s: %s", url, resp.Status)
	}

	return resp.Body, nil
}

func IsRemoteURL(uri string) bool {
	uri = strings.TrimSpace(uri)
	return (len(uri) >= len("http://") && strings.EqualFold(uri[:len("http://")], "http://")) ||
		(len(uri) >= len("https://") && strings.EqualFold(uri[:len("https://")], "https://"))
}

func GetIgnoreIPType(onlyIPType IPType) IgnoreIPOption {
	switch onlyIPType {
	case IPv4:
		return IgnoreIPv6
	case IPv6:
		return IgnoreIPv4
	}

	return nil
}

type WantedListExtended struct {
	TypeSlice []string
	TypeMap   map[string][]string
}

func (w *WantedListExtended) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	slice := make([]string, 0)
	mapMap := make(map[string][]string, 0)

	err := json.Unmarshal(data, &slice)
	if err != nil {
		err2 := json.Unmarshal(data, &mapMap)
		if err2 != nil {
			return err2
		}
	}

	w.TypeSlice = slice
	w.TypeMap = mapMap

	return nil
}
