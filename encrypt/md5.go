package encrypt

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

type md5Impl struct {}

func (this *md5Impl) Encode(value []byte) string {
	m := md5.New()
	m.Write(value)
	return hex.EncodeToString(m.Sum(nil))
}

func (this *md5Impl) EncodeWithTimestamp(value []byte, timestamp int64) string {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	var buf bytes.Buffer
	buf.Write(value)
	buf.WriteString(strconv.FormatInt(timestamp, 10))
	m := md5.New()
	m.Write(buf.Bytes())
	return hex.EncodeToString(m.Sum(nil))
}

func (this *md5Impl) HmacEncode(key, value []byte) string {
	hash := hmac.New(md5.New,key)
	hash.Write(value)
	return hex.EncodeToString(hash.Sum(nil))
}

