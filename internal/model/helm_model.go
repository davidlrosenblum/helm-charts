package model

import (
	"fmt"
	"os"
	"strings"
)

type ReleaseName interface {
	String() string
	Namespace() Namespace
	DiskName() PersistentDiskName
	PodName() string
	ShortName() string
	EnvConfigMapName() string
	UserConfigMapName() string
	InternalServiceName() string
	DefaultConfigMapName() string
	UserLogsConfigMapName() string
	ServerLogsConfigMapName() string
}

func NewReleaseName(name string) ReleaseName {
	r := releaseName(name)
	return &r
}

type releaseName string

func (r *releaseName) String() string {
	return string(*r)
}

func (r *releaseName) Namespace() Namespace {
	return Namespace("neo4j-" + string(*r))
}

func (r *releaseName) DiskName() PersistentDiskName {
	return PersistentDiskName(fmt.Sprintf("neo4j-data-disk-%s", *r))
}

func (r *releaseName) PodName() string {
	return string(*r) + "-0"
}

func (r *releaseName) EnvConfigMapName() string {
	return string(*r) + "-env"
}

func (r *releaseName) DefaultConfigMapName() string {
	return string(*r) + "-default-config"
}

func (r *releaseName) UserConfigMapName() string {
	return string(*r) + "-user-config"
}

func (r *releaseName) UserLogsConfigMapName() string {
	return string(*r) + "-user-logs-config"
}

func (r *releaseName) ServerLogsConfigMapName() string {
	return string(*r) + "-server-logs-config"
}
func (r *releaseName) InternalServiceName() string {
	return string(*r) + "-internals"
}
func (r *releaseName) ShortName() string {
	len := len(string(*r)) / 2
	return string(*r)[0:len]
}

func NewCoreReleaseName(clusterName ReleaseName, number int) ReleaseName {
	r := clusterMemberReleaseName{clusterName, releaseName(fmt.Sprintf("%s-core-%d", clusterName, number))}
	return &r
}

func NewLoadBalancerReleaseName(clusterName ReleaseName) ReleaseName {
	r := clusterMemberReleaseName{clusterName, releaseName(fmt.Sprintf("%s-loadbalancer", clusterName))}
	return &r
}

func NewHeadlessServiceReleaseName(clusterName ReleaseName) ReleaseName {
	r := clusterMemberReleaseName{clusterName, releaseName(fmt.Sprintf("%s-headless", clusterName))}
	return &r
}

type clusterMemberReleaseName struct {
	clusterName ReleaseName
	memberName  releaseName
}

func (r *clusterMemberReleaseName) String() string {
	return r.memberName.String()
}

func (r *clusterMemberReleaseName) Namespace() Namespace {
	return r.clusterName.Namespace()
}

func (r *clusterMemberReleaseName) DiskName() PersistentDiskName {
	return r.memberName.DiskName()
}

func (r *clusterMemberReleaseName) PodName() string {
	return r.memberName.PodName()
}

func (r *clusterMemberReleaseName) EnvConfigMapName() string {
	return r.memberName.EnvConfigMapName()
}

func (r *clusterMemberReleaseName) UserConfigMapName() string {
	return r.memberName.UserConfigMapName()
}

func (r *clusterMemberReleaseName) UserLogsConfigMapName() string {
	return r.memberName.UserLogsConfigMapName()
}
func (r *clusterMemberReleaseName) ServerLogsConfigMapName() string {
	return r.memberName.ServerLogsConfigMapName()
}

func (r *clusterMemberReleaseName) InternalServiceName() string {
	return r.memberName.InternalServiceName()
}

func (r *clusterMemberReleaseName) DefaultConfigMapName() string {
	return r.memberName.DefaultConfigMapName()
}

func (r *clusterMemberReleaseName) ShortName() string {
	len := len(string(r.memberName)) / 2
	return string(r.memberName)[0:len]
}

type Namespace string
type PersistentDiskName string

var DefaultEnterpriseValues = HelmValues{
	Neo4J: Neo4J{
		Name:                   "test",
		AcceptLicenseAgreement: "yes",
		Edition:                "enterprise",
	},
	Volumes: Volumes{
		Data: Data{
			Mode:           "selector",
			DisableSubPath: false,
		},
	},
}

var DefaultNeo4jBackupValues = Neo4jBackupValues{
	ConsistencyCheck: ConsistencyCheck{
		Enable:              true,
		CheckIndexes:        true,
		CheckGraph:          true,
		CheckCounts:         true,
		CheckPropertyOwners: true,
		Verbose:             true,
	},
	Neo4J: Neo4jBackupNeo4j{
		Image:                      strings.Split(os.Getenv("NEO4J_DOCKER_BACKUP_IMG"), ":")[0],
		ImageTag:                   strings.Split(os.Getenv("NEO4J_DOCKER_BACKUP_IMG"), ":")[1],
		JobSchedule:                "* * * * *",
		SuccessfulJobsHistoryLimit: 3,
		FailedJobsHistoryLimit:     3,
		BackoffLimit:               3,
	},
	TempVolume: map[string]interface{}{
		"emptyDir": nil,
	},
	Resources: Neo4jBackupResources{
		Requests: Neo4jBackupRequests{
			EphemeralStorage: "2Gi",
		},
		Limits: Neo4jBackupLimits{
			EphemeralStorage: "2Gi",
		},
	},
	SecurityContext: SecurityContext{
		RunAsNonRoot:        true,
		RunAsUser:           7474,
		RunAsGroup:          7474,
		FsGroup:             7474,
		FsGroupChangePolicy: "Always",
	},
	ContainerSecurityContext: ContainerSecurityContext{
		RunAsNonRoot:             true,
		RunAsUser:                7474,
		RunAsGroup:               7474,
		ReadOnlyRootFilesystem:   false,
		AllowPrivilegeEscalation: false,
		Capabilities: Capabilities{
			Drop: []string{"ALL"},
		},
	},
}

var DefaultNeo4jReverseProxyValues = Neo4jReverseProxyValues{
	ReverseProxy: ReverseProxy{
		Image: os.Getenv("NEO4J_REVERSE_PROXY_IMG"),
		Ingress: Ingress{
			Enabled: true,
			TLS: TLS{
				Enabled: true,
				Config: []Config{
					{
						SecretName: "ingress-secret",
					},
				},
			},
		},
	},
}

var DefaultCommunityValues = HelmValues{
	Neo4J: Neo4J{
		Name:    "test",
		Edition: "community",
	},
	Volumes: Volumes{
		Data: Data{
			Mode:           "selector",
			DisableSubPath: false,
		},
	},
}
