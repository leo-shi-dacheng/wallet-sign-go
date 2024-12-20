package flags

import "github.com/urfave/cli/v2"

const evnVarPrefix = "SIGNATURE"

func prefixEnvVars(name string) []string {
	return []string{evnVarPrefix + "_" + name}
}

var (
	// LevelDbPathFlag Database
	LevelDbPathFlag = &cli.StringFlag{
		Name:    "master-db-host",
		Usage:   "The path of the leveldb",
		EnvVars: prefixEnvVars("LEVEL_DB_PATH"),
		Value:   "./",
	}

	// RpcHostFlag RPC Service
	RpcHostFlag = &cli.StringFlag{
		Name:     "rpc-host",
		Usage:    "The host of the rpc",
		EnvVars:  prefixEnvVars("RPC_HOST"),
		Required: true,
	}
	RpcPortFlag = &cli.IntFlag{
		Name:     "rpc-port",
		Usage:    "The port of the rpc",
		EnvVars:  prefixEnvVars("RPC_PORT"),
		Value:    8987,
		Required: true,
	}

	HsmEnable = &cli.BoolFlag{
		Name:    "hsm-enable",
		Usage:   "Hsm enable",
		EnvVars: prefixEnvVars("HSM_ENABLE"),
		Value:   false,
	}
	CredentialsFileFlag = &cli.StringFlag{
		Name:    "credentials-file",
		Usage:   "the credentials file of cloud hsm",
		EnvVars: prefixEnvVars("CREDENTIALS_FILE"),
	}
	KeyNameFlag = &cli.StringFlag{
		Name:    "key-name",
		Usage:   "The key name of cloud hsm",
		EnvVars: prefixEnvVars("KEY_NAME"),
	}
)

var requireFlags = []cli.Flag{
	RpcHostFlag,
	RpcPortFlag,
	LevelDbPathFlag,
}

var optionalFlags = []cli.Flag{
	CredentialsFileFlag,
	KeyNameFlag,
	HsmEnable,
}

func init() {
	Flags = append(requireFlags, optionalFlags...)
}

var Flags []cli.Flag
