// Package service provides domain primitives for managing todo lists.
package service

import "github.com/adrianpk/cirrustodo/internal/base"

type (
	Todo struct {
		base.Worker
	}
)
