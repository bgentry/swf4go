package swf

/*WorkflowProtocol*/
type StartWorkflowRequest struct {
	ChildPolicy                  string       `json:"childPolicy"`
	Domain                       string       `json:"domain"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout"`
	Input                        string       `json:"input"`
	TagList                      []string     `json:"tagList"`
	TaskList                     TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout"`
	WorkflowId                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}

type StartWorkflowResponse struct {
	RunId string `json:"runId"`
}

type RequestCancelWorkflowExecution struct {
	Domain     string `json:"domain"`
	RunId      string `json:"runId"`
	WorkflowId string `json:"workflowId"`
}

type SignalWorkflowRequest struct {
	Domain     string `json:"domain"`
	Input      string `json:"input"`
	RunId      string `json:"runId"`
	SignalName string `json:"signalName"`
	WorkflowId string `json:"workflowId"`
}

type ListWorkflowTypesRequest struct {
	Domain             string  `json:"domain"`
	MaximumPageSize    int     `json:"maximumPageSize"`
	Name               *string `json:"name"`
	NextPageToken      *string `json:"nextPageToken"`
	RegistrationStatus string  `json:"registrationStatus"`
	ReverseOrder       bool    `json:"reverseOrder"`
}

type ListWorkflowTypesResponse struct {
	NextPageToken string     `json:"nextPageToken"`
	TypeInfos     []TypeInfo `json:"typeInfos"`
}

type TerminateWorkflowExecution struct {
	ChildPolicy string `json:"childPolicy"`
	Details     string `json:"details"`
	Domain      string `json:"domain"`
	Reason      string `json:"reason"`
	RunId       string `json:"runId"`
	WorkflowId  string `json:"workflowId"`
}

/*DecisionWorkerProtocol*/

type PollForDecisionTaskRequest struct {
	Domain          string   `json:"domain"`
	Identity        string   `json:"identity"`
	MaximumPageSize int      `json:"maximumPageSize"`
	NextPageToken   *string  `json:"nextPageToken"`
	ReverseOrder    bool     `json:"reverseOrder"`
	TaskList        TaskList `json:"taskList"`
}

