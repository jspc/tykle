package tyk

type ApisCreate struct {
	Cors struct {
		AllowCredentials   bool          `json:"allow_credentials"`
		AllowedHeaders     []interface{} `json:"allowed_headers"`
		AllowedMethods     []interface{} `json:"allowed_methods"`
		AllowedOrigins     []interface{} `json:"allowed_origins"`
		Debug              bool          `json:"debug"`
		Enable             bool          `json:"enable"`
		ExposedHeaders     []interface{} `json:"exposed_headers"`
		MaxAge             int64         `json:"max_age"`
		OptionsPassthrough bool          `json:"options_passthrough"`
	} `json:"CORS"`
	Active     bool          `json:"active"`
	AllowedIps []interface{} `json:"allowed_ips"`
	APIID      string        `json:"api_id"`
	Auth       struct {
		AuthHeaderName string `json:"auth_header_name"`
		UseCookie      bool   `json:"use_cookie"`
		UseParam       bool   `json:"use_param"`
	} `json:"auth"`
	AuthProvider struct {
		Meta          struct{} `json:"meta"`
		Name          string   `json:"name"`
		StorageEngine string   `json:"storage_engine"`
	} `json:"auth_provider"`
	CacheOptions struct {
		CacheAllSafeRequests       bool  `json:"cache_all_safe_requests"`
		CacheTimeout               int64 `json:"cache_timeout"`
		EnableCache                bool  `json:"enable_cache"`
		EnableUpstreamCacheControl bool  `json:"enable_upstream_cache_control"`
	} `json:"cache_options"`
	CustomMiddleware struct {
		Post     []interface{} `json:"post"`
		Pre      []interface{} `json:"pre"`
		Response []interface{} `json:"response"`
	} `json:"custom_middleware"`
	Definition struct {
		Key      string `json:"key"`
		Location string `json:"location"`
	} `json:"definition"`
	Domain                    string `json:"domain"`
	DontSetQuotaOnCreate      bool   `json:"dont_set_quota_on_create"`
	EnableBatchRequestSupport bool   `json:"enable_batch_request_support"`
	EnableIPWhitelisting      bool   `json:"enable_ip_whitelisting"`
	EnableJwt                 bool   `json:"enable_jwt"`
	EnableSignatureChecking   bool   `json:"enable_signature_checking"`
	EventHandlers             struct {
		Events struct{} `json:"events"`
	} `json:"event_handlers"`
	ExpireAnalyticsAfter int64  `json:"expire_analytics_after"`
	HmacAllowedClockSkew int64  `json:"hmac_allowed_clock_skew"`
	ID                   string `json:"id"`
	JwtSigningMethod     string `json:"jwt_signing_method"`
	Name                 string `json:"name"`
	Notifications        struct {
		OauthOnKeychangeURL string `json:"oauth_on_keychange_url"`
		SharedSecret        string `json:"shared_secret"`
	} `json:"notifications"`
	OauthMeta struct {
		AllowedAccessTypes    []interface{} `json:"allowed_access_types"`
		AllowedAuthorizeTypes []interface{} `json:"allowed_authorize_types"`
		AuthLoginRedirect     string        `json:"auth_login_redirect"`
	} `json:"oauth_meta"`
	OrgID string `json:"org_id"`
	Proxy struct {
		CheckHostAgainstUptimeTests bool   `json:"check_host_against_uptime_tests"`
		EnableLoadBalancing         bool   `json:"enable_load_balancing"`
		ListenPath                  string `json:"listen_path"`
		ServiceDiscovery            struct {
			CacheTimeout        int64  `json:"cache_timeout"`
			DataPath            string `json:"data_path"`
			EndpointReturnsList bool   `json:"endpoint_returns_list"`
			ParentDataPath      string `json:"parent_data_path"`
			PortDataPath        string `json:"port_data_path"`
			QueryEndpoint       string `json:"query_endpoint"`
			UseDiscoveryService bool   `json:"use_discovery_service"`
			UseNestedQuery      bool   `json:"use_nested_query"`
			UseTargetList       bool   `json:"use_target_list"`
		} `json:"service_discovery"`
		StripListenPath bool          `json:"strip_listen_path"`
		TargetList      []interface{} `json:"target_list"`
		TargetURL       string        `json:"target_url"`
	} `json:"proxy"`
	ResponseProcessors []interface{} `json:"response_processors"`
	SessionLifetime    int64         `json:"session_lifetime"`
	SessionProvider    struct {
		Meta          interface{} `json:"meta"`
		Name          string      `json:"name"`
		StorageEngine string      `json:"storage_engine"`
	} `json:"session_provider"`
	Slug        string        `json:"slug"`
	Tags        []interface{} `json:"tags"`
	UptimeTests struct {
		CheckList []interface{} `json:"check_list"`
		Config    struct {
			ExpireUtimeAfter int64 `json:"expire_utime_after"`
			RecheckWait      int64 `json:"recheck_wait"`
			ServiceDiscovery struct {
				CacheTimeout        int64  `json:"cache_timeout"`
				DataPath            string `json:"data_path"`
				EndpointReturnsList bool   `json:"endpoint_returns_list"`
				ParentDataPath      string `json:"parent_data_path"`
				PortDataPath        string `json:"port_data_path"`
				QueryEndpoint       string `json:"query_endpoint"`
				UseDiscoveryService bool   `json:"use_discovery_service"`
				UseNestedQuery      bool   `json:"use_nested_query"`
				UseTargetList       bool   `json:"use_target_list"`
			} `json:"service_discovery"`
		} `json:"config"`
	} `json:"uptime_tests"`
	UseBasicAuth bool `json:"use_basic_auth"`
	UseKeyless   bool `json:"use_keyless"`
	UseOauth2    bool `json:"use_oauth2"`
	VersionData  struct {
		NotVersioned bool `json:"not_versioned"`
		Versions     struct {
			Default struct {
				Expires       string `json:"expires"`
				ExtendedPaths struct {
					BlackList                []interface{} `json:"black_list"`
					Cache                    []interface{} `json:"cache"`
					CircuitBreakers          []interface{} `json:"circuit_breakers"`
					HardTimeouts             []interface{} `json:"hard_timeouts"`
					Ignored                  []interface{} `json:"ignored"`
					SizeLimits               []interface{} `json:"size_limits"`
					Transform                []interface{} `json:"transform"`
					TransformHeaders         []interface{} `json:"transform_headers"`
					TransformResponse        []interface{} `json:"transform_response"`
					TransformResponseHeaders []interface{} `json:"transform_response_headers"`
					URLRewrites              []interface{} `json:"url_rewrites"`
					Virtual                  []interface{} `json:"virtual"`
					WhiteList                []interface{} `json:"white_list"`
				} `json:"extended_paths"`
				GlobalHeaders       struct{}      `json:"global_headers"`
				GlobalHeadersRemove []interface{} `json:"global_headers_remove"`
				GlobalSizeLimit     int64         `json:"global_size_limit"`
				Name                string        `json:"name"`
				Paths               struct {
					BlackList []interface{} `json:"black_list"`
					Ignored   []interface{} `json:"ignored"`
					WhiteList []interface{} `json:"white_list"`
				} `json:"paths"`
				UseExtendedPaths bool `json:"use_extended_paths"`
			} `json:"Default"`
		} `json:"versions"`
	} `json:"version_data"`
}

var (
	apisEndpoint = "/apis"
	apisMethod   = "post"
)

// Do wraps and returns the post'ing of ApisCreate to /apis
func (o ApisCreate) Do() (i interface{}, err error) {
	return Client.Do(apisMethod, apisEndpoint, o)
}
