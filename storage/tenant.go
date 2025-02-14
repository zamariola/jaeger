// Copyright (c) 2022 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import "context"

// tenantKeyType is a custom type for the key "tenant", following context.Context convention
type tenantKeyType string

const (
	// tenantKey holds tenancy for spans
	tenantKey = tenantKeyType("tenant")
)

// WithTenant creates a Context with a tenant association
func WithTenant(ctx context.Context, tenant string) context.Context {
	return context.WithValue(context.Background(), tenantKey, tenant)
}

// GetTenant retrieves a tenant associated with a Context
func GetTenant(ctx context.Context) string {
	tenant := ctx.Value(tenantKey)
	if tenant == nil {
		return ""
	}

	if s, ok := tenant.(string); ok {
		return s
	}
	return ""
}
