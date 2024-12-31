package job

import (
	"bytes"
	"encoding/json"
)

// Implements JSON unmarshalling for JobActionGetRegisterSchema with UseNumber.
func asJobActionGetRegisterSchema(action *JobActionSchema) (JobActionGetRegisterSchema, error) {
	tmp := JobActionGetRegisterSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetRegister != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetRegister"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetPeriodicalProfileSchema with UseNumber.
func asJobActionGetPeriodicalProfileSchema(action *JobActionSchema) (JobActionGetPeriodicalProfileSchema, error) {
	tmp := JobActionGetPeriodicalProfileSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetPeriodicalProfile != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetPeriodicalProfile"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetIrregularProfileSchema with UseNumber.
func asJobActionGetIrregularProfileSchema(action *JobActionSchema) (JobActionGetIrregularProfileSchema, error) {
	tmp := JobActionGetIrregularProfileSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetIrregularProfile != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetIrregularProfile"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetEventsSchema with UseNumber.
func asJobActionGetEventsSchema(action *JobActionSchema) (JobActionGetEventsSchema, error) {
	tmp := JobActionGetEventsSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetEvents != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetEvents"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetClockSchema with UseNumber.
func asJobActionGetClockSchema(action *JobActionSchema) (JobActionGetClockSchema, error) {
	tmp := JobActionGetClockSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetClock != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetClock"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionSyncClockSchema with UseNumber.
func asJobActionSyncClockSchema(action *JobActionSchema) (JobActionSyncClockSchema, error) {
	tmp := JobActionSyncClockSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeSyncClock != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeSyncClock"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetRelayStateSchema with UseNumber.
func asJobActionGetRelayStateSchema(action *JobActionSchema) (JobActionGetRelayStateSchema, error) {
	tmp := JobActionGetRelayStateSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetRelayState != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetRelayState"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionSetRelayStateSchema with UseNumber.
func asJobActionSetRelayStateSchema(action *JobActionSchema) (JobActionSetRelayStateSchema, error) {
	tmp := JobActionSetRelayStateSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeSetRelayState != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeSetRelayState"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetDisconnectorStateSchema with UseNumber.
func asJobActionGetDisconnectorStateSchema(action *JobActionSchema) (JobActionGetDisconnectorStateSchema, error) {
	tmp := JobActionGetDisconnectorStateSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetDisconnectorState != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetDisconnectorState"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionSetDisconnectorStateSchema with UseNumber.
func asJobActionSetDisconnectorStateSchema(action *JobActionSchema) (JobActionSetDisconnectorStateSchema, error) {
	tmp := JobActionSetDisconnectorStateSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeSetDisconnectorState != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeSetDisconnectorState"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetTouSchema with UseNumber.
func asJobActionGetTouSchema(action *JobActionSchema) (JobActionGetTouSchema, error) {
	tmp := JobActionGetTouSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetTou != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetTou"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionSetTouSchema with UseNumber.
func asJobActionSetTouSchema(action *JobActionSchema) (JobActionSetTouSchema, error) {
	tmp := JobActionSetTouSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeSetTou != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeSetTou"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionGetLimiterSchema with UseNumber.
func asJobActionGetLimiterSchema(action *JobActionSchema) (JobActionGetLimiterSchema, error) {
	tmp := JobActionGetLimiterSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeGetLimiter != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeGetLimiter"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionSetLimiterSchema with UseNumber.
func asJobActionSetLimiterSchema(action *JobActionSchema) (JobActionSetLimiterSchema, error) {
	tmp := JobActionSetLimiterSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeSetLimiter != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeSetLimiter"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionResetBillingPeriodSchema with UseNumber.
func asJobActionResetBillingPeriodSchema(action *JobActionSchema) (JobActionResetBillingPeriodSchema, error) {
	tmp := JobActionResetBillingPeriodSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeResetBillingPeriod != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeResetBillingPeriod"}
	}
	return tmp, err
}

// Implements JSON unmarshalling for JobActionFwUpdateSchema with UseNumber.
func asJobActionFwUpdateSchema(action *JobActionSchema) (JobActionFwUpdateSchema, error) {
	tmp := JobActionFwUpdateSchema{}
	d := json.NewDecoder(bytes.NewReader(action.Union))
	d.UseNumber()
	d.DisallowUnknownFields()
	err := d.Decode(&tmp)
	if err == nil && tmp.TypeFwUpdate != 1 {
		err = &json.UnmarshalTypeError{Field: "TypeFwUpdate"}
	}
	return tmp, err
}
