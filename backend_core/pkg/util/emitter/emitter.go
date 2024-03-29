// Package emitter provides an event emitter.
package emitter

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
)

/*
 Copyright (C) 2022 Eunimart Omnichannel Pvt Ltd. (www.eunimart.com)
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Lesser General Public License v3.0 as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Lesser General Public License v3.0 for more details.
 You should have received a copy of the GNU Lesser General Public License v3.0
 along with this program.  If not, see <https://www.gnu.org/licenses/lgpl-3.0.html/>.
*/
// Default number of maximum listeners for an event.
const DefaultMaxListeners = 1000

// Error presented when an invalid argument is provided as a listener function
var ErrNoneFunction = errors.New("Kind of Value for listener is not Func.")

// Emitter object
var obj *Emitter

// RecoveryListener ...
type RecoveryListener func(interface{}, interface{}, error)

// Emitter ...
type Emitter struct {
	// Mutex to prevent race conditions within the Emitter.
	*sync.Mutex
	// Map of event to a slice of listener function's reflect Values.
	events map[interface{}][]reflect.Value
	// Optional RecoveryListener to call when a panic occurs.
	recoverer RecoveryListener
	// Maximum listeners for debugging potential memory leaks.
	maxListeners int
	// Map used to remove Listeners wrapped in a Once func
	onces map[reflect.Value]reflect.Value
}

// AddListener appends the listener argument to the event arguments slice
// in the Emitter's events map. If the number of listeners for an event
// is greater than the Emitter's maximum listeners then a warning is printed.
// If the relect Value of the listener does not have a Kind of Func then
// AddListener panics. If a RecoveryListener has been set then it is called
// recovering from the panic.
func (emitter *Emitter) AddListener(event, listener interface{}) *Emitter {
	emitter.Lock()
	defer emitter.Unlock()

	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		if nil == emitter.recoverer {
			panic(ErrNoneFunction)
		} else {
			emitter.recoverer(event, listener, ErrNoneFunction)
		}
	}

	if emitter.maxListeners != -1 && emitter.maxListeners < len(emitter.events[event])+1 {
		fmt.Fprintf(os.Stdout, "Warning: event `%v` has exceeded the maximum "+
			"number of listeners of %d.\n", event, emitter.maxListeners)
	}

	emitter.events[event] = append(emitter.events[event], fn)

	return emitter
}

// On is an alias for AddListener.
func (emitter *Emitter) On(event, listener interface{}) *Emitter {
	return emitter.AddListener(event, listener)
}

// RemoveListener removes the listener argument from the event arguments slice
// in the Emitter's events map.  If the reflect Value of the listener does not
// have a Kind of Func then RemoveListener panics. If a RecoveryListener has
// been set then it is called after recovering from the panic.
func (emitter *Emitter) RemoveListener(event, listener interface{}) *Emitter {
	emitter.Lock()
	defer emitter.Unlock()

	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		if nil == emitter.recoverer {
			panic(ErrNoneFunction)
		} else {
			emitter.recoverer(event, listener, ErrNoneFunction)
		}
	}

	if events, ok := emitter.events[event]; ok {
		if _, ok = emitter.onces[fn]; ok {
			fn = emitter.onces[fn]
		}

		newEvents := []reflect.Value{}

		for _, listener := range events {
			if fn.Pointer() != listener.Pointer() {
				newEvents = append(newEvents, listener)
			}
		}

		emitter.events[event] = newEvents
	}

	return emitter
}

// Off is an alias for RemoveListener.
func (emitter *Emitter) Off(event, listener interface{}) *Emitter {
	return emitter.RemoveListener(event, listener)
}

// Once generates a new function which invokes the supplied listener
// only once before removing itself from the event's listener slice
// in the Emitter's events map. If the reflect Value of the listener
// does not have a Kind of Func then Once panics. If a RecoveryListener
// has been set then it is called after recovering from the panic.
func (emitter *Emitter) Once(event, listener interface{}) *Emitter {
	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		if nil == emitter.recoverer {
			panic(ErrNoneFunction)
		} else {
			emitter.recoverer(event, listener, ErrNoneFunction)
		}
	}

	var run func(...interface{})

	run = func(arguments ...interface{}) {
		defer emitter.RemoveListener(event, run)

		var values []reflect.Value

		for i := 0; i < len(arguments); i++ {
			values = append(values, reflect.ValueOf(arguments[i]))
		}

		fn.Call(values)
	}

	// Lock before changing onces
	emitter.Lock()
	emitter.onces[fn] = reflect.ValueOf(run)
	emitter.Unlock()

	emitter.AddListener(event, run)
	return emitter
}

