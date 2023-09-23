package collectors

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mytlogos/netbox_application_controller/models"
)

type ApplicationCollector struct {
}

func (c *ApplicationCollector) GetName() string {
	return "Application"
}

func (c *ApplicationCollector) GetEndpoint() string {
	return "/api/agent/application"
}

func (c *ApplicationCollector) groupContainerInProjects(containers []types.Container) map[string]*Project {
	dockerAppGroups := map[string]*Project{}

	for _, container := range containers {
		var name string

		if len(container.Names) == 0 {
			panic("container without any names: " + container.ID)
		} else {
			if len(container.Names) > 1 {
				log.Printf("Container with multiple names: %v\n", container.Names)
			}
			name = strings.TrimLeft(container.Names[0], "/")
		}

		dir, ok := container.Labels["com.docker.compose.project.working_dir"]

		if ok {
			project, ok := dockerAppGroups[dir]

			if !ok {
				composeFile := container.Labels["com.docker.compose.project.config_files"]
				name := container.Labels["com.docker.compose.project"]

				project = &Project{
					Name:        name,
					WorkingDir:  dir,
					ComposeFile: composeFile,
					Services:    make(map[string]types.Container),
					IsCompose:   true,
				}
				dockerAppGroups[dir] = project
			}
			service := container.Labels["com.docker.compose.service"]
			project.Services[service] = container
		} else {
			dockerAppGroups[container.ID] = &Project{
				Name:        name,
				WorkingDir:  "",
				ComposeFile: "",
				Services: map[string]types.Container{
					name: container,
				},
				IsCompose: false,
			}
		}
	}
	return dockerAppGroups
}

func (c *ApplicationCollector) getDockerApps() []models.ApplicationGroup {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	dockerAppGroups := c.groupContainerInProjects(containers)
	groups := []models.ApplicationGroup{}

	for _, project := range dockerAppGroups {
		apps := []models.Application{}

		for _, container := range project.Services {
			inspect, err := cli.ContainerInspect(ctx, container.ID)

			if err != nil {
				panic(err)
			}

			ports := []models.ApplicationPort{}

			for port, hostBindings := range inspect.HostConfig.PortBindings {
				for _, binding := range hostBindings {
					portNumber, err := strconv.ParseInt(binding.HostPort, 10, 64)

					if err != nil {
						panic(err)
					}

					ports = append(ports, models.ApplicationPort{
						Port:              int(portNumber),
						TransportProtocol: models.TransportProtocol(port.Proto()),
						InterfaceAddress:  binding.HostIP,
					})
				}
			}
			apps = append(apps, models.Application{
				Name:         container.Names[0],
				AppManager:   models.ManagerDocker,
				Status:       models.ApplicationStatus(container.State),
				Ports:        ports,
				CpuLimits:    0,
				MemoryLimits: uint64(inspect.HostConfig.Memory),
				AdditionalInfo: models.DockerInfo{
					ContainerId:   container.ID,
					ContainerName: container.Names[0],
				},
			})
		}

		groups = append(groups, models.ApplicationGroup{
			Name:         project.Name,
			Applications: apps,
			AdditionalInfo: models.DockerGroupInfo{
				IsCompose:   project.IsCompose,
				ComposeFile: project.ComposeFile,
				WorkingDir:  project.WorkingDir,
			},
		})
	}
	return groups
}

