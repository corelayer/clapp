/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package clapp

// func TestNewConfiguration(t *testing.T) {
// 	filename := "config.yaml"
// 	path := "$PWD"
//
// 	var err error
// 	_, err = NewConfiguration(filename, path, nil, false, false)
// 	if err != nil {
// 		t.Errorf("failed to create new configuration with message %s", err)
// 	}
// }
//
// func TestConfiguration_Initialize(t *testing.T) {
// 	filename := "config.yaml"
// 	path := "$PWD"
//
// 	var c *Configuration
// 	var err error
// 	c, err = NewConfiguration(filename, path, nil, false, false)
//
// 	if err != nil {
// 		t.Errorf("failed to create new configuration with message %s", err)
// 	}
//
// 	c.Initialize()
//
// 	if c.IsInitialized() != false {
// 		t.Errorf("Initialize() failed with autoLoad set to %t", false)
// 	}
// }
//
// func TestConfiguration_Load(t *testing.T) {
// 	filename := "config.yaml"
// 	path := "./"
//
// 	var c *Configuration
// 	var err error
//
// 	// Create dummy config file
// 	var f *os.File
// 	f, err = os.Create(filepath.Join(path, filename))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	time.Sleep(5 * time.Second)
// 	defer func(name string) {
// 		err := os.Remove(name)
// 		if err != nil {
//
// 		}
// 	}(f.Name()) // clean up
//
// 	c, err = NewConfiguration(filename, path, nil, false, false)
//
// 	if err != nil {
// 		t.Errorf("failed to create new configuration with message %s", err)
// 	}
//
// 	err = c.Load()
// 	if err != nil {
// 		t.Errorf("Load() failed with message %s", err)
// 	}
// }
