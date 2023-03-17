package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeycloakRealmSpec defines the desired state of KeycloakRealm.
// +k8s:openapi-gen=true
type KeycloakRealmSpec struct {
	// When set to true, this KeycloakRealm will be marked as unmanaged and not be managed by this operator.
	// It can then be used for targeting purposes.
	// +optional
	Unmanaged bool `json:"unmanaged,omitempty"`
	// Selector for looking up Keycloak Custom Resources.
	// +kubebuilder:validation:Required
	InstanceSelector *metav1.LabelSelector `json:"instanceSelector,omitempty"`
	// Keycloak Realm REST object.
	// +kubebuilder:validation:Required
	Realm *KeycloakAPIRealm `json:"realm"`
	// A list of overrides to the default Realm behavior.
	// +listType=atomic
	RealmOverrides []*RedirectorIdentityProviderOverride `json:"realmOverrides,omitempty"`
}

type KeycloakAPIRealm struct {
	// +kubebuilder:validation:Required
	// +optional
	ID string `json:"id,omitempty"`
	// Realm name.
	// +kubebuilder:validation:Required
	Realm string `json:"realm"`
	// Realm enabled flag.
	// +optional
	Enabled *bool `json:"enabled"`
	// Realm display name.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// +optional
	NotBefore int `json:"notBefore,omitempty"`
	// Realm default signature algorithm
	// +optional
	DefaultSignatureAlgorithm string `json:"defaultSignatureAlgorithm,omitempty"`
	// Realm HTML display name.
	// +optional
	DisplayNameHTML string `json:"displayNameHtml,omitempty"`
	// Realm Password Policy
	// +optional
	PasswordPolicy string `json:"passwordPolicy,omitempty"`
	// A set of Keycloak Users.
	// +optional
	Users []*KeycloakAPIUser `json:"users,omitempty"`
	// A set of Keycloak Clients.
	// +optional
	Clients []*KeycloakAPIClient `json:"clients,omitempty"`
	// A set of Identity Providers.
	// +optional
	IdentityProviders []*KeycloakIdentityProvider `json:"identityProviders,omitempty"`
	// A set of Identity Provider Mappers.
	// +optional
	IdentityProviderMappers []*KeycloakIdentityProviderMapper `json:"identityProviderMappers,omitempty"`
	// A set of Event Listeners.
	// +optional
	EventsListeners []string `json:"eventsListeners,omitempty"`
	// Enable events recording
	// TODO: change to values and use kubebuilder default annotation once supported
	// +optional
	EventsEnabled *bool `json:"eventsEnabled,omitempty"`
	// Enabled event types
	// +optional
	EnabledEventTypes []string `json:"enabledEventTypes,omitempty"`
	// Enable events recording
	// TODO: change to values and use kubebuilder default annotation once supported
	// +optional
	AdminEventsEnabled *bool `json:"adminEventsEnabled,omitempty"`
	// Enable admin events details
	// TODO: change to values and use kubebuilder default annotation once supported
	// +optional
	AdminEventsDetailsEnabled *bool `json:"adminEventsDetailsEnabled,omitempty"`

	// Keycloak version
	// +optional
	KeycloakVersion string `json:"keycloakVersion,omitempty"`

	// Attributes
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`

	// Client scopes
	// +optional
	ClientScopes []KeycloakClientScope `json:"clientScopes,omitempty"`

	// Default client scopes to add to all new clients
	// +optional
	DefaultDefaultClientScopes []string `json:"defaultDefaultClientScopes,omitempty"`

	// Default optional client scopes
	// +optional
	DefaultOptionalClientScopes []string `json:"defaultOptionalClientScopes,omitempty"`

	// Browser security headers
	// +optional
	BrowserSecurityHeaders map[string]string `json:"browserSecurityHeaders,omitempty"`

	// Authentication flows
	// +optional
	AuthenticationFlows []KeycloakAPIAuthenticationFlow `json:"authenticationFlows,omitempty"`

	// Authenticator config
	// +optional
	AuthenticatorConfig []KeycloakAPIAuthenticatorConfig `json:"authenticatorConfig,omitempty"`

	// Point keycloak to an external user provider to validate
	// credentials or pull in identity information.
	// +optional
	UserFederationProviders []KeycloakAPIUserFederationProvider `json:"userFederationProviders,omitempty"`

	// User federation mappers are extension points triggered by the
	// user federation at various points.
	// +optional
	UserFederationMappers []KeycloakAPIUserFederationMapper `json:"userFederationMappers,omitempty"`

	// User registration
	// +optional
	RegistrationAllowed *bool `json:"registrationAllowed,omitempty"`
	// Email as username
	// +optional
	RegistrationEmailAsUsername *bool `json:"registrationEmailAsUsername,omitempty"`
	// Edit username
	// +optional
	EditUsernameAllowed *bool `json:"editUsernameAllowed,omitempty"`
	// Forgot password
	// +optional
	ResetPasswordAllowed *bool `json:"resetPasswordAllowed,omitempty"`
	// Remember me
	// +optional
	RememberMe *bool `json:"rememberMe,omitempty"`
	// Verify email
	// +optional
	VerifyEmail *bool `json:"verifyEmail,omitempty"`
	// Login with email
	// +optional
	LoginWithEmailAllowed *bool `json:"loginWithEmailAllowed,omitempty"`
	// Duplicate emails
	// +optional
	DuplicateEmailsAllowed *bool `json:"duplicateEmailsAllowed,omitempty"`
	// Require SSL
	// +optional
	SslRequired string `json:"sslRequired,omitempty"`

	// Brute Force Detection
	// +optional
	BruteForceProtected *bool `json:"bruteForceProtected,omitempty"`
	// Permanent Lockout
	// +optional
	PermanentLockout *bool `json:"permanentLockout,omitempty"`
	// Max Login Failures
	// +optional
	FailureFactor *int32 `json:"failureFactor,omitempty"`
	// Wait Increment
	// +optional
	WaitIncrementSeconds *int32 `json:"waitIncrementSeconds,omitempty"`
	// Quick Login Check Milli Seconds
	// +optional
	QuickLoginCheckMilliSeconds *int64 `json:"quickLoginCheckMilliSeconds,omitempty"`
	// Minimum Quick Login Wait
	// +optional
	MinimumQuickLoginWaitSeconds *int32 `json:"minimumQuickLoginWaitSeconds,omitempty"`
	// Max Wait
	// +optional
	MaxFailureWaitSeconds *int32 `json:"maxFailureWaitSeconds,omitempty"`
	// Failure Reset Time
	// +optional
	MaxDeltaTimeSeconds *int32 `json:"maxDeltaTimeSeconds,omitempty"`

	// Email
	// +optional
	SMTPServer map[string]string `json:"smtpServer,omitempty"`

	// Login Theme
	// +optional
	LoginTheme string `json:"loginTheme,omitempty"`
	// Account Theme
	// +optional
	AccountTheme string `json:"accountTheme,omitempty"`
	// Admin Console Theme
	// +optional
	AdminTheme string `json:"adminTheme,omitempty"`
	// Email Theme
	// +optional
	EmailTheme string `json:"emailTheme,omitempty"`
	// Internationalization Enabled
	// +optional
	InternationalizationEnabled *bool `json:"internationalizationEnabled,omitempty"`
	// Supported Locales
	// +optional
	SupportedLocales []string `json:"supportedLocales,omitempty"`
	// Default Locale
	// +optional
	DefaultLocale string `json:"defaultLocale,omitempty"`

	// Roles
	// +optional
	Roles *RolesRepresentation `json:"roles,omitempty"`

	// Default role
	// +optional
	DefaultRole *RoleRepresentation `json:"defaultRole,omitempty"`

	// Scope Mappings
	// +optional
	ScopeMappings []ScopeMappingRepresentation `json:"scopeMappings,omitempty"`
	// Client Scope Mappings
	// +optional
	ClientScopeMappings map[string]ScopeMappingRepresentationArray `json:"clientScopeMappings,omitempty"`

	// Access Token Lifespan For Implicit Flow
	// +optional
	AccessTokenLifespanForImplicitFlow *int32 `json:"accessTokenLifespanForImplicitFlow,omitempty"`
	// Access Token Lifespan
	// +optional
	AccessTokenLifespan *int32 `json:"accessTokenLifespan,omitempty"`

	// Access code lifespan
	// +optional
	AccessCodeLifespan *int32 `json:"accessCodeLifespan,omitempty"`
	// Access code lifespan for user action
	// +optional
	AccessCodeLifeSpanUserAction *int32 `json:"accessCodeLifeSpanUserAction,omitempty"`
	// Access code lifespan for login
	// +optional
	AccessCodeLifespanLogin *int32 `json:"accessCodeLifespanLogin,omitempty"`

	// Access code lifespan
	// +optional
	ActionTokenGeneratedByAdminLifespan *int32 `json:"actionTokenGeneratedByAdminLifespan,omitempty"`
	// Access code lifespan for user action
	// +optional
	ActionTokenGeneratedByUserLifespan *int32 `json:"actionTokenGeneratedByUserLifespan,omitempty"`

	// OAuth2 device code lifespan
	// +optional
	Oauth2DeviceCodeLifespan *int32 `json:"oauth2DeviceCodeLifespan,omitempty"`
	//  OAuth2 device code lifespan
	// +optional
	Oauth2DevicePollingInterval *int32 `json:"oauth2DevicePollingInterval,omitempty"`

	// SSO session idle timeout
	// +optional
	SsoSessionIdleTimeout *int32 `json:"ssoSessionIdleTimeout,omitempty"`
	// SSO session max lifespan
	// +optional
	SsoSessionMaxLifespan *int32 `json:"ssoSessionMaxLifespan,omitempty"`
	// SSO session idle timeout (remember me option)
	// +optional
	SsoSessionIdleTimeoutRememberMe *int32 `json:"ssoSessionIdleTimeoutRememberMe,omitempty"`
	// SSO session max lifespan (remember me option)
	// +optional
	SsoSessionMaxLifespanRememberMe *int32 `json:"ssoSessionMaxLifespanRememberMe,omitempty"`

	// Offline session idle timeout
	// +optional
	OfflineSessionIdleTimeout *int32 `json:"offlineSessionIdleTimeout,omitempty"`
	// Offline session max lifespan enabled
	// +optional
	OfflineSessionMaxLifespanEnabled *bool `json:"offlineSessionMaxLifespanEnabled,omitempty"`
	// Offline session max lifespan
	// +optional
	OfflineSessionMaxLifespan *int32 `json:"offlineSessionMaxLifespan,omitempty"`

	// Client session idle timeout
	// +optional
	ClientSessionIdleTimeout *int32 `json:"clientSessionIdleTimeout,omitempty"`
	// Client session max lifespan
	// +optional
	ClientSessionMaxLifespan *int32 `json:"clientSessionMaxLifespan,omitempty"`
	// Client offline session idle timeout
	// +optional
	ClientOfflineSessionIdleTimeout *int32 `json:"clientOfflineSessionIdleTimeout,omitempty"`
	// Client offline session max lifespan
	// +optional
	ClientOfflineSessionMaxLifespan *int32 `json:"clientOfflineSessionMaxLifespan,omitempty"`

	// Required credentials
	// +optional
	RequiredCredentials []string `json:"requiredCredentials,omitempty"`

	// User Managed Access Allowed
	// +optional
	UserManagedAccessAllowed *bool `json:"userManagedAccessAllowed,omitempty"`

	// OTP Policy Algorithm
	// +optional
	OtpPolicyAlgorithm string `json:"otpPolicyAlgorithm,omitempty"`

	// OTP Policy Digits
	// +optional
	OtpPolicyDigits *int32 `json:"otpPolicyDigits,omitempty"`

	// OTP Policy Initial Counter
	// +optional
	OtpPolicyInitialCounter *int32 `json:"otpPolicyInitialCounter,omitempty"`

	// OTP Policy Look Ahead Window
	// +optional
	OtpPolicyLookAheadWindow *int32 `json:"otpPolicyLookAheadWindow,omitempty"`

	// OTP Policy Period
	// +optional
	OtpPolicyPeriod *int32 `json:"otpPolicyPeriod,omitempty"`

	// OTP Policy Type
	// +optional
	OtpPolicyType string `json:"otpPolicyType,omitempty"`

	// OTP Supported Applications
	// +optional
	OtpSupportedApplications []string `json:"otpSupportedApplications,omitempty"`

	// OTP policy code reusable?
	// +optional
	OtpPolicyCodeReusable *bool `json:"otpPolicyCodeReusable,omitempty"`

	// WebAuthn Policy Acceptable Aaguids
	// +optional
	WebAuthnPolicyAcceptableAaguids []string `json:"webAuthnPolicyAcceptableAaguids,omitempty"`

	// WebAuthn Policy Attestation Conveyance Preference
	// +optional
	WebAuthnPolicyAttestationConveyancePreference string `json:"webAuthnPolicyAttestationConveyancePreference,omitempty"`

	// WebAuthn Policy Authenticator Attachment
	// +optional
	WebAuthnPolicyAuthenticatorAttachment string `json:"webAuthnPolicyAuthenticatorAttachment,omitempty"`

	// WebAuthn Policy Avoid Same Authenticator Register
	// +optional
	WebAuthnPolicyAvoidSameAuthenticatorRegister *bool `json:"webAuthnPolicyAvoidSameAuthenticatorRegister,omitempty"`

	// WebAuthn Policy Create Timeout
	// +optional
	WebAuthnPolicyCreateTimeout *int32 `json:"webAuthnPolicyCreateTimeout,omitempty"`

	// WebAuthn Policy Passwordless Acceptable Aaguids
	// +optional
	WebAuthnPolicyPasswordlessAcceptableAaguids []string `json:"webAuthnPolicyPasswordlessAcceptableAaguids,omitempty"`

	// WebAuthn Policy Passwordless Attestation Conveyance Preference
	// +optional
	WebAuthnPolicyPasswordlessAttestationConveyancePreference string `json:"webAuthnPolicyPasswordlessAttestationConveyancePreference,omitempty"`

	// WebAuthn Policy Passwordless Authenticator Attachment
	// +optional
	WebAuthnPolicyPasswordlessAuthenticatorAttachment string `json:"webAuthnPolicyPasswordlessAuthenticatorAttachment,omitempty"`

	// WebAuthn Policy Passwordless Avoid Same Authenticator Register
	// +optional
	WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister *bool `json:"webAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister,omitempty"`

	// WebAuthn Policy Passwordless Create Timeout
	// +optional
	WebAuthnPolicyPasswordlessCreateTimeout *int32 `json:"webAuthnPolicyPasswordlessCreateTimeout,omitempty"`

	// WebAuthn Policy Passwordless Require Resident Key
	// +optional
	WebAuthnPolicyPasswordlessRequireResidentKey string `json:"webAuthnPolicyPasswordlessRequireResidentKey,omitempty"`

	// WebAuthn Policy Passwordless RP Entity Name
	// +optional
	WebAuthnPolicyPasswordlessRpEntityName string `json:"webAuthnPolicyPasswordlessRpEntityName,omitempty"`

	// WebAuthn Policy Passwordless RP ID
	// +optional
	WebAuthnPolicyPasswordlessRpID string `json:"webAuthnPolicyPasswordlessRpId,omitempty"`

	// WebAuthn Policy Passwordless Signature Algorithms
	// +optional
	WebAuthnPolicyPasswordlessSignatureAlgorithms []string `json:"webAuthnPolicyPasswordlessSignatureAlgorithms,omitempty"`

	// WebAuthn Policy Passwordless User Verification Requirement
	// +optional
	WebAuthnPolicyPasswordlessUserVerificationRequirement string `json:"webAuthnPolicyPasswordlessUserVerificationRequirement,omitempty"`

	// WebAuthn Policy Require Resident Key
	// +optional
	WebAuthnPolicyRequireResidentKey string `json:"webAuthnPolicyRequireResidentKey,omitempty"`

	// WebAuthn Policy RP Entity Name
	// +optional
	WebAuthnPolicyRpEntityName string `json:"webAuthnPolicyRpEntityName,omitempty"`

	// WebAuthn Policy RP ID
	// +optional
	WebAuthnPolicyRpID string `json:"webAuthnPolicyRpId,omitempty"`

	// WebAuthn Policy Signature Algorithms
	// +optional
	WebAuthnPolicySignatureAlgorithms []string `json:"webAuthnPolicySignatureAlgorithms,omitempty"`

	// WebAuthn Policy User Verification Requirement
	// +optional
	WebAuthnPolicyUserVerificationRequirement string `json:"webAuthnPolicyUserVerificationRequirement,omitempty"`

	// Browser authentication flow
	// +optional
	BrowserFlow string `json:"browserFlow,omitempty"`

	// Direct Grant authentication flow
	// +optional
	DirectGrantFlow string `json:"directGrantFlow,omitempty"`

	// Client authentication flow
	// +optional
	ClientAuthenticationFlow string `json:"clientAuthenticationFlow,omitempty"`

	// Reset Credentials authentication flow
	// +optional
	ResetCredentialsFlow string `json:"resetCredentialsFlow,omitempty"`

	// Registration flow
	// +optional
	RegistrationFlow string `json:"registrationFlow,omitempty"`

	// Docker Authentication flow
	// +optional
	DockerAuthenticationFlow string `json:"dockerAuthenticationFlow,omitempty"`

	// Client profiles
	// +optional
	ClientProfiles apiextensionsv1.JSON `json:"clientProfiles,omitempty"`
	// Client policies
	// +optional
	ClientPolicies apiextensionsv1.JSON `json:"clientPolicies,omitempty"`

	// Required actions
	// +optional
	RequiredActions []RequiredActionProviderRepresentation `json:"requiredActions,omitempty"`

	// Components
	// +optional
	Components map[string][]ComponentExportRepresentation `json:"components,omitempty"`
}
type ComponentExportRepresentation struct {
	Config        map[string][]string                        `json:"config,omitempty"`
	Id            string                                     `json:"id,omitempty"`
	Name          string                                     `json:"name,omitempty"`
	ProviderId    string                                     `json:"providerId,omitempty"`
	SubComponents map[string][]ComponentExportRepresentation `json:"subComponents,omitempty"`
}

type RequiredActionProviderRepresentation struct {
	DefaultAction bool              `json:"defaultAction,omitEmpty"`
	ProviderID    string            `json:"providerId,omitEmpty"`
	Name          string            `json:"name,omitEmpty"`
	Alias         string            `json:"alias,omitEmpty"`
	Priority      int32             `json:"priority,omitEmpty"`
	Config        map[string]string `json:"config,omitEmpty"`
	Enabled       bool              `json:"enabled,omitEmpty"`
}

type RoleRepresentationArray []RoleRepresentation

// https://www.keycloak.org/docs-api/11.0/rest-api/index.html#_rolesrepresentation
type RolesRepresentation struct {
	// Client Roles
	// +optional
	Client map[string]RoleRepresentationArray `json:"client,omitempty"`

	// Realm Roles
	// +optional
	Realm []RoleRepresentation `json:"realm,omitempty"`
}

// https://www.keycloak.org/docs-api/11.0/rest-api/index.html#_rolerepresentation
type RoleRepresentation struct {
	// Role Attributes
	// +optional
	Attributes map[string][]string `json:"attributes,omitempty"`

	// Client Role
	// +optional
	ClientRole *bool `json:"clientRole,omitempty"`

	// Composite
	// +optional
	Composite *bool `json:"composite,omitempty"`

	// Composites
	// +optional
	Composites *RoleRepresentationComposites `json:"composites,omitempty"`

	// Container Id
	// +optional
	ContainerID string `json:"containerId,omitempty"`

	// Description
	// +optional
	Description string `json:"description,omitempty"`

	// Id
	// +optional
	ID string `json:"id,omitempty"`

	// Name
	Name string `json:"name"`
}

type ScopeMappingRepresentationArray []ScopeMappingRepresentation

// https://www.keycloak.org/docs-api/11.0/rest-api/index.html#_scopemappingrepresentation
type ScopeMappingRepresentation struct {
	// Client
	// +optional
	Client string `json:"client,omitempty"`

	// Client Scope
	// +optional
	ClientScope string `json:"clientScope,omitempty"`

	// Roles
	// +optional
	Roles []string `json:"roles,omitempty"`

	// Self
	// +optional
	Self string `json:"self,omitempty"`
}

// https://www.keycloak.org/docs-api/11.0/rest-api/index.html#_rolerepresentation-composites
type RoleRepresentationComposites struct {
	// Map client => []role
	// +optional
	Client map[string][]string `json:"client,omitempty"`

	// Realm roles
	// +optional
	Realm []string `json:"realm,omitempty"`
}

// https://www.keycloak.org/docs-api/10.0/rest-api/index.html#_userfederationproviderrepresentation
type KeycloakAPIUserFederationProvider struct {
	// lastSync int32

	// +optional
	ChangedSyncPeriod *int32 `json:"changedSyncPeriod,omitempty"`

	// User federation provider config.
	// +optional
	Config map[string]string `json:"config,omitempty"`

	// The display name of this provider instance.
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// +optional
	FullSyncPeriod *int32 `json:"fullSyncPeriod,omitempty"`

	// The ID of this provider
	// +optional
	ID string `json:"id,omitempty"`

	// The priority of this provider when looking up users or adding a user.
	// +optional
	Priority *int32 `json:"priority,omitempty"`

	// The name of the user provider, such as "ldap", "kerberos" or a custom SPI.
	// +optional
	ProviderName string `json:"providerName,omitempty"`
}

// https://www.keycloak.org/docs/11.0/server_admin/#_ldap_mappers
// https://www.keycloak.org/docs-api/11.0/rest-api/index.html#_userfederationmapperrepresentation
type KeycloakAPIUserFederationMapper struct {
	// User federation mapper config.
	// +optional
	Config map[string]string `json:"config,omitempty"`

	// +optional
	Name string `json:"name,omitempty"`

	// +optional
	ID string `json:"id,omitempty"`

	// +optional
	FederationMapperType string `json:"federationMapperType,omitempty"`

	// The displayName for the user federation provider this mapper applies to.
	FederationProviderDisplayName string `json:"federationProviderDisplayName,omitempty"`
}

type KeycloakAPIAuthenticationFlow struct {
	// Alias
	Alias string `json:"alias"`

	// Authentication executions
	AuthenticationExecutions []KeycloakAPIAuthenticationExecution `json:"authenticationExecutions"`

	// Built in
	// +optional
	BuiltIn bool `json:"builtIn,omitempty"`

	// Description
	// +optional
	Description string `json:"description,omitempty"`

	// ID
	// +optional
	ID string `json:"id,omitempty"`

	// Provider ID
	// +optional
	ProviderID string `json:"providerId,omitempty"`

	// Top level
	// +optional
	TopLevel bool `json:"topLevel,omitempty"`
}

type KeycloakAPIAuthenticationExecution struct {
	// Authenticator
	Authenticator string `json:"authenticator,omitempty"`

	// Authenticator Config
	// +optional
	AuthenticatorConfig string `json:"authenticatorConfig,omitempty"`

	// Authenticator flow
	// +optional
	AuthenticatorFlow bool `json:"authenticatorFlow,omitempty"`

	// Authenticator flow (deprecated property; has spelling error)
	// (Please use `authenticatorFlow` instead)
	// +optional
	AutheticatorFlow bool `json:"autheticatorFlow,omitempty"`

	// Flow Alias
	// +optional
	FlowAlias string `json:"flowAlias,omitempty"`

	// Priority
	// +optional
	Priority int32 `json:"priority,omitempty"`

	// Requirement [REQUIRED, OPTIONAL, ALTERNATIVE, DISABLED]
	Requirement string `json:"requirement,omitempty"`

	// User setup allowed
	// +optional
	UserSetupAllowed bool `json:"userSetupAllowed,omitempty"`
}

type KeycloakAPIAuthenticatorConfig struct {
	// Alias
	Alias string `json:"alias"`

	// Config
	// +optional
	Config map[string]string `json:"config,omitempty"`

	// ID
	// +optional
	ID string `json:"id,omitempty"`
}

type RedirectorIdentityProviderOverride struct {
	// Identity Provider to be overridden.
	IdentityProvider string `json:"identityProvider"`
	// Flow to be overridden.
	// +optional
	ForFlow string `json:"forFlow,omitempty"`
}

type KeycloakClientScope struct {
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	ID string `json:"id,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// Protocol Mappers.
	// +optional
	ProtocolMappers []KeycloakProtocolMapper `json:"protocolMappers,omitempty"`
}

