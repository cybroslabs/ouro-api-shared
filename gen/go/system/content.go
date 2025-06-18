package system

import (
	embed "embed"
)

//go:embed content/license.pem
var Content embed.FS
