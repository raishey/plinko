/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
package state

import "gitlab.com/stevehebert/plinko"

func WithName(name string) func(*plinko.StateConfig) {
	return func(c *plinko.StateConfig) {
		c.Name = name
	}
}

func WithDescription(description string) func(*plinko.StateConfig) {
	return func(c *plinko.StateConfig) {
		c.Description = description
	}
}
