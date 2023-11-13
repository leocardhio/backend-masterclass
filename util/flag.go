package util

import "flag"

type FlagArgs struct {
	Env *string
}

func DeclareFlag() FlagArgs {
	return FlagArgs{
		Env: flag.String("env", "dev", "define in what environment does the application runs"),
	}
}
