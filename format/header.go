package format

type Header struct {
	Version string
	Layout  struct {
		Regions []RegionInfo
	}
}
