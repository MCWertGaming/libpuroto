/* <Libpuroto - a shared codebase for Puroto's services>
   Copyright (C) 2022  PurotoApp

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package libpuroto

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ConfigRouter(router *gin.Engine) {

	if os.Getenv("GIN_MODE") == "release" {
		// turn on proxy support
		// TODO: allow users to specify trusted proxies
		// TODO: what if proxy behind proxy
		// TODO: what if no value specified
		ErrorFatal("Router", router.SetTrustedProxies(nil))
	} else {
		// turn off proxy support for debugging
		ErrorFatal("Router", router.SetTrustedProxies(nil))
	}
	// set health status route
	router.GET("/health", getHealth)
}

func getHealth(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", []byte(`{"status":"ok"}`))
}

// redirects to the given url
func Redirect(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, url)
	}
}

// returns true, if the client requested json format, also sets the response to 406, if not
func JsonRequested(c *gin.Context) bool {
	if c.GetHeader("Content-Type") != "application/json" {
		c.AbortWithStatus(http.StatusNotAcceptable)
		LogEvent("authfox", "Received request with wrong Content-Type header")
		return false
	}
	return true
}
