/*
 * Copyright (c) 2019 Omkar Yadav. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package rules

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/google/cel-go/cel"
	"log"
	"sync"
)

var cache map[string]*cel.Program
var mutex sync.Mutex

func genSha() string {
	var bv []byte
	hasher := sha1.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

// Generate the rule
func RuleGenerator(id, program string, env cel.Env) (error, *cel.Program) {
	// Check the value present in cache
	if val, ok := cache[id]; ok {
		return nil, val
	}

	// First time compile
	mutex.Lock()
	if val, ok := cache[id]; ok {
		return nil, val
	}
	defer mutex.Unlock()

	// Lets parse the program
	parsed, issues := env.Parse(program)
	if issues != nil && issues.Err() != nil {
		return issues.Err(), nil
	}

	// Check the progradocumentm
	checked, issues := env.Check(parsed)
	if issues != nil && issues.Err() != nil {
		return issues.Err(), nil
	}

	// Program construction
	prg, err := env.Program(checked)
	if err != nil {
		log.Fatalf("program construction error: %s", err)
	}

	// Store into the cache
	cache[id] = &prg

	return nil, &prg
}
