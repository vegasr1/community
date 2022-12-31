// Package nvhighwaysigns provides details for the NV Highway Signs applet.
package nvhighwaysigns

import (
	_ "embed"

	"tidbyt.dev/community/apps/manifest"
)

//go:embed nv_highway_signs.star
var source []byte

// New creates a new instance of the Nevada Highway Signs applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "nv-highway-signs",
		Name:        "NV Highway Signs",
		Author:      "Richard Du Bois",
		Summary:     "Displays NV highway signs",
		Desc:        "Displays messages from overhead signs on Nevada highways.",
		FileName:    "nv_highway_signs.star",
		PackageName: "nvhighwaysigns",
		Source:      source,
	}
}