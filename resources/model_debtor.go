/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Debtor struct {
	Key
	Attributes DebtorAttributes `json:"attributes"`
}
type DebtorResponse struct {
	Data     Debtor   `json:"data"`
	Included Included `json:"included"`
}

type DebtorListResponse struct {
	Data     []Debtor `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustDebtor - returns Debtor from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDebtor(key Key) *Debtor {
	var debtor Debtor
	if c.tryFindEntry(key, &debtor) {
		return &debtor
	}
	return nil
}