type HistoryEvent struct {
	ActivityTaskCancelRequestedEventAttributes                     *ActivityTaskCancelRequestedEventAttributes                     `json:"activityTaskCancelRequestedEventAttributes"`
	ActivityTaskCanceledEventAttributes                            *ActivityTaskCanceledEventAttributes                            `json:"activityTaskCanceledEventAttributes"`
	ActivityTaskCompletedEventAttributes                           *ActivityTaskCompletedEventAttributes                           `json:"activityTaskCompletedEventAttributes"`
	ActivityTaskFailedEventAttributes                              *ActivityTaskFailedEventAttributes                              `json:"activityTaskFailedEventAttributes"`
	ActivityTaskScheduledEventAttributes                           *ActivityTaskScheduledEventAttributes                           `json:"activityTaskScheduledEventAttributes"`
	ActivityTaskStartedEventAttributes                             *ActivityTaskStartedEventAttributes                             `json:"activityTaskStartedEventAttributes"`
	ActivityTaskTimedOutEventAttributes                            *ActivityTaskTimedOutEventAttributes                            `json:"activityTaskTimedOutEventAttributes"`
	CancelTimerFailedEventAttributes                               *CancelTimerFailedEventAttributes                               `json:"cancelTimerFailedEventAttributes"`
	CancelWorkflowExecutionFailedEventAttributes                   CancelWorkflowExecutionFailedEventAttributes                    `json:"cancelWorkflowExecutionFailedEventAttributes"`
	ChildWorkflowExecutionCanceledEventAttributes                  *ChildWorkflowExecutionCanceledEventAttributes                  `json:"childWorkflowExecutionCanceledEventAttributes"`
	ChildWorkflowExecutionCompletedEventAttributes                 *ChildWorkflowExecutionCompletedEventAttributes                 `json:"childWorkflowExecutionCompletedEventAttributes"`
	ChildWorkflowExecutionFailedEventAttributes                    *CancelWorkflowExecutionFailedEventAttributes                   `json:"childWorkflowExecutionFailedEventAttributes"`
	ChildWorkflowExecutionStartedEventAttributes                   *ChildWorkflowExecutionStartedEventAttributes                   `json:"childWorkflowExecutionStartedEventAttributes"`
	ChildWorkflowExecutionTerminatedEventAttributes                *ChildWorkflowExecutionTerminatedEventAttributes                `json:"childWorkflowExecutionTerminatedEventAttributes"`
	ChildWorkflowExecutionTimedOutEventAttributes                  *ChildWorkflowExecutionTimedOutEventAttributes                  `json:"childWorkflowExecutionTimedOutEventAttributes"`
	CompleteWorkflowExecutionFailedEventAttributes                 *CompleteWorkflowExecutionFailedEventAttributes                 `json:"completeWorkflowExecutionFailedEventAttributes"`
	ContinueAsNewWorkflowExecutionFailedEventAttributes            *ContinueAsNewWorkflowExecutionFailedEventAttributes            `json:"continueAsNewWorkflowExecutionFailedEventAttributes"`
	DecisionTaskCompletedEventAttributes                           *DecisionTaskCompletedEventAttributes                           `json:"decisionTaskCompletedEventAttributes"`
	DecisionTaskScheduledEventAttributes                           *DecisionTaskScheduledEventAttributes                           `json:"decisionTaskScheduledEventAttributes"`
	DecisionTaskStartedEventAttributes                             *DecisionTaskStartedEventAttributes                             `json:"decisionTaskStartedEventAttributes"`
	DecisionTaskTimedOutEventAttributes                            *DecisionTaskTimedOutEventAttributes                            `json:"decisionTaskTimedOutEventAttributes"`
	EventId                                                        int                                                             `json:"eventId"`
	EventTimestamp                                                 float32                                                         `json:"eventTimestamp"`
	EventType                                                      string                                                          `json:"eventType"`
	ExternalWorkflowExecutionCancelRequestedEventAttributes        *ExternalWorkflowExecutionCancelRequestedEventAttributes        `json:"externalWorkflowExecutionCancelRequestedEventAttributes"`
	ExternalWorkflowExecutionSignaledEventAttributes               *ExternalWorkflowExecutionSignaledEventAttributes               `json:"externalWorkflowExecutionSignaledEventAttributes"`
	FailWorkflowExecutionFailedEventAttributes                     *FailWorkflowExecutionFailedEventAttributes                     `json:"failWorkflowExecutionFailedEventAttributes"`
	MarkerRecordedEventAttributes                                  *MarkerRecordedEventAttributes                                  `json:"markerRecordedEventAttributes"`
	RecordMarkerFailedEventAttributes                              *RecordMarkerFailedEventAttributes                              `json:"recordMarkerFailedEventAttributes"`
	RequestCancelActivityTaskFailedEventAttributes                 *RequestCancelActivityTaskFailedEventAttributes                 `json:"requestCancelActivityTaskFailedEventAttributes"`
	RequestCancelExternalWorkflowExecutionFailedEventAttributes    *RequestCancelExternalWorkflowExecutionFailedEventAttributes    `json:"requestCancelExternalWorkflowExecutionFailedEventAttributes"`
	RequestCancelExternalWorkflowExecutionInitiatedEventAttributes *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes `json:"requestCancelExternalWorkflowExecutionInitiatedEventAttributes"`
	ScheduleActivityTaskFailedEventAttributes                      *ScheduleActivityTaskFailedEventAttributes                      `json:"scheduleActivityTaskFailedEventAttributes"`
	SignalExternalWorkflowExecutionFailedEventAttributes           *SignalExternalWorkflowExecutionFailedEventAttributes           `json:"signalExternalWorkflowExecutionFailedEventAttributes"`
	SignalExternalWorkflowExecutionInitiatedEventAttributes        *SignalExternalWorkflowExecutionInitiatedEventAttributes        `json:"signalExternalWorkflowExecutionInitiatedEventAttributes"`
	StartChildWorkflowExecutionFailedEventAttributes               *StartChildWorkflowExecutionFailedEventAttributes               `json:"startChildWorkflowExecutionFailedEventAttributes"`
	StartChildWorkflowExecutionInitiatedEventAttributes            *StartChildWorkflowExecutionInitiatedEventAttributes            `json:"startChildWorkflowExecutionInitiatedEventAttributes"`
	StartTimerFailedEventAttributes                                *StartTimerFailedEventAttributes                                `json:"startTimerFailedEventAttributes"`
	TimerCanceledEventAttributes                                   *TimerCanceledEventAttributes                                   `json:"timerCanceledEventAttributes"`
	TimerFiredEventAttributes                                      *TimerFiredEventAttributes                                      `json:"timerFiredEventAttributes"`
	TimerStartedEventAttributes                                    *TimerStartedEventAttributes                                    `json:"timerStartedEventAttributes"`
	WorkflowExecutionCancelRequestedEventAttributes                *WorkflowExecutionCancelRequestedEventAttributes                `json:"workflowExecutionCancelRequestedEventAttributes"`
	WorkflowExecutionCanceledEventAttributes                       *WorkflowExecutionCanceledEventAttributes                       `json:"workflowExecutionCanceledEventAttributes"`
	WorkflowExecutionCompletedEventAttributes                      *WorkflowExecutionCompletedEventAttributes                      `json:"workflowExecutionCompletedEventAttributes"`
	WorkflowExecutionContinuedAsNewEventAttributes                 *WorkflowExecutionContinuedAsNewEventAttributes                 `json:"workflowExecutionContinuedAsNewEventAttributes"`
	WorkflowExecutionFailedEventAttributes                         *WorkflowExecutionFailedEventAttributes                         `json:"workflowExecutionFailedEventAttributes"`
	WorkflowExecutionSignaledEventAttributes                       *WorkflowExecutionSignaledEventAttributes                       `json:"workflowExecutionSignaledEventAttributes"`
	WorkflowExecutionStartedEventAttributes                        *WorkflowExecutionStartedEventAttributes                        `json:"workflowExecutionStartedEventAttributes"`
	WorkflowExecutionTerminatedEventAttributes                     *WorkflowExecutionTerminatedEventAttributes                     `json:"workflowExecutionTerminatedEventAttributes"`
	WorkflowExecutionTimedOutEventAttributes                       *WorkflowExecutionTimedOutEventAttributes                       `json:"workflowExecutionTimedOutEventAttributes"`
}

