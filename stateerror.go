/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
package plinko

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
