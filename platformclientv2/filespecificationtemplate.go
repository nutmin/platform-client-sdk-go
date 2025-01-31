package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Filespecificationtemplate
type Filespecificationtemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the File Specification template.
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

	// Description - Description of the file specification template
	Description *string `json:"description,omitempty"`

	// Format - File format
	Format *string `json:"format,omitempty"`

	// NumberOfHeadingLinesSkipped - Number of heading lines to be skipped
	NumberOfHeadingLinesSkipped *int `json:"numberOfHeadingLinesSkipped,omitempty"`

	// NumberOfTrailingLinesSkipped - Number of trailing lines to be skipped
	NumberOfTrailingLinesSkipped *int `json:"numberOfTrailingLinesSkipped,omitempty"`

	// Header - If true indicates that delimited file has a header row, which can provide column names
	Header *bool `json:"header,omitempty"`

	// Delimiter - Kind of delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// DelimiterValue - Delimiter character, used only when delimiter=\"Custom\"
	DelimiterValue *string `json:"delimiterValue,omitempty"`

	// ColumnInformation - Columns specification
	ColumnInformation *[]Column `json:"columnInformation,omitempty"`

	// PreprocessingRules - Preprocessing rules
	PreprocessingRules *[]Preprocessingrule `json:"preprocessingRules,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Filespecificationtemplate) SetField(field string, fieldValue interface{}) {
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

func (o Filespecificationtemplate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Filespecificationtemplate
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Format *string `json:"format,omitempty"`
		
		NumberOfHeadingLinesSkipped *int `json:"numberOfHeadingLinesSkipped,omitempty"`
		
		NumberOfTrailingLinesSkipped *int `json:"numberOfTrailingLinesSkipped,omitempty"`
		
		Header *bool `json:"header,omitempty"`
		
		Delimiter *string `json:"delimiter,omitempty"`
		
		DelimiterValue *string `json:"delimiterValue,omitempty"`
		
		ColumnInformation *[]Column `json:"columnInformation,omitempty"`
		
		PreprocessingRules *[]Preprocessingrule `json:"preprocessingRules,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Description: o.Description,
		
		Format: o.Format,
		
		NumberOfHeadingLinesSkipped: o.NumberOfHeadingLinesSkipped,
		
		NumberOfTrailingLinesSkipped: o.NumberOfTrailingLinesSkipped,
		
		Header: o.Header,
		
		Delimiter: o.Delimiter,
		
		DelimiterValue: o.DelimiterValue,
		
		ColumnInformation: o.ColumnInformation,
		
		PreprocessingRules: o.PreprocessingRules,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Filespecificationtemplate) UnmarshalJSON(b []byte) error {
	var FilespecificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &FilespecificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FilespecificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FilespecificationtemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := FilespecificationtemplateMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := FilespecificationtemplateMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := FilespecificationtemplateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Description, ok := FilespecificationtemplateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Format, ok := FilespecificationtemplateMap["format"].(string); ok {
		o.Format = &Format
	}
    
	if NumberOfHeadingLinesSkipped, ok := FilespecificationtemplateMap["numberOfHeadingLinesSkipped"].(float64); ok {
		NumberOfHeadingLinesSkippedInt := int(NumberOfHeadingLinesSkipped)
		o.NumberOfHeadingLinesSkipped = &NumberOfHeadingLinesSkippedInt
	}
	
	if NumberOfTrailingLinesSkipped, ok := FilespecificationtemplateMap["numberOfTrailingLinesSkipped"].(float64); ok {
		NumberOfTrailingLinesSkippedInt := int(NumberOfTrailingLinesSkipped)
		o.NumberOfTrailingLinesSkipped = &NumberOfTrailingLinesSkippedInt
	}
	
	if Header, ok := FilespecificationtemplateMap["header"].(bool); ok {
		o.Header = &Header
	}
    
	if Delimiter, ok := FilespecificationtemplateMap["delimiter"].(string); ok {
		o.Delimiter = &Delimiter
	}
    
	if DelimiterValue, ok := FilespecificationtemplateMap["delimiterValue"].(string); ok {
		o.DelimiterValue = &DelimiterValue
	}
    
	if ColumnInformation, ok := FilespecificationtemplateMap["columnInformation"].([]interface{}); ok {
		ColumnInformationString, _ := json.Marshal(ColumnInformation)
		json.Unmarshal(ColumnInformationString, &o.ColumnInformation)
	}
	
	if PreprocessingRules, ok := FilespecificationtemplateMap["preprocessingRules"].([]interface{}); ok {
		PreprocessingRulesString, _ := json.Marshal(PreprocessingRules)
		json.Unmarshal(PreprocessingRulesString, &o.PreprocessingRules)
	}
	
	if SelfUri, ok := FilespecificationtemplateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Filespecificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
