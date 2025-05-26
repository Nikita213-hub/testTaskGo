package helpers

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/Nikita213-hub/testTaskGo/flags"
)

func GetHostAddr() (address, port string) {
	if envAddr, ok := os.LookupEnv("ADDRESS"); ok {
		splt := strings.Split(envAddr, ":")
		if len(splt) != 2 {
			panic(errors.New("incorrect env var"))
		}
		address = splt[0]
		port = splt[1]
	} else {
		addr := flags.NewAddress("localhost", "8080")
		_ = flag.Value(addr)
		flag.Var(addr, "a", "Address in host:port fmt")
		flag.Parse()
		address = addr.GetHost()
		port = addr.GetPort()
	}
	return address, port
}
