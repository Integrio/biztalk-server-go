package client

import (
	"github.com/google/uuid"
)

// CertificateInfo Encryption certificate
type CertificateInfo struct {
	// CommonName Common Name
	CommonName string `json:"CommonName,omitempty"`

	// Thumbprint Thumbprint of the certificate.
	Thumbprint string `json:"Thumbprint,omitempty"`
}

// HostInstance The host instance.
type HostInstance struct {
	// HostName Name of the BizTalk Host this BizTalk Host instance belongs to
	HostName string `json:"HostName,omitempty"`

	// HostType Host type
	HostType string `json:"HostType,omitempty"`

	// IsDisabled Used to enable or disable the BizTalk Host instance. It can only be changed when the BizTalk Host instance is not started.
	IsDisabled *bool `json:"IsDisabled,omitempty"`

	// Logon Logon that this BizTalk Host instance is using.
	Logon string `json:"Logon,omitempty"`

	// NTGroupName NT group Name
	NTGroupName string `json:"NTGroupName,omitempty"`

	// Name Host Name
	Name string `json:"Name,omitempty"`

	// RunningServer Name of the server this BizTalk Host instance is running on.
	RunningServer string `json:"RunningServer,omitempty"`

	// ServiceState State of the given BizTalk Host instance.
	ServiceState string `json:"ServiceState,omitempty"`
}

// Instance The instance.
type Instance struct {
	// Adapter Adapter
	Adapter string `json:"Adapter,omitempty"`

	// Application Application
	Application string `json:"Application,omitempty"`

	// Class Class
	Class string `json:"Class,omitempty"`

	// CreationTime Creation Time
	CreationTime string `json:"CreationTime,omitempty"`

	// ErrorCode Error Code
	ErrorCode string `json:"ErrorCode,omitempty"`

	// ErrorDescription Error Description
	ErrorDescription string `json:"ErrorDescription,omitempty"`

	// HostName Host Name
	HostName string `json:"HostName,omitempty"`

	// Id Id
	Id *uuid.UUID `json:"Id,omitempty"`

	// InstanceStatus Instance Status
	InstanceStatus string `json:"InstanceStatus,omitempty"`

	// MessageBoxDb MessageBox database
	MessageBoxDb string `json:"MessageBoxDb,omitempty"`

	// MessageBoxServer MessageBox Server
	MessageBoxServer string `json:"MessageBoxServer,omitempty"`

	// PendingJobSubmitTime Pending Job Submit Time
	PendingJobSubmitTime string `json:"PendingJobSubmitTime,omitempty"`

	// PendingOperation PendingO peration
	PendingOperation string `json:"PendingOperation,omitempty"`

	// ProcessingServer ProcessingS erver
	ProcessingServer string `json:"ProcessingServer,omitempty"`

	// ServiceType Service Type
	ServiceType string `json:"ServiceType,omitempty"`

	// ServiceTypeID ServiceType Id
	ServiceTypeID *uuid.UUID `json:"ServiceTypeID,omitempty"`

	// SuspendTime Suspend Time
	SuspendTime string `json:"SuspendTime,omitempty"`

	// Uri Uri
	Uri string `json:"Uri,omitempty"`
}

// MessageBodyTracking Message body tracking
type MessageBodyTracking struct {
	// AfterReceivePipeline Flag indicating whether body to be tracked after port processing of response message
	AfterReceivePipeline *bool `json:"AfterReceivePipeline,omitempty"`

	// AfterSendPipeline Flag indicating whether body to be tracked after port processing of request message
	AfterSendPipeline *bool `json:"AfterSendPipeline,omitempty"`

	// BeforeReceivePipeline Flag indicating whether body to be tracked before port processing of response message
	BeforeReceivePipeline *bool `json:"BeforeReceivePipeline,omitempty"`

	// BeforeSendPipeline Flag indicating whether body to be tracked before port processing of request message
	BeforeSendPipeline *bool `json:"BeforeSendPipeline,omitempty"`
}