type ActivityTaskCancelRequestedEventAttributes struct {
	ActivityId                   string `json:"activityId"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type ActivityTaskCanceledEventAttributes struct {
	Details                      string `json:"details"`
	LatestCancelRequestedEventId string `json:"latestCancelRequestedEventId"`
	ScheduledEventId             string `json:"scheduledEventId"`
	StartedEventId               string `json:"startedEventId"`
}
type ActivityTaskCompletedEventAttributes struct {
	Result           string `json:"result"`
	ScheduledEventId string `json:"scheduledEventId"`
	StartedEventId   string `json:"startedEventId"`
}
type ActivityTaskFailedEventAttributes struct {
	Details          string `json:"details"`
	Reason           string `json:"reason"`
	ScheduledEventId string `json:"scheduledEventId"`
	StartedEventId   string `json:"startedEventId"`
}
type ActivityTaskScheduledEventAttributes struct {
	ActivityId                   string       `json:"activityId"`
	ActivityType                 ActivityType `json:"activityType"`
	Control                      string       `json:"control"`
	DecisionTaskCompletedEventId string       `json:"decisionTaskCompletedEventId"`
	HeartbeatTimeout             string       `json:"heartbeatTimeout"`
	Input                        string       `json:"input"`
	ScheduleToCloseTimeout       string       `json:"scheduleToCloseTimeout"`
	ScheduleToStartTimeout       string       `json:"scheduleToStartTimeout"`
	StartToCloseTimeout          string       `json:"startToCloseTimeout"`
	TaskList                     TaskList     `json:"taskList"`
}
type ActivityTaskStartedEventAttributes struct {
	Identity         string `json:"identity"`
	ScheduledEventId string `json:"scheduledEventId"`
}
type ActivityTaskTimedOutEventAttributes struct {
	Details          string `json:"details"`
	ScheduledEventId string `json:"scheduledEventId"`
	StartedEventId   string `json:"startedEventId"`
	TimeoutType      string `json:"timeoutType"`
}
type CancelTimerFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	TimerId                      string `json:"timerId"`
}
type CancelWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type ChildWorkflowExecutionCanceledEventAttributes struct {
	Details           string            `json:"details"`
	InitiatedEventId  string            `json:"initiatedEventId"`
	StartedEventId    string            `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}
