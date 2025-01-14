package errorx_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/herodot"

	"github.com/ory/hive/errorx"
	"github.com/ory/hive/internal"
	"github.com/ory/hive/x"
)

func TestHandler(t *testing.T) {
	_, reg := internal.NewMemoryRegistry(t)
	h := errorx.NewHandler(reg)

	router := x.NewRouterPublic()
	h.RegisterPublicRoutes(router)
	ts := httptest.NewServer(router)
	defer ts.Close()

	for k, tc := range []struct {
		gave []error
	}{
		{
			gave: []error{
				herodot.ErrNotFound.WithReason("foobar"),
			},
		},
		// {
		// 	gave: []error{
		// 		&x.ErrAuthnPasswordCredentialsInvalid,
		// 	},
		// },
		// {
		// 	gave: []error{
		// 		x.ErrAuthnPasswordCredentialsInvalid.WithMessage("whatever"),
		// 	},
		// },
		{
			gave: []error{
				herodot.ErrNotFound.WithReason("foobar"),
				herodot.ErrNotFound.WithReason("foobar"),
			},
		},
		{
			gave: []error{
				errors.WithStack(herodot.ErrNotFound.WithReason("foobar")),
			},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			id, err := reg.ErrorManager().Add(tc.gave...)
			require.NoError(t, err)

			res, err := ts.Client().Get(ts.URL + "/errors?error=" + id)
			require.NoError(t, err)
			defer res.Body.Close()
			assert.EqualValues(t, http.StatusOK, res.StatusCode)

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			gg := make([]error, len(tc.gave))
			for k, g := range tc.gave {
				gg[k] = errors.Cause(g)
			}

			gave, err := json.Marshal(gg)
			require.NoError(t, err)

			assert.JSONEq(t, string(gave), string(body))
			assert.NotEqual(t, "[{}]", string(body))
			t.Logf("%s", body)
		})
	}
}
