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

func TestCreatePlinkoTriggerError(t *testing.T) {
	var e *PlinkoTriggerError
	err := CreatePlinkoTriggerError("foo", "set")

	if errors.As(err, &e) {
		assert.Equal(t, Trigger("foo"), e.Trigger)
		assert.Equal(t, "set", e.Error())
	} else {
		assert.Fail(t, "error not returning properly")
	}
}