// Copyright © 2011-12 Qtrac Ltd.
// 
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
    "fmt"
    "github.com/carlvine500/golang-examples/goeg/src/fuzzy_value/fuzzybool"
)

func main() {
    var a fuzzybool.FuzzyBool  // Same as: a, _ := fuzzybool.New(0)
    b, _ := fuzzybool.New(.25) // known valid values; must check if using
    c, _ := fuzzybool.New(.75) // variables though.
    d := c.Copy()
    if err := d.Set(1); err != nil {
        fmt.Println(err)
    }
    process(a, b, c, d)
    s := []fuzzybool.FuzzyBool{a, b, c, d}
    fmt.Println(s)
    m := map[fuzzybool.FuzzyBool]string{a: "a", b: "b", c: "c", c: "d"}
    fmt.Println(m)
}

func process(a, b, c, d fuzzybool.FuzzyBool) {
    fmt.Println("Original:", a, b, c, d)
    fmt.Println("Not:     ", a.Not(), b.Not(), c.Not(), d.Not())
    fmt.Println("Not Not: ", a.Not().Not(), b.Not().Not(), c.Not().Not(),
        d.Not().Not())
    fmt.Print("0.And(.25)→", a.And(b), "• .25.And(.75)→", b.And(c),
        "• .75.And(1)→", c.And(d), " • .25.And(.75,1)→", b.And(c, d), "\n")
    fmt.Print("0.Or(.25)→", a.Or(b), "• .25.Or(.75)→", b.Or(c),
        "• .75.Or(1)→", c.Or(d), " • .25.Or(.75,1)→", b.Or(c, d), "\n")
    fmt.Println("a < c, a == c, a > c:", a.Less(c), a.Equal(c), c.Less(a))
    fmt.Println("Bool:    ", a.Bool(), b.Bool(), c.Bool(), d.Bool())
    fmt.Println("Float:   ", a.Float(), b.Float(), c.Float(), d.Float())
}
