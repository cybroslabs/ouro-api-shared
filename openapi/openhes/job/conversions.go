package job

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/cybroslabs/hes-2-apis/openapi/openhes/attribute"
	driverdata "github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdataproxy"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmaster"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrInvalidJobStatus     = errors.New("invalid job status")
	ErrInvalidActionType    = errors.New("invalid action type")
	ErrUnknownJobActionType = errors.New("unknown job action type")
)

const (
	DefaultPriority    = pbdriver.JobPriority_PRIORITY_0
	DefaultMaxDuration = int64(5 * 60 * 1000)
	DefaultRetryDelay  = int64(60 * 1000)
	DefaultAttempts    = int32(1)
	DefaultDeferStart  = uint64(0)
)

// Converts the job action list - Rest API to gRPC
func R2GJobActions(actions *JobActionListSchema) ([]*pbdriver.JobAction, error) {
	if actions == nil {
		return nil, nil
	}

	result := make([]*pbdriver.JobAction, len(*actions))
	for i := range *actions {
		tmp, err := R2GJobAction(&(*actions)[i])
		if err != nil {
			return nil, err
		}
		result[i] = tmp
	}
	return result, nil
}

// Converts the job action list - gRPC to Rest API
func G2RJobActions(actions []*pbdriver.JobAction) (*JobActionListSchema, error) {
	if actions == nil {
		return nil, nil
	}

	result := make(JobActionListSchema, len(actions))
	for i := range actions {
		err := G2RJobAction(actions[i], &result[i])
		if err != nil {
			return nil, err
		}
	}
	return &result, nil
}

// Converts the job action - Rest API to gRPC
func R2GJobAction(action *JobActionSchema) (*pbdriver.JobAction, error) {
	if tmp, err := asJobActionGetRegisterSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetRegister{
				GetRegister: &pbdriver.ActionGetRegister{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetPeriodicalProfileSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		var param_tfrom *timestamppb.Timestamp
		if tmp.ParamTfrom != nil {
			param_tfrom = timestamppb.New(*tmp.ParamTfrom)
		}
		var param_tto *timestamppb.Timestamp
		if tmp.ParamTfrom != nil {
			param_tto = timestamppb.New(*tmp.ParamTfrom)
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetPeriodicalProfile{
				GetPeriodicalProfile: &pbdriver.ActionGetPeriodicalProfile{
					From: param_tfrom,
					To:   param_tto,
				},
			},
		}, nil
	}
	if tmp, err := asJobActionGetIrregularProfileSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetIrregularProfile{
				GetIrregularProfile: &pbdriver.ActionGetIrregularProfile{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetEventsSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetEvents{
				GetEvents: &pbdriver.ActionGetEvents{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetClockSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetClock{
				GetClock: &pbdriver.ActionGetClock{},
			},
		}, nil
	}
	if tmp, err := asJobActionSyncClockSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_SyncClock{
				SyncClock: &pbdriver.ActionSyncClock{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetRelayStateSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetRelayState{
				GetRelayState: &pbdriver.ActionGetRelayState{},
			},
		}, nil
	}
	if tmp, err := asJobActionSetRelayStateSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_SetRelayState{
				SetRelayState: &pbdriver.ActionSetRelayState{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetDisconnectorStateSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetDisconnectorState{
				GetDisconnectorState: &pbdriver.ActionGetDisconnectorState{},
			},
		}, nil
	}
	if tmp, err := asJobActionSetDisconnectorStateSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_SetDisconnectorState{
				SetDisconnectorState: &pbdriver.ActionSetDisconnectorState{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetTouSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetTou{
				GetTou: &pbdriver.ActionGetTou{},
			},
		}, nil
	}
	if tmp, err := asJobActionSetTouSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_SetTou{
				SetTou: &pbdriver.ActionSetTou{},
			},
		}, nil
	}
	if tmp, err := asJobActionGetLimiterSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_GetLimiter{
				GetLimiter: &pbdriver.ActionGetLimiter{},
			},
		}, nil
	}
	if tmp, err := asJobActionSetLimiterSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_SetLimiter{
				SetLimiter: &pbdriver.ActionSetLimiter{},
			},
		}, nil
	}
	if tmp, err := asJobActionResetBillingPeriodSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_ResetBillingPeriod{
				ResetBillingPeriod: &pbdriver.ActionResetBillingPeriod{},
			},
		}, nil
	}
	if tmp, err := asJobActionFwUpdateSchema(action); err == nil {
		attr, err := attribute.R2GAttributes(tmp.Attributes)
		if err != nil {
			return nil, err
		}
		return &pbdriver.JobAction{
			ActionId:   tmp.Id.String(),
			Attributes: attr,
			Action: &pbdriver.JobAction_FwUpdate{
				FwUpdate: &pbdriver.ActionFwUpdate{},
			},
		}, nil
	}
	return nil, ErrUnknownJobActionType
}

