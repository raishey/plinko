/*
- This source code is licensed under the MIT license found in the
- LICENSE file in the root directory of this source tree.
*/
package plinko

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