type ChildWorkflowExecutionCompletedEventAttributes struct {
	InitiatedEventId  string            `json:"initiatedEventId"`
	Result            string            `json:"result"`
	StartedEventId    string            `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}
type ChildWorkflowExecutionFailedEventAttributes struct {
	Details           string            `json:"details"`
	InitiatedEventId  string            `json:"initiatedEventId"`
	Reason            string            `json:"reason"`
	StartedEventId    string            `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}
type ChildWorkflowExecutionStartedEventAttributes struct {
	InitiatedEventId  string            `json:"initiatedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}
type ChildWorkflowExecutionTerminatedEventAttributes struct {
	InitiatedEventId  string            `json:"initiatedEventId"`
	StartedEventId    string            `json:"startedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}
type ChildWorkflowExecutionTimedOutEventAttributes struct {
	InitiatedEventId  string            `json:"initiatedEventId"`
	StartedEventId    string            `json:"startedEventId"`
	TimeoutType       string            `json:"timeoutType"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
	WorkflowType      WorkflowType      `json:"workflowType"`
}
type CompleteWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type ContinueAsNewWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type DecisionTaskCompletedEventAttributes struct {
	ExecutionContext string `json:"executionContext"`
	ScheduledEventId string `json:"scheduledEventId"`
	StartedEventId   string `json:"startedEventId"`
}
type DecisionTaskScheduledEventAttributes struct {
	StartToCloseTimeout string   `json:"startToCloseTimeout"`
	TaskList            TaskList `json:"taskList"`
}
type DecisionTaskStartedEventAttributes struct {
	Identity         string `json:"identity"`
	ScheduledEventId string `json:"scheduledEventId"`
}
type DecisionTaskTimedOutEventAttributes struct {
	ScheduledEventId string `json:"scheduledEventId"`
	StartedEventId   string `json:"startedEventId"`
	TimeoutType      string `json:"timeoutType"`
}
type ExternalWorkflowExecutionCancelRequestedEventAttributes struct {
	InitiatedEventId  string            `json:"initiatedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
}
type ExternalWorkflowExecutionSignaledEventAttributes struct {
	InitiatedEventId  string            `json:"initiatedEventId"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
}
type FailWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type MarkerRecordedEventAttributes struct {
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details"`
	MarkerName                   string `json:"markerName"`
}
type RecordMarkerFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	MarkerName                   string `json:"markerName"`
}
type RequestCancelActivityTaskFailedEventAttributes struct {
	ActivityId                   string `json:"activityId"`
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type RequestCancelExternalWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	Control                      string `json:"control"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	InitiatedEventId             string `json:"initiatedEventId"`
	RunId                        string `json:"runId"`
	WorkflowId                   string `json:"workflowId"`
}
type RequestCancelExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control                      string `json:"control"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	RunId                        string `json:"runId"`
	WorkflowId                   string `json:"workflowId"`
}
type ScheduleActivityTaskFailedEventAttributes struct {
	ActivityId                   string `json:"activityId"`
	ActivityType                 ActivityType
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
}
type SignalExternalWorkflowExecutionFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	Control                      string `json:"control"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	InitiatedEventId             string `json:"initiatedEventId"`
	RunId                        string `json:"runId"`
	WorkflowId                   string `json:"workflowId"`
}
type SignalExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control                      string `json:"control"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	Input                        string `json:"input"`
	RunId                        string `json:"runId"`
	SignalName                   string `json:"signalName"`
	WorkflowId                   string `json:"workflowId"`
}
type StartChildWorkflowExecutionFailedEventAttributes struct {
	Cause                        string       `json:"cause"`
	Control                      string       `json:"control"`
	DecisionTaskCompletedEventId string       `json:"decisionTaskCompletedEventId"`
	InitiatedEventId             string       `json:"initiatedEventId"`
	WorkflowId                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}
