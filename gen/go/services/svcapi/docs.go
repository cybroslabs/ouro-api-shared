package svcapi

import (
	"embed"
)

//go:embed docs/* raw.binpb
var Content embed.FS
