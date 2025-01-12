// Package nvhighwaysigns provides details for the NV Highway Signs applet.
package nvroads

import (
	_ "embed"

	"tidbyt.dev/community/apps/manifest"
)

//go:embed nv_highway_signs.star
var source []byte

// New creates a new instance of the Nevada Highway Signs applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "nvroads",
		Name:        "NVRoads",
		Author:      "iAmVegas",
		Summary:     "Displays NV Road signs",
		Desc:        "Displays messages from overhead signs on Nevada highways.",
		FileName:    "nvroads.star",
		PackageName: "nvroads",
		Source:      source,
	}
}