// MessagePropertyTracking Message Property Tracking
type MessagePropertyTracking struct {
	// AfterReceivePipeline Flag indicating whether body to be tracked after port processing of response message
	AfterReceivePipeline *bool `json:"AfterReceivePipeline,omitempty"`

	// AfterSendPipeline Flag indicating whether body to be tracked after port processing of request message
	AfterSendPipeline *bool `json:"AfterSendPipeline,omitempty"`

	// BeforeReceivePipeline Flag indicating whether body to be tracked before port processing of response message
	BeforeReceivePipeline *bool `json:"BeforeReceivePipeline,omitempty"`

	// BeforeSendPipeline Flag indicating whether body to be tracked before port processing of request message
	BeforeSendPipeline *bool `json:"BeforeSendPipeline,omitempty"`
}

// Orchestration Model representing Orchestration
type Orchestration struct {
	// AnalyticsEnabled Determines whether analytics is enabled
	AnalyticsEnabled *bool `json:"AnalyticsEnabled,omitempty"`

	// ApplicationName Application Name
	ApplicationName string `json:"ApplicationName,omitempty"`

	// AssemblyName Assembly Name
	AssemblyName string `json:"AssemblyName,omitempty"`

	// Description Description
	Description string `json:"Description,omitempty"`

	// FullName Orchestration Name
	FullName string `json:"FullName,omitempty"`

	// Host Host Name
	Host string `json:"Host,omitempty"`

	// ImplementedRoles Implemented Roles
	ImplementedRoles *[]string `json:"ImplementedRoles,omitempty"`

	// InboundPorts Inbound Ports
	InboundPorts *[]OrchestrationInboundPort `json:"InboundPorts,omitempty"`

	// InvokedOrchestrations Invoked Orchestrations
	InvokedOrchestrations *[]string `json:"InvokedOrchestrations,omitempty"`

	// OutboundPorts Outbound Ports
	OutboundPorts *[]OrchestrationOutboundPort `json:"OutboundPorts,omitempty"`

	// Status Status
	Status string `json:"Status,omitempty"`

	// Tracking Model representing Tracking Options for Orchestration
	Tracking *OrchestrationTrackingOptions `json:"Tracking,omitempty"`

	// UsedRoles Used Roles
	UsedRoles *[]string `json:"UsedRoles,omitempty"`
}

// OrchestrationInboundPort Model representing an Inbound Port
type OrchestrationInboundPort struct {
	// Binding Port Binding Type
	Binding string `json:"Binding,omitempty"`

	// Name Port Name
	Name string `json:"Name,omitempty"`

	// PortType Port Type
	PortType string `json:"PortType,omitempty"`

	// ReceivePort Physical Receive Port bound to Port
	ReceivePort string `json:"ReceivePort,omitempty"`
}

// OrchestrationOutboundPort Model representing an Outbound Port
type OrchestrationOutboundPort struct {
	// Binding Port Binding Type
	Binding string `json:"Binding,omitempty"`

	// Name Port Name
	Name string `json:"Name,omitempty"`

	// PortType Port Type
	PortType string `json:"PortType,omitempty"`

	// SendPort Physical SendPort bound to Port
	SendPort string `json:"SendPort,omitempty"`

	// SendPortGroup Physical SendPortGroup bound to Port
	SendPortGroup string `json:"SendPortGroup,omitempty"`
}

// OrchestrationTrackingOptions Model representing Tracking Options for Orchestration
type OrchestrationTrackingOptions struct {
	// InboundMessageBody Inbound Message Body
	InboundMessageBody *bool `json:"InboundMessageBody,omitempty"`

	// MessageSendReceive Message Send Receive
	MessageSendReceive *bool `json:"MessageSendReceive,omitempty"`

	// OrchestartionEvents Orchestration Events
	OrchestartionEvents *bool `json:"OrchestartionEvents,omitempty"`

	// OutboundMessageBody Outbound Message Body
	OutboundMessageBody *bool `json:"OutboundMessageBody,omitempty"`

	// ServiceStartEnd Service Start End
	ServiceStartEnd *bool `json:"ServiceStartEnd,omitempty"`

	// TrackPropertiesForIncomingMessages Track Properties for Incoming Messages
	TrackPropertiesForIncomingMessages *bool `json:"TrackPropertiesForIncomingMessages,omitempty"`

	// TrackPropertiesForOutgoingMessages Track Properties for Outgoing Messages
	TrackPropertiesForOutgoingMessages *bool `json:"TrackPropertiesForOutgoingMessages,omitempty"`
}

