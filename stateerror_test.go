/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
package plinko

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlinkoStateError(t *testing.T) {
	var e *PlinkoStateError
	err := CreatePlinkoStateError("foo", "set")

	if errors.As(err, &e) {
		assert.Equal(t, State("foo"), e.State)
		assert.Equal(t, "set", e.Error())
	} else {
		assert.Fail(t, "error not returning properly")
	}

}
