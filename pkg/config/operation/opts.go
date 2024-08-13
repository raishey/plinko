/**
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
package operation

import (
	"gitlab.com/stevehebert/plinko"
)

func WithName(name string) func(*plinko.OperationConfig) {
	return func(c *plinko.OperationConfig) {
		c.Name = name
	}
}
