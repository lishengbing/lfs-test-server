package rice

// LocateMethod defines how a box is located.
type LocateMethod int

const (
	LocateFS       = LocateMethod(iota) // Locate on the filesystem.
	LocateAppended                      // Locate boxes appended to the executable.
	LocateEmbedded                      // Locate embedded boxes.
)

// Config allows customizing the box lookup behavior.
type Config struct {
	// LocateOrder defines the priority order that boxes are searched for. By
	// default, the package global FindBox searches for embedded boxes first,
	// then appended boxes, and then finally boxes on the filesystem.  That
	// search order may be customized by provided the ordered list here. Leaving
	// out a particular method will omit that from the search space. For
	// example, []LocateMethod{LocateEmbedded, LocateAppended} will never search
	// the filesystem for boxes.
	LocateOrder []LocateMethod
}

// FindBox searches for boxes using the LocateOrder of the config.
func (c *Config) FindBox(boxName string) (*Box, error) {
	return findBox(boxName, c.LocateOrder)
}
