package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicmessage
type Conversationeventtopicmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

	// Direction - Whether a message is inbound or outbound.
	Direction *string `json:"direction,omitempty"`

	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`

	// ErrorInfo - Detailed information about an error response.
	ErrorInfo *Conversationeventtopicerrordetails `json:"errorInfo,omitempty"`

	// Provider - The source provider of the email.
	Provider *string `json:"provider,omitempty"`

	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// StartHoldTime - The timestamp the email was placed on hold in the cloud clock if the email is currently on hold.
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// ToAddress - Address and name data for a call endpoint.
	ToAddress *Conversationeventtopicaddress `json:"toAddress,omitempty"`

	// FromAddress - Address and name data for a call endpoint.
	FromAddress *Conversationeventtopicaddress `json:"fromAddress,omitempty"`

	// Messages - The messages sent on this communication channel.
	Messages *[]Conversationeventtopicmessagedetails `json:"messages,omitempty"`

	// MessagesTranscriptUri - the messages transcript file uri.
	MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`

	// VarType - Indicates the type of message platform from which the message originated.
	VarType *string `json:"type,omitempty"`

	// RecipientCountry - Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
	RecipientCountry *string `json:"recipientCountry,omitempty"`

	// RecipientType - The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
	RecipientType *string `json:"recipientType,omitempty"`

	// JourneyContext - A subset of the Journey System's data relevant to a part of a conversation (for external linkage and internal usage/context).
	JourneyContext *Conversationeventtopicjourneycontext `json:"journeyContext,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`

	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

	// AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

	// ByoSmsIntegrationId
	ByoSmsIntegrationId *string `json:"byoSmsIntegrationId,omitempty"`

	// QueueMediaSettings - Represents the queue setting for this media.
	QueueMediaSettings *Conversationeventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationeventtopicmessage) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Conversationeventtopicmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartHoldTime","ConnectedTime","DisconnectedTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicmessage
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		ErrorInfo *Conversationeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		ToAddress *Conversationeventtopicaddress `json:"toAddress,omitempty"`
		
		FromAddress *Conversationeventtopicaddress `json:"fromAddress,omitempty"`
		
		Messages *[]Conversationeventtopicmessagedetails `json:"messages,omitempty"`
		
		MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		RecipientCountry *string `json:"recipientCountry,omitempty"`
		
		RecipientType *string `json:"recipientType,omitempty"`
		
		JourneyContext *Conversationeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		ByoSmsIntegrationId *string `json:"byoSmsIntegrationId,omitempty"`
		
		QueueMediaSettings *Conversationeventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Direction: o.Direction,
		
		Held: o.Held,
		
		ErrorInfo: o.ErrorInfo,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		ToAddress: o.ToAddress,
		
		FromAddress: o.FromAddress,
		
		Messages: o.Messages,
		
		MessagesTranscriptUri: o.MessagesTranscriptUri,
		
		VarType: o.VarType,
		
		RecipientCountry: o.RecipientCountry,
		
		RecipientType: o.RecipientType,
		
		JourneyContext: o.JourneyContext,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AgentAssistantId: o.AgentAssistantId,
		
		ByoSmsIntegrationId: o.ByoSmsIntegrationId,
		
		QueueMediaSettings: o.QueueMediaSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationeventtopicmessage) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicmessageMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := ConversationeventtopicmessageMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := ConversationeventtopicmessageMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Direction, ok := ConversationeventtopicmessageMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Held, ok := ConversationeventtopicmessageMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if ErrorInfo, ok := ConversationeventtopicmessageMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Provider, ok := ConversationeventtopicmessageMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := ConversationeventtopicmessageMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := ConversationeventtopicmessageMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if DisconnectType, ok := ConversationeventtopicmessageMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startHoldTimeString, ok := ConversationeventtopicmessageMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := ConversationeventtopicmessageMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := ConversationeventtopicmessageMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if ToAddress, ok := ConversationeventtopicmessageMap["toAddress"].(map[string]interface{}); ok {
		ToAddressString, _ := json.Marshal(ToAddress)
		json.Unmarshal(ToAddressString, &o.ToAddress)
	}
	
	if FromAddress, ok := ConversationeventtopicmessageMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if Messages, ok := ConversationeventtopicmessageMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if MessagesTranscriptUri, ok := ConversationeventtopicmessageMap["messagesTranscriptUri"].(string); ok {
		o.MessagesTranscriptUri = &MessagesTranscriptUri
	}
    
	if VarType, ok := ConversationeventtopicmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if RecipientCountry, ok := ConversationeventtopicmessageMap["recipientCountry"].(string); ok {
		o.RecipientCountry = &RecipientCountry
	}
    
	if RecipientType, ok := ConversationeventtopicmessageMap["recipientType"].(string); ok {
		o.RecipientType = &RecipientType
	}
    
	if JourneyContext, ok := ConversationeventtopicmessageMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if Wrapup, ok := ConversationeventtopicmessageMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := ConversationeventtopicmessageMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := ConversationeventtopicmessageMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if AgentAssistantId, ok := ConversationeventtopicmessageMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if ByoSmsIntegrationId, ok := ConversationeventtopicmessageMap["byoSmsIntegrationId"].(string); ok {
		o.ByoSmsIntegrationId = &ByoSmsIntegrationId
	}
    
	if QueueMediaSettings, ok := ConversationeventtopicmessageMap["queueMediaSettings"].(map[string]interface{}); ok {
		QueueMediaSettingsString, _ := json.Marshal(QueueMediaSettings)
		json.Unmarshal(QueueMediaSettingsString, &o.QueueMediaSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
