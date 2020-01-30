/*
Copyright 2020 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package userconfig

const (
	NameKey         = "name"
	ModelKey        = "model"
	TypeKey         = "type"
	PathKey         = "path"
	PredictorKey    = "predictor"
	EndpointKey     = "endpoint"
	SignatureKeyKey = "signature_key"
	TrackerKey      = "tracker"
	ModelTypeKey    = "model_type"
	KeyKey          = "key"
	ConfigKey       = "config"
	PythonPathKey   = "python_path"
	EnvKey          = "env"

	// Compute
	ComputeKey              = "compute"
	MinReplicasKey          = "min_replicas"
	MaxReplicasKey          = "max_replicas"
	InitReplicasKey         = "init_replicas"
	TargetCPUUtilizationKey = "target_cpu_utilization"
	CPUKey                  = "cpu"
	GPUKey                  = "gpu"
	MemKey                  = "mem"
	MaxSurgeKey             = "max_surge"
	MaxUnavailableKey       = "max_unavailable"

	// Networking
	NetworkingKey             = "networking"
	TimeoutKey                = "timeout"
	LoadBalancerKey           = "load_balancer"
	APIGatewayKey             = "api_gateway"
	APIGatewayConfigKey       = "api_gateway_config"
	AuthKey                   = "auth"
	RequestsPerSecondLimitKey = "requests_per_second_limit"
	BurstLimitKey             = "burst_limit"
)
