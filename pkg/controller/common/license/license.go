// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package license

import logf "sigs.k8s.io/controller-runtime/pkg/log"

var log = logf.Log.WithName("license")

const (
	// FileName is the name used in the license secret to point to the license data.
	FileName = "license"
)
