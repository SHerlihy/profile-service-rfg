package database

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var DBClient *redis.Client

func Connect() {
//	dsn := os.Getenv("DSN")
//	opt, err := redis.ParseURL(dsn)

    //cert, caCertPool := tlsCert()

    opt := &redis.Options{
        Addr: os.Getenv("DATABASE_ADDRESS"),
        Username: os.Getenv("DATABASE_USER"),
        Password: os.Getenv("DATABASE_PASSWORD"),
        DB: 0,
    }
     //   TLSConfig: &tls.Config{
     //    MinVersion:   tls.VersionTLS12,
     //    Certificates: []tls.Certificate{cert},
     //    RootCAs:      caCertPool,
     //   },
//	if err != nil {
//		panic(err)
//	}

	DBClient = redis.NewClient(opt)
}

func tlsCert()(tls.Certificate, *x509.CertPool){
// Load client cert
cert, err := tls.LoadX509KeyPair("redis_user.crt", "redis_user_private.key")
if err != nil {
    log.Fatal(err)
}

// Load CA cert
caCert, err := os.ReadFile("redis_ca.pem")
if err != nil {
    log.Fatal(err)
}
caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

return cert, caCertPool
}
