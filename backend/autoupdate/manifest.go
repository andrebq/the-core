package autoupdate

type (
	// Manifest contains information about packages to be updated
	Manifest struct {
		Application string   `json:"application"`
		Version     string   `json:"version"`
		Bundles     []Bundle `json:"bundles"`
	}

	// Bundle indicates a downloadable content
	Bundle struct {
		Name        string `json:"name"`
		Kind        string `json:"kind"`
		OS          string `json:"os"`
		Arch        string `json:"arch"`
		DownloadURL string `json:"downloadUrl"`
		Checksum    string `json:"checksum"`
	}
)

func (m *Manifest) getBundle(name, os, arch string) (Bundle, bool) {
	for _, b := range m.Bundles {
		if b.Name == name && b.OS == os && b.Arch == arch {
			return b, true
		}
	}
	return Bundle{}, false
}
