package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"net"
	"testing"
	"time"
)

func TestCa(t *testing.T) {
	var days time.Duration = 1024
	max := new(big.Int).Lsh(big.NewInt(1),128)  //把 1 左移 128 位，返回给 big.Int
	serialNumber, _ := rand.Int(rand.Reader, max)   //返回在 [0, max) 区间均匀随机分布的一个随机值
	subject := pkix.Name{   //Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略。
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "169.264.169.254",
	}
	tpl := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(days * 24 *time.Hour),
		BasicConstraintsValid: true,
		IsCA: 				   true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		IPAddresses:   		   []net.IP{net.ParseIP("127.0.0.1")},
	}


	// privateKey generate
	bits := 4096
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatalf("rsa.GenerateKey err:" + err.Error())
	}
	derCert1, err := x509.CreateCertificate(rand.Reader, &tpl, &tpl, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("x509.CreateCertificate err:" + err.Error())
	}
	derCert2, err := x509.CreateCertificate(rand.Reader, &tpl, &tpl, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("x509.CreateCertificate err:" + err.Error())
	}

	// is equal
	if string(derCert1) == string(derCert2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}


