// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"fmt"
)

const (
	// CassandraService const
	CassandraService = "cassandra"

	// CassandraPort const
	CassandraPort = "9042"

	// CassandraDockerImage const
	CassandraDockerImage = "cassandra"

	// CassandraDockerImageVersion const
	CassandraDockerImageVersion = "4.0"

	// CassandraRestartPolicy const
	CassandraRestartPolicy = "unless-stopped"

	// ClusterDefaultNodes const
	ClusterDefaultNodes = 3

	// ClusterDefaultDC const
	ClusterDefaultDC = "DC1"

	// ClusterDefaultRack const
	ClusterDefaultRack = "RAC1"
)

// GetCassandraConfig gets yaml definition object
func GetCassandraConfig(name string, version string, nodes int, dc, rack string) DockerComposeConfig {
	services := make(map[string]Service)

	if version == "" {
		version = CassandraDockerImageVersion
	}

	if nodes == 0 {
		nodes = ClusterDefaultNodes
	}

	if dc == "" {
		dc = ClusterDefaultDC
	}

	if rack == "" {
		rack = ClusterDefaultRack
	}

	for i := 1; i <= nodes; i++ {
		if i == 1 {
			services[fmt.Sprintf("%s%d", name, i)] = Service{
				Image:   fmt.Sprintf("%s:%s", CassandraDockerImage, version),
				Restart: CassandraRestartPolicy,
				Ports:   []string{CassandraPort},
				Environment: []string{
					fmt.Sprintf("CASSANDRA_SEEDS=%s", fmt.Sprintf("%s%d", name, i)),
					fmt.Sprintf("CASSANDRA_ENDPOINT_SNITCH=GossipingPropertyFileSnitch"),
					fmt.Sprintf("CASSANDRA_DC=%s", ClusterDefaultDC),
					fmt.Sprintf("CASSANDRA_RACK=%s", ClusterDefaultRack),
				},
			}
		} else {
			services[fmt.Sprintf("%s%d", name, i)] = Service{
				Image:   fmt.Sprintf("%s:%s", CassandraDockerImage, version),
				Restart: CassandraRestartPolicy,
				Ports:   []string{CassandraPort},
				Environment: []string{
					fmt.Sprintf("CASSANDRA_SEEDS=%s", fmt.Sprintf("%s%d", name, i)),
					fmt.Sprintf("CASSANDRA_ENDPOINT_SNITCH=GossipingPropertyFileSnitch"),
					fmt.Sprintf("CASSANDRA_DC=%s", ClusterDefaultDC),
					fmt.Sprintf("CASSANDRA_RACK=%s", ClusterDefaultRack),
				},
				DependsOn: []string{
					fmt.Sprintf("%s%d", name, 1),
				},
				Networks: []string{
					name,
				},
			}
		}
	}

	return DockerComposeConfig{
		Version:  "3",
		Services: services,
		Networks: map[string]string{name: "omitempty"},
	}
}
