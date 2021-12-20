/*
 * Copyright 2021 DADi590
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Package GlobalUtils_APU contains constants useful inside and outside this module
package GlobalUtils_APU

// VERSION is the constant to check to know the version of the compiled module. The format is
// 		"yyyy-MM-dd -- HH:mm:ss.SSSSSS ([timezone taken from the system])"
const VERSION string = "2021-12-06 -- 22:02:04.192837 (Hora padrão de GMT)"

// ASSISTANT_NAME is the constant that has the assistant name used in the entire module.
const ASSISTANT_NAME string = "V.I.S.O.R."

// APU_ERR_PREFIX is the prefix to be used after MOD_RET_ERR_PREFIX and its additions to return a custom error (that is,
// an error that is not from third-party/Go libraries). A string must be appended with the format "X: Y", in which X is
// a unique error identifier (float - 1 or 1.1) for the submodule, and Y is an error description. A result example
// might be
// 		"APU error 1: Some description here"
const APU_ERR_PREFIX = "APU error "
