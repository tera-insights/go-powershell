// Copyright (c) 2017 Gorillalabs. All rights reserved.

package middleware

// Middleware ...
type Middleware interface {
	Execute(cmd string) (string, string, error)
	Exit()
}
