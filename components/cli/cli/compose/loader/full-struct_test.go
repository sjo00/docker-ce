package loader

import (
	"time"

	"github.com/docker/cli/cli/compose/types"
)

func fullExampleConfig(workingDir, homeDir string) *types.Config {
	return &types.Config{
		Services: services(workingDir, homeDir),
		Networks: networks(),
		Volumes:  volumes(),
	}
}

func services(workingDir, homeDir string) []types.ServiceConfig {
	return []types.ServiceConfig{
		{
			Name: "foo",

			Build: types.BuildConfig{
				Context:    "./dir",
				Dockerfile: "Dockerfile",
				Args:       map[string]*string{"foo": strPtr("bar")},
				Target:     "foo",
				Network:    "foo",
				CacheFrom:  []string{"foo", "bar"},
				Labels:     map[string]string{"FOO": "BAR"},
			},
			CapAdd:        []string{"ALL"},
			CapDrop:       []string{"NET_ADMIN", "SYS_ADMIN"},
			CgroupParent:  "m-executor-abcd",
			Command:       []string{"bundle", "exec", "thin", "-p", "3000"},
			ContainerName: "my-web-container",
			DependsOn:     []string{"db", "redis"},
			Deploy: types.DeployConfig{
				Mode:     "replicated",
				Replicas: uint64Ptr(6),
				Labels:   map[string]string{"FOO": "BAR"},
				UpdateConfig: &types.UpdateConfig{
					Parallelism:     uint64Ptr(3),
					Delay:           time.Duration(10 * time.Second),
					FailureAction:   "continue",
					Monitor:         time.Duration(60 * time.Second),
					MaxFailureRatio: 0.3,
					Order:           "start-first",
				},
				Resources: types.Resources{
					Limits: &types.Resource{
						NanoCPUs:    "0.001",
						MemoryBytes: 50 * 1024 * 1024,
					},
					Reservations: &types.Resource{
						NanoCPUs:    "0.0001",
						MemoryBytes: 20 * 1024 * 1024,
						GenericResources: []types.GenericResource{
							{
								DiscreteResourceSpec: &types.DiscreteGenericResource{
									Kind:  "gpu",
									Value: 2,
								},
							},
							{
								DiscreteResourceSpec: &types.DiscreteGenericResource{
									Kind:  "ssd",
									Value: 1,
								},
							},
						},
					},
				},
				RestartPolicy: &types.RestartPolicy{
					Condition:   "on-failure",
					Delay:       durationPtr(5 * time.Second),
					MaxAttempts: uint64Ptr(3),
					Window:      durationPtr(2 * time.Minute),
				},
				Placement: types.Placement{
					Constraints: []string{"node=foo"},
					Preferences: []types.PlacementPreferences{
						{
							Spread: "node.labels.az",
						},
					},
				},
				EndpointMode: "dnsrr",
			},
			Devices:    []string{"/dev/ttyUSB0:/dev/ttyUSB0"},
			DNS:        []string{"8.8.8.8", "9.9.9.9"},
			DNSSearch:  []string{"dc1.example.com", "dc2.example.com"},
			DomainName: "foo.com",
			Entrypoint: []string{"/code/entrypoint.sh", "-p", "3000"},
			Environment: map[string]*string{
				"FOO": strPtr("foo_from_env_file"),
				"BAR": strPtr("bar_from_env_file_2"),
				"BAZ": strPtr("baz_from_service_def"),
				"QUX": strPtr("qux_from_environment"),
			},
			EnvFile: []string{
				"./example1.env",
				"./example2.env",
			},
			Expose: []string{"3000", "8000"},
			ExternalLinks: []string{
				"redis_1",
				"project_db_1:mysql",
				"project_db_1:postgresql",
			},
			ExtraHosts: []string{
				"somehost:162.242.195.82",
				"otherhost:50.31.209.229",
			},
			HealthCheck: &types.HealthCheckConfig{
				Test:        types.HealthCheckTest([]string{"CMD-SHELL", "echo \"hello world\""}),
				Interval:    durationPtr(10 * time.Second),
				Timeout:     durationPtr(1 * time.Second),
				Retries:     uint64Ptr(5),
				StartPeriod: durationPtr(15 * time.Second),
			},
			Hostname: "foo",
			Image:    "redis",
			Ipc:      "host",
			Labels: map[string]string{
				"com.example.description": "Accounting webapp",
				"com.example.number":      "42",
				"com.example.empty-label": "",
			},
			Links: []string{
				"db",
				"db:database",
				"redis",
			},
			Logging: &types.LoggingConfig{
				Driver: "syslog",
				Options: map[string]string{
					"syslog-address": "tcp://192.168.0.42:123",
				},
			},
			MacAddress:  "02:42:ac:11:65:43",
			NetworkMode: "container:0cfeab0f748b9a743dc3da582046357c6ef497631c1a016d28d2bf9b4f899f7b",
			Networks: map[string]*types.ServiceNetworkConfig{
				"some-network": {
					Aliases:     []string{"alias1", "alias3"},
					Ipv4Address: "",
					Ipv6Address: "",
				},
				"other-network": {
					Ipv4Address: "172.16.238.10",
					Ipv6Address: "2001:3984:3989::10",
				},
				"other-other-network": nil,
			},
			Pid: "host",
			Ports: []types.ServicePortConfig{
				//"3000",
				{
					Mode:     "ingress",
					Target:   3000,
					Protocol: "tcp",
				},
				{
					Mode:     "ingress",
					Target:   3001,
					Protocol: "tcp",
				},
				{
					Mode:     "ingress",
					Target:   3002,
					Protocol: "tcp",
				},
				{
					Mode:     "ingress",
					Target:   3003,
					Protocol: "tcp",
				},
				{
					Mode:     "ingress",
					Target:   3004,
					Protocol: "tcp",
				},
				{
					Mode:     "ingress",
					Target:   3005,
					Protocol: "tcp",
				},
				//"8000:8000",
				{
					Mode:      "ingress",
					Target:    8000,
					Published: 8000,
					Protocol:  "tcp",
				},
				//"9090-9091:8080-8081",
				{
					Mode:      "ingress",
					Target:    8080,
					Published: 9090,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    8081,
					Published: 9091,
					Protocol:  "tcp",
				},
				//"49100:22",
				{
					Mode:      "ingress",
					Target:    22,
					Published: 49100,
					Protocol:  "tcp",
				},
				//"127.0.0.1:8001:8001",
				{
					Mode:      "ingress",
					Target:    8001,
					Published: 8001,
					Protocol:  "tcp",
				},
				//"127.0.0.1:5000-5010:5000-5010",
				{
					Mode:      "ingress",
					Target:    5000,
					Published: 5000,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5001,
					Published: 5001,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5002,
					Published: 5002,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5003,
					Published: 5003,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5004,
					Published: 5004,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5005,
					Published: 5005,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5006,
					Published: 5006,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5007,
					Published: 5007,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5008,
					Published: 5008,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5009,
					Published: 5009,
					Protocol:  "tcp",
				},
				{
					Mode:      "ingress",
					Target:    5010,
					Published: 5010,
					Protocol:  "tcp",
				},
			},
			Privileged: true,
			ReadOnly:   true,
			Restart:    "always",
			SecurityOpt: []string{
				"label=level:s0:c100,c200",
				"label=type:svirt_apache_t",
			},
			StdinOpen:       true,
			StopSignal:      "SIGUSR1",
			StopGracePeriod: durationPtr(time.Duration(20 * time.Second)),
			Tmpfs:           []string{"/run", "/tmp"},
			Tty:             true,
			Ulimits: map[string]*types.UlimitsConfig{
				"nproc": {
					Single: 65535,
				},
				"nofile": {
					Soft: 20000,
					Hard: 40000,
				},
			},
			User: "someone",
			Volumes: []types.ServiceVolumeConfig{
				{Target: "/var/lib/mysql", Type: "volume"},
				{Source: "/opt/data", Target: "/var/lib/mysql", Type: "bind"},
				{Source: workingDir, Target: "/code", Type: "bind"},
				{Source: workingDir + "/static", Target: "/var/www/html", Type: "bind"},
				{Source: homeDir + "/configs", Target: "/etc/configs/", Type: "bind", ReadOnly: true},
				{Source: "datavolume", Target: "/var/lib/mysql", Type: "volume"},
				{Source: workingDir + "/opt", Target: "/opt", Consistency: "cached", Type: "bind"},
				{Target: "/opt", Type: "tmpfs", Tmpfs: &types.ServiceVolumeTmpfs{
					Size: int64(10000),
				}},
			},
			WorkingDir: "/code",
		},
	}
}

