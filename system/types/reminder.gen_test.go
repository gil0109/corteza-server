package types

import (
	"testing"

	"errors"

	"github.com/stretchr/testify/require"
)

// 	Hello! This file is auto-generated.

func TestReminderSetWalk(t *testing.T) {
	var (
		value = make(ReminderSet, 3)
		req   = require.New(t)
	)

	// check walk with no errors
	{
		err := value.Walk(func(*Reminder) error {
			return nil
		})
		req.NoError(err)
	}

	// check walk with error
	req.Error(value.Walk(func(*Reminder) error { return errors.New("walk error") }))

}

func TestReminderSetFilter(t *testing.T) {
	var (
		value = make(ReminderSet, 3)
		req   = require.New(t)
	)

	// filter nothing
	{
		set, err := value.Filter(func(*Reminder) (bool, error) {
			return true, nil
		})
		req.NoError(err)
		req.Equal(len(set), len(value))
	}

	// filter one item
	{
		found := false
		set, err := value.Filter(func(*Reminder) (bool, error) {
			if !found {
				found = true
				return found, nil
			}
			return false, nil
		})
		req.NoError(err)
		req.Len(set, 1)
	}

	// filter error
	{
		_, err := value.Filter(func(*Reminder) (bool, error) {
			return false, errors.New("filter error")
		})
		req.Error(err)
	}
}

func TestReminderSetIDs(t *testing.T) {
	var (
		value = make(ReminderSet, 3)
		req   = require.New(t)
	)

	// construct objects
	value[0] = new(Reminder)
	value[1] = new(Reminder)
	value[2] = new(Reminder)
	// set ids
	value[0].ID = 1
	value[1].ID = 2
	value[2].ID = 3

	// Find existing
	{
		val := value.FindByID(2)
		req.Equal(uint64(2), val.ID)
	}

	// Find non-existing
	{
		val := value.FindByID(4)
		req.Nil(val)
	}

	// List IDs from set
	{
		val := value.IDs()
		req.Equal(len(val), len(value))
	}
}
