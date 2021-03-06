package enforcer

import (
	"sync"

	"github.com/aporeto-inc/trireme/constants"
	"github.com/aporeto-inc/trireme/enforcer/lookup"
	"github.com/aporeto-inc/trireme/policy"
)

// A PolicyEnforcer is implementing the enforcer that will modify//analyze the capture packets
type PolicyEnforcer interface {

	// Enforce starts enforcing policies for the given policy.PUInfo.
	Enforce(contextID string, puInfo *policy.PUInfo) error

	// Unenforce stops enforcing policy for the given IP.
	Unenforce(contextID string) error

	// GetFilterQueue returns the current FilterQueueConfig.
	GetFilterQueue() *FilterQueue

	// Start starts the PolicyEnforcer.
	Start() error

	// Stop stops the PolicyEnforcer.
	Stop() error
}

// PublicKeyAdder register a publicKey for a Node.
type PublicKeyAdder interface {

	// PublicKeyAdd adds the given cert for the given host.
	PublicKeyAdd(host string, cert []byte) error
}

// PacketProcessor is an interface implemented to stitch into our enforcer
type PacketProcessor interface {

	// PreProcessTCPAppPacket will be called for application packets and return value of false means drop packet.
	PreProcessTCPAppPacket(pkt interface{}) bool

	// PostProcessTCPAppPacket will be called for application packets and return value of false means drop packet.
	PostProcessTCPAppPacket(pkt interface{}, action interface{}) bool

	// PreProcessTCPNetPacket will be called for network packets and return value of false means drop packet
	PreProcessTCPNetPacket(pkt interface{}) bool

	// PostProcessTCPNetPacket will be called for network packets and return value of false means drop packet
	PostProcessTCPNetPacket(pkt interface{}, action interface{}) bool
}

// PUContext holds data indexed by the PU ID
type PUContext struct {
	ID             string
	ManagementID   string
	Identity       *policy.TagsMap
	Annotations    *policy.TagsMap
	AcceptTxtRules *lookup.PolicyDB
	RejectTxtRules *lookup.PolicyDB
	AcceptRcvRules *lookup.PolicyDB
	RejectRcvRules *lookup.PolicyDB
	Extension      interface{}
	IP             string
	Mark           string
	Ports          []string
	PUType         constants.PUType
	sync.Mutex
}
