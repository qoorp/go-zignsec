package zignsec

// ZSInitConfig defines the configuration which are sent along with the init request
// UserNonVisibleData is signed, if present. UserVisibleData is then only for user guidance.
// Both must be base64 encoded.
// If decoded UserVisibleData contains certain bytes (integer value > 127?),
// Init() will fail with INVALID_PARAMETERS.
type ZSInitConfig struct {
	Personalnumber     string `json:"personalnumber,omitempty"`
	UserVisibleData    string `json:"userVisibleData,omitempty"`
	UserNonVisibleData string `json:"userNonVisibleData,omitempty"`
	Relaystate         string `json:"relaystate,omitempty"`
	Configid           string `json:"configid,omitempty"`
	Target             string `json:"target,omitempty"`
}

// ZSInitRespBody defines the response Body of the init request
type ZSInitRespBody struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	RedirectURL string `json:"redirect_url"`
}

// ZSVerifyRespBody defines the response Body of the verify request
type ZSVerifyRespBody struct {
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

const (
	// CollectProgressStatusComplete completed and validated
	CollectProgressStatusComplete = "COMPLETE"
	// CollectProgressStatusUserSign signed and validated
	CollectProgressStatusUserSign = "USER_SIGN"
	// CollectProgressStatusOutstanding waiting for user to complete login
	CollectProgressStatusOutstanding = "OUTSTANDING_TRANSACTION"
)

// CollectResponse defines the response Body of the S2S collect request
type CollectResponse struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	ProgressStatus string `json:"progressStatus"`
	UserInfo       struct {
		GivenName      string `json:"givenName"`
		Surname        string `json:"surname"`
		Name           string `json:"name"`
		PersonalNumber string `json:"personalNumber"`
		NotBefore      string `json:"notBefore"`
		NotAfter       string `json:"notAfter"`
		IPAddress      string `json:"ipAddress"`
	} `json:"userInfo"`
	Signature string `json:"signature"`
	OCSP      string `json:"ocspResponse"`
}

// InitResponse defines the response Body of the S2S init request
type InitResponse struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	OrderRef       string `json:"orderRef"`
	AutoStartToken string `json:"autoStartToken"`
}
