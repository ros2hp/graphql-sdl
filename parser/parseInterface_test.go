package parser

import (
	"fmt"
	"testing"

	"github.com/ros2hp/graphql-sdl/internal/db"
	"github.com/ros2hp/graphql-sdl/lexer"
)

func TestImplements1(t *testing.T) {

	input := `
interface ValuedEntity {
  value: Int
}

type Person implements NamedEntity {
  name: String
  age: Int
}

`

	var expectedErr [1]string
	expectedErr[0] = `"NamedEntity" does not exist in document "DefaultDoc" at line: 6 column: 24`

	err := db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplements2(t *testing.T) {

	input := `
type NamedEntity {
  value: Int
}

type Person implements NamedEntity {
  name: String
  age: Int
}

`
	var expectedErr []string = []string{
		`"NamedEntity" is not an interface type, at line: 6 column: 24`,
	}

	l := lexer.New(input)
	p := New(l)

	_, errs := p.ParseDocument()
	for _, v := range errs {
		fmt.Println("errs:  ", v.Error())
	}
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplements3(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
  name2: Int

}
interface ValuedEntity {
  value: Int
}

type Person implements NamedEntity & ValuedEntity2 {
  name: String
  age: Int
}
`
	err := db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	var expectedErr []string = []string{
		`"ValuedEntity2"  does not exist in document "DefaultDoc" at line: 11 column: 38`,
		//	`Type "Person" does not implement interface "NamedEntity", missing  "name2"`, // error is caught in phase 3 which is aborted because of resolve error
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func Test4a4x(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
  name2: Int2

}
interface ValuedEntity {
  value: Int
  value2: FLoat
  value3: Boolean
  value4: Bool
}

type Person implements NamedEntity & ValuedEntity {
  name: String
  age: In
}
`

	var expectedErr []string = []string{
		`"Int2" does not exist in document "DefaultDoc" at line: 4 column: 10`,
		`"Bool" does not exist in document "DefaultDoc" at line: 11 column: 11`,
		`"FLoat" does not exist in document "DefaultDoc" at line: 9 column: 11`,
		`"In" does not exist in document "DefaultDoc" at line: 16 column: 8`,
	}

	err := db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Int2")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Bool")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("In")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplements4a(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
  name2: Int2

}
interface ValuedEntity {
  value: Int
  value2: FLoat
  value3: Boolean
  value4: Bool
}

type Int2 {
	x: Int
}

type In {
	Age: Int
}

type Bool {
	z: Boolean
}


type Person implements NamedEntity & ValuedEntity {
  name: String
  age: Int
  value: Int
}
`

	var expectedErr []string = []string{
		`Type "Person" does not implement interface "NamedEntity", missing  "name2"`,
		`Type "Person" does not implement interface "ValuedEntity", missing  "value2" "value3" "value4"`,
		`"FLoat" does not exist in document "DefaultDoc" at line: 9 column: 11`,
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplement6x(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
}

interface ValuedEntity {
  value: Int
}

type Person implements NamedEntity {
  name: String
  age: Int
}

type Business implements NamedEntity & ValuedEntity & NamedEntity {
  name: String
  value: Int
  employeeCount: Int
}
`
	var expectedErr []string = []string{
		`Duplicate interface name at line: 15 column: 55`,
	}
	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplements6x(t *testing.T) {

	input := `
	
	type Business implements NamedEntity & ValuedEntity {
	  name: [[String!]!]!
	  value: Int
	  employeeCount: Int
	}
	
	interface NamedEntity {
	  name: [[String!]!]!
	}
	
	type Person implements NamedEntity {
	  name: [[String!]!]!
	  age: Int
	}
	
	interface ValuedEntity {
	  value: Int
	}

	`
	// replace entities with their above definitions.
	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	// expectedDoc := `type Business implements NamedEntity & ValuedEntity {name:[[String!]!]!value:Intemployee Count:Int}
	// 				interface NamedEntity {name:[[String!]!]!}
	// 				type Person implements NamedEntity{name:[[String!]!]! age:Int}
	// 				interface ValuedEntity {value:Int}`

	l := lexer.New(input)
	p := New(l)
	d, errs := p.ParseDocument()
	//fmt.Println(d.String())
	if len(errs) > 0 {
		t.Errorf("Unexpected, should be 0 errors, got %d", len(errs))
		for _, v := range errs {
			t.Errorf(`Unexpected error: %s`, v.Error())
		}
	}
	if compare(d.String(), input) {
		t.Errorf("Got:      [%s] \n", trimWS(d.String()))
		t.Errorf("Expected: [%s] \n", trimWS(input))
		t.Errorf(`Unexpected: program.String() wrong. `)
	}

}

func TestImplements6a(t *testing.T) {

	input := `
	interface NamedEntity {
	  name: [[String!]!]!
	}

	interface ValuedEntity {
	  value: Int
	}

	type Person implements NamedEntity {
	  name: [[String!]!]!
	  age: Int
	}

	type Business implements NamedEntity & ValuedEntity {
	  name: [[String!]]!
	  value: Int
	  employeeCount: Int
	}
	`

	var expectedErr [1]string
	expectedErr[0] = `Type "Business" does not implement interface "NamedEntity", missing "name"`

	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplementsInterfaceBadKeyword(t *testing.T) {

	input := `
	interfacei NamedEntity6b {
	  name: String!
	}

	interface ValuedEntity6b {
	  value: Int
	}

	type Person6b implements NamedEntity6b {
	  name:  String!
	  age: Int
	}

	type Business6b implements NamedEntity6b & ValuedEntity6b {
	  name:  String!
	  employeeCount: Int
	}
	`

	var expectedErr [1]string
	expectedErr[0] = `Parse aborted. "interfacei" is not a statement keyword at line: 2, column: 2`

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplements7(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
}

interface ValuedEntity {
  value: Int
}

type Person implements NamedEntity {
  name: String
  age: Int
}

type Business implements NamedEntity & ValuedEntity {
  name: String
  value: Int
  employeeCount: Int
}
`
	expectedDoc := `type Business implements NamedEntity & ValuedEntity {name:String value:IntemployeeCount:Int}
				 interface NamedEntity {name:String} 
				 type Person implements NamedEntity {name:String age:Int} 
				 interface ValuedEntity {value:Int}`

	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	d, errs := p.ParseDocument()
	//fmt.Println(d.String())
	if len(errs) > 0 {
		t.Errorf("Unexpected, should be 0 errors, got %d", len(errs))
		for _, v := range errs {
			t.Errorf(`Unexpected error: %s`, v.Error())
		}
	}
	if compare(d.String(), expectedDoc) {
		t.Errorf("Got:      [%s] \n", trimWS(d.String()))
		t.Errorf("Expected: [%s] \n", trimWS(expectedDoc))
		t.Errorf(`Unexpected: program.String() wrong. `)
	}
}

func TestImplementsNotAllFieldsx(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
  
}

interface ValuedEntity {
  value: Int
  size: String
  length: Float
}

type Person implements NamedEntity {
  name: String
  age: Int
}

type Business implements & NamedEntity & ValuedEntity {
  name: String
  value: Int
  employeeCount: Int
}
`

	var expectedErr [1]string
	expectedErr[0] = `Type "Business" does not implement interface "ValuedEntity", missing "size" "length" `

	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Person")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	fmt.Println("Error Count: ", len(errs))
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplementsNotAllFields2(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
  XXX: Boolean
  
}

interface ValuedEntity {
  value: Int
  size: [String]
  length: Float
}

type Person implements NamedEntity {
  name: String
  age: Int
}

type Business implements & NamedEntity & ValuedEntity {
  name: String
  value: Int
  length: String
  employeeCount: Int
}

type Business2 implements & NamedEntity & ValuedEntity {
  name: String
  XXX: Boolean
  size: String
  length: Float
  value: Int
  employeeCount: Int
}
`
	expectedErr := []string{
		`Type "Person" does not implement interface "NamedEntity", missing  "XXX"`,
		`Type "Business" does not implement interface "NamedEntity", missing  "XXX"`,
		`Type "Business" does not implement interface "ValuedEntity", missing  "size" "length"`,
		`Type "Business2" does not implement interface "ValuedEntity", missing  "size"`,
	}
	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business2")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplementsNotAllFields3(t *testing.T) {

	input := `
interface NamedEntity {
  name: String
  XXX: Boolean
  
}

interface ValuedEntity {
  value: Int
  size: String
  length: Float
}

type Person implements NamedEntity {
  name: String
  age: [[Int!]]!
}

type Business implements & NamedEntity & ValuedEntity {
  name: String
  value: Int
  length: String
  employeeCount: Int
}

type Business2 implements & NamedEntity & ValuedEntity {
  name: String
    age: [[Int!]]!
  XXX: Boolean
  size: String
  length: Float
  value: Int
  employeeCount: Int
}
`
	var expectedErr [3]string
	expectedErr[0] = `Type "Person" does not implement interface "NamedEntity", missing  "XXX"`
	expectedErr[1] = `Type "Business" does not implement interface "NamedEntity", missing  "XXX"`
	expectedErr[2] = `Type "Business" does not implement interface "ValuedEntity", missing  "size" "length"`

	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business2")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}

func TestImplemenSetupFragments(t *testing.T) {

	input := `
enum Episode {
  NEWHOPE
  EMPIRE
  JEDI
}

type Starship {
  id: ID!
  name: String!
  length(unit: LengthUnit = METER): Float
}

interface Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
}

type Human implements Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
  starships: [Starship]
  totalCredits: Int
}

type Droid implements Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
  primaryFunction: String
}

enum LengthUnit{
METER
CENTERMETER
MILLIMETER
KILOMETER
}
`

	expectedErr := []string{}
	err := db.DeleteType("NamedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("ValuedEntity")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}
	err = db.DeleteType("Business2")
	if err != nil {
		t.Errorf(`Not expected Error =[%q]`, err.Error())
	}

	l := lexer.New(input)
	p := New(l)
	_, errs := p.ParseDocument()
	for _, v := range errs {
		fmt.Println("Err: ", v)
	}
	for _, ex := range expectedErr {
		found := false
		for _, err := range errs {
			if trimWS(err.Error()) == trimWS(ex) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Expected Error = [%q]`, ex)
		}
	}
	for _, got := range errs {
		found := false
		for _, exp := range expectedErr {
			if trimWS(got.Error()) == trimWS(exp) {
				found = true
			}
		}
		if !found {
			t.Errorf(`Unexpected Error = [%q]`, got.Error())
		}
	}
}
