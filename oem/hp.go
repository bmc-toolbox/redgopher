package oem

// Hpe declares HPE OEM fields.
type Hpe struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType      string `json:"@odata.type"`
	LoginName      string `json:"LoginName"`
	ServiceAccount bool   `json:"ServiceAccount"`
	Privileges     `json:"Privileges"`
}

type Privileges struct {
	HostBIOSConfigPriv       bool `json:"HostBIOSConfigPriv"`
	HostNICConfigPriv        bool `json:"HostNICConfigPriv"`
	HostStorageConfigPriv    bool `json:"HostStorageConfigPriv"`
	LoginPriv                bool `json:"LoginPriv"`
	RemoveConsolePriv        bool `json:"RemoteConsolePriv"`
	SystemRecoveryConfigPriv bool `json:"SystemRecoveryConfigPriv"`
	UserConfigPriv           bool `json:"UserConfigPriv"`
	VirtualMediaPriv         bool `json:"VirtualMediaPriv"`
	VirtualPowerAndResetPriv bool `json:"VirtualPowerAndResetPriv"`
	ILOConfigPriv            bool `json:"ILOConfigPriv"`
}
