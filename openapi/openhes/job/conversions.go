package job

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/cybroslabs/hes-2-apis/openapi/openhes/attribute"
	odevice "github.com/cybroslabs/hes-2-apis/openapi/openhes/device"
	"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver"
	driverdata "github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdataproxy"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmaster"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrInvalidJobStatus      = errors.New("invalid job status")
	ErrUnknownJobActionType  = errors.New("unknown job action type")
	ErrInvalidDeviceList     = errors.New("invalid device list")
	ErrInvalidConnectionInfo = errors.New("invalid connection info")
)

const (
	// Default job priority (lowest)
	DefaultPriority = pbdriver.JobPriority_JOB_PRIORITY_0
	// Default job duration in milliseconds
	DefaultMaxDuration = int64(5 * 60 * 1000)
	// Default job retry delay in milliseconds
	DefaultRetryDelay = int64(60 * 1000)
	// Default job attempts
	DefaultAttempts = int32(1)
	// Default job defer start in milliseconds
	DefaultDeferStart = uint64(0)
)

var (
	// Default job attempts in a list form
	defaultAttemptsList = []int32{DefaultAttempts}
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

// Converts the job action list - gRPC to Rest API
func G2RJobActionsTo(actions []*pbdriver.JobAction, result *JobActionListSchema) error {
	if actions == nil {
		*result = nil
		return nil
	}

	*result = make(JobActionListSchema, len(actions))
	for i := range actions {
		err := G2RJobAction(actions[i], &(*result)[i])
		if err != nil {
			return err
		}
	}
	return nil
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
		if tmp.ParamTto != nil {
			param_tto = timestamppb.New(*tmp.ParamTto)
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
			ParamTfrom: t_from,
			ParamTto:   t_to,
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetIrregularProfile(); tmp != nil {
		err := result.FromJobActionGetIrregularProfileSchema(JobActionGetIrregularProfileSchema{
			Id:                      action_id,
			TypeGetIrregularProfile: 1,
			Attributes:              attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetEvents(); tmp != nil {
		err := result.FromJobActionGetEventsSchema(JobActionGetEventsSchema{
			Id:            action_id,
			TypeGetEvents: 1,
			Attributes:    attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetClock(); tmp != nil {
		err := result.FromJobActionGetClockSchema(JobActionGetClockSchema{
			Id:           action_id,
			TypeGetClock: 1,
			Attributes:   attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSyncClock(); tmp != nil {
		err := result.FromJobActionSyncClockSchema(JobActionSyncClockSchema{
			Id:            action_id,
			TypeSyncClock: 1,
			Attributes:    attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetRelayState(); tmp != nil {
		err := result.FromJobActionGetRelayStateSchema(JobActionGetRelayStateSchema{
			Id:                action_id,
			TypeGetRelayState: 1,
			Attributes:        attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetRelayState(); tmp != nil {
		err := result.FromJobActionSetRelayStateSchema(JobActionSetRelayStateSchema{
			Id:                action_id,
			TypeSetRelayState: 1,
			Attributes:        attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetDisconnectorState(); tmp != nil {
		err := result.FromJobActionGetDisconnectorStateSchema(JobActionGetDisconnectorStateSchema{
			Id:                       action_id,
			TypeGetDisconnectorState: 1,
			Attributes:               attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetDisconnectorState(); tmp != nil {
		err := result.FromJobActionSetDisconnectorStateSchema(JobActionSetDisconnectorStateSchema{
			Id:                       action_id,
			TypeSetDisconnectorState: 1,
			Attributes:               attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetTou(); tmp != nil {
		err := result.FromJobActionGetTouSchema(JobActionGetTouSchema{
			Id:         action_id,
			TypeGetTou: 1,
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetTou(); tmp != nil {
		err := result.FromJobActionSetTouSchema(JobActionSetTouSchema{
			Id:         action_id,
			TypeSetTou: 1,
			Attributes: attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetLimiter(); tmp != nil {
		err := result.FromJobActionGetLimiterSchema(JobActionGetLimiterSchema{
			Id:             action_id,
			TypeGetLimiter: 1,
			Attributes:     attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetLimiter(); tmp != nil {
		err := result.FromJobActionSetLimiterSchema(JobActionSetLimiterSchema{
			Id:             action_id,
			TypeSetLimiter: 1,
			Attributes:     attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetResetBillingPeriod(); tmp != nil {
		err := result.FromJobActionResetBillingPeriodSchema(JobActionResetBillingPeriodSchema{
			Id:                     action_id,
			TypeResetBillingPeriod: 1,
			Attributes:             attribute.G2RAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetFwUpdate(); tmp != nil {
		err := result.FromJobActionFwUpdateSchema(JobActionFwUpdateSchema{
			Id:           action_id,
			TypeFwUpdate: 1,
			Attributes:   attribute.G2RAttributes(action.Attributes),
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
	result := &pbdriver.JobSettings{}

	if pr := settings.Priority; pr != nil {
		if *pr < 0 || *pr > 9 {
			return nil, fmt.Errorf("error while converting priority %v, value out of range", *pr)
		}
		result.Priority = (pbdriver.JobPriority)(*pr)
	} else {
		result.Priority = DefaultPriority
	}

	if pr := settings.MaxDuration; pr != nil {
		// REST is in seconds, gRPC is in milliseconds
		result.MaxDuration = *pr * 1000
	} else {
		result.MaxDuration = DefaultMaxDuration
	}

	if pr := settings.RetryDelay; pr != nil {
		// REST is in seconds, gRPC is in milliseconds
		result.RetryDelay = *pr * 1000
	} else {
		result.RetryDelay = DefaultRetryDelay
	}

	if pr := settings.DeferStart; pr != nil {
		result.DeferStart = uint64(*pr * 1000)
	} else {
		result.DeferStart = DefaultDeferStart
	}

	if ts := settings.ExpiresAt; ts != nil {
		result.ExpiresAt = timestamppb.New(*ts)
	}

	if pr := settings.Attempts; pr != nil {
		result.Attempts = *pr
	} else {
		result.Attempts = defaultAttemptsList
	}

	return result, nil
}

// Converts the bulk spec - gRPC to Rest API
func G2RBulkSpec(spec *pbdataproxy.BulkSpec) (*BulkSpecSchema, error) {
	var err error

	items := make([]JobCustomDeviceSchema, len(spec.Devices))
	for i, device := range spec.Devices {
		target := &items[i]
		target.Id, err = uuid.Parse(device.Id)
		if err != nil {
			return nil, err
		}
		target.ExternalID = device.ExternalId
		target.DeviceAttributes = attribute.G2RAttributes(device.DeviceAttributes)
		target.Timezone = device.Timezone
		target.ConnectionInfo = make([]odevice.ConnectionInfoSchema, len(device.ConnectionInfo))
		for j, ci := range device.ConnectionInfo {
			if tcp := ci.GetTcpip(); tcp != nil {
				err := target.ConnectionInfo[j].FromConnectionTypeTcpIpSchema(odevice.ConnectionTypeTcpIpSchema{
					Host: tcp.Host,
					Port: tcp.Port,
				})
				if err != nil {
					return nil, err
				}
			} else if modem := ci.GetModemPool(); modem != nil {
				var pool_id uuid.UUID
				pool_id, err = uuid.Parse(modem.PoolId)
				if err != nil {
					return nil, err
				}
				err = target.ConnectionInfo[j].FromConnectionTypePhoneLineSchema(odevice.ConnectionTypePhoneLineSchema{
					Number: modem.Number,
					PoolId: pool_id,
				})
				if err != nil {
					return nil, err
				}
			} else if controller_serial := ci.GetSerialOverIp(); controller_serial != nil {
				if moxa := controller_serial.GetMoxa(); moxa != nil {
					err = target.ConnectionInfo[j].FromConnectionTypeSerialMoxaSchema(odevice.ConnectionTypeSerialMoxaSchema{
						Host:        moxa.Host,
						DataPort:    moxa.DataPort,
						CommandPort: moxa.CommandPort,
					})
					if err != nil {
						return nil, err
					}
				} else if direct := controller_serial.GetDirect(); direct != nil {
					err = target.ConnectionInfo[j].FromConnectionTypeSerialDirectSchema(odevice.ConnectionTypeSerialDirectSchema{
						Host: direct.Host,
						Port: direct.Port,
					})
					if err != nil {
						return nil, err
					}
				} else {
					return nil, ErrInvalidConnectionInfo
				}
			} else {
				return nil, ErrInvalidConnectionInfo
			}
		}
	}

	wrapper := JobCustomDeviceListTypedSchema{
		Items: &items,
	}

	var id uuid.UUID
	id, err = uuid.Parse(spec.BulkId)
	if err != nil {
		return nil, err
	}

	var corr_id *string
	if spec.CorrelationId != "" {
		corr_id = &spec.CorrelationId
	}

	var settings *JobSettingsSchema
	settings, err = G2RJobSettings(spec.Settings)
	if err != nil {
		return nil, err
	}

	result := &BulkSpecSchema{
		Id:            id,
		CorrelationID: corr_id,
		DriverType:    spec.DriverType,
		Settings:      settings,
		WebhookURL:    spec.WebhookUrl,
	}

	err = G2RJobActionsTo(spec.JobActions, &result.Actions)
	if err != nil {
		return nil, err
	}

	err = result.Devices.FromJobCustomDeviceListTypedSchema(wrapper)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func decodeJobCustomDeviceList(raw json.RawMessage) (JobCustomDeviceListSchema, error) {
	var devices JobCustomDeviceListTypedSchema
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	if err := d.Decode(&devices); err != nil {
		return nil, err
	}
	result := devices.Items
	if result == nil {
		return make(JobCustomDeviceListSchema, 0), nil
	}
	return *result, nil
}

func decodeJobDeviceList(raw json.RawMessage) (JobDeviceListSchema, error) {
	var devices JobDeviceListTypedSchema
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	if err := d.Decode(&devices); err != nil {
		return nil, err
	}
	result := devices.Items
	if result == nil {
		return make(JobDeviceListSchema, 0), nil
	}
	return *result, nil
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

	list_type, _ := spec.Devices.Discriminator()

	actions, err := R2GJobActions(&spec.Actions)
	if err != nil {
		return nil, err
	}

	var devices []*pbtaskmaster.JobDevice
	if list_type == string(JOBDEVICESFULL) {
		if device_list, err := decodeJobCustomDeviceList(spec.Devices.Union); err == nil && device_list != nil {
			devices = make([]*pbtaskmaster.JobDevice, len(device_list))
			for i, device := range device_list {
				device_attributes, err := attribute.R2GAttributes(device.DeviceAttributes)
				if err != nil {
					return nil, err
				}

				connection_info_list := make([]*pbdriver.ConnectionInfo, len(device.ConnectionInfo))
				for j, ci := range device.ConnectionInfo {
					link_proto, err := driver.R2GDataLinkProtocol(ci.LinkProtocol)
					if err != nil {
						return nil, err
					}

					connection_info := &pbdriver.ConnectionInfo{
						LinkProtocol: link_proto,
					}
					connection_info_list[j] = connection_info

					if tcp, err := ci.AsConnectionTypeTcpIpSchema(); err == nil {
						connection_info.Connection = &pbdriver.ConnectionInfo_Tcpip{
							Tcpip: &pbdriver.ConnectionTypeDirectTcpIp{
								Host: tcp.Host,
								Port: uint32(tcp.Port),
							},
						}
					} else if phone, err := ci.AsConnectionTypePhoneLineSchema(); err == nil {
						connection_info.Connection = &pbdriver.ConnectionInfo_ModemPool{
							ModemPool: &pbdriver.ConnectionTypeModemPool{
								Number: phone.Number,
								PoolId: phone.PoolId.String(),
							},
						}
					} else if moxa, err := ci.AsConnectionTypeSerialMoxaSchema(); err == nil {
						connection_info.Connection = &pbdriver.ConnectionInfo_SerialOverIp{
							SerialOverIp: &pbdriver.ConnectionTypeControlledSerial{
								Converter: &pbdriver.ConnectionTypeControlledSerial_Moxa{
									Moxa: &pbdriver.ConnectionTypeSerialMoxa{
										Host:        moxa.Host,
										DataPort:    uint32(moxa.DataPort),
										CommandPort: uint32(moxa.CommandPort),
									},
								},
							},
						}
					}
				}

				app_protocol, err := driver.R2GAppProtocol(device.ApplicationProtocol)
				if err != nil {
					return nil, err
				}

				devices[i] = &pbtaskmaster.JobDevice{
					Id:               device.Id.String(),
					DeviceAttributes: device_attributes,
					ExternalId:       device.ExternalID,
					ConnectionInfo:   connection_info_list,
					AppProtocol:      app_protocol,
					Timezone:         device.Timezone,
				}
			}
		} else {
			return nil, ErrInvalidDeviceList
		}
	} else if list_type == string(JOBDEVICESID) {
		if device_list, err := decodeJobDeviceList(spec.Devices.Union); err == nil && device_list != nil {
			devices = make([]*pbtaskmaster.JobDevice, len(device_list))
			for i, device := range device_list {
				dev_id := device.DeviceId.String()
				devices[i] = &pbtaskmaster.JobDevice{
					Id:       device.Id.String(),
					DeviceId: &dev_id,
				}
			}
		} else {
			return nil, ErrInvalidDeviceList
		}
	} else {
		return nil, ErrInvalidDeviceList
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
		BulkId:        bulk_id,
		CorrelationId: corr_id,
		DriverType:    spec.DriverType,
		Settings:      settings,
		Devices:       devices,
		JobActions:    actions,
		WebhookUrl:    webhook_url,
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
							case *pbdriver.MeasuredValue_StrValue:
								err = tmp_value.FromGenericValueStatusSchemaValue0(vt.StrValue)
								if err != nil {
									return nil, err
								}
								tmp_valueinfo.Value = tmp_value
								tmp[i].Value = tmp_valueinfo
							case *pbdriver.MeasuredValue_TimestampValue:
								err = tmp_value.FromGenericValueStatusSchemaValue1(vt.TimestampValue.AsTime().UnixMilli())
								if err != nil {
									return nil, err
								}
								tmp_valueinfo.Value = tmp_value
								tmp[i].Value = tmp_valueinfo
							case *pbdriver.MeasuredValue_TimestampTzValue:
								err = tmp_value.FromGenericValueStatusSchemaValue0(vt.TimestampTzValue)
								if err != nil {
									return nil, err
								}
								tmp_valueinfo.Value = tmp_value
								tmp[i].Value = tmp_valueinfo
							case *pbdriver.MeasuredValue_BoolValue:
								err = tmp_value.FromGenericValueStatusSchemaValue3(vt.BoolValue)
								if err != nil {
									return nil, err
								}
								tmp_valueinfo.Value = tmp_value
								tmp[i].Value = tmp_valueinfo

							default:
								return nil, fmt.Errorf("unknown value type: %v", vt)
							}
						}
						err = ar.Data.FromExternalRef2DeviceRegistersDataSchema(tmp)
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
						err = ar.Data.FromExternalRef2DeviceProfileDataSchema(tmp)
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
