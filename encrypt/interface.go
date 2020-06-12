package encrypt

import "sync"

type IMd5 interface {
	Encode(value []byte) string
	EncodeWithTimestamp(value []byte, timestamp int64) string
	HmacEncode(key, value []byte) string
}

type IRsa interface {
	Encrypt(publicKey string, data []byte) ([]byte, error)
	Decrypt(privateKey string, data []byte) ([]byte, error)
	Sign(privateKey string, data []byte) ([]byte, error)
	Verify(publicKey string, data []byte, signedData []byte) error
}

type IAes interface {
	EncryptCBC(origData []byte, key []byte) ([]byte, error)
	DecryptCBC(encrypted []byte, key []byte) ([]byte, error)
	GenerateKey() []byte
}

var (
	_md5Once     sync.Once
	_md5Instance IMd5

	_rsaOnce     sync.Once
	_rsaInstance IRsa

	_aesOnce     sync.Once
	_aesInstance IAes
)

func GetMd5() IMd5 {
	_md5Once.Do(func() {
		_md5Instance = new(md5Impl)
	})
	return _md5Instance
}

func GetRsa() IRsa {
	_rsaOnce.Do(func() {
		_rsaInstance = new(rsaImpl)
	})
	return _rsaInstance
}

func GetAes() IAes {
	_aesOnce.Do(func() {
		_aesInstance = new(aesImpl)
	})
	return _aesInstance
}