package data

const (
	// CHAR_MAX represents the maximum value of an EO char (1-byte encoded integer)
	CHAR_MAX int = 253

	// SHORT_MAX represents the maximum value of an EO short (2-byte encoded integer)
	SHORT_MAX int = CHAR_MAX * CHAR_MAX

	// THREE_MAX represents the maximum value of an EO three (3-byte encoded integer)
	THREE_MAX int = CHAR_MAX * CHAR_MAX * CHAR_MAX

	// INT_MAX represents the maximum value of an EO int (4-byte encoded integer)
	INT_MAX int = SHORT_MAX * SHORT_MAX
)
