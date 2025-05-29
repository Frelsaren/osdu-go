package osdu

type SchemaService service

type SchemaStatus string

const (
	PUBLISHED   SchemaStatus = "PUBLISHED"
	OBSOLETE    SchemaStatus = "OBSOLETE"
	DEVELOPMENT SchemaStatus = "DEVELOPMENT"
)

type SchemaScope string

const (
	INTERNAL SchemaScope = "INTERNAL"
	SHARED   SchemaScope = "SHARED"
)

type WriteSchemaBody struct {
	SchemaInfo SchemaInfo `json:"schemaInfo"`

	Schema interface{} `json:"schema"`
}

type SchemaIdentity struct {
	Authority          string `json:"authority"`
	Source             string `json:"source"`
	EntityType         string `json:"entityType"`
	SchemaVersionMajor int    `json:"schemaVersionMajor"`
	SchemaVersionMinor int    `json:"schemaVersionMinor"`
	SchemaVersionPatch int    `json:"schemaVersionPatch"`
	ID                 string `json:"id"`
}

type SchemaInfo struct {
	SchemaIdentity SchemaIdentity `json:"schemaIdentity"`
	CreatedBy      string         `json:"createdBy"`
	DateCreated    string         `json:"dateCreated"`
	Status         string         `json:"status"`
	Scope          string         `json:"scope"`
	SupersededBy   SchemaIdentity `json:"supersededBy"`
}

type WriteSchemaRespone struct {
	Code       int16
	SchemaInfo SchemaInfo
}
