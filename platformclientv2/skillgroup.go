package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgroup
type Skillgroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The group name.
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description - Group description
	Description *string `json:"description,omitempty"`


	// MemberCount - Estimated number of members in this group. It can take some time for the count to be updated after expressions are changed.
	MemberCount *int `json:"memberCount,omitempty"`


	// DateModified - Last modified date/time of the skill group. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DateCreated - Created date/time of the skill group. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// SkillConditions - Conditions for this group
	SkillConditions *[]Skillgroupcondition `json:"skillConditions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Skillgroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillgroup
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		SkillConditions *[]Skillgroupcondition `json:"skillConditions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		MemberCount: o.MemberCount,
		
		DateModified: DateModified,
		
		DateCreated: DateCreated,
		
		SkillConditions: o.SkillConditions,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillgroup) UnmarshalJSON(b []byte) error {
	var SkillgroupMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SkillgroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SkillgroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := SkillgroupMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := SkillgroupMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if MemberCount, ok := SkillgroupMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if dateModifiedString, ok := SkillgroupMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateCreatedString, ok := SkillgroupMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if SkillConditions, ok := SkillgroupMap["skillConditions"].([]interface{}); ok {
		SkillConditionsString, _ := json.Marshal(SkillConditions)
		json.Unmarshal(SkillConditionsString, &o.SkillConditions)
	}
	
	if SelfUri, ok := SkillgroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
