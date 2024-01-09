package common

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func StrToBase64(str string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(str))

	return sEnc
}

func CommandParameterAdd(data ...interface{}) {
	for i := 0; i < len(data); i++ {
		CommandParams[data[i].(string)] = data[i+1]
		i = i + 1
	}

}

func CommandParameterGet(key string) interface{} {
	return CommandParams[key]
}

func IntToStr(i int) string {
	s := strconv.Itoa(i)

	return s
}

func GetRandomString(n int) string {
	s := fmt.Sprintf("%x", sha256.Sum256([]byte(uuid.NewV4().String()+strconv.FormatInt(time.Now().UnixNano(), 10))))

	randBytes := make([]byte, len(s)/2)
	rand.Read(randBytes)
	s1 := fmt.Sprintf("%x", randBytes)

	return s[:n-3] + s1[15:18]
}
