package job

import (
	"fmt"
	"time"

	"github.com/cybroslabs/hes-2-apis/openapi/openhes/attribute"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrUnknownJobActionType = fmt.Errorf("unknown job action type")
)

func APIv1ActionToGRPCAction(action *JobActionSchema) (*pbdriver.JobAction, error) {
	if tmp, err := action.AsJobActionGetRegisterSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetPeriodicalProfileSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetIrregularProfileSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetEventsSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetClockSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionSyncClockSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetRelayStateSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionSetRelayStateSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetDisconnectorStateSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionSetDisconnectorStateSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetTouSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionSetTouSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionGetLimiterSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionSetLimiterSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionResetBillingPeriodSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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
	if tmp, err := action.AsJobActionFwUpdateSchema(); err == nil {
		attr, err := attribute.AttributesToAPIv1Attributes(tmp.Attributes)
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

func GRPCActionToAPIv1Action(action *pbdriver.JobAction, result *JobActionSchema) error {
	if action == nil {
		return ErrUnknownJobActionType
	}

	action_id, _ := uuid.Parse(action.ActionId)

	if tmp := action.GetGetRegister(); tmp != nil {
		err := result.FromJobActionGetRegisterSchema(JobActionGetRegisterSchema{
			Id:         action_id,
			Type:       string(GETREGISTER),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
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
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetIrregularProfile(); tmp != nil {
		err := result.FromJobActionGetIrregularProfileSchema(JobActionGetIrregularProfileSchema{
			Id:         action_id,
			Type:       string(GETIRREGULARPROFILE),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetEvents(); tmp != nil {
		err := result.FromJobActionGetEventsSchema(JobActionGetEventsSchema{
			Id:         action_id,
			Type:       string(GETEVENTS),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetClock(); tmp != nil {
		err := result.FromJobActionGetClockSchema(JobActionGetClockSchema{
			Id:         action_id,
			Type:       string(GETCLOCK),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSyncClock(); tmp != nil {
		err := result.FromJobActionSyncClockSchema(JobActionSyncClockSchema{
			Id:         action_id,
			Type:       string(SYNCCLOCK),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetRelayState(); tmp != nil {
		err := result.FromJobActionGetRelayStateSchema(JobActionGetRelayStateSchema{
			Id:         action_id,
			Type:       string(GETRELAYSTATE),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetRelayState(); tmp != nil {
		err := result.FromJobActionSetRelayStateSchema(JobActionSetRelayStateSchema{
			Id:         action_id,
			Type:       string(SETRELAYSTATE),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetDisconnectorState(); tmp != nil {
		err := result.FromJobActionGetDisconnectorStateSchema(JobActionGetDisconnectorStateSchema{
			Id:         action_id,
			Type:       string(GETDISCONNECTORSTATE),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetDisconnectorState(); tmp != nil {
		err := result.FromJobActionSetDisconnectorStateSchema(JobActionSetDisconnectorStateSchema{
			Id:         action_id,
			Type:       string(SETDISCONNECTORSTATE),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetTou(); tmp != nil {
		err := result.FromJobActionGetTouSchema(JobActionGetTouSchema{
			Id:         action_id,
			Type:       string(GETTOU),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetTou(); tmp != nil {
		err := result.FromJobActionSetTouSchema(JobActionSetTouSchema{
			Id:         action_id,
			Type:       string(SETTOU),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetGetLimiter(); tmp != nil {
		err := result.FromJobActionGetLimiterSchema(JobActionGetLimiterSchema{
			Id:         action_id,
			Type:       string(GETLIMITER),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetSetLimiter(); tmp != nil {
		err := result.FromJobActionSetLimiterSchema(JobActionSetLimiterSchema{
			Id:         action_id,
			Type:       string(SETLIMITER),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetResetBillingPeriod(); tmp != nil {
		err := result.FromJobActionResetBillingPeriodSchema(JobActionResetBillingPeriodSchema{
			Id:         action_id,
			Type:       string(RESETBILLINGPERIOD),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	} else if tmp := action.GetFwUpdate(); tmp != nil {
		err := result.FromJobActionFwUpdateSchema(JobActionFwUpdateSchema{
			Id:         action_id,
			Type:       string(FWUPDATE),
			Attributes: attribute.APIv1AttributesToAttributes(action.Attributes),
		})
		return err
	}

	return ErrUnknownJobActionType
}
