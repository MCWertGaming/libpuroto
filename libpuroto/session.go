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
	"crypto/subtle"
	"time"

	"github.com/go-redis/redis"
)

// returns true if the session of the given redis DB is valid
func SessionValid(uid, token *string, redisClient *redis.Client) (bool, error) {
	var res string
	var err error

	// the UUID session extension is part of the session, so no work is needed
	res, err = redisClient.Get(*uid).Result()

	if err == redis.Nil {
		// the uid has no session stored in redis, it's not valid therefore
		return false, nil
	} else if err != nil {
		return false, err
	} else if subtle.ConstantTimeCompare([]byte(res), []byte(*token)) == 1 {
		// session and token match, the session is valid
		// We'll increase the TTL to keep the session alive
		redisClient.Expire(*uid, time.Hour*24*7)

		return true, nil
	}
	// the session seems to not being valid
	return false, nil
}
