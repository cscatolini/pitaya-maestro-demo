package main

import (
	"flag"
	"fmt"

	"strings"

	"github.com/cscatolini/pitaya-maestro-demo/servers"

	"github.com/topfreegames/extensions/jaeger"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/route"
	"github.com/topfreegames/pitaya/serialize/json"
	"github.com/topfreegames/pitaya/session"
)

func configureJaeger(svType string) {
	opts := jaeger.Options{
		Disabled:    false,
		Probability: 1.0,
		ServiceName: svType,
	}

	_, err := jaeger.Configure(opts)
	if err != nil {
		fmt.Printf("failed to configure jaeger: %s\n", err.Error())
	} else {
		fmt.Printf("configured jaeger for server: %s\n", svType)
	}
}

func configureBackend() {
	configureJaeger("room")
	room := servers.NewRoom()
	pitaya.Register(room,
		component.WithName("roomhandler"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.RegisterRemote(room,
		component.WithName("roomremote"),
		component.WithNameFunc(strings.ToLower),
	)
}

func configureFrontend(port int) {
	configureJaeger("connector")
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))

	pitaya.Register(&servers.Connector{},
		component.WithName("connectorhandler"),
		component.WithNameFunc(strings.ToLower),
	)
	pitaya.RegisterRemote(&servers.ConnectorRemote{},
		component.WithName("connectorremote"),
		component.WithNameFunc(strings.ToLower),
	)

	err := pitaya.AddRoute("room", func(
		session *session.Session,
		route *route.Route,
		payload []byte,
		servers map[string]*cluster.Server,
	) (*cluster.Server, error) {
		// will return the first server
		for k := range servers {
			return servers[k], nil
		}
		return nil, nil
	})

	if err != nil {
		fmt.Printf("error adding route %s\n", err.Error())
	}

	pitaya.AddAcceptor(tcp)
}

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")

	flag.Parse()

	defer pitaya.Shutdown()

	pitaya.SetSerializer(json.NewSerializer())

	if !*isFrontend {
		configureBackend()
	} else {
		configureFrontend(*port)
	}

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, map[string]string{})
	pitaya.Start()
}
