// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package handler_test

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bangumi/server/config"
	"github.com/bangumi/server/internal/test"
	"github.com/bangumi/server/mocks"
	"github.com/bangumi/server/model"
	"github.com/bangumi/server/web/res"
	"github.com/bangumi/server/web/session"
)

func TestHandler_GetCurrentUser_private(t *testing.T) {
	t.Parallel()
	const uid model.UIDType = 7
	const sessionID = "ss"

	u := mocks.NewUserRepo(t)
	u.EXPECT().GetByID(mock.Anything, uid).Return(model.User{ID: uid}, nil)

	s := mocks.NewSessionManager(t)
	s.EXPECT().Get(mock.Anything, sessionID).Return(session.Session{UserID: uid}, nil)

	app := test.GetWebApp(t,
		test.Mock{
			UserRepo:       u,
			SessionManager: s,
		},
	)

	var r res.User
	resp := test.New(t).Get("/p/me").Cookie(session.Key, sessionID).
		Header(fiber.HeaderOrigin, config.FrontendOrigin).
		Execute(app).JSON(&r)

	require.Equal(t, http.StatusOK, resp.StatusCode, resp.BodyString())
	require.Equal(t, uid, r.ID, resp.BodyString())
}