// ReceiveLocation Receive location model
type ReceiveLocation struct {
	// Address Address property for the adapter.
	Address string `json:"Address,omitempty"`

	// CustomData Custom data
	CustomData string `json:"CustomData,omitempty"`

	// Description Description for the receive location.
	Description string `json:"Description,omitempty"`

	// Enable Status of the receive location.
	Enable *bool `json:"Enable,omitempty"`

	// EncryptionCert Encryption certificate
	EncryptionCert *CertificateInfo `json:"EncryptionCert,omitempty"`

	// FragmentMessages Field to fragment messages
	FragmentMessages string `json:"FragmentMessages,omitempty"`

	// IsPrimary Determines whether the receive location is primary receive location of the receive port
	IsPrimary *bool `json:"IsPrimary,omitempty"`

	// Name Name of the receive location.
	Name string `json:"Name,omitempty"`

	// PublicAddress Public Address property for the adapter.
	PublicAddress string `json:"PublicAddress,omitempty"`

	// ReceiveHandler Receive handler for this eceive location.
	ReceiveHandler string `json:"ReceiveHandler,omitempty"`

	// ReceivePipeline Receive pipeline used to receive a response when a message is sent out through this receive location.
	ReceivePipeline string `json:"ReceivePipeline,omitempty"`

	// ReceivePipelineData Custom configuration of receive pipeline specific to this instance of the usage of the pipeline.
	ReceivePipelineData string `json:"ReceivePipelineData,omitempty"`

	// ReceivePortName Receive port name to which this receive location belongs to
	ReceivePortName string `json:"ReceivePortName,omitempty"`

	// Schedule Service window
	Schedule *ReceiveLocationSchedule `json:"Schedule,omitempty"`

	// SendPipeline Send pipeline used to send data sent through this receive location.
	SendPipeline string `json:"SendPipeline,omitempty"`

	// SendPipelineData Custom configuration of send pipeline specific to this instance of the usage of the pipeline.
	SendPipelineData string `json:"SendPipelineData,omitempty"`

	// TransportType Protocol for the adapter.
	TransportType string `json:"TransportType,omitempty"`

	// TransportTypeData Configuration specific to the adapter.
	TransportTypeData string `json:"TransportTypeData,omitempty"`
}

