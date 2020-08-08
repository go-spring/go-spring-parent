/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package SpringError_test

import (
	"fmt"
	"testing"

	"github.com/go-spring/go-spring-parent/spring-error"
	"github.com/go-spring/go-spring-parent/spring-utils"
)

func TestRpcPanic(t *testing.T) {

	r := SpringError.SUCCESS.Data("123")
	fmt.Println(SpringUtils.ToJson(r))

	t.Run("Panic", func(t *testing.T) {

		defer func() {
			err := recover()
			fmt.Println(SpringUtils.ToJson(err))
		}()

		err := fmt.Errorf("reason: %s", "panic")
		SpringError.ERROR.Panic(err).When(err != nil)
	})

	SpringError.ERROR.Panicf("reason: %s", "panicf").When(false)

	t.Run("Panicf", func(t *testing.T) {

		defer func() {
			err := recover()
			fmt.Println(SpringUtils.ToJson(err))
		}()

		SpringError.ERROR.Panicf("reason: %s", "panicf").When(true)
	})
}
