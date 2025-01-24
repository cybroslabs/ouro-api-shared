package svcapi

import (
	"embed"
)

//go:embed docs/*
var Content embed.FS