type KeycloakIdentityProvider struct {
	// Identity Provider Alias.
	// +optional
	Alias string `json:"alias,omitempty"`
	// Identity Provider Display Name.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// Identity Provider Internal ID.
	// +optional
	InternalID string `json:"internalId,omitempty"`
	// Identity Provider ID.
	// +optional
	ProviderID string `json:"providerId,omitempty"`
	// Identity Provider enabled flag.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// Identity Provider Trust Email.
	// +optional
	TrustEmail bool `json:"trustEmail,omitempty"`
	// Identity Provider Store to Token.
	// +optional
	StoreToken bool `json:"storeToken,omitempty"`
	// Adds Read Token role when creating this Identity Provider.
	// +optional
	AddReadTokenRoleOnCreate bool `json:"addReadTokenRoleOnCreate,omitempty"`
	// Identity Provider First Broker Login Flow Alias.
	// +optional
	FirstBrokerLoginFlowAlias string `json:"firstBrokerLoginFlowAlias,omitempty"`
	// Identity Provider Post Broker Login Flow Alias.
	// +optional
	PostBrokerLoginFlowAlias string `json:"postBrokerLoginFlowAlias,omitempty"`
	// Identity Provider Link Only setting.
	// +optional
	LinkOnly bool `json:"linkOnly,omitempty"`
	// Identity Provider config.
	// +optional
	Config map[string]string `json:"config,omitempty"`
}

