package db

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"

	"github.com/globalsign/mgo"
)

const (
	clientCertPath = "../../db/certs/client.pem"
	caCertPath     = "../../db/certs/ca-chain.pem"
)

type sequenceRDN struct {
	seq pkix.RDNSequence
}

func (s *sequenceRDN) getMap() (map[string]string, error) {
	ss := s.seq.String()
	elems := strings.Split(ss, ",")

	res := make(map[string]string, 0)
	for _, elem := range elems {
		r := strings.Split(elem, "=")
		if len(r) != 2 {
			return nil, fmt.Errorf("failed to parse RDN pair: expected 2, got %d", len(r))
		}
		res[r[0]] = r[1]
	}

	return res, nil
}

func getRDNSequenceFromFile(path string) (*sequenceRDN, error) {
	rest, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var block *pem.Block

	for {
		block, rest = pem.Decode(rest)
		if block == nil && rest == nil {
			return nil, fmt.Errorf("failed to parse certificate")
		}

		if block.Type == "CERTIFICATE" {
			break
		}
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return &sequenceRDN{seq: cert.Subject.ToRDNSequence()}, nil
}

func getSSLDialInfo(addrs string, database string) (*mgo.DialInfo, error) {

	p, err := filepath.Abs("./db/certs/")
	if err != nil {
		return nil, fmt.Errorf("not able to create tls connection: %v", err)
	}

	// CA certificate
	rootCerts := x509.NewCertPool()
	ca, err := ioutil.ReadFile(filepath.Join(p, "ca-chain.pem"))
	if err != nil {
		return nil, fmt.Errorf("not able to create tls connection: %v", err)
	}
	rootCerts.AppendCertsFromPEM(ca)

	// Client certificates
	clientCerts := []tls.Certificate{}
	cc, err := tls.LoadX509KeyPair(filepath.Join(p, "client.cert.pem"),
		filepath.Join(p, "client.key.pem"))
	if err != nil {
		return nil, fmt.Errorf("not able to create tls connection: %v", err)
	}
	clientCerts = append(clientCerts, cc)

	// get user from cert
	seq, err := getRDNSequenceFromFile(filepath.Join(p, "client.pem"))
	if err != nil {
		return nil, fmt.Errorf("failed to create mongodb client: %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create mongodb client: %v", err)
	}

	dialInfo := &mgo.DialInfo{
		Addrs:     strings.Split(addrs, ","),
		Database:  database,
		Username:  seq.seq.String(),
		Mechanism: "MONGODB-X509",
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), &tls.Config{
				RootCAs:      rootCerts,
				Certificates: clientCerts,
			})
			return conn, err
		},
	}

	fmt.Println(dialInfo)

	return dialInfo, nil
}