// ReceiveLocationSchedule Service window
type ReceiveLocationSchedule struct {
	// AutoAdjustToDaylightSaving If auto adjust to day light saving
	AutoAdjustToDaylightSaving *bool `json:"AutoAdjustToDaylightSaving,omitempty"`

	// DaysOfWeek Days of week if recurrence schedule type is Weekly
	DaysOfWeek string `json:"DaysOfWeek,omitempty"`

	// EndDate Endt date of the service window
	EndDate string `json:"EndDate,omitempty"`

	// EndDateEnabled Flag indiacting if end date of the service window is enabled.
	EndDateEnabled *bool `json:"EndDateEnabled,omitempty"`

	// FromTime Start time for the service window.
	FromTime string `json:"FromTime,omitempty"`

	// LastDayOfMonth If schedule on last day Of month
	LastDayOfMonth *bool `json:"LastDayOfMonth,omitempty"`

	// MonthDays Days of month if recurrence schedule type is Months
	MonthDays string `json:"MonthDays,omitempty"`

	// Months Months of year if recurrence schedule type is Months
	Months string `json:"Months,omitempty"`

	// OrdinalDayOfWeek Ordinal Day Of Week if recurrence schedule type is Months
	OrdinalDayOfWeek string `json:"OrdinalDayOfWeek,omitempty"`

	// OrdinalSchedule Ordinal type if recurrence schedule type is Months
	OrdinalSchedule string `json:"OrdinalSchedule,omitempty"`

	// RecurFrom Recur From
	RecurFrom string `json:"RecurFrom,omitempty"`

	// RecurInterval Recurrence Interval
	RecurInterval *int32 `json:"RecurInterval,omitempty"`

	// RecurrenceSchType Recurrence Schedule Type
	RecurrenceSchType string `json:"RecurrenceSchType,omitempty"`

	// ScheduleIsOrdinal if Schedule Is Ordinal
	ScheduleIsOrdinal *bool `json:"ScheduleIsOrdinal,omitempty"`

	// ScheduleTimeZone Time zone
	ScheduleTimeZone string `json:"ScheduleTimeZone,omitempty"`

	// ServiceWindowEnabled Flag that specifies whether the service window is enabled.
	ServiceWindowEnabled *bool `json:"ServiceWindowEnabled,omitempty"`

	// StartDate Start date of the service window
	StartDate string `json:"StartDate,omitempty"`

	// StartDateEnabled Flag indicating if start date of the service window is enabled.
	StartDateEnabled *bool `json:"StartDateEnabled,omitempty"`

	// ToTime End time for the service window.
	ToTime string `json:"ToTime,omitempty"`
}

// ReceivePort Model for Receive port
type ReceivePort struct {
	// AnalyticsEnabled Determines whether analytics is enabled
	AnalyticsEnabled *bool `json:"AnalyticsEnabled,omitempty"`

	// ApplicationName Application name to which receive port belongs to
	ApplicationName string `json:"ApplicationName,omitempty"`

	// CustomData Custom data
	CustomData string `json:"CustomData,omitempty"`

	// Description Description for the receive port.
	Description string `json:"Description,omitempty"`

	// InboundTransforms Collection of inbound transforms of receive port
	InboundTransforms *[]string `json:"InboundTransforms,omitempty"`

	// IsTwoWay Determines whether the port is two-way
	IsTwoWay *bool `json:"IsTwoWay,omitempty"`

	// Name Name of the recieve port.
	Name string `json:"Name,omitempty"`

	// OutboundTransforms Collection of outbound transforms of a two-way receive port
	OutboundTransforms *[]string `json:"OutboundTransforms,omitempty"`

	// PrimaryReceiveLocation Primary receive location name
	PrimaryReceiveLocation string `json:"PrimaryReceiveLocation,omitempty"`

	// ReceiveLocations List of receive locations of the receive port
	ReceiveLocations *[]string `json:"ReceiveLocations,omitempty"`

	// Tracking Port tracking details
	Tracking *Tracking `json:"Tracking,omitempty"`
}

// Schedule Service window
type Schedule struct {
	// FromTime Start time for the service window.
	FromTime string `json:"FromTime,omitempty"`

	// ServiceWindowEnabled Flag that specifies whether the service window is enabled.
	ServiceWindowEnabled *bool `json:"ServiceWindowEnabled,omitempty"`

	// ToTime End time for the service window.
	ToTime string `json:"ToTime,omitempty"`
}