type StartChildWorkflowExecutionInitiatedEventAttributes struct {
	ChildPolicy                  string       `json:"childPolicy"`
	Control                      string       `json:"control"`
	DecisionTaskCompletedEventId string       `json:"decisionTaskCompletedEventId"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout"`
	Input                        string       `json:"input"`
	TagList                      []string     `json:"tagList"`
	TaskList                     TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout"`
	WorkflowId                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}
type StartTimerFailedEventAttributes struct {
	Cause                        string `json:"cause"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	TimerId                      string `json:"timerId"`
}
type TimerCanceledEventAttributes struct {
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	StartedEventId               string `json:"startedEventId"`
	TimerId                      string `json:"timerId"`
}
type TimerFiredEventAttributes struct {
	StartedEventId string `json:"startedEventId"`
	TimerId        string `json:"timerId"`
}
type TimerStartedEventAttributes struct {
	Control                      string `json:"control"`
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	StartToFireTimeout           string `json:"startToFireTimeout"`
	TimerId                      string `json:"timerId"`
}
type WorkflowExecutionCancelRequestedEventAttributes struct {
	Cause                     string            `json:"cause"`
	ExternalInitiatedEventId  string            `json:"externalInitiatedEventId"`
	ExternalWorkflowExecution WorkflowExecution `json:"externalWorkflowExecution"`
}
type WorkflowExecutionCanceledEventAttributes struct {
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details"`
}
type WorkflowExecutionCompletedEventAttributes struct {
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	Result                       string `json:"result"`
}
type WorkflowExecutionContinuedAsNewEventAttributes struct {
	ChildPolicy                  string       `json:"childPolicy"`
	DecisionTaskCompletedEventId string       `json:"decisionTaskCompletedEventId"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout"`
	Input                        string       `json:"input"`
	NewExecutionRunId            string       `json:"newExecutionRunId"`
	TagList                      []string     `json:"tagList"`
	TaskList                     TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}
type WorkflowExecutionFailedEventAttributes struct {
	DecisionTaskCompletedEventId string `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details"`
	Reason                       string `json:"reason"`
}
type WorkflowExecutionSignaledEventAttributes struct {
	ExternalInitiatedEventId  string            `json:"externalInitiatedEventId"`
	ExternalWorkflowExecution WorkflowExecution `json:"externalWorkflowExecution"`
	Input                     string            `json:"input"`
	SignalName                string            `json:"signalName"`
}
type WorkflowExecutionStartedEventAttributes struct {
	ChildPolicy                  string            `json:"childPolicy"`
	ContinuedExecutionRunId      string            `json:"continuedExecutionRunId"`
	ExecutionStartToCloseTimeout string            `json:"executionStartToCloseTimeout"`
	Input                        string            `json:"input"`
	ParentInitiatedEventId       string            `json:"parentInitiatedEventId"`
	ParentWorkflowExecution      WorkflowExecution `json:"parentWorkflowExecution"`
	TagList                      []string          `json:"tagList"`
	TaskList                     TaskList          `json:"taskList"`
	TaskStartToCloseTimeout      string            `json:"taskStartToCloseTimeout"`
	WorkflowType                 WorkflowType      `json:"workflowType"`
}
type WorkflowExecutionTerminatedEventAttributes struct {
	Cause       string `json:"cause"`
	ChildPolicy string `json:"childPolicy"`
	Details     string `json:"details"`
	Reason      string `json:"reason"`
}
type WorkflowExecutionTimedOutEventAttributes struct {
	ChildPolicy string `json:"childPolicy"`
	TimeoutType string `json:"timeoutType"`
}

