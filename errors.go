/*
- This source code is licensed under the MIT license found in the
- LICENSE file in the root directory of this source tree.
*/

package plinko

import "fmt"

type PlinkoStateError struct {
	State
	ErrorMessage string
}

func (e *PlinkoStateError) Error() string {
	return e.ErrorMessage
}

func CreatePlinkoStateError(state State, errorMessage string) error {
	return &PlinkoStateError{
		State:        state,
		ErrorMessage: errorMessage,
	}
}

type PlinkoTriggerError struct {
	Trigger
	ErrorMessage string
}

func (e *PlinkoTriggerError) Error() string {
	return e.ErrorMessage
}

func CreatePlinkoTriggerError(trigger Trigger, errorMessage string) error {
	return &PlinkoTriggerError{
		Trigger:      trigger,
		ErrorMessage: errorMessage,
	}
}

type PlinkoPanicError struct {
	TransitionInfo
	StepNumber        int
	StepName          string
	InnerError        error
	UnknownInnerError interface{}
	Stack             string
}

func (ce *PlinkoPanicError) Error() string {
	return fmt.Sprintf("%+v", *ce)
}

func CreatePlinkoPanicError(pn interface{}, t TransitionInfo, step int, name string, stack string) error {
	if err, ok := pn.(error); ok {
		return &PlinkoPanicError{
			TransitionInfo: t,
			StepNumber:     step,
			StepName:       name,
			InnerError:     err,
			Stack:          stack,
		}
	}

	return &PlinkoPanicError{
		TransitionInfo:    t,
		StepNumber:        step,
		StepName:          name,
		UnknownInnerError: pn,
		Stack:             stack,
	}
}
