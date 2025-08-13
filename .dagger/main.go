// A generated module for Ci functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/ci/internal/dagger"
	"strings"
)

type Ci struct{}

func (m *Ci) PublishTtl(ctx context.Context, src *dagger.Directory) (string, error) {
	ref, err := src.DockerBuild().
		Publish(ctx, "ttl.sh/caddy-proxy")
	if err != nil {
		return "", err
	}
	return ref, nil
}

func (m *Ci) PublishGhcr(ctx context.Context,
	src *dagger.Directory,
	username string,
	password *dagger.Secret,
) (string, error) {
	registry := "ghcr.io"
	ref, err := src.DockerBuild().
		WithRegistryAuth(registry, username, password).
		Publish(ctx, registry+"/"+strings.ToLower(username)+"/caddy-proxy")
	if err != nil {
		return "", err
	}
	return ref, nil
}
