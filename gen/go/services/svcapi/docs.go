package api

import (
	"embed"
)

//go:embed docs/*
var Content embed.FS