type KeycloakIdentityProviderMapper struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Identity Provider Alias.
	// +optional
	IdentityProviderAlias string `json:"identityProviderAlias,omitempty"`
	// Identity Provider Mapper.
	// +optional
	IdentityProviderMapper string `json:"identityProviderMapper,omitempty"`
	// Identity Provider Mapper config.
	// +optional
	Config map[string]string `json:"config,omitempty"`
}

type KeycloakUserRole struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Composite   bool   `json:"composite,omitempty"`
	ClientRole  bool   `json:"clientRole,omitempty"`
	ContainerID string `json:"containerId,omitempty"`
}

type AuthenticatorConfig struct {
	// Authenticator Config Alias.
	// +optional
	Alias string `json:"alias,omitempty"`
	// Authenticator config.
	// +optional
	Config map[string]string `json:"config,omitempty"`
	// Authenticator ID.
	// +optional
	ID string `json:"id,omitempty"`
}

type KeycloakAPIPasswordReset struct {
	// Password Reset Type.
	// +optional
	Type string `json:"type"`
	// Password Reset Value.
	// +optional
	Value string `json:"value"`
	// True if this Password Reset object is temporary.
	// +optional
	Temporary bool `json:"temporary"`
}

type AuthenticationExecutionInfo struct {
	// Authentication Execution Info Alias.
	// +optional
	Alias string `json:"alias,omitempty"`
	// Authentication Execution Info Config.
	// +optional
	AuthenticationConfig string `json:"authenticationConfig,omitempty"`
	// True if Authentication Flow is enabled.
	// +optional
	AuthenticationFlow bool `json:"authenticationFlow,omitempty"`
	// True if Authentication Execution Info is configurable.
	// +optional
	Configurable bool `json:"configurable,omitempty"`
	// Authentication Execution Info Display Name.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// Authentication Execution Info Flow ID.
	// +optional
	FlowID string `json:"flowId,omitempty"`
	// Authentication Execution Info ID.
	// +optional
	ID string `json:"id,omitempty"`
	// Authentication Execution Info Index.
	// +optional
	Index int32 `json:"index,omitempty"`
	// Authentication Execution Info Level.
	// +optional
	Level int32 `json:"level,omitempty"`
	// Authentication Execution Info Provider ID.
	// +optional
	ProviderID string `json:"providerId,omitempty"`
	// Authentication Execution Info Requirement.
	// +optional
	Requirement string `json:"requirement,omitempty"`
	// Authentication Execution Info Requirement Choices.
	// +optional
	RequirementChoices []string `json:"requirementChoices,omitempty"`
}

