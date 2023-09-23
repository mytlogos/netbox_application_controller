package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"

	"github.com/mytlogos/netbox_application_controller/agent/collectors"
	"github.com/mytlogos/netbox_application_controller/controllers"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getConfig() controllers.AgentConfig {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(path.Join(wd, "config.json"))

	if err != nil {
		log.Fatal(err)
	}

	var config controllers.AgentConfig
	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	if !IsUrl(config.Server) {
		log.Println(string(data))
		log.Fatalf("invalid config, '%s' is not a valid url", config.Server)
	}

	if _, err := uuid.Parse(config.Uuid); err != nil {
		log.Println(string(data))
		log.Fatalf("invalid config, '%s' is not a valid uuid", config.Uuid)
	}

	return config
}

func collect(collector collectors.Collector, config controllers.AgentConfig) {
	update, err := collector.Collect(config.Uuid)

	if err != nil {
		// TODO: report error
		log.Printf("%s could not gather data - skipping\n", collector.GetName())
		return
	}

	data, err := json.Marshal(update)

	if err != nil {
		// TODO: report error
		log.Println("could not marshal data - skipping")
		return
	}

	log.Printf("%s: Sending update: %s\n", collector.GetName(), string(data))
	resp, err := http.Post(config.Server+collector.GetEndpoint(), "application/json", bytes.NewReader(data))

	if err != nil {
		// TODO: report error
		log.Printf("%s: response error: %s - skipping\n", collector.GetName(), err)
		return
	}

	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)

	if err != nil {
		// TODO: report error
		log.Printf("%s: response read error: %s - skipping\n", collector.GetName(), err)
		return
	}

	log.Printf("%s Update Response: %s\n", collector.GetName(), string(data))
}

func simpleCollectAndPrint(collector collectors.Collector) {
	update, err := collector.Collect("none")

	if err != nil {
		// TODO: report error
		log.Printf("%s could not gather data - skipping\n", collector.GetName())
		return
	}

	data, err := json.Marshal(update)

	if err != nil {
		// TODO: report error
		log.Println("could not marshal data - skipping")
		return
	}

	log.Printf("%s: Sending update: %s\n", collector.GetName(), string(data))
}

func cliArgs() *cli.App {
	collectorList := []collectors.Collector{
		&collectors.DeviceCollector{},
		&collectors.NetworkCollector{},
		&collectors.RamCollector{},
		&collectors.ApplicationCollector{},
	}
	collectorNames := []string{}
	for _, c := range collectorList {
		collectorNames = append(collectorNames, strings.ToLower(c.GetName()))
	}
	app := &cli.App{
		Name: "netbox_application_agent",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "collector",
				Usage:    "use this flag to enable only specified collectors",
				Required: false,
				Action: func(ctx *cli.Context, s []string) error {
					for _, cliName := range s {
						found := false
						for _, c := range collectorNames {
							if strings.EqualFold(c, cliName) {
								found = true
								break
							}
						}
						if !found {
							return fmt.Errorf("unknown collector value '%s', known collectors: %v", s, collectorNames)
						}
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:  "print-only",
				Usage: "do not read config or contact server, only collect and print result as json",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "list-collector",
				Usage: "list all collectors",
				Action: func(ctx *cli.Context) error {
					for _, name := range collectorNames {
						fmt.Println(name)
					}
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			specifiedCollectorNames := ctx.StringSlice("collector")

			if len(specifiedCollectorNames) == 0 {
				specifiedCollectorNames = collectorNames
			}

			specifiedCollectors := []collectors.Collector{}

			for _, collector := range collectorList {
				for _, name := range specifiedCollectorNames {
					if strings.EqualFold(name, collector.GetName()) {
						specifiedCollectors = append(specifiedCollectors, collector)
						break
					}
				}
			}

			if ctx.Bool("print-only") {
				for _, collector := range specifiedCollectors {
					simpleCollectAndPrint(collector)
				}
			} else {
				config := getConfig()

				for _, collector := range specifiedCollectors {
					collect(collector, config)
				}
			}
			return nil
		},
	}
	return app
}

func main() {
	app := cliArgs()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