type PollForDecisionTaskResponse struct {
	Events                 []HistoryEvent    `json:"events"`
	NextPageToken          string            `json:"nextPageToken"`
	PreviousStartedEventId string            `json:"previousStartedEventId"`
	StartedEventId         string            `json:"startedEventId"`
	TaskToken              string            `json:"taskToken"`
	WorkflowExecution      WorkflowExecution `json:"workflowExecution"`
	WorkflowType           WorkflowType      `json:"workflowType"`
}

type Decision struct {
	CancelTimerDecisionAttributes                            *CancelTimerDecisionAttributes                            `json:"cancelTimerDecisionAttributes"`
	CancelWorkflowExecutionDecisionAttributes                *CancelWorkflowExecutionDecisionAttributes                `json:"cancelWorkflowExecutionDecisionAttributes"`
	CompleteWorkflowExecutionDecisionAttributes              *CompleteWorkflowExecutionDecisionAttributes              `json:"completeWorkflowExecutionDecisionAttributes"`
	ContinueAsNewWorkflowExecutionDecisionAttributes         *ContinueAsNewWorkflowExecutionDecisionAttributes         `json:"continueAsNewWorkflowExecutionDecisionAttributes"`
	DecisionType                                             string                                                    `json:"decisionType"`
	FailWorkflowExecutionDecisionAttributes                  *FailWorkflowExecutionDecisionAttributes                  `json:"failWorkflowExecutionDecisionAttributes"`
	RecordMarkerDecisionAttributes                           *RecordMarkerDecisionAttributes                           `json:"recordMarkerDecisionAttributes"`
	RequestCancelActivityTaskDecisionAttributes              *RequestCancelActivityTaskDecisionAttributes              `json:"requestCancelActivityTaskDecisionAttributes"`
	RequestCancelExternalWorkflowExecutionDecisionAttributes *RequestCancelExternalWorkflowExecutionDecisionAttributes `json:"requestCancelExternalWorkflowExecutionDecisionAttributes"`
	ScheduleActivityTaskDecisionAttributes                   *ScheduleActivityTaskDecisionAttributes                   `json:"scheduleActivityTaskDecisionAttributes"`
	SignalExternalWorkflowExecutionDecisionAttributes        *SignalExternalWorkflowExecutionDecisionAttributes        `json:"signalExternalWorkflowExecutionDecisionAttributes"`
	StartChildWorkflowExecutionDecisionAttributes            *StartChildWorkflowExecutionDecisionAttributes            `json:"startChildWorkflowExecutionDecisionAttributes"`
	StartTimerDecisionAttributes                             *StartTimerDecisionAttributes                             `json:"startTimerDecisionAttributes"`
}

