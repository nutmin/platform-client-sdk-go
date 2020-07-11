package platformclientv2
import (
	"encoding/json"
)

// Notificationtemplatebody - Template body object
type Notificationtemplatebody struct { 
	// Text - Body text. For WhatsApp, ignored
	Text *string `json:"text,omitempty"`


	// Parameters - Template parameters for placeholders in template
	Parameters *[]Notificationtemplateparameter `json:"parameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplatebody) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
