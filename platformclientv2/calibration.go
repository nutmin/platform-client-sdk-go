package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Calibration
type Calibration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Conversation
	Conversation *Conversationreference `json:"conversation,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// AverageScore
	AverageScore *int `json:"averageScore,omitempty"`


	// HighScore
	HighScore *int `json:"highScore,omitempty"`


	// LowScore
	LowScore *int `json:"lowScore,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Evaluations
	Evaluations *[]Evaluation `json:"evaluations,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// ScoringIndex
	ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`


	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Calibration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Calibration
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Calibrator *User `json:"calibrator,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		AverageScore *int `json:"averageScore,omitempty"`
		
		HighScore *int `json:"highScore,omitempty"`
		
		LowScore *int `json:"lowScore,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Evaluations *[]Evaluation `json:"evaluations,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`
		
		ExpertEvaluator *User `json:"expertEvaluator,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Calibrator: o.Calibrator,
		
		Agent: o.Agent,
		
		Conversation: o.Conversation,
		
		EvaluationForm: o.EvaluationForm,
		
		ContextId: o.ContextId,
		
		AverageScore: o.AverageScore,
		
		HighScore: o.HighScore,
		
		LowScore: o.LowScore,
		
		CreatedDate: CreatedDate,
		
		Evaluations: o.Evaluations,
		
		Evaluators: o.Evaluators,
		
		ScoringIndex: o.ScoringIndex,
		
		ExpertEvaluator: o.ExpertEvaluator,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Calibration) UnmarshalJSON(b []byte) error {
	var CalibrationMap map[string]interface{}
	err := json.Unmarshal(b, &CalibrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CalibrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CalibrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Calibrator, ok := CalibrationMap["calibrator"].(map[string]interface{}); ok {
		CalibratorString, _ := json.Marshal(Calibrator)
		json.Unmarshal(CalibratorString, &o.Calibrator)
	}
	
	if Agent, ok := CalibrationMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Conversation, ok := CalibrationMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if EvaluationForm, ok := CalibrationMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if ContextId, ok := CalibrationMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if AverageScore, ok := CalibrationMap["averageScore"].(float64); ok {
		AverageScoreInt := int(AverageScore)
		o.AverageScore = &AverageScoreInt
	}
	
	if HighScore, ok := CalibrationMap["highScore"].(float64); ok {
		HighScoreInt := int(HighScore)
		o.HighScore = &HighScoreInt
	}
	
	if LowScore, ok := CalibrationMap["lowScore"].(float64); ok {
		LowScoreInt := int(LowScore)
		o.LowScore = &LowScoreInt
	}
	
	if createdDateString, ok := CalibrationMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if Evaluations, ok := CalibrationMap["evaluations"].([]interface{}); ok {
		EvaluationsString, _ := json.Marshal(Evaluations)
		json.Unmarshal(EvaluationsString, &o.Evaluations)
	}
	
	if Evaluators, ok := CalibrationMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if ScoringIndex, ok := CalibrationMap["scoringIndex"].(map[string]interface{}); ok {
		ScoringIndexString, _ := json.Marshal(ScoringIndex)
		json.Unmarshal(ScoringIndexString, &o.ScoringIndex)
	}
	
	if ExpertEvaluator, ok := CalibrationMap["expertEvaluator"].(map[string]interface{}); ok {
		ExpertEvaluatorString, _ := json.Marshal(ExpertEvaluator)
		json.Unmarshal(ExpertEvaluatorString, &o.ExpertEvaluator)
	}
	
	if SelfUri, ok := CalibrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Calibration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