func G2RJobAction(action *pbdriver.JobAction, result *JobActionSchema) error {
	if action == nil {
		return ErrUnknownJobActionType
	}

	action_id, _ := uuid.Parse(action.ActionId)

	if tmp := action.GetGetRegister(); tmp != nil {
		err := result.FromJobActionGetRegisterSchema(JobActionGetRegisterSchema{
			Id:         action_id,
			Type:       string(GETREGISTER),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetPeriodicalProfile(); tmp != nil {
		var t_from *time.Time
		if tmp.From != nil {
			t := tmp.From.AsTime()
			t_from = &t
		}
		var t_to *time.Time
		if tmp.To != nil {
			t := tmp.To.AsTime()
			t_to = &t
		}
		err := result.FromJobActionGetPeriodicalProfileSchema(JobActionGetPeriodicalProfileSchema{
			Id:         action_id,
			Type:       string(GETPERIODICALPROFILE),
			ParamTfrom: t_from,
			ParamTto:   t_to,
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetIrregularProfile(); tmp != nil {
		err := result.FromJobActionGetIrregularProfileSchema(JobActionGetIrregularProfileSchema{
			Id:         action_id,
			Type:       string(GETIRREGULARPROFILE),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetEvents(); tmp != nil {
		err := result.FromJobActionGetEventsSchema(JobActionGetEventsSchema{
			Id:         action_id,
			Type:       string(GETEVENTS),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetClock(); tmp != nil {
		err := result.FromJobActionGetClockSchema(JobActionGetClockSchema{
			Id:         action_id,
			Type:       string(GETCLOCK),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSyncClock(); tmp != nil {
		err := result.FromJobActionSyncClockSchema(JobActionSyncClockSchema{
			Id:         action_id,
			Type:       string(SYNCCLOCK),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetRelayState(); tmp != nil {
		err := result.FromJobActionGetRelayStateSchema(JobActionGetRelayStateSchema{
			Id:         action_id,
			Type:       string(GETRELAYSTATE),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetRelayState(); tmp != nil {
		err := result.FromJobActionSetRelayStateSchema(JobActionSetRelayStateSchema{
			Id:         action_id,
			Type:       string(SETRELAYSTATE),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetDisconnectorState(); tmp != nil {
		err := result.FromJobActionGetDisconnectorStateSchema(JobActionGetDisconnectorStateSchema{
			Id:         action_id,
			Type:       string(GETDISCONNECTORSTATE),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetDisconnectorState(); tmp != nil {
		err := result.FromJobActionSetDisconnectorStateSchema(JobActionSetDisconnectorStateSchema{
			Id:         action_id,
			Type:       string(SETDISCONNECTORSTATE),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetTou(); tmp != nil {
		err := result.FromJobActionGetTouSchema(JobActionGetTouSchema{
			Id:         action_id,
			Type:       string(GETTOU),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetTou(); tmp != nil {
		err := result.FromJobActionSetTouSchema(JobActionSetTouSchema{
			Id:         action_id,
			Type:       string(SETTOU),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetLimiter(); tmp != nil {
		err := result.FromJobActionGetLimiterSchema(JobActionGetLimiterSchema{
			Id:         action_id,
			Type:       string(GETLIMITER),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetLimiter(); tmp != nil {
		err := result.FromJobActionSetLimiterSchema(JobActionSetLimiterSchema{
			Id:         action_id,
			Type:       string(SETLIMITER),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetResetBillingPeriod(); tmp != nil {
		err := result.FromJobActionResetBillingPeriodSchema(JobActionResetBillingPeriodSchema{
			Id:         action_id,
			Type:       string(RESETBILLINGPERIOD),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetFwUpdate(); tmp != nil {
		err := result.FromJobActionFwUpdateSchema(JobActionFwUpdateSchema{
			Id:         action_id,
			Type:       string(FWUPDATE),
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	}

	return ErrUnknownJobActionType
}

// Converts the job status code - gRPC to Rest API
func G2RJobStatusCode(status pbtaskmaster.JobStatusCode) (JobStatusCodeEnumSchema, error) {
	switch status {
	case pbtaskmaster.JobStatusCode_JOB_STATUS_QUEUED:
		return JobStatusCodeEnumSchemaQUEUED, nil
	case pbtaskmaster.JobStatusCode_JOB_STATUS_RUNNING:
		return JobStatusCodeEnumSchemaRUNNING, nil
	case pbtaskmaster.JobStatusCode_JOB_STATUS_COMPLETED:
		return JobStatusCodeEnumSchemaCOMPLETED, nil
	case pbtaskmaster.JobStatusCode_JOB_STATUS_CANCELLED:
		return JobStatusCodeEnumSchemaCANCELLED, nil
	case pbtaskmaster.JobStatusCode_JOB_STATUS_EXPIRED:
		return JobStatusCodeEnumSchemaEXPIRED, nil
	default:
		return JobStatusCodeEnumSchemaQUEUED, ErrInvalidJobStatus
	}
}

// Converts the job error code - gRPC to Rest API
func G2RJobErrorCode(code pbdriver.JobErrorCode) (JobErrorCodeSchema, error) {
	switch code {
	case pbdriver.JobErrorCode_JOB_ERROR_CODE_NONE:
		return JobErrorCodeSchemaNONE, nil
	case pbdriver.JobErrorCode_JOB_ERROR_CODE_ERROR:
		return JobErrorCodeSchemaERROR, nil
	case pbdriver.JobErrorCode_JOB_ERROR_CODE_FATAL:
		return JobErrorCodeSchemaFATAL, nil
	default:
		return JobErrorCodeSchemaFATAL, ErrInvalidJobStatus
	}
}

// Converts the action result code - gRPC to Rest API
func G2RActionResultCode(status pbdriver.ActionResultCode) (JobActionResultSchemaCode, error) {
	switch status {
	case pbdriver.ActionResultCode_ERROR_CODE_ACTION_OK:
		return JobActionResultSchemaCodeOK, nil
	case pbdriver.ActionResultCode_ERROR_CODE_ACTION_ERROR:
		return JobActionResultSchemaCodeERROR, nil
	case pbdriver.ActionResultCode_ERROR_CODE_ACTION_UNSUPPORTED:
		return JobActionResultSchemaCodeUNSUPPORTED, nil
	default:
		return JobActionResultSchemaCodeERROR, ErrInvalidJobStatus
	}
}

// Converts the action type - Rest API to gRPC
func R2GActionType(actionType ActionTypeSchema) (pbdriver.ActionType, error) {
	action_name := "ACTION_TYPE_" + string(actionType)
	no, ok := pbdriver.ActionType_value[action_name]
	if !ok {
		return -1, ErrInvalidActionType
	}
	return pbdriver.ActionType(no), nil
}

// Converts the action type - gRPC to Rest API
func G2RActionType(actionType pbdriver.ActionType) (ActionTypeSchema, error) {
	no := int32(actionType.Number())
	action_name, ok := pbdriver.ActionType_name[no]
	if !ok {
		return "", ErrInvalidActionType
	}
	result, ok := strings.CutPrefix(action_name, "ACTION_TYPE_")
	if !ok {
		return "", ErrInvalidActionType
	}
	return ActionTypeSchema(result), nil
}

// Converts the job settings - gRPC to Rest API
func G2RJobSettings(settings *pbdriver.JobSettings) (*JobSettingsSchema, error) {
	intPriority := int32(settings.Priority)

	// gRPC is in milliseconds, REST is in seconds
	max_duration := settings.MaxDuration / 1000
	retry_delay := settings.RetryDelay / 1000
	defer_start := int64(settings.DeferStart / 1000)

	var expires_at *time.Time = nil
	if ts := settings.ExpiresAt; ts != nil {
		t := ts.AsTime()
		expires_at = &t
	}

	result := &JobSettingsSchema{
		Attempts:    &settings.Attempts,
		MaxDuration: &max_duration,
		Priority:    &intPriority,
		RetryDelay:  &retry_delay,
		DeferStart:  &defer_start,
		ExpiresAt:   expires_at,
	}

	return result, nil
}

// Converts the job settings - Rest API to gRPC
func R2GJobSettings(settings *JobSettingsSchema) (*pbdriver.JobSettings, error) {
	job_priority := DefaultPriority
	if pr := settings.Priority; pr != nil {
		if *pr < 0 || *pr > 9 {
			return nil, fmt.Errorf("error while converting priority %v, value out of range", *pr)
		}
		job_priority = (pbdriver.JobPriority)(*pr)
	}

	max_duration := DefaultMaxDuration
	if pr := settings.MaxDuration; pr != nil {
		// REST is in seconds, gRPC is in milliseconds
		max_duration = *pr * 1000
	}

	retry_delay := DefaultRetryDelay
	if pr := settings.RetryDelay; pr != nil {
		// REST is in seconds, gRPC is in milliseconds
		retry_delay = *pr * 1000
	}

	attempts := DefaultAttempts
	if pr := settings.Attempts; pr != nil {
		attempts = *pr
	}

	defer_start := DefaultDeferStart
	if pr := settings.DeferStart; pr != nil {
		defer_start = uint64(*pr * 1000)
	}

	var expires_at *timestamppb.Timestamp = nil
	if ts := settings.ExpiresAt; ts != nil {
		expires_at = timestamppb.New(*ts)
	}

	return &pbdriver.JobSettings{
		Attempts:    attempts,
		MaxDuration: max_duration,
		Priority:    job_priority,
		RetryDelay:  retry_delay,
		DeferStart:  defer_start,
		ExpiresAt:   expires_at,
	}, nil
}

// Converts the bulk spec - gRPC to Rest API
func G2RBulkSpec(spec *pbdataproxy.BulkSpec) (*BulkSpecSchema, error) {
	actions, err := G2RJobActions(spec.JobActions)
	if err != nil {
		return nil, err
	}

	devices := make(JobDeviceListSchema, len(spec.Devices))
	for i, device := range spec.Devices {
		target := &devices[i]
		target.Id, err = uuid.Parse(device.Id)
		if err != nil {
			return nil, err
		}
		target.ExternalID = &device.ExternalId
		target.CommunicationUnitAttributes = attribute.G2RAttributes(device.ConnectionInfo.CommunicationUnitAttributes)
		target.DeviceAttributes = attribute.G2RAttributes(device.ConnectionInfo.DeviceAttributes)
		target.Endpoint = device.ConnectionInfo.Hostname
	}

	id, err := uuid.Parse(spec.BulkId)
	if err != nil {
		return nil, err
	}

	var corr_id *string
	if spec.CorrelationId != "" {
		corr_id = &spec.CorrelationId
	}

	settings, err := G2RJobSettings(spec.Settings)
	if err != nil {
		return nil, err
	}

	result := &BulkSpecSchema{
		Id:               id,
		CorrelationID:    corr_id,
		DeviceDriverType: spec.DeviceDriverType,
		Settings:         settings,
		Devices:          devices,
		Actions:          *actions,
		WebhookURL:       spec.WebhookUrl,
	}

	return result, nil
}

// Converts the bulk spec - Rest API to gRPC
func R2GBulkSpec(spec *BulkSpecSchema) (*pbdataproxy.BulkSpec, error) {
	webhook_url := spec.WebhookURL
	if webhook_url != nil {
		if *webhook_url == "" {
			webhook_url = nil
		} else {
			uri, err := url.ParseRequestURI(*webhook_url)
			if err != nil {
				return nil, fmt.Errorf("error while parsing webhook url: %v", err)
			}
			if uri.Scheme != "http" && uri.Scheme != "https" {
				return nil, fmt.Errorf("invalid WebhookURL scheme, given %s, only http or https is accepted", uri.Scheme)
			}
		}
	}

	actions, err := R2GJobActions(&spec.Actions)
	if err != nil {
		return nil, err
	}

	devices := make([]*pbtaskmaster.JobDevice, len(spec.Devices))
	for i, device := range spec.Devices {
		communication_unit_attributes, err := attribute.R2GAttributes(device.CommunicationUnitAttributes)
		if err != nil {
			return nil, err
		}

		device_attributes, err := attribute.R2GAttributes(device.DeviceAttributes)
		if err != nil {
			return nil, err
		}

		external_id := ""
		if device.ExternalID != nil {
			external_id = *device.ExternalID
		}

		devices[i] = &pbtaskmaster.JobDevice{
			Id: device.Id.String(),
			ConnectionInfo: &pbdriver.ConnectionInfo{
				Hostname:                    device.Endpoint,
				DeviceAttributes:            device_attributes,
				CommunicationUnitAttributes: communication_unit_attributes,
			},
			ExternalId: external_id,
		}
	}

	bulk_id := spec.Id.String()

	var corr_id string
	if spec.CorrelationID != nil {
		corr_id = *spec.CorrelationID
	}

	settings, err := R2GJobSettings(spec.Settings)
	if err != nil {
		return nil, err
	}

	return &pbdataproxy.BulkSpec{
		BulkId:           bulk_id,
		CorrelationId:    corr_id,
		DeviceDriverType: spec.DeviceDriverType,
		Settings:         settings,
		Devices:          devices,
		JobActions:       actions,
		WebhookUrl:       webhook_url,
	}, nil
}

// Converts the bulk status code - gRPC to Rest API
func G2RBulkStatusCode(status pbdataproxy.BulkStatusCode) (BulkStatusEnumSchema, error) {
	switch status {
	case pbdataproxy.BulkStatusCode_BULK_STATUS_QUEUED:
		return BulkStatusEnumSchemaQUEUED, nil
	case pbdataproxy.BulkStatusCode_BULK_STATUS_RUNNING:
		return BulkStatusEnumSchemaRUNNING, nil
	case pbdataproxy.BulkStatusCode_BULK_STATUS_COMPLETED:
		return BulkStatusEnumSchemaCOMPLETED, nil
	case pbdataproxy.BulkStatusCode_BULK_STATUS_CANCELLED:
		return BulkStatusEnumSchemaCANCELLED, nil
	case pbdataproxy.BulkStatusCode_BULK_STATUS_EXPIRED:
		return BulkStatusEnumSchemaEXPIRED, nil
	default:
		return BulkStatusEnumSchemaQUEUED, ErrInvalidJobStatus
	}
}

// Converts the bulk status - gRPC to Rest API
func G2RBulkStatus(status *pbdataproxy.BulkStatus) (*BulkStatusSchema, error) {
	status_code, err := G2RBulkStatusCode(status.Status)
	if err != nil {
		return nil, err
	}

	result := &BulkStatusSchema{
		StartedAt:  nil,
		FinishedAt: nil,
		Status:     status_code,
	}
	return result, nil
}

// Converts the job status - gRPC to Rest API
func G2RJobStatus(status *pbtaskmaster.JobStatus) (*JobStatusSchema, error) {
	status_code, err := G2RJobStatusCode(status.Status)
	if err != nil {
		return nil, err
	}

	error_code, err := G2RJobErrorCode(*status.Code)
	if err != nil {
		return nil, err
	}

	var started_at *time.Time = nil
	if ts := status.StartedAt; ts != nil {
		t := ts.AsTime()
		started_at = &t
	}

	var finished_at *time.Time = nil
	if ts := status.FinishedAt; ts != nil {
		t := ts.AsTime()
		finished_at = &t
	}

	var results_ptr *[]JobActionResultSchema = nil
	if len(status.Results) > 0 {
		results := make([]JobActionResultSchema, len(status.Results))
		for i, result := range status.Results {
			results[i].Id, err = uuid.Parse(result.ActionId)
			if err != nil {
				return nil, err
			}

			results[i].Code, err = G2RActionResultCode(status.Results[i].Status)
			if err != nil {
				return nil, err
			}

			ar := &results[i]
			if rd := result.Data; rd != nil {
				// Handle BV & LP data
				switch d := rd.Data.(type) {
				case *pbdriver.ActionData_Billings:
					if dbv := d.Billings; dbv != nil {
						v := dbv.Values
						tmp := make(driverdata.DeviceRegistersDataSchema, len(v))
						for i, v := range v {
							if ts := v.Timestamp; ts != nil {
								ts_time := ts.AsTime()
								tmp[i].Timestamp = &ts_time
							}
							tmp[i].Unit = &v.Unit

							tmp_valueinfo := &driverdata.GenericValueStatusSchema{
								Exponent: &v.Value.Exponent,
								Status:   v.Value.Status,
							}
							tmp_value := &driverdata.GenericValueStatusSchema_Value{}
							switch vt := v.Value.Value.(type) {
							case *pbdriver.MeasuredValue_IntValue:
								err = tmp_value.FromGenericValueStatusSchemaValue1(vt.IntValue)
								if err != nil {
									return nil, err
								}
								tmp_valueinfo.Value = tmp_value
								tmp[i].Value = tmp_valueinfo
							case *pbdriver.MeasuredValue_DoubleValue:
								err = tmp_value.FromGenericValueStatusSchemaValue2(vt.DoubleValue)
								if err != nil {
									return nil, err
								}
								tmp_valueinfo.Value = tmp_value
								tmp[i].Value = tmp_valueinfo
							default:
								return nil, fmt.Errorf("unknown value type: %v", vt)
							}
						}
						err = ar.Data.FromExternalRef0DeviceRegistersDataSchema(tmp)
						if err != nil {
							return nil, err
						}
					}
				case *pbdriver.ActionData_Profile:
					if dlp := d.Profile; dlp != nil {
						blocks := dlp.Blocks

						tmp_blocks := make([]driverdata.DeviceProfileBlockSchema, len(blocks))
						tmp := driverdata.DeviceProfileDataSchema{
							Unit:   dlp.Unit,
							Period: dlp.Period,
							Blocks: tmp_blocks,
						}
						for i, vb := range blocks {
							if ts := vb.StartTimestamp; ts != nil {
								ts_time := ts.AsTime()
								tmp_blocks[i].Start = &ts_time
							}
							tmp_values := make([]driverdata.GenericValueStatusSchema, len(vb.Values))
							for j, v := range vb.Values {
								tmp_valueinfo := &tmp_values[j]
								tmp_valueinfo.Exponent = &v.Exponent
								tmp_valueinfo.Status = v.Status
								tmp_value := &driverdata.GenericValueStatusSchema_Value{}
								switch vt := v.Value.(type) {
								case *pbdriver.MeasuredValue_IntValue:
									err = tmp_value.FromGenericValueStatusSchemaValue1(vt.IntValue)
									if err != nil {
										return nil, err
									}
									tmp_valueinfo.Value = tmp_value
								case *pbdriver.MeasuredValue_DoubleValue:
									err = tmp_value.FromGenericValueStatusSchemaValue2(vt.DoubleValue)
									if err != nil {
										return nil, err
									}
									tmp_valueinfo.Value = tmp_value
								default:
									return nil, fmt.Errorf("unknown value type: %v", vt)
								}
							}
							tmp_blocks[i].Values = &tmp_values
						}
						err = ar.Data.FromExternalRef0DeviceProfileDataSchema(tmp)
						if err != nil {
							return nil, err
						}
					}
				}
			}
		}
		results_ptr = &results
	}

	result := &JobStatusSchema{
		StartedAt:  started_at,
		FinishedAt: finished_at,
		Results:    results_ptr,
		Status:     status_code,
		Code:       error_code,
	}

	return result, nil
}
