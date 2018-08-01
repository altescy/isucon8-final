package main

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/ken39arg/isucon2018-final/bench"
)

var (
	appep        = flag.String("appep", "http://127.0.0.1:12510", "app endpoint")
	bankep       = flag.String("bankep", "http://127.0.0.1:5515", "isubank endpoint")
	logep        = flag.String("logep", "http://127.0.0.1:5516", "isulog endpoint")
	internalbank = flag.String("internalbank", "http://127.0.0.1:5515", "isubank endpoint")
	log          = bench.NewLogger(os.Stderr)
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bc, err := bench.NewContext(os.Stderr, *appep, *bankep, *logep, *internalbank)
	if err != nil {
		return err
	}
	bm := bench.NewRunner(bc, time.Minute, time.Second*1)
	if err = bm.Run(context.Background()); err != nil {
		return err
	}
	bm.Result()
	return nil
}

func init() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
}