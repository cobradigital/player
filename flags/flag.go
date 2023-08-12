package flags

var AppName = "COBRA Core Service"
var AppVersion = "1.0.0"
var AppCommitHash = "N/A"

const (
	// Prefix for environment variables
	EnvPrefix = "COBRA"

	// Content Type Headers
	HeaderKeyContentType        = "Content-Type"
	HeaderKeyCOBRAAuthorization = "Authorization"
	HeaderKeyCOBRAAccessToken   = "X-COBRA-Access-Token"
	HeaderKeyCOBRATokenExpired  = "X-COBRA-Token-Expired"
	HeaderKeyCOBRASubject       = "X-COBRA-Subject"

	// Content Type Value
	ContentTypeJSON = "application/json; charset=utf-8"
	ContentTypeXML  = "application/xml; charset=utf-8"
	ContentTypeHTML = "text/html; charset=utf-8"

	// ACL
	ACLAuthenticatedAdmin     = "0"
	ACLAuthenticatedUser      = "1"
	ACLAuthenticatedAnonymous = "2"
	ACLEveryone               = "3"

	//Type Deposit
	DepositTopup  = "topup"
	DepositDebit  = "debit"
	DepositCredit = "credit"

	// player filter
	FilterUsername       = "username"
	FilterBank           = "bank"
	FilterNamaRekening   = "nama_rekening"
	FilterNoRekening     = "no_rekening"
	FilterDeposit        = "deposit"
	FilterStartCreatedAt = "start_created_at"
	FilterEndCreatedAt   = "end_created_at"

	JwtToken = "jwt_token"
)