func networks() map[string]types.NetworkConfig {
	return map[string]types.NetworkConfig{
		"some-network": {},

		"other-network": {
			Driver: "overlay",
			DriverOpts: map[string]string{
				"foo": "bar",
				"baz": "1",
			},
			Ipam: types.IPAMConfig{
				Driver: "overlay",
				Config: []*types.IPAMPool{
					{Subnet: "172.16.238.0/24"},
					{Subnet: "2001:3984:3989::/64"},
				},
			},
		},

		"external-network": {
			Name:     "external-network",
			External: types.External{External: true},
		},

		"other-external-network": {
			Name:     "my-cool-network",
			External: types.External{External: true},
		},
	}
}

func volumes() map[string]types.VolumeConfig {
	return map[string]types.VolumeConfig{
		"some-volume": {},
		"other-volume": {
			Driver: "flocker",
			DriverOpts: map[string]string{
				"foo": "bar",
				"baz": "1",
			},
		},
		"another-volume": {
			Name:   "user_specified_name",
			Driver: "vsphere",
			DriverOpts: map[string]string{
				"foo": "bar",
				"baz": "1",
			},
		},
		"external-volume": {
			Name:     "external-volume",
			External: types.External{External: true},
		},
		"other-external-volume": {
			Name:     "my-cool-volume",
			External: types.External{External: true},
		},
		"external-volume3": {
			Name:     "this-is-volume3",
			External: types.External{External: true},
		},
	}
}