func (c *ApplicationCollector) getOpenPorts() ([]OpenPort, error) {
	cmd := exec.Command("ss", "-tulnpH")
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}
	// normalize separating whitespace to two spaces
	whitespaceSplit := "  "
	output = regexp.MustCompile(`[^\S\r\n]+`).ReplaceAll(output, []byte(whitespaceSplit))
	scanner := bufio.NewScanner(bytes.NewReader(output))
	results := []OpenPort{}
	pidPattern := regexp.MustCompile(`users:\(\("[\w-]+",pid=(\d+),fd=\d+\)\)`)
	interfacePortPattern := regexp.MustCompile(`((\*)|(\d{1,4}(\.\d{1,4}){3})|(\[[^[\]]+\])):(\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(strings.TrimSpace(line), whitespaceSplit)

		if len(columns) != 7 {
			log.Printf("line does not contain expected number of 7 columns: [%s]\n", strings.Join(columns, ","))
			continue
		}

		match := interfacePortPattern.FindStringSubmatch(columns[4])

		if match == nil {
			log.Printf("could not match interface and port from column: %s\n", columns[4])
			continue
		}

		interfaceAddress := match[1]
		portString := match[6]
		port, err := strconv.ParseInt(portString, 10, 64)

		if err != nil {
			log.Printf("could not parse int from port: %s\n", portString)
			continue
		}

		match = pidPattern.FindStringSubmatch(columns[6])

		if match == nil {
			log.Printf("could not parse pid from column: %s\n", columns[6])
			continue
		}
		pid, err := strconv.ParseInt(match[1], 10, 64)

		if err != nil {
			log.Printf("could not parse int from pid: %s\n", match[1])
			continue
		}

		if interfaceAddress == "*" {
			results = append(results, OpenPort{
				Transport: models.TransportProtocol(columns[0]),
				Interface: "0.0.0.0",
				Port:      int(port),
				Pid:       int(pid),
			})
			results = append(results, OpenPort{
				Transport: models.TransportProtocol(columns[0]),
				Interface: "::",
				Port:      int(port),
				Pid:       int(pid),
			})
		} else {
			results = append(results, OpenPort{
				Transport: models.TransportProtocol(columns[0]),
				Interface: strings.Trim(interfaceAddress, "[]"),
				Port:      int(port),
				Pid:       int(pid),
			})
		}

	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return results, nil
}

func (c *ApplicationCollector) getSystemdApps() []models.ApplicationGroup {
	openports, err := c.getOpenPorts()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	con, err := dbus.NewSystemdConnectionContext(ctx)

	if err != nil {
		panic(err)
	}

	defer con.Close()

	groupedByUnit := map[string][]OpenPort{}

	for _, openport := range openports {
		unit, err := con.GetUnitNameByPID(ctx, uint32(openport.Pid))

		if err != nil {
			log.Printf("no systemd unit found for pid %d\n", openport.Pid)
			continue
		}
		if unit == "docker.service" {
			// ignore docker as it listens on ports for the containers
			continue
		}

		ports, ok := groupedByUnit[unit]

		if ok {
			groupedByUnit[unit] = append(ports, openport)
		} else {
			groupedByUnit[unit] = []OpenPort{openport}
		}
	}

	groups := []models.ApplicationGroup{}

	for unit, openports := range groupedByUnit {
		properties, err := con.GetAllPropertiesContext(ctx, unit)

		if err != nil {
			log.Println(err)
			continue
		}

		names := properties["Names"].([]string)
		state := properties["SubState"].(string)
		memoryHigh := properties["MemoryHigh"].(uint64)
		memoryMax := properties["MemoryMax"].(uint64)

		var memoryLimit uint64 = 0

		if memoryMax != math.MaxUint64 {
			memoryLimit = memoryMax
		} else if memoryHigh != math.MaxUint64 {
			memoryLimit = memoryHigh
		}

		appPorts := []models.ApplicationPort{}

		for _, openport := range openports {
			appPorts = append(appPorts, models.ApplicationPort{
				Port:              openport.Port,
				TransportProtocol: openport.Transport,
				InterfaceAddress:  openport.Interface,
			})
		}

		groups = append(groups, models.ApplicationGroup{
			Name: names[0],
			Applications: []models.Application{
				{
					Name:           names[0],
					AppManager:     models.ManagerSystemd,
					Status:         models.ApplicationStatus(state),
					Ports:          appPorts,
					CpuLimits:      0,
					MemoryLimits:   memoryLimit,
					AdditionalInfo: nil,
				},
			},
		})
	}

	return groups
}

func (c *ApplicationCollector) Collect(uuid string) (interface{}, error) {
	groups := c.getDockerApps()
	ag2 := c.getSystemdApps()
	groups = append(groups, ag2...)

	return models.AppUpdate{
		BasicUpdate: models.BasicUpdate{
			Uuid: uuid,
		},
		App: models.App{
			Applications: groups,
		},
	}, nil
}

type Project struct {
	Name        string
	IsCompose   bool
	WorkingDir  string
	ComposeFile string
	Services    map[string]types.Container
}

type OpenPort struct {
	Transport models.TransportProtocol
	Interface string
	Port      int
	Pid       int
}
