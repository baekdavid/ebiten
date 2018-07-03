// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stb

//go:generate bash build.sh
//go:generate file2byteslice -package=stb -input=decode.js -output=decodejs_file_js.go -var=decode_js
//go:generate file2byteslice -package=stb -input=stbvorbis.js -output=stbvorbisjs_file_js.go -var=stbvorbis_js
//go:generate file2byteslice -package=stb -input=wasm.js -output=wasmjs_file_js.go -var=wasm_js