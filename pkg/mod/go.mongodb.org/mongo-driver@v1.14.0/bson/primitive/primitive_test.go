// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package primitive

import (
	"encoding/json"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/internal/assert"
	"go.mongodb.org/mongo-driver/internal/require"
)

// The same interface as bsoncodec.Zeroer implemented for tests.
type zeroer interface {
	IsZero() bool
}

func TestCompareTimestamp(t *testing.T) {
	testcases := []struct {
		name     string
		tp       Timestamp
		tp2      Timestamp
		expected int
	}{
		{"equal", Timestamp{T: 12345, I: 67890}, Timestamp{T: 12345, I: 67890}, 0},
		{"T greater than", Timestamp{T: 12345, I: 67890}, Timestamp{T: 2345, I: 67890}, 1},
		{"I greater than", Timestamp{T: 12345, I: 67890}, Timestamp{T: 12345, I: 7890}, 1},
		{"T less than", Timestamp{T: 12345, I: 67890}, Timestamp{T: 112345, I: 67890}, -1},
		{"I less than", Timestamp{T: 12345, I: 67890}, Timestamp{T: 12345, I: 167890}, -1},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := CompareTimestamp(tc.tp, tc.tp2)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestTimestamp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description     string
		tp              Timestamp
		tp2             Timestamp
		expectedAfter   bool
		expectedBefore  bool
		expectedEqual   bool
		expectedCompare int
	}{
		{
			description:     "equal",
			tp:              Timestamp{T: 12345, I: 67890},
			tp2:             Timestamp{T: 12345, I: 67890},
			expectedBefore:  false,
			expectedAfter:   false,
			expectedEqual:   true,
			expectedCompare: 0,
		},
		{
			description:     "T greater than",
			tp:              Timestamp{T: 12345, I: 67890},
			tp2:             Timestamp{T: 2345, I: 67890},
			expectedBefore:  false,
			expectedAfter:   true,
			expectedEqual:   false,
			expectedCompare: 1,
		},
		{
			description:     "I greater than",
			tp:              Timestamp{T: 12345, I: 67890},
			tp2:             Timestamp{T: 12345, I: 7890},
			expectedBefore:  false,
			expectedAfter:   true,
			expectedEqual:   false,
			expectedCompare: 1,
		},
		{
			description:     "T less than",
			tp:              Timestamp{T: 12345, I: 67890},
			tp2:             Timestamp{T: 112345, I: 67890},
			expectedBefore:  true,
			expectedAfter:   false,
			expectedEqual:   false,
			expectedCompare: -1,
		},
		{
			description:     "I less than",
			tp:              Timestamp{T: 12345, I: 67890},
			tp2:             Timestamp{T: 12345, I: 167890},
			expectedBefore:  true,
			expectedAfter:   false,
			expectedEqual:   false,
			expectedCompare: -1,
		},
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable.

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.expectedAfter, tc.tp.After(tc.tp2), "expected After results to be the same")
			assert.Equal(t, tc.expectedBefore, tc.tp.Before(tc.tp2), "expected Before results to be the same")
			assert.Equal(t, tc.expectedEqual, tc.tp.Equal(tc.tp2), "expected Equal results to be the same")
			assert.Equal(t, tc.expectedCompare, tc.tp.Compare(tc.tp2), "expected Compare result to be the same")
		})
	}
}

func TestPrimitiveIsZero(t *testing.T) {
	testcases := []struct {
		name    string
		zero    zeroer
		nonzero zeroer
	}{
		{"binary", Binary{}, Binary{Data: []byte{0x01, 0x02, 0x03}, Subtype: 0xFF}},
		{"decimal128", Decimal128{}, NewDecimal128(1, 2)},
		{"objectID", ObjectID{}, NewObjectID()},
		{"regex", Regex{}, Regex{Pattern: "foo", Options: "bar"}},
		{"dbPointer", DBPointer{}, DBPointer{DB: "foobar", Pointer: ObjectID{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C}}},
		{"timestamp", Timestamp{}, Timestamp{T: 12345, I: 67890}},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require.True(t, tc.zero.IsZero())
			require.False(t, tc.nonzero.IsZero())
		})
	}
}

func TestRegexCompare(t *testing.T) {
	testcases := []struct {
		name string
		r1   Regex
		r2   Regex
		eq   bool
	}{
		{"equal", Regex{Pattern: "foo1", Options: "bar1"}, Regex{Pattern: "foo1", Options: "bar1"}, true},
		{"not equal", Regex{Pattern: "foo1", Options: "bar1"}, Regex{Pattern: "foo2", Options: "bar2"}, false},
		{"not equal", Regex{Pattern: "foo1", Options: "bar1"}, Regex{Pattern: "foo1", Options: "bar2"}, false},
		{"not equal", Regex{Pattern: "foo1", Options: "bar1"}, Regex{Pattern: "foo2", Options: "bar1"}, false},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require.True(t, tc.r1.Equal(tc.r2) == tc.eq)
		})
	}
}

func TestDateTime(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		t.Run("round trip", func(t *testing.T) {
			original := DateTime(1000)
			jsonBytes, err := json.Marshal(original)
			assert.Nil(t, err, "Marshal error: %v", err)

			var unmarshalled DateTime
			err = json.Unmarshal(jsonBytes, &unmarshalled)
			assert.Nil(t, err, "Unmarshal error: %v", err)

			assert.Equal(t, original, unmarshalled, "expected DateTime %v, got %v", original, unmarshalled)
		})
		t.Run("decode null", func(t *testing.T) {
			jsonBytes := []byte("null")
			var dt DateTime
			err := json.Unmarshal(jsonBytes, &dt)
			assert.Nil(t, err, "Unmarshal error: %v", err)
			assert.Equal(t, DateTime(0), dt, "expected DateTime value to be 0, got %v", dt)
		})
		t.Run("UTC", func(t *testing.T) {
			dt := DateTime(1681145535123)
			jsonBytes, err := json.Marshal(dt)
			assert.Nil(t, err, "Marshal error: %v", err)
			assert.Equal(t, `"2023-04-10T16:52:15.123Z"`, string(jsonBytes))
		})
	})
	t.Run("NewDateTimeFromTime", func(t *testing.T) {
		t.Run("range is not limited", func(t *testing.T) {
			// If the implementation internally calls time.Time.UnixNano(), the constructor cannot handle times after
			// the year 2262.

			timeFormat := "2006-01-02T15:04:05.999Z07:00"
			timeString := "3001-01-01T00:00:00Z"
			tt, err := time.Parse(timeFormat, timeString)
			assert.Nil(t, err, "Parse error: %v", err)

			dt := NewDateTimeFromTime(tt)
			assert.True(t, dt > 0, "expected a valid DateTime greater than 0, got %v", dt)
		})
	})
}