type TokenResponse struct {
	// Token Response Access Token.
	// +optional
	AccessToken string `json:"access_token"`
	// Token Response Expired In setting.
	// +optional
	ExpiresIn int `json:"expires_in"`
	// Token Response Refresh Expires In setting.
	// +optional
	RefreshExpiresIn int `json:"refresh_expires_in"`
	// Token Response Refresh Token.
	// +optional
	RefreshToken string `json:"refresh_token"`
	// Token Response Token Type.
	// +optional
	TokenType string `json:"token_type"`
	// Token Response Not Before Policy setting.
	// +optional
	NotBeforePolicy int `json:"not-before-policy"`
	// Token Response Session State.
	// +optional
	SessionState string `json:"session_state"`
	// Token Response Error.
	// +optional
	Error string `json:"error"`
	// Token Response Error Description.
	// +optional
	ErrorDescription string `json:"error_description"`
}

// KeycloakRealmStatus defines the observed state of KeycloakRealm
// +k8s:openapi-gen=true
type KeycloakRealmStatus struct {
	// Current phase of the operator.
	Phase StatusPhase `json:"phase"`
	// Human-readable message indicating details about current operator phase or error.
	Message string `json:"message"`
	// True if all resources are in a ready state and all work is done.
	Ready bool `json:"ready"`
	// A map of all the secondary resources types and names created for this CR. e.g "Deployment": [ "DeploymentName1", "DeploymentName2" ]
	SecondaryResources map[string][]string `json:"secondaryResources,omitempty"`
	// TODO
	LoginURL string `json:"loginURL"`
}

// KeycloakRealm is the Schema for the keycloakrealms API
// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KeycloakRealm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeycloakRealmSpec   `json:"spec,omitempty"`
	Status KeycloakRealmStatus `json:"status,omitempty"`
}

// KeycloakRealmList contains a list of KeycloakRealm
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KeycloakRealmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeycloakRealm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeycloakRealm{}, &KeycloakRealmList{})
}

func (i *KeycloakRealm) UpdateStatusSecondaryResources(kind string, resourceName string) {
	i.Status.SecondaryResources = UpdateStatusSecondaryResources(i.Status.SecondaryResources, kind, resourceName)
}