// Emit attempts to use the reflect package to Call each listener stored
// in the Emitter's events map with the supplied arguments. Each listener
// is called within its own go routine. The reflect package will panic if
// the agruments supplied do not align the parameters of a listener function.
// If a RecoveryListener has been set then it is called after recovering from
// the panic.
func (emitter *Emitter) Emit(event interface{}, arguments ...interface{}) *Emitter {
	var (
		listeners []reflect.Value
		ok        bool
	)

	// Lock the mutex when reading from the Emitter's
	// events map.
	emitter.Lock()

	if listeners, ok = emitter.events[event]; !ok {
		// If the Emitter does not include the event in its
		// event map, it has no listeners to Call yet.
		emitter.Unlock()
		return emitter
	}

	// Unlock the mutex immediately following the read
	// instead of deferring so that listeners registered
	// with Once can aquire the mutex for removal.
	emitter.Unlock()

	var wg sync.WaitGroup

	wg.Add(len(listeners))

	for _, fn := range listeners {
		go func(fn reflect.Value) {
			defer wg.Done()

			// Recover from potential panics, supplying them to a
			// RecoveryListener if one has been set, else allowing
			// the panic to occur.
			if nil != emitter.recoverer {
				defer func() {
					if r := recover(); nil != r {
						err := fmt.Errorf("%v", r)
						emitter.recoverer(event, fn.Interface(), err)
					}
				}()
			}

			var values []reflect.Value

			for i := 0; i < len(arguments); i++ {
				if arguments[i] == nil {
					values = append(values, reflect.New(fn.Type().In(i)).Elem())
				} else {
					values = append(values, reflect.ValueOf(arguments[i]))
				}
			}

			fn.Call(values)
		}(fn)
	}

	wg.Wait()
	return emitter
}

// EmitSync attempts to use the reflect package to Call each listener stored
// in the Emitter's events map with the supplied arguments. Each listener
// is called synchronously. The reflect package will panic if
// the agruments supplied do not align the parameters of a listener function.
// If a RecoveryListener has been set then it is called after recovering from
// the panic.
func (emitter *Emitter) EmitSync(event interface{}, arguments ...interface{}) *Emitter {
	var (
		listeners []reflect.Value
		ok        bool
	)

	// Lock the mutex when reading from the Emitter's
	// events map.
	emitter.Lock()

	if listeners, ok = emitter.events[event]; !ok {
		// If the Emitter does not include the event in its
		// event map, it has no listeners to Call yet.
		emitter.Unlock()
		return emitter
	}

	// Unlock the mutex immediately following the read
	// instead of deferring so that listeners registered
	// with Once can aquire the mutex for removal.
	emitter.Unlock()

	for _, fn := range listeners {
		// Recover from potential panics, supplying them to a
		// RecoveryListener if one has been set, else allowing
		// the panic to occur.
		if nil != emitter.recoverer {
			defer func() {
				if r := recover(); nil != r {
					err := fmt.Errorf("%v", r)
					emitter.recoverer(event, fn.Interface(), err)
				}
			}()
		}

		var values []reflect.Value

		for i := 0; i < len(arguments); i++ {
			if arguments[i] == nil {
				values = append(values, reflect.New(fn.Type().In(i)).Elem())
			} else {
				values = append(values, reflect.ValueOf(arguments[i]))
			}
		}

		fn.Call(values)
	}

	return emitter
}

// RecoverWith sets the listener to call when a panic occurs, recovering from
// panics and attempting to keep the application from crashing.
func (emitter *Emitter) RecoverWith(listener RecoveryListener) *Emitter {
	emitter.recoverer = listener
	return emitter
}

// SetMaxListeners sets the maximum number of listeners per
// event for the Emitter. If -1 is passed as the maximum,
// all events may have unlimited listeners. By default, each
// event can have a maximum number of 10 listeners which is
// useful for finding memory leaks.
func (emitter *Emitter) SetMaxListeners(max int) *Emitter {
	emitter.Lock()
	defer emitter.Unlock()

	emitter.maxListeners = max
	return emitter
}

// GetListenerCount gets count of listeners for a given event.
func (emitter *Emitter) GetListenerCount(event interface{}) (count int) {
	emitter.Lock()
	if listeners, ok := emitter.events[event]; ok {
		count = len(listeners)
	}
	emitter.Unlock()
	return
}

// NewEmitter returns a new Emitter object, defaulting the
// number of maximum listeners per event to the DefaultMaxListeners
// constant and initializing its events map.
func NewEmitter() (emitter *Emitter) {
	emitter = new(Emitter)
	emitter.Mutex = new(sync.Mutex)
	emitter.events = make(map[interface{}][]reflect.Value)
	emitter.maxListeners = DefaultMaxListeners
	emitter.onces = make(map[reflect.Value]reflect.Value)
	return
}

// func CreateObj() {
// 	obj = NewEmitter()
// }

func GetObj() (emitter *Emitter) {
	if obj == nil {
		obj = NewEmitter()
	}
	return obj
}