// SendPort Model for Send port
type SendPort struct {
	// AnalyticsEnabled Determines whether analytics is enabled
	AnalyticsEnabled *bool `json:"AnalyticsEnabled,omitempty"`

	// ApplicationName Application name to which sendport belongs to
	ApplicationName string `json:"ApplicationName,omitempty"`

	// CustomData Custom data
	CustomData string `json:"CustomData,omitempty"`

	// Description Description for the send port.
	Description string `json:"Description,omitempty"`

	// EncryptionCert Encryption certificate
	EncryptionCert *CertificateInfo `json:"EncryptionCert,omitempty"`

	// Filter Optional filter expression used for this send port.
	Filter string `json:"Filter,omitempty"`

	// InboundTransforms Collection of inbound transforms of two-way send port
	InboundTransforms *[]string `json:"InboundTransforms,omitempty"`

	// IsDynamic Whether the send port is dynamic or static
	IsDynamic *bool `json:"IsDynamic,omitempty"`

	// IsTwoWay Determines whether the port is two-way
	IsTwoWay *bool `json:"IsTwoWay,omitempty"`

	// Name Name of the send port.
	Name string `json:"Name,omitempty"`

	// OrderPerAddress Flag indicating whether enforce order across outbound locations for dynamic send port.
	OrderPerAddress *bool `json:"OrderPerAddress,omitempty"`

	// OrderedDelivery Flag specifying whether or not the send port orders the delivery of messages.
	OrderedDelivery *bool `json:"OrderedDelivery,omitempty"`

	// OutboundTransforms Collection of outbound transforms of send port
	OutboundTransforms *[]string `json:"OutboundTransforms,omitempty"`

	// PrimaryTransport The transport info.
	PrimaryTransport *TransportInfo `json:"PrimaryTransport,omitempty"`

	// Priority Priority of the send port.
	Priority *int32 `json:"Priority,omitempty"`

	// ReceivePipeline Receive pipeline used to receive a response when a message is sent out through this port.
	ReceivePipeline string `json:"ReceivePipeline,omitempty"`

	// ReceivePipelineData Custom configuration information specific to the current instance of the receive pipeline.
	ReceivePipelineData string `json:"ReceivePipelineData,omitempty"`

	// RouteFailedMessage Gets or sets a value indicating whether or not failed messages are routed to failed message subscribers.
	RouteFailedMessage *bool `json:"RouteFailedMessage,omitempty"`

	// SecondaryTransport The transport info.
	SecondaryTransport *TransportInfo `json:"SecondaryTransport,omitempty"`

	// SendPipeline Send pipeline used to send data sent through this port.
	SendPipeline string `json:"SendPipeline,omitempty"`

	// SendPipelineData Custom configuration specific to this instance of the usage of the pipeline.
	SendPipelineData string `json:"SendPipelineData,omitempty"`

	// Status Status of the send port.
	Status string `json:"Status,omitempty"`

	// StopSendingOnFailure Flag indicating whether or not the send port stops sending messages on a failure.
	StopSendingOnFailure *bool `json:"StopSendingOnFailure,omitempty"`

	// Tracking Port tracking details
	Tracking *Tracking `json:"Tracking,omitempty"`
}

// SendPortGroup Model for SendPortGroup
type SendPortGroup struct {
	// ApplicationName Application Name
	ApplicationName string `json:"ApplicationName,omitempty"`

	// CustomData Custom Data
	CustomData string `json:"CustomData,omitempty"`

	// Description Description
	Description string `json:"Description,omitempty"`

	// Filter Filter
	Filter string `json:"Filter,omitempty"`

	// Name Name of the SendPortGroup
	Name string `json:"Name,omitempty"`

	// SendPorts List of SendPorts in the group
	SendPorts *[]string `json:"SendPorts,omitempty"`

	// Status Port Status
	Status string `json:"Status,omitempty"`
}

