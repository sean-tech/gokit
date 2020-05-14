package encrypt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"time"
)

// KeyPairWithPin 返回 PEM证书 and PEM-Key 和SKPI(PIN码)
// 公共证书的指纹
func GenCertKeyPairWithPin(isCA bool, parentCert *x509.Certificate) (pemCert []byte, pemKey []byte, pin []byte, err error) {

	var days time.Duration = 1024
	max := new(big.Int).Lsh(big.NewInt(1),128)  //把 1 左移 128 位，返回给 big.Int
	serialNumber, _ := rand.Int(rand.Reader, max)   //返回在 [0, max) 区间均匀随机分布的一个随机值
	subject := pkix.Name{   //Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略。
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "169.264.169.254",
	}
	// SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
	// KeyUsage 与 ExtKeyUsage 用来表明该证书是用来做服务器认证的
	tpl := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(days * 24 *time.Hour),
		BasicConstraintsValid: true,
		IsCA: 				   isCA,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		IPAddresses:   		   []net.IP{net.ParseIP("127.0.0.1")},
	}

	// privateKey generate
	// 生成一对具有指定字位数的RSA密钥
	bits := 4096
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, nil, errors.New("rsa.GenerateKey err:" + err.Error())
	}

	// cert generate
	if parentCert == nil {
		parentCert = &tpl
	}
	//CreateCertificate基于模板创建一个新的证书
	//第二个第三个参数相同，则证书是自签名的
	//返回的切片是DER编码的证书
	derCert, err := x509.CreateCertificate(rand.Reader, &tpl, parentCert, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, nil, errors.New("x509.CreateCertificate err:" + err.Error())
	}

	// cert encode
	buf := &bytes.Buffer{}
	err = pem.Encode(buf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derCert,
	})
	if err != nil {
		return nil, nil, nil, errors.New("pem.Encode err:" + err.Error())
	}
	pemCert = buf.Bytes()

	// key encode
	buf = &bytes.Buffer{}
	err = pem.Encode(buf, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return nil, nil, nil, errors.New("pem.Encode err:" + err.Error())
	}
	pemKey = buf.Bytes()

	// pin
	cert, err := x509.ParseCertificate(derCert)
	if err != nil {
		return nil, nil, nil, errors.New("x509.ParseCertificate err:" + err.Error())
	}
	pubDER, err := x509.MarshalPKIXPublicKey(cert.PublicKey.(*rsa.PublicKey))
	if err != nil {
		return nil, nil, nil, errors.New("x509.MarshalPKIXPublicKey err:" + err.Error())
	}
	sum := sha256.Sum256(pubDER)
	pin = make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
	base64.StdEncoding.Encode(pin, sum[:])

	// return
	return pemCert, pemKey, pin, nil
}