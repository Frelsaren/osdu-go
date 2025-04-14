package osdu

// StorageService provides access to the Storage API
type StorageService service

type Record struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
	ACL     struct {
		Owners  []string `json:"owners"`
		Viewers []string `json:"viewers"`
	} `json:"acl"`
	Legal struct {
		LegalTags                  []string `json:"legaltags"`
		OtherRelevantDataCountries []string `json:"otherRelevantDataCountries"`
		Status                     string   `json:"status"`
	} `json:"legal"`
	Ancestry struct {
		Parents []string `json:"parents"`
	} `json:"ancestry"`
	Meta       []map[string]interface{} `json:"meta"`
	Tags       map[string]string        `json:"tags"`
	CreateUser string                   `json:"createUser"`
	CreateTime string                   `json:"createTime"`
	ModifyUser string                   `json:"modifyUser"`
	ModifyTime string                   `json:"modifyTime"`
	Data       interface{}              `json:"data"`
}
