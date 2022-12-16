package api

type Namespace struct {
	Name   string
	Status string
	Code   string
}

type Pod struct {
	Name   string
	Labels string
	Status string
}

type Deployment struct {
	Name   string
	Labels string
	Ready  string
}

type Service struct {
	Name       string
	Ports      string
	TargetPort string
	Type       string
}

type Secret struct {
}

type ConfigMap struct {
}

type IngressRule struct {
	Name     string
	HostPath string
	HostName string
	Service  *Service
}

type NetworkPolicy struct {
	Egress  *Egress
	Ingress *Ingress
}

type Egress struct {
}

type Ingress struct {
}

type ServiceAccount struct {
}

type Clusterrole struct {
}

type ClusterroleBinding struct {
}

type Role struct {
}

type RoleBinding struct {
}

type PV struct {
}

type PVC struct {
}

type storageClass struct {
}