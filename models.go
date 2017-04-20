package zignsec

// ZSInitConfig defines the configuration which are sent along with the init request
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

// CollectResponse defines the response Body of the S2S collect request
type CollectResponse struct {
	ID     string `json:"id"`
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
	Status string `json:"progressStatus"`
	User   struct {
		FirstName      string `json:"givenName"`
		LastName       string `json:"surname"`
		Name           string `json:"name"`
		PersonalNumber string `json:"personalNumber"`
		NotBefore      string `json:"notBefore"`
		NotAfter       string `json:"notAfter"`
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
	Order     string `json:"orderRef"`
	AutoStart string `json:"autoStartToken"`
}
