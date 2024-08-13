/**
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
package renderers_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stevehebert/plinko/internal/renderers"
	"gitlab.com/stevehebert/plinko/pkg/config"
)

func Test_CreateUML(t *testing.T) {
	p := config.CreatePlinkoDefinition()

	p.Configure(NewOrder).
		Permit("Submit", "PublishedOrder").
		Permit("Review", "UnderReview")

	p.Configure("PublishedOrder")

	p.Configure("UnderReview").
		Permit("CompleteReview", "PublishedOrder").
		Permit("RejectOrder", "RejectedOrder")

	p.Configure("RejectedOrder")

	buf := bytes.NewBufferString("")

	err := p.Render(renderers.NewUML(buf))
	assert.Nil(t, err)
	assert.Contains(t, buf.String(), "UnderReview --> PublishedOrder : CompleteReview")
}
