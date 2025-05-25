package types

const (
	// Event types
	EventTypeVerificationRequestCreated   = "verification_request_created"
	EventTypeVerificationRequestApproved  = "verification_request_approved"
	EventTypeVerificationRequestRejected  = "verification_request_rejected"

	// Attribute keys
	AttributeKeyRequestID            = "request_id"
	AttributeKeyUserAddress         = "user_address"
	AttributeKeyInstitutionAddress  = "institution_address"
	AttributeKeyStatus             = "status"
	AttributeKeyEvidence           = "evidence"
) 