# go-zignsec
Implements the ZignSec api

This is an interface to both http://docs.zignsec.com/api/web-based/ and http://docs.zignsec.com/api/s2s/. It has some simple functions and data structures.

EXAMPLE

s2s, signing:

	s2s := zignsec.NewS2SClient(zignsec.S2SURLTest, "Your API key")
	ic := zignsec.ZSInitConfig{Personalnumber: "1234567890",
		UserVisibleData:    "VGVzdGFy"}
	ir, err := s2s.Init("Sign", ic)
	cr, err := s2s.Collect(ir.OrderRef)
	if cr.ProgressStatus == zignsec.CollectProgressStatusComplete {
		return cr.Signature
	}

	PACKAGE DOCUMENTATION

	package zignsec
	    import "github.com/qoorp/go-zignsec"

	CONSTANTS

	const (
	    // APIHostBase is the production endpoint
	    APIHostBase = "https://api.zignsec.com/v2/eid"
	    // APIHostBaseTest is the test endpoint
	    APIHostBaseTest = "https://test.zignsec.com/v2/eid"
	)

	const (
	    // CollectProgressStatusComplete completed and validated
	    CollectProgressStatusComplete = "COMPLETE"
	    // CollectProgressStatusUserSign signed and validated
	    CollectProgressStatusUserSign = "USER_SIGN"
	    // CollectProgressStatusOutstanding waiting for user to complete login
	    CollectProgressStatusOutstanding = "OUTSTANDING_TRANSACTION"
	)

	const (
	    // S2SURL is the production endpoint
	    S2SURL = "https://api.zignsec.com/v2/BankIDSE"
	    // S2SURLTest is the test endpoint
	    S2SURLTest = "https://test.zignsec.com/v2/BankIDSE"
	)

	TYPES

	type Client struct {
	    APIHostBase string
	    APIKey      string
	}
	    Client is a zignsec client

	func New(APIHostBase string, APIKey string) *Client
	    New create a new client

	func (c *Client) Initiate(method string, config ZSInitConfig) (*ZSInitRespBody, error)
	    Initiate a login or sign request

	func (c *Client) Verify(uuid string) (*ZSVerifyRespBody, error)
	    Verify a login or signature

	type ZSInitConfig struct {
	    Personalnumber     string `json:"personalnumber,omitempty"`
	    UserVisibleData    string `json:"userVisibleData,omitempty"`
	    UserNonVisibleData string `json:"userNonVisibleData,omitempty"`
	    Relaystate         string `json:"relaystate,omitempty"`
	    Configid           string `json:"configid,omitempty"`
	    Target             string `json:"target,omitempty"`
	}
	    ZSInitConfig defines the configuration which are sent along with both
	    web-based Initiate() and s2s Init().
	    UserNonVisibleData is signed, if present. UserVisibleData
	    is then only for user guidance. Both must be base64 encoded. If decoded
	    UserVisibleData contains certain bytes (integer value > 127?), Init()
	    will fail with INVALID_PARAMETERS.

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
	    ZSVerifyRespBody defines the response Body of Verify()

	type S2SClient struct {
	    // contains filtered or unexported fields
	}
	    S2SClient is a Zignsec server to server client

	func NewS2SClient(baseURL string, key string) *S2SClient
	    NewS2SClient create a new client. Key is the Zignsec Authorization.

	func (c *S2SClient) Collect(order string) (*CollectResponse, error)
	    Collect the state of an authenticate or sign request. Order is from the
	    Init() response. The guidelines recommend polling for results every 2
	    seconds. When the response Status is COMPLETE, then further use of
	    Collect() will receive an error.

	func (c *S2SClient) Init(method string, config ZSInitConfig) (*InitResponse, error)
	    Init a authenticate or sign request. Use Collect() to get the answer
	    afterwards.

	type ZSInitRespBody struct {
	    ID     string `json:"id"`
	    Errors []struct {
	        Code        string `json:"code"`
	        Description string `json:"description"`
	    } `json:"errors"`
	    RedirectURL string `json:"redirect_url"`
	}
	    ZSInitRespBody defines the response Body of Initiate()

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
	    CollectResponse defines the response Body of S2S Collect()

	type InitResponse struct {
	    ID     string `json:"id"`
	    Errors []struct {
	        Code        string `json:"code"`
	        Description string `json:"description"`
	    } `json:"errors"`
	    OrderRef       string `json:"orderRef"`
	    AutoStartToken string `json:"autoStartToken"`
	}
	    InitResponse defines the response Body of S2S Init()
