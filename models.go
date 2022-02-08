package zignsec

// ZWInitConfig defines the configuration which are sent along with web-based Initiate().
// UserNonVisibleData is signed, if present. UserVisibleData is then only for user guidance.
// Both must be base64 encoded.
// If decoded UserVisibleData contains certain bytes (integer value > 127?),
// Init() will fail with INVALID_PARAMETERS.
type ZWInitConfig struct {
	Personalnumber     string `json:"personalnumber,omitempty"`
	UserVisibleData    string `json:"userVisibleData,omitempty"`
	UserNonVisibleData string `json:"userNonVisibleData,omitempty"`
	Relaystate         string `json:"relaystate,omitempty"`
	Configid           string `json:"configid,omitempty"`
	Target             string `json:"target,omitempty"`
}

// ZWInitRespBody defines the response Body of Initiate()
type ZWInitRespBody struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	RedirectURL string `json:"redirect_url"`
}

// ZWVerifyRespBody defines the response Body of Verify()
type ZWVerifyRespBody struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
	Result struct {
		Identity struct {
			State string `json:"state"`
		} `json:"identity"`
	}
	Identity struct {
		CountryCode    string `json:"CountryCode"`
		FirstName      string `json:"FirstName"`
		LastName       string `json:"LastName"`
		PersonalNumber string `json:"PersonalNumber"`
		DateOfBirth    string `json:"DateOfBirth"`
		Age            int    `json:"Age"`
	} `json:"identity"`
	Signature string `json:"signature"`
}

// ZSInitConfig defines the configuration which are sent along with s2s Initiate().
// UserNonVisibleData is signed, if present. UserVisibleData is then only for user guidance.
// Both must be base64 encoded.
// If decoded UserVisibleData contains certain bytes (integer value > 127?),
// Init() will fail with INVALID_PARAMETERS.
type ZSInitConfig struct {
	Personalnumber      *string `json:"personalnumber,omitempty"`
	EndUserInfo         *string `json:"endUserInfo,omitempty"`
	UserVisibleData     *string `json:"userVisibleData,omitempty"`
	UserNonVisibleData  *string `json:"userNonVisibleData,omitempty"`
	LookupPersonAddress *string `json:"lookupPersonAddress,omitempty"`
	Relaystate          *string `json:"relaystate,omitempty"`
	Webhook             *string `json:"webhook,omitempty"`
	Requirement         *struct {
		AutoStartTokenRequired string `json:"AutoStartTokenRequired"`
	} `json:"requirement,omitempty"`
	EndUserIp string `json:"endUserIp,omitempty"`
}

// ZSCancelBody is used when canceling a pending auth/sign request
type ZSCancelBody struct {
	OrderRef string `json:"orderRef"`
}

// ZSCancelResponse defines the response Body of S2S Cancel()
type ZSCancelResponse struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
}

const (
	// CollectProgressStatusOutstanding ...
	CollectProgressStatusOutstanding = "OUTSTANDING_TRANSACTION"
	// CollectProgressStatusNoClient ...
	CollectProgressStatusNoClient = "NO_CLIENT"
	// CollectProgressStatusStarted ...
	CollectProgressStatusStarted = "STARTED"
	// CollectProgressStatusUserSign ...
	CollectProgressStatusUserSign = "USER_SIGN"
	// CollectProgressStatusComplete ...
	CollectProgressStatusComplete = "COMPLETE"
)

const (
	// CollectErrorInvalidParameters ...
	CollectErrorInvalidParameters = "INVALID_PARAMETERS"
	// CollectErrorReqPrecond ...
	CollectErrorReqPrecond = "REQ_PRECOND"
	// CollectErrorReqError ...
	CollectErrorReqError = "REQ_ERROR"
	// CollectErrorReqBlocked ...
	CollectErrorReqBlocked = "REQ_BLOCKED"
	// CollectErrorInternalError ...
	CollectErrorInternalError = "INTERNAL_ERROR"
	// CollectErrorRetry ...
	CollectErrorRetry = "RETRY"
	// CollectErrorAccessDeniedRP ...
	CollectErrorAccessDeniedRP = "ACCESS_DENIED_RP"
	// CollectErrorClientErr ...
	CollectErrorClientErr = "CLIENT_ERR"
	// CollectErrorExpiredTransaction ...
	CollectErrorExpiredTransaction = "EXPIRED_TRANSACTION"
	// CollectErrorCertificateErr ...
	CollectErrorCertificateErr = "CERTIFICATE_ERR"
	// CollectErrorUSerCancel ...
	CollectErrorUSerCancel = "USER_CANCEL"
	// CollectErrorCancelled ...
	CollectErrorCancelled = "CANCELLED"
	// CollectErrorStartFailed ...
	CollectErrorStartFailed = "START_FAILED"
	// CollectErrorAleadyCollected ...
	CollectErrorAleadyCollected = "ALREADY_COLLECTED"
)

// ZSCollectResponse defines the response Body of S2S Collect()
type ZSCollectResponse struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	ProgressStatus string `json:"progressStatus"`
	Signature      string `json:"signature"`
	OCSPResponse   string `json:"ocspResponse"`
	UserInfo       struct {
		PersonalNumber string `json:"personalNumber"`
		GivenName      string `json:"givenName"`
		Surname        string `json:"surname"`
		Name           string `json:"name"`
		NotBefore      string `json:"notBefore"`
		NotAfter       string `json:"notAfter"`
		IPAddress      string `json:"ipAddress"`
	} `json:"userInfo"`
	RelayState string `json:"relayState"`
}

// InitResponse defines the response Body of S2S Init()
type ZSInitResponse struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	OrderRef       string  `json:"orderRef"`
	AutoStartToken string  `json:"autoStartToken"`
	QRCodeLink     *string `json:"qrCodeLink,omitempty"`
}