type CancelTimerDecisionAttributes struct {
	TimerId string `json:"timerId"`
}
type CancelWorkflowExecutionDecisionAttributes struct {
	Details string `json:"details"`
}
type CompleteWorkflowExecutionDecisionAttributes struct {
	Result string `json:"result"`
}
type ContinueAsNewWorkflowExecutionDecisionAttributes struct {
	ChildPolicy                  string   `json:"childPolicy"`
	ExecutionStartToCloseTimeout string   `json:"executionStartToCloseTimeout"`
	Input                        string   `json:"input"`
	TagList                      []string `json:"tagList"`
	TaskList                     TaskList `json:"taskList"`
	TaskStartToCloseTimeout      string   `json:"taskStartToCloseTimeout"`
	WorkflowTypeVersion          string   `json:"workflowTypeVersion"`
}
type FailWorkflowExecutionDecisionAttributes struct {
	Details string `json:"details"`
	Reason  string `json:"reason"`
}
type RecordMarkerDecisionAttributes struct {
	Details    string `json:"details"`
	MarkerName string `json:"markerName"`
}
type RequestCancelActivityTaskDecisionAttributes struct {
	ActivityId string `json:"activityId"`
}
type RequestCancelExternalWorkflowExecutionDecisionAttributes struct {
	Control    string `json:"control"`
	RunId      string `json:"runId"`
	WorkflowId string `json:"workflowId"`
}
type ScheduleActivityTaskDecisionAttributes struct {
	ActivityId             string       `json:"activityId"`
	ActivityType           ActivityType `json:"activityType"`
	Control                string       `json:"control"`
	HeartbeatTimeout       string       `json:"heartbeatTimeout"`
	Input                  string       `json:"input"`
	ScheduleToCloseTimeout string       `json:"scheduleToCloseTimeout"`
	ScheduleToStartTimeout string       `json:"scheduleToStartTimeout"`
	StartToCloseTimeout    string       `json:"startToCloseTimeout"`
	TaskList               TaskList     `json:"taskList"`
}
type SignalExternalWorkflowExecutionDecisionAttributes struct {
	Control    string `json:"control"`
	Input      string `json:"input"`
	RunId      string `json:"runId"`
	SignalName string `json:"signalName"`
	WorkflowId string `json:"workflowId"`
}
type StartChildWorkflowExecutionDecisionAttributes struct {
	ChildPolicy                  string       `json:"childPolicy"`
	Control                      string       `json:"control"`
	ExecutionStartToCloseTimeout string       `json:"executionStartToCloseTimeout"`
	Input                        string       `json:"input"`
	TagList                      []string     `json:"tagList"`
	TaskList                     TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string       `json:"taskStartToCloseTimeout"`
	WorkflowId                   string       `json:"workflowId"`
	WorkflowType                 WorkflowType `json:"workflowType"`
}
type StartTimerDecisionAttributes struct {
	Control            string `json:"control"`
	StartToFireTimeout string `json:"startToFireTimeout"`
	TimerId            string `json:"timerId"`
}

type RespondDecisionTaskCompletedRequest struct {
	Decisions        []Decision `json:"decisions"`
	ExecutionContext string     `json:"executionContext"`
	TaskToken        string     `json:"taskToken"`
}

/*ActivityWorkerProtocol*/

type PollForActivityTaskRequest struct {
	Domain   string   `json:"domain"`
	Identity string   `json:"identity"`
	TaskList TaskList `json:"taskList"`
}

type PollForActivityTaskResponse struct {
	ActivityId        string            `json:"activityId"`
	ActivityType      ActivityType      `json:"activityType"`
	Input             string            `json:"input"`
	StartedEventId    string            `json:"startedEventId"`
	TaskToken         string            `json:"taskToken"`
	WorkflowExecution WorkflowExecution `json:"workflowExecution"`
}

type RespondActivityTaskCompletedRequest struct {
	Result    string `json:"result"`
	TaskToken string `json:"taskToken"`
}

type RespondActivityTaskFailedRequest struct {
	Details   string `json:"details"`
	Reason    string `json:"reason"`
	TaskToken string `json:"taskToken"`
}

type RespondActivityTaskCanceledRequest struct {
	Details   string `json:"details"`
	TaskToken string `json:"taskToken"`
}

type RecordActivityTaskHeartbeatRequest struct {
	Details   string `json:"details"`
	TaskToken string `json:"taskToken"`
}

type RecordActivityTaskHeartbeatResponse struct {
	CancelRequested string `json:"cancelRequested"`
}

/*common types*/

type TaskList struct {
	Name string `json:"name"`
}
type WorkflowType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type WorkflowExecution struct {
	RunId      string `json:"runId"`
	WorkflowId string `json:"workflowId"`
}

type ActivityType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type TypeInfo struct {
	CreationDate    float32 `json:"creationDate"`
	DeprecationDate float32 `json:"deprecationDate"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
	WorkflowType    WorkflowType
}
