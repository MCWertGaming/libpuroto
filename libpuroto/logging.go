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
	"log"
	"os"
	"runtime"
)

func redBoldColor(message string) string {
	if _, exists := os.LookupEnv("DISABLE_COLOR"); exists {
		return message
	} else {
		return "\033[31m\033[1m" + message + "\033[m"
	}
}
func boldColor(message string) string {
	if _, exists := os.LookupEnv("DISABLE_COLOR"); exists {
		return message
	} else {
		return "\033[1m" + message + "\033[m"
	}
}
func ErrorFatal(name string, err error) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Fatalf("[%v] [%v] [%v:%v] %v", name, redBoldColor("FATAL"), filename, line, redBoldColor(err.Error()))
	}
}
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func LogEvent(name string, message string) {
	_, filename, line, _ := runtime.Caller(1)
	log.Printf("[%v] [%v] [%v:%v] %v", name, boldColor("LOG"), filename, line, boldColor(message))
}
func LogError(name string, err error) {
	_, filename, line, _ := runtime.Caller(1)
	log.Printf("[%v] [%v] [%v:%v] %v", name, redBoldColor("ERROR"), filename, line, redBoldColor(err.Error()))
}
