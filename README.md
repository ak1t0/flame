# flame
[![Go Report Card](https://goreportcard.com/badge/github.com/ak1t0/flame)](https://goreportcard.com/report/github.com/ak1t0/flame)

flame is a onion service crawler

## Installation

flame requires Go(>=1.7.4) and tor.


    go get github.com/ak1t0/flame
    sudo apt install tor

## Usage


    flame -f "log.json"


## Target JSON
```json
{"domain":"facebookcorewwwi.onion"}
{"domain":"deepdot35wvmeyd5.onion"}
...
```

## Output
```
2017/08/01 22:06:02 [{facebookcorewwwi.onion 0 available} {deepdot35wvmeyd5.onion 0 timeout} ...]
```

## TODO
- response data preservation
- response data analysis


## License

Copyright (c) 2017, ak1t0

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.

    * Redistributions in binary form must reproduce the above
      copyright notice, this list of conditions and the following
      disclaimer in the documentation and/or other materials provided
      with the distribution.

    * Neither the name of ak1t0 nor the names of other
      contributors may be used to endorse or promote products derived
      from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.