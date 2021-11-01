/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Component struct {
	Key
	Attributes ComponentAttributes `json:"attributes"`
}
type ComponentResponse struct {
	Data     Component `json:"data"`
	Included Included  `json:"included"`
}

type ComponentListResponse struct {
	Data     []Component `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustComponent - returns Component from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustComponent(key Key) *Component {
	var component Component
	if c.tryFindEntry(key, &component) {
		return &component
	}
	return nil
}
