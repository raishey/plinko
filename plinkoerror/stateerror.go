/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
package plinkoerror

import "gitlab.com/stevehebert/plinko"

type PlinkoStateError struct {
	plinko.State
	ErrorMessage string
}

func (e *PlinkoStateError) Error() string {
	return e.ErrorMessage
}

func CreatePlinkoStateError(state plinko.State, errorMessage string) error {
	return &PlinkoStateError{
		State:        state,
		ErrorMessage: errorMessage,
	}
}
