package acquisition

type StatusBits uint64

const (
	// These bits (0..7) follow IDIS specifications, originating from the device

	// Bit 0: Critical error (hardware failure, checksum error)
	StatusBitsCriticalError StatusBits = 1 << 0

	// Bit 1: Clock invalid (power reserve exhausted)
	StatusBitsClockInvalid StatusBits = 1 << 1

	// Bit 2: Data not valid (requires further validation due to special event)
	StatusBitsDataNotValid StatusBits = 1 << 2

	// Bit 3: Daylight saving time is active
	StatusBitsDaylightSavingTime StatusBits = 1 << 3

	// Bit 4: Reserved for future use
	StatusBitsReserved05 StatusBits = 1 << 4

	// Bit 5: Clock adjusted beyond synchronization limit
	StatusBitsClockAdjusted StatusBits = 1 << 5

	// Bit 6: Reserved for future use
	StatusBitsReserved07 StatusBits = 1 << 6

	// Bit 7: Power down — total outage detected during the value recording
	StatusBitsPowerDown StatusBits = 1 << 7

	// These bits (8..15) are used for typical status bits, originating from the device

	// Bit 8: Billing period reset during the value recording
	StatusBitsBillingPeriodReset StatusBits = 1 << 8

	// Bit 9: Parametrization or firmware changed or errorneously set during the value recording
	StatusBitsParametrizationChanged StatusBits = 1 << 9

	// Bit 10: Counter overflow detected (e.g. due to exceeding the maximum value)
	StatusBitsCounterOverflow StatusBits = 1 << 10

	// Bits 11–15: Reserved for future use
	StatusBitsReserved11 StatusBits = 1 << 11
	StatusBitsReserved12 StatusBits = 1 << 12
	StatusBitsReserved13 StatusBits = 1 << 13
	StatusBitsReserved14 StatusBits = 1 << 14
	StatusBitsReserved15 StatusBits = 1 << 15

	// System-originated status bits (bits 16–54)

	// Bit 16: Value originated from user input
	StatusBitsUserValue StatusBits = 1 << 16

	// Bit 17: Value imported from another system or device
	StatusBitsImportedValue StatusBits = 1 << 17

	// Bit 18: Value is estimated (not measured directly)
	StatusBitsEstimatedValue StatusBits = 1 << 18

	// Bit 19: Intermediate value (not final or confirmed), probably incompletely processed
	StatusBitsIntermediateValue StatusBits = 1 << 19

	// Bit 20: Substituted value (measured by a backup device), used when there are multiple measurement sources and this value is not from the primary source
	StatusBitsSubstitutedValue StatusBits = 1 << 20

	// Bit 21: Value was validated by system or operator and found to be valid
	StatusBitsValidValue StatusBits = 1 << 21

	// Bit 22: Value was validated by system or operator but found to be invalid
	StatusBitsInvalidValue StatusBits = 1 << 22

	// Bit 23: Value has been read from a concentrator (e.g., a data concentrator or gateway device), not directly from the measurement device
	StatusBitsDataQualityConcentrator StatusBits = 1 << 23

	// Bit 24: Reserved for future data quality extensions
	StatusBitsDataQualityStatus24 StatusBits = 1 << 24

	// Bit 25: Reserved for future data quality extensions
	StatusBitsDataQualityStatus25 StatusBits = 1 << 25

	StatusBitsReserved26 StatusBits = 1 << 26
	StatusBitsReserved27 StatusBits = 1 << 27
	StatusBitsReserved28 StatusBits = 1 << 28
	StatusBitsReserved29 StatusBits = 1 << 29
	StatusBitsReserved30 StatusBits = 1 << 30
	StatusBitsReserved31 StatusBits = 1 << 31
	StatusBitsReserved32 StatusBits = 1 << 32
	StatusBitsReserved33 StatusBits = 1 << 33
	StatusBitsReserved34 StatusBits = 1 << 34
	StatusBitsReserved35 StatusBits = 1 << 35
	StatusBitsReserved36 StatusBits = 1 << 36
	StatusBitsReserved37 StatusBits = 1 << 37
	StatusBitsReserved38 StatusBits = 1 << 38
	StatusBitsReserved39 StatusBits = 1 << 39
	StatusBitsReserved40 StatusBits = 1 << 40
	StatusBitsReserved41 StatusBits = 1 << 41
	StatusBitsReserved42 StatusBits = 1 << 42
	StatusBitsReserved43 StatusBits = 1 << 43
	StatusBitsReserved44 StatusBits = 1 << 44
	StatusBitsReserved45 StatusBits = 1 << 45
	StatusBitsReserved46 StatusBits = 1 << 46
	StatusBitsReserved47 StatusBits = 1 << 47
	StatusBitsReserved48 StatusBits = 1 << 48
	StatusBitsReserved49 StatusBits = 1 << 49
	StatusBitsReserved50 StatusBits = 1 << 50
	StatusBitsReserved51 StatusBits = 1 << 51
	StatusBitsReserved52 StatusBits = 1 << 52
	StatusBitsReserved53 StatusBits = 1 << 53
	StatusBitsReserved54 StatusBits = 1 << 54

	// User-defined status bits (bits 55–62)
	StatusBitsUserDefined55 StatusBits = 1 << 55
	StatusBitsUserDefined56 StatusBits = 1 << 56
	StatusBitsUserDefined57 StatusBits = 1 << 57
	StatusBitsUserDefined58 StatusBits = 1 << 58
	StatusBitsUserDefined59 StatusBits = 1 << 59
	StatusBitsUserDefined60 StatusBits = 1 << 60
	StatusBitsUserDefined61 StatusBits = 1 << 61
	StatusBitsUserDefined62 StatusBits = 1 << 62

	// Bit 63 must always be 0
)
