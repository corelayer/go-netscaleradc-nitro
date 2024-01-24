/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package config

type LbVserver struct {
	Name                                string    `json:"name,omitempty" nitro:"permission=readwrite"`
	ServiceType                         string    `json:"servicetype,omitempty" nitro:"permission=readwrite"`
	Ipv46                               string    `json:"ipv46,omitempty" nitro:"permission=readwrite"`
	IpPattern                           string    `json:"ippattern,omitempty" nitro:"permission=readwrite"`
	IpMask                              string    `json:"ipmask,omitempty" nitro:"permission=readwrite"`
	Port                                int       `json:"port" nitro:"permission=readwrite"`
	IpSet                               string    `json:"ipset,omitempty" nitro:"permission=readwrite"`
	Range                               string    `json:"range,omitempty" nitro:"permission=readwrite"`
	PersistenceType                     string    `json:"persistencetype,omitempty" nitro:"permission=readwrite"`
	Timeout                             float64   `json:"timeout,omitempty" nitro:"permission=readwrite"`
	PersistenceBackup                   string    `json:"persistencebackup,omitempty" nitro:"permission=readwrite"`
	BackupPersistenceTimeout            float64   `json:"backuppersistencetimeout,omitempty" nitro:"permission=readwrite"`
	LbMethod                            string    `json:"lbmethod,omitempty" nitro:"permission=readwrite"`
	HashLength                          float64   `json:"hashlength,omitempty" nitro:"permission=readwrite"`
	Netmask                             string    `json:"netmask,omitempty" nitro:"permission=readwrite"`
	Ipv6NetmaskLength                   float64   `json:"v6netmasklen,omitempty" nitro:"permission=readwrite"`
	BackupLbMethod                      string    `json:"backuplbmethod,omitempty" nitro:"permission=readwrite"`
	CookieName                          string    `json:"cookiename,omitempty" nitro:"permission=readwrite"`
	Rule                                string    `json:"rule,omitempty" nitro:"permission=readwrite"`
	ListenPolicy                        string    `json:"listenpolicy,omitempty" nitro:"permission=readwrite"`
	ListenPriority                      float64   `json:"listenpriority,omitempty" nitro:"permission=readwrite"`
	PersistenceRule                     string    `json:"resrule,omitempty" nitro:"permission=readwrite"`
	PersistenceMask                     string    `json:"persistmask,omitempty" nitro:"permission=readwrite"`
	Ipv6PersistenceMaskLength           string    `json:"v6persistmasklen,omitempty" nitro:"permission=readwrite"`
	RtspNat                             string    `json:"rtspnat,omitempty" nitro:"permission=readwrite"`
	RedirectionMode                     string    `json:"m,omitempty" nitro:"permission=readwrite"`
	TosId                               float64   `json:"tosid,omitempty" nitro:"permission=readwrite"`
	DataLength                          string    `json:"datalength,omitempty" nitro:"permission=readwrite"`
	DataOffset                          string    `json:"dataoffset,omitempty" nitro:"permission=readwrite"`
	Sessionless                         string    `json:"sessionless,omitempty" nitro:"permission=readwrite"`
	TrofsPersistence                    string    `json:"trofspersistence,omitempty" nitro:"permission=readwrite"`
	State                               string    `json:"state,omitempty" nitro:"permission=readwrite"`
	ConnectionFailover                  string    `json:"connfailover,omitempty" nitro:"permission=readwrite"`
	RedirectUrl                         string    `json:"redirurl,omitempty" nitro:"permission=readwrite"`
	Cacheable                           string    `json:"cacheable,omitempty" nitro:"permission=readwrite"`
	ClientTimeout                       string    `json:"clttimeout,omitempty" nitro:"permission=readwrite"`
	SpilloverMethod                     string    `json:"somethod,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistence                string    `json:"sopersistence,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistenceTimeout         string    `json:"sopersistencetimeout,omitempty" nitro:"permission=readwrite"`
	HealthThreshold                     string    `json:"healththreshold,omitempty" nitro:"permission=readwrite"`
	SpilloverThreshold                  float64   `json:"sothreshold,omitempty" nitro:"permission=readwrite"`
	SpilloverBackupAction               string    `json:"sobackupaction,omitempty" nitro:"permission=readwrite"`
	RedirectPortRewrite                 string    `json:"redirectportrewrite,omitempty" nitro:"permission=readwrite"`
	DownstateFlush                      string    `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	BackupVserver                       string    `json:"backupvserver,omitempty" nitro:"permission=readwrite"`
	DisablePrimaryOnDown                string    `json:"disableprimaryondown,omitempty" nitro:"permission=readwrite"`
	InsertVserverIpPort                 string    `json:"insertvserveripport,omitempty" nitro:"permission=readwrite"`
	VipHeader                           string    `json:"vipheader,omitempty" nitro:"permission=readwrite"`
	AuthenticationHost                  string    `json:"authenticationhost,omitempty" nitro:"permission=readwrite"`
	Authentication                      string    `json:"authentication,omitempty" nitro:"permission=readwrite"`
	Authentication401                   string    `json:"authn401,omitempty" nitro:"permission=readwrite"`
	AuthenticationVserverName           string    `json:"authnvsname,omitempty" nitro:"permission=readwrite"`
	Push                                string    `json:"push,omitempty" nitro:"permission=readwrite"`
	PushVserver                         string    `json:"pushvserver,omitempty" nitro:"permission=readwrite"`
	PushLabel                           string    `json:"pushlabel,omitempty" nitro:"permission=readwrite"`
	PushMultipleClientConnections       string    `json:"pushmulticlients,omitempty" nitro:"permission=readwrite"`
	TcpProfileName                      string    `json:"tcpprofilename,omitempty" nitro:"permission=readwrite"`
	HttpProfileName                     string    `json:"httpprofilename,omitempty" nitro:"permission=readwrite"`
	DatabaseProfileName                 string    `json:"dbprofilename,omitempty" nitro:"permission=readwrite"`
	Comment                             string    `json:"comment,omitempty" nitro:"permission=readwrite"`
	L2ConnectionParamenters             string    `json:"l2conn,omitempty" nitro:"permission=readwrite"`
	OracleServerVersion                 string    `json:"oracleserverversion,omitempty" nitro:"permission=readwrite"`
	MssqlServerVersion                  string    `json:"mssqlserverversion,omitempty" nitro:"permission=readwrite"`
	MySqlProtocalVersion                string    `json:"mysqlprotocolversion,omitempty" nitro:"permission=readwrite"`
	MysqlServerVersion                  string    `json:"mysqlserverversion,omitempty" nitro:"permission=readwrite"`
	MysqlCharacterSet                   string    `json:"mysqlcharacterset,omitempty" nitro:"permission=readwrite"`
	AppFlowLog                          string    `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	NetProfile                          string    `json:"netprofile,omitempty" nitro:"permission=readwrite"`
	IcmpVserverResponse                 string    `json:"icmpvserverresponse,omitempty" nitro:"permission=readwrite"`
	RouthHealthInjectionState           string    `json:"rhistate,omitempty" nitro:"permission=readwrite"`
	NewServiceRequest                   float64   `json:"newservicerequest,omitempty" nitro:"permission=readwrite"`
	NewServiceRequestUnit               string    `json:"newservicerequestunit,omitempty" nitro:"permission=readwrite"`
	NewServiceRequestIncrementInterval  float64   `json:"newservicerequestincrementinterval,omitempty" nitro:"permission=readwrite"`
	MinimumAutoscaleMembers             string    `json:"minautoscalemembers,omitempty" nitro:"permission=readwrite"`
	MaximumAutoscaleMembers             string    `json:"maxautoscalemembers,omitempty" nitro:"permission=readwrite"`
	PersistAvpNumber                    []float64 `json:"persistavpno,omitempty" nitro:"permission=readwrite"`
	SkipPersistency                     string    `json:"skippersistency,omitempty" nitro:"permission=readwrite"`
	TrafficDomain                       string    `json:"td,omitempty" nitro:"permission=readwrite"`
	AuthenticationProfile               string    `json:"authnprofile,omitempty" nitro:"permission=readwrite"`
	MacModeRetainVlan                   string    `json:"macmoderetainvlan,omitempty" nitro:"permission=readwrite"`
	DatabaseSpecificLb                  string    `json:"dbslb,omitempty" nitro:"permission=readwrite"`
	Dns64                               string    `json:"dns64,omitempty" nitro:"permission=readwrite"`
	ByPassAaaaQueries                   string    `json:"bypassaaaa,omitempty" nitro:"permission=readwrite"`
	RecursionAvailable                  string    `json:"recursionavailable,omitempty" nitro:"permission=readwrite"`
	ProcessLocal                        string    `json:"processlocal,omitempty" nitro:"permission=readwrite"`
	DnsProfileName                      string    `json:"dnsprofilename,omitempty" nitro:"permission=readwrite"`
	LbProfileName                       string    `json:"lbprofilename,omitempty" nitro:"permission=readwrite"`
	RedirectFromPort                    int       `json:"redirectfromport,omitempty" nitro:"permission=readwrite"`
	HttpRedirectUrl                     string    `json:"httpsredirecturl,omitempty" nitro:"permission=readwrite"`
	RetainConnectionsOnCluster          string    `json:"retainconnectionsoncluster,omitempty" nitro:"permission=readwrite"`
	AdfsProxyProfile                    string    `json:"adfsproxyprofile,omitempty" nitro:"permission=readwrite"`
	TcpProbePort                        int       `json:"tcpprobeport,omitempty" nitro:"permission=readwrite"`
	QuicProfileName                     string    `json:"quicprofilename,omitempty" nitro:"permission=readwrite"`
	QuicBridgeProfileName               string    `json:"quicbridgeprofilename,omitempty" nitro:"permission=readwrite"`
	ProbeProtocol                       string    `json:"probeprotocol,omitempty" nitro:"permission=readwrite"`
	ProbeSuccessResponseCode            string    `json:"probesuccessresponsecode,omitempty" nitro:"permission=readwrite"`
	ProbePort                           int       `json:"probeport,omitempty" nitro:"permission=readwrite"`
	Weight                              float64   `json:"weight,omitempty" nitro:"permission=readwrite"`
	ServiceName                         string    `json:"servicename,omitempty" nitro:"permission=readwrite"`
	RedirectUrlFlags                    bool      `json:"redirurlflags,omitempty" nitro:"permission=readwrite"`
	NewName                             string    `json:"newname,omitempty" nitro:"permission=readwrite"`
	IpMapping                           string    `json:"ipmapping,omitempty" nitro:"permission=readonly"`
	NodegroupName                       string    `json:"ngname,omitempty" nitro:"permission=readonly"`
	Type                                string    `json:"type,omitempty" nitro:"permission=readonly"`
	CurrentState                        string    `json:"curstate,omitempty" nitro:"permission=readonly"`
	EffectiveState                      string    `json:"effectivestate,omitempty" nitro:"permission=readonly"`
	Status                              int       `json:"status,omitempty" nitro:"permission=readonly"`
	LbRrReason                          int       `json:"lbrrreason,omitempty" nitro:"permission=readonly"`
	Redirect                            string    `json:"redirect,omitempty" nitro:"permission=readonly"`
	Precedence                          string    `json:"precedense,omitempty" nitro:"permission=readonly"`
	Homepage                            string    `json:"homepage,omitempty" nitro:"permission=readonly"`
	DnsVserverName                      string    `json:"dnsvservername,omitempty" nitro:"permission=readonly"`
	Domain                              string    `json:"domain,omitempty" nitro:"permission=readonly"`
	CacheVserver                        string    `json:"cachevserver,omitempty" nitro:"permission=readonly"`
	Health                              string    `json:"health,omitempty" nitro:"permission=readonly"`
	RuleType                            string    `json:"ruletype,omitempty" nitro:"permission=readonly"`
	GroupName                           string    `json:"groupname,omitempty" nitro:"permission=readonly"`
	CookieDomain                        string    `json:"cookiedomain,omitempty" nitro:"permission=readonly"`
	Map                                 string    `json:"map,omitempty" nitro:"permission=readonly"`
	GreaterThan2GBTransactions          string    `json:"gt2gb,omitempty" nitro:"permission=readonly"`
	ConsolidatedLeastConnectionStats    string    `json:"consolidatedlconn,omitempty" nitro:"permission=readonly"`
	ConsolidatedLeastConnectionGlobal   string    `json:"consolidatedlconngbl,omitempty" nitro:"permission=readonly"`
	ThresholdValue                      int       `json:"thresholdvalue,omitempty" nitro:"permission=readonly"`
	Bindpoint                           string    `json:"bindpoint,omitempty" nitro:"permission=readonly"`
	Version                             int       `json:"version,omitempty" nitro:"permission=readonly"`
	TotalServices                       float64   `json:"totalservice,omitempty" nitro:"permission=readonly"`
	ActiveServices                      string    `json:"activeservices,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSeconds              string    `json:"statechangetimeseconds,omitempty" nitro:"permission=readonly"`
	StateChangeTimeMilliSeconds         string    `json:"statechangetimemsec,omitempty" nitro:"permission=readonly"`
	TicksSinceLastStateChange           string    `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	IsGslb                              bool      `json:"isgslb,omitempty" nitro:"permission=readonly"`
	SpilloverDynamicConnectionThreshold string    `json:"vsvrdynconnsothreshold,omitempty" nitro:"permission=readonly"`
	BackupVserverStatus                 string    `json:"backupvserverstatus,omitempty" nitro:"permission=readonly"`
	NoDefaultBindings                   string    `json:"nodefaultbindings,omitempty" nitro:"permission=readonly"`
	Count                               float64   `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r LbVserver) GetTypeName() string {
	return "lbvserver"
}

func NewLbVserverAddRequest(name string, servicetype string, ipaddress string, port int) LbVserver {
	return LbVserver{
		Name:        name,
		ServiceType: servicetype,
		Ipv46:       ipaddress,
		Port:        port,
	}
}

func NewLbVserverDisableRequest(name string) LbVserver {
	return LbVserver{
		Name: name,
	}
}

func NewLbVserverEnableRequest(name string) LbVserver {
	return LbVserver{
		Name: name,
	}
}

func NewLbVserverRenameRequest(oldName string, newName string) LbVserver {
	return LbVserver{
		Name:    oldName,
		NewName: newName,
	}
}
