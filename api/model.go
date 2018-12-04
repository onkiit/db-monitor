package api

import (
	"context"
)

type Store interface {
	GetVersion(context.Context) (*DBVersion, error)
	GetActiveClient(context.Context) (*DBActiveClient, error)
	GetHealth(context.Context) (*DBHealth, error)
}

type DBVersion struct {
	Version string `json:"version,omitempty"`
}

type DBActiveClient struct {
	ActiveClient int `json:"active_client,omitempty"`
}

type DBHealth struct {
	PsqlHealth  PsqlHealth  `json:"psql_health,omitempty"`
	RedisHealth RedisHealth `json:"redis_health,omitempty"`
}

type PsqlHealth struct {
	DBInformation    DBInformation      `json:"db_information,omitempty"`
	TableInformation []TableInformation `json:"table_information,omitempty"`
}

type DBInformation struct {
	DBName string `json:"db_name,omitempty"`
	DBSize string `json:"db_size,omitempty"`
}

type TableInformation struct {
	SchemaName string `json:"schema_name,omitempty"`
	TableName  string `json:"table_name,omitempty"`
	TableSize  string `json:"table_size,omitempty"`
	IndexSize  string `json:"index_size,omitempty"`
}

type RedisHealth struct {
	AvailableKey int         `json:"available_key,omitempty"`
	MemoryUsage  string      `json:"memory_usage,omitempty"`
	ExpiredKeys  string      `json:"expired_key,omitempty"`
	EvictedKeys  string      `json:"evicted_key,omitempty"`
	SlowlogCount int         `json:"slow_count"`
	MemoryStats  MemoryStats `json:"memory_stats,omitempty"`
}

type MemoryStats struct {
	PeakAllocated    int64 `json:"peak_allocated,omitempty"`
	TotalAllowed     int64 `json:"total_allowed,omitempty"`
	StartupAllocated int64 `json:"startup_allocated,omitempty"`
	PeakPercentage   int64 `json:"peak_percentage,omitempty"`
	Fragmentation    int64 `json:"fragmentation,omitempty"`
}