// TrackedMessageEvent TrackedMessageEvent
type TrackedMessageEvent struct {
	// ActivityId Activity Id
	ActivityId *uuid.UUID `json:"ActivityId,omitempty"`

	// Adapter Adapter
	Adapter string `json:"Adapter,omitempty"`

	// DecryptionCertificate Decryption Certificate
	DecryptionCertificate string `json:"DecryptionCertificate,omitempty"`

	// EventType Event Type
	EventType string `json:"EventType,omitempty"`

	// Id Id
	Id *uuid.UUID `json:"Id,omitempty"`

	// MessageInstanceId Message Instance Id
	MessageInstanceId *uuid.UUID `json:"MessageInstanceId,omitempty"`

	// PartCount Part Count
	PartCount *int32 `json:"PartCount,omitempty"`

	// PartyName Party Name
	PartyName string `json:"PartyName,omitempty"`

	// PortName Port Name
	PortName string `json:"PortName,omitempty"`

	// SchemaName Schema Name
	SchemaName string `json:"SchemaName,omitempty"`

	// ServiceClass ServiceClass
	ServiceClass string `json:"ServiceClass,omitempty"`

	// ServiceClassId Service Class Id
	ServiceClassId *uuid.UUID `json:"ServiceClassId,omitempty"`

	// ServiceInstanceId Service Instance Id
	ServiceInstanceId *uuid.UUID `json:"ServiceInstanceId,omitempty"`

	// ServiceName Service Name
	ServiceName string `json:"ServiceName,omitempty"`

	// Signature Signature
	Signature string `json:"Signature,omitempty"`

	// Size Size
	Size *int64 `json:"Size,omitempty"`

	// TimeStamp TimeStamp
	TimeStamp string `json:"TimeStamp,omitempty"`

	// Uri Uri
	Uri string `json:"Uri,omitempty"`
}

// TrackedServiceInstance TrackedServiceInstance
type TrackedServiceInstance struct {
	// AssemblyName Assembly Name
	AssemblyName string `json:"AssemblyName,omitempty"`

	// AssemblyVersion Assembly Version
	AssemblyVersion string `json:"AssemblyVersion,omitempty"`

	// DeploymentTime Deployment Time
	DeploymentTime string `json:"DeploymentTime,omitempty"`

	// Duration Duration
	Duration *int32 `json:"Duration,omitempty"`

	// EndTime EndTime
	EndTime string `json:"EndTime,omitempty"`

	// ErrorCode Error Code
	ErrorCode *int64 `json:"ErrorCode,omitempty"`

	// ErrorDescription Error Description
	ErrorDescription string `json:"ErrorDescription,omitempty"`

	// HostName Host Name
	HostName string `json:"HostName,omitempty"`

	// Id Id
	Id *uuid.UUID `json:"Id,omitempty"`

	// ServiceClass Service Class
	ServiceClass string `json:"ServiceClass,omitempty"`

	// ServiceClassId Service Class Id
	ServiceClassId *uuid.UUID `json:"ServiceClassId,omitempty"`

	// ServiceId Service Id
	ServiceId *uuid.UUID `json:"ServiceId,omitempty"`

	// ServiceInstanceId Service Instance Id
	ServiceInstanceId *uuid.UUID `json:"ServiceInstanceId,omitempty"`

	// ServiceName Service Name
	ServiceName string `json:"ServiceName,omitempty"`

	// ServiceVersionId Service Version Id
	ServiceVersionId *uuid.UUID `json:"ServiceVersionId,omitempty"`

	// StartTime StartTime
	StartTime string `json:"StartTime,omitempty"`

	// State State
	State string `json:"State,omitempty"`
}

// Tracking Port tracking details
type Tracking struct {
	// Body Message body tracking
	Body *MessageBodyTracking `json:"Body,omitempty"`

	// Property Message Property Tracking
	Property *MessagePropertyTracking `json:"Property,omitempty"`
}

// TransportInfo The transport info.
type TransportInfo struct {
	// Address Address property for the adapter.
	Address string `json:"Address,omitempty"`

	// OrderedDelivery Flag that specifies whether the BizTalk Message Queuing adapter supports ordered delivery.
	OrderedDelivery *bool `json:"OrderedDelivery,omitempty"`

	// RetryCount Retry count for the adapter.
	RetryCount *int32 `json:"RetryCount,omitempty"`

	// RetryInterval Retry interval for the adapter.
	RetryInterval *int32 `json:"RetryInterval,omitempty"`

	// Schedule Service window
	Schedule *Schedule `json:"Schedule,omitempty"`

	// SendHandler Send handler for this transport.
	SendHandler string `json:"SendHandler,omitempty"`

	// TransportType Protocol for the adapter.
	TransportType string `json:"TransportType,omitempty"`

	// TransportTypeData Configuration specific to the adapter.
	TransportTypeData string `json:"TransportTypeData,omitempty"`
}
