/*
 * Copyright [there's no copyright notice for this. I'll just leave this empty]
 */

package Tcf

///////////////////////////////////////
// Try / Catch / Finally

// Credits: https://dzone.com/articles/try-and-catch-in-golang

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Tcf) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

type Tcf struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

/* Original example:
Tcf {
	Try: func() {
		fmt.Println("I tried")
		Throw("Oh,...sh...")
	},
	Catch: func(e Exception) {
		fmt.Printf("Caught %v\n", e)
	},
	Finally: func() {
		fmt.Println("Finally...")
	},
}.Do()
*/
///////////////////////////////////////
