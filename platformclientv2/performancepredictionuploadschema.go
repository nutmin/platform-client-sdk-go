package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Performancepredictionuploadschema
type Performancepredictionuploadschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CalculationStartDate - Date as an ISO-8601 string, corresponding to the beginning of the performance prediction results
	CalculationStartDate *time.Time `json:"calculationStartDate,omitempty"`

	// OnQueueTimes - List of agent on queue times by management unit
	OnQueueTimes *[]Muagentqueuetimerequest `json:"onQueueTimes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Performancepredictionuploadschema) SetField(field string, fieldValue interface{}) {
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

func (o Performancepredictionuploadschema) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CalculationStartDate", }
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
	type Alias Performancepredictionuploadschema
	
	CalculationStartDate := new(string)
	if o.CalculationStartDate != nil {
		
		*CalculationStartDate = timeutil.Strftime(o.CalculationStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CalculationStartDate = nil
	}
	
	return json.Marshal(&struct { 
		CalculationStartDate *string `json:"calculationStartDate,omitempty"`
		
		OnQueueTimes *[]Muagentqueuetimerequest `json:"onQueueTimes,omitempty"`
		Alias
	}{ 
		CalculationStartDate: CalculationStartDate,
		
		OnQueueTimes: o.OnQueueTimes,
		Alias:    (Alias)(o),
	})
}

func (o *Performancepredictionuploadschema) UnmarshalJSON(b []byte) error {
	var PerformancepredictionuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &PerformancepredictionuploadschemaMap)
	if err != nil {
		return err
	}
	
	if calculationStartDateString, ok := PerformancepredictionuploadschemaMap["calculationStartDate"].(string); ok {
		CalculationStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", calculationStartDateString)
		o.CalculationStartDate = &CalculationStartDate
	}
	
	if OnQueueTimes, ok := PerformancepredictionuploadschemaMap["onQueueTimes"].([]interface{}); ok {
		OnQueueTimesString, _ := json.Marshal(OnQueueTimes)
		json.Unmarshal(OnQueueTimesString, &o.OnQueueTimes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Performancepredictionuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
