package model

type Neo4jBackupValues struct {
	NameOverride             string                   `yaml:"nameOverride,omitempty"`
	FullnameOverride         string                   `yaml:"fullnameOverride,omitempty"`
	DisableLookups           bool                     `yaml:"disableLookups" default:"false"`
	Neo4J                    Neo4jBackupNeo4j         `yaml:"neo4j"`
	Backup                   Backup                   `yaml:"backup"`
	ConsistencyCheck         ConsistencyCheck         `yaml:"consistencyCheck"`
	ServiceAccountName       string                   `yaml:"serviceAccountName"`
	TempVolume               map[string]interface{}   `yaml:"tempVolume"`
	SecurityContext          SecurityContext          `yaml:"securityContext"`
	ContainerSecurityContext ContainerSecurityContext `yaml:"containerSecurityContext,omitempty"`
	NodeSelector             map[string]string        `yaml:"nodeSelector,omitempty"`
	Resources                Neo4jBackupResources     `yaml:"resources,omitempty"`
	Tolerations              []Toleration             `yaml:"tolerations,omitempty"`
	Affinity                 Affinity                 `yaml:"affinity,omitempty"`
}

type Neo4jBackupResources struct {
	Requests Neo4jBackupRequests `yaml:"requests,omitempty"`
	Limits   Neo4jBackupLimits   `yaml:"limits,omitempty"`
}
type Neo4jBackupRequests struct {
	EphemeralStorage string `yaml:"ephemeralStorage,omitempty"`
	CPU              string `yaml:"cpu,omitempty"`
	Memory           string `yaml:"memory,omitempty"`
}
type Neo4jBackupLimits struct {
	EphemeralStorage string `yaml:"ephemeralStorage,omitempty"`
	CPU              string `yaml:"cpu,omitempty"`
	Memory           string `yaml:"memory,omitempty"`
}

type Affinity struct {
	PodAffinity PodAffinity `yaml:"podAffinity"`
}
type PodAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution []RequiredDuringSchedulingIgnoredDuringExecution `yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}
type RequiredDuringSchedulingIgnoredDuringExecution struct {
	LabelSelector LabelSelector `yaml:"labelSelector,omitempty"`
	TopologyKey   string        `yaml:"topologyKey,omitempty"`
}
type LabelSelector struct {
	MatchExpressions []MatchExpressions `yaml:"matchExpressions,omitempty"`
	MatchLabels      map[string]string  `yaml:"matchLabels,omitempty"`
}

type MatchExpressions struct {
	Key      string   `yaml:"key"`
	Operator string   `yaml:"operator"`
	Values   []string `yaml:"values"`
}

type Neo4jBackupNeo4j struct {
	Image                      string            `yaml:"image" default:"neo4jbuildservice/helm-charts"`
	ImageTag                   string            `yaml:"imageTag" default:"backup"`
	PodLabels                  map[string]string `yaml:"podLabels,omitempty"`
	PodAnnotations             map[string]string `yaml:"podAnnotations,omitempty"`
	JobSchedule                string            `yaml:"jobSchedule" default:"* * * * *"`
	SuccessfulJobsHistoryLimit int               `yaml:"successfulJobsHistoryLimit" default:"3"`
	FailedJobsHistoryLimit     int               `yaml:"failedJobsHistoryLimit" default:"1"`
	BackoffLimit               int               `yaml:"backoffLimit" default:"6"`
	Labels                     map[string]string `yaml:"labels,omitempty"`
}

type Backup struct {
	BucketName               string          `yaml:"bucketName,omitempty"`
	DatabaseAdminServiceName string          `yaml:"databaseAdminServiceName,omitempty"`
	DatabaseAdminServiceIP   string          `yaml:"databaseAdminServiceIP,omitempty"`
	DatabaseNamespace        string          `yaml:"databaseNamespace,omitempty" default:"default"`
	DatabaseBackupPort       string          `yaml:"databaseBackupPort,omitempty" default:"6362"`
	DatabaseClusterDomain    string          `yaml:"databaseClusterDomain,omitempty" default:"cluster.local"`
	DatabaseBackupEndpoints  string          `yaml:"databaseBackupEndpoints,omitempty"`
	Database                 string          `yaml:"database,omitempty"`
	AzureStorageAccountName  string          `yaml:"azureStorageAccountName,omitempty"`
	CloudProvider            string          `yaml:"cloudProvider,omitempty"`
	MinioEndpoint            string          `yaml:"minioEndpoint,omitempty"`
	SecretName               string          `yaml:"secretName,omitempty"`
	SecretKeyName            string          `yaml:"secretKeyName,omitempty"`
	PageCache                string          `yaml:"pageCache,omitempty"`
	HeapSize                 string          `yaml:"heapSize,omitempty"`
	FallbackToFull           bool            `yaml:"fallbackToFull" default:"true"`
	IncludeMetadata          string          `yaml:"includeMetadata,omitempty"`
	Type                     string          `yaml:"type,omitempty"`
	KeepFailed               bool            `yaml:"keepFailed" default:"false"`
	ParallelRecovery         bool            `yaml:"parallelRecovery" default:"false"`
	KeepBackupFiles          bool            `yaml:"keepBackupFiles" default:"true"`
	Verbose                  bool            `yaml:"verbose" default:"true"`
	AggregateBackup          AggregateBackup `yaml:"aggregate,omitempty"`
}

type AggregateBackup struct {
	Enabled          bool   `yaml:"enabled" default:"false"`
	Verbose          bool   `yaml:"verbose" default:"true"`
	KeepOldBackup    bool   `yaml:"keepOldBackup" default:"false"`
	ParallelRecovery bool   `yaml:"parallelRecovery" default:"false"`
	FromPath         string `yaml:"fromPath"`
	Database         string `yaml:"database"`
}

type ConsistencyCheck struct {
	Enable              bool   `yaml:"enable" default:"false"`
	Database            string `yaml:"database,omitempty"`
	CheckIndexes        bool   `yaml:"checkIndexes" default:"true"`
	CheckGraph          bool   `yaml:"checkGraph" default:"true"`
	CheckCounts         bool   `yaml:"checkCounts" default:"true"`
	CheckPropertyOwners bool   `yaml:"checkPropertyOwners" default:"true"`
	MaxOffHeapMemory    string `yaml:"maxOffHeapMemory,omitempty"`
	Threads             string `yaml:"threads,omitempty"`
	Verbose             bool   `yaml:"verbose" default:"true"`
}

type Toleration struct {
	Key      string `yaml:"key,omitempty"`
	Operator string `yaml:"operator,omitempty"`
	Value    string `yaml:"value,omitempty"`
	Effect   string `yaml:"effect,omitempty"`
}

type ContainerSecurityContext struct {
	RunAsNonRoot             bool         `yaml:"runAsNonRoot,omitempty"`
	RunAsUser                int64        `yaml:"runAsUser,omitempty"`
	RunAsGroup               int64        `yaml:"runAsGroup,omitempty"`
	ReadOnlyRootFilesystem   bool         `yaml:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation bool         `yaml:"allowPrivilegeEscalation,omitempty"`
	Capabilities             Capabilities `yaml:"capabilities,omitempty"`
}
