package internal

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"yadro/client/service"
)

func CallerFromFlagSet(flags *pflag.FlagSet) (service.Caller, error) {
	addr, _ := flags.GetString("addr")
	timeout, _ := flags.GetDuration("timeout")
	rest, _ := flags.GetBool("rest")

	return service.NewCaller(addr, timeout, rest)
}

func Fatalf(format string, v ...any) {
	fmt.Printf(format, v...)
	fmt.Printf("\n")
	os.Exit(1)
}
