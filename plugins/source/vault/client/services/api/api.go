// Code generated by codegen; DO NOT EDIT.
package api

import (
	"context"
	"github.com/hashicorp/vault/api"
)

//go:generate mockgen -package=mocks -destination=../mocks/api.go -source=api.go Sys
type Sys interface {
	GetPlugin(*api.GetPluginInput) (*api.GetPluginResponse, error)
	GetPluginWithContext(context.Context, *api.GetPluginInput) (*api.GetPluginResponse, error)
	GetPolicy(string) (string, error)
	GetPolicyWithContext(context.Context, string) (string, error)
	ListAudit() (map[string]*api.Audit, error)
	ListAuditWithContext(context.Context) (map[string]*api.Audit, error)
	ListAuth() (map[string]*api.MountOutput, error)
	ListAuthWithContext(context.Context) (map[string]*api.MountOutput, error)
	ListMounts() (map[string]*api.MountOutput, error)
	ListMountsWithContext(context.Context) (map[string]*api.MountOutput, error)
	ListPlugins(*api.ListPluginsInput) (*api.ListPluginsResponse, error)
	ListPluginsWithContext(context.Context, *api.ListPluginsInput) (*api.ListPluginsResponse, error)
	ListPolicies() ([]string, error)
	ListPoliciesWithContext(context.Context) ([]string, error)
}
