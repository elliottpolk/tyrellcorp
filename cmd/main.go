package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/elliottpolk/tyrellcorp/config"
	"github.com/elliottpolk/tyrellcorp/grpc"
	"github.com/elliottpolk/tyrellcorp/rest"

	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v2"
	altsrc "gopkg.in/urfave/cli.v2/altsrc"
)

const (
	PanicLevel string = "panic"
	FatalLevel string = "fatal"
	ErrorLevel string = "error"
	WarnLevel  string = "warn"
	InfoLevel  string = "info"
	DebugLevel string = "debug"
	TraceLevel string = "trace"
)

var (
	version  string
	compiled string = fmt.Sprint(time.Now().Unix())

	RpcPortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "rpc-port",
		Value:   "7000",
		Usage:   "RPC port to listen on",
		EnvVars: []string{"TYRELLCORP_RPC_PORT"},
	})

	HttpPortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "http-port",
		Value:   "8080",
		Usage:   "HTTP port to listen on",
		EnvVars: []string{"TYRELLCORP_HTTP_PORT"},
	})

	HttpsPortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "tls-port",
		Value:   "8443",
		Usage:   "HTTPS port to listen on",
		EnvVars: []string{"TYRELLCORP_HTTPS_PORT"},
	})

	TlsCertFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "tls-cert",
		Usage:   "TLS certificate file for HTTPS",
		EnvVars: []string{"TYRELLCORP_TLS_CERT"},
	})

	TlsKeyFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "tls-key",
		Usage:   "TLS key file for HTTPS",
		EnvVars: []string{"TYRELLCORP_TLS_KEY"},
	})

	DatastoreAddrFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-addr",
		Aliases: []string{"ds-addr", "dsa"},
		Usage:   "Database address",
	})

	DatastorePortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-port",
		Aliases: []string{"ds-port", "dsp"},
		Value:   "27017",
		Usage:   "Database port",
	})

	DatastoreNameFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-name",
		Aliases: []string{"ds-name", "dsn"},
		Value:   "tyrellco",
		Usage:   "Database name",
	})

	DatastoreUserFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-user",
		Aliases: []string{"ds-user", "dsu"},
		Usage:   "Database user",
	})

	DatastorePasswordFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-password",
		Aliases: []string{"ds-password", "dspwd"},
		Usage:   "Database password",
	})

	GitLabTokenFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "gitlab-access-token",
		Aliases: []string{"gl-token"},
		Usage:   "private access token for GitLab",
	})

	GitLabSSHKeyFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "gitlab-ssh-key",
		Aliases: []string{"gl-ssh-key"},
		Usage:   "private ssh key for GitLab",
	})

	GitLabAddrFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "gitlab-addr",
		Aliases: []string{"gl-addr"},
		Usage:   "URL for GitLab for source code management",
	})

	GitLabGroupFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "gitlab-group",
		Aliases: []string{"gitlab-namespace", "gl-group", "gl-namespace"},
		Usage:   "GitLab group the source code will live in",
	})

	CfgFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c", "cfg", "confg"},
		Usage:   "optional path to config file",
	})

	LogVerbosityFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "log-verbosity",
		Aliases: []string{"verbosity", "verbose", "log-level", "lvl", "ll"},
		Value:   InfoLevel,
		Usage:   "logging level",
	})

	flags = []cli.Flag{
		RpcPortFlag,
		HttpPortFlag,
		HttpsPortFlag,
		TlsCertFlag,
		TlsKeyFlag,
		CfgFlag,
		DatastoreAddrFlag,
		DatastorePortFlag,
		DatastoreNameFlag,
		DatastoreUserFlag,
		DatastorePasswordFlag,
		GitLabTokenFlag,
		GitLabSSHKeyFlag,
		GitLabAddrFlag,
		GitLabGroupFlag,
		LogVerbosityFlag,
	}
)

func main() {
	ct, err := strconv.ParseInt(compiled, 0, 0)
	if err != nil {
		panic(err)
	}

	app := cli.App{
		Name:      "tyrellcorp",
		Copyright: "Copyright © 2019 Elliott Polk",
		Version:   version,
		Compiled:  time.Unix(ct, -1),
		Flags:     flags,
		Before: func(ctx *cli.Context) error {
			if len(ctx.String(CfgFlag.Name)) > 0 {
				return altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))(ctx)
			}
			return nil
		},
		Action: func(ctx *cli.Context) error {
			// set logging level
			log.SetLevel(getLvl(ctx.String(LogVerbosityFlag.Name)))

			// read in the configuration
			comp := &config.Composition{
				Server: &config.ServerCfg{
					RpcPort:   ctx.String(RpcPortFlag.Name),
					HttpPort:  ctx.String(HttpPortFlag.Name),
					HttpsPort: ctx.String(HttpsPortFlag.Name),
					TlsCert:   ctx.String(TlsCertFlag.Name),
					TlsKey:    ctx.String(TlsKeyFlag.Name),
				},
				Db: &config.DbCfg{
					Addr:     ctx.String(DatastoreAddrFlag.Name),
					Port:     ctx.String(DatastorePortFlag.Name),
					DbName:   ctx.String(DatastoreNameFlag.Name),
					User:     ctx.String(DatastoreUserFlag.Name),
					Password: ctx.String(DatastorePasswordFlag.Name),
				},
				GitLab: &config.GitLabCfg{
					APIToken: ctx.String(GitLabTokenFlag.Name),
					SSHKey:   ctx.String(GitLabSSHKeyFlag.Name),
					Addr:     ctx.String(GitLabAddrFlag.Name),
					Group:    ctx.String(GitLabGroupFlag.Name),
				},
			}

			// run in a non-blocking goroutine since it is blocking
			go func() {
				if err := rest.Serve(context.Background(), comp); err != nil {
					log.Fatal(err)
				}
			}()

			// use this one to block and prevent exiting
			if err := grpc.Serve(context.Background(), comp); err != nil {
				return cli.Exit(err, 1)
			}

			return nil
		},
	}

	app.Run(os.Args)
}

func getLvl(want string) log.Level {
	lvl := log.InfoLevel

	switch want {
	case PanicLevel:
		lvl = log.PanicLevel

	case FatalLevel:
		lvl = log.FatalLevel

	case ErrorLevel:
		lvl = log.ErrorLevel

	case WarnLevel:
		lvl = log.WarnLevel

	case DebugLevel:
		lvl = log.DebugLevel

	case TraceLevel:
		lvl = log.TraceLevel

	default:
		lvl = log.InfoLevel
	}

	return lvl
}
