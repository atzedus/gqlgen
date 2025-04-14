// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package singlefile

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Animal interface {
	IsAnimal()
	GetSpecies() string
	GetSize() *Size
}

type ContentChild interface {
	IsContentChild()
}

type Mammalian interface {
	IsAnimal()
	IsMammalian()
	GetSpecies() string
	GetSize() *Size
}

type TestUnion interface {
	IsTestUnion()
}

type A struct {
	ID string `json:"id"`
}

func (A) IsTestUnion() {}

type AIt struct {
	ID string `json:"id"`
}

type AbIt struct {
	ID string `json:"id"`
}

type B struct {
	ID string `json:"id"`
}

func (B) IsTestUnion() {}

type Cat struct {
	Species  string `json:"species"`
	Size     *Size  `json:"size"`
	CatBreed string `json:"catBreed"`
}

func (Cat) IsAnimal()               {}
func (this Cat) GetSpecies() string { return this.Species }
func (this Cat) GetSize() *Size     { return this.Size }

type CheckIssue896 struct {
	ID *int `json:"id,omitempty"`
}

type ContentPost struct {
	Foo *string `json:"foo,omitempty"`
}

func (ContentPost) IsContentChild() {}

type ContentUser struct {
	Foo *string `json:"foo,omitempty"`
}

func (ContentUser) IsContentChild() {}

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type DefaultInput struct {
	FalsyBoolean  *bool `json:"falsyBoolean,omitempty"`
	TruthyBoolean *bool `json:"truthyBoolean,omitempty"`
}

type DefaultParametersMirror struct {
	FalsyBoolean  *bool `json:"falsyBoolean,omitempty"`
	TruthyBoolean *bool `json:"truthyBoolean,omitempty"`
}

type DeferModel struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type Dog struct {
	Species  string `json:"species"`
	Size     *Size  `json:"size"`
	DogBreed string `json:"dogBreed"`
}

func (Dog) IsAnimal()               {}
func (this Dog) GetSpecies() string { return this.Species }
func (this Dog) GetSize() *Size     { return this.Size }

type EmbeddedDefaultScalar struct {
	Value *string `json:"value,omitempty"`
}

type FieldsOrderPayload struct {
	FirstFieldValue *string `json:"firstFieldValue,omitempty"`
}

type Horse struct {
	Species    string `json:"species"`
	Size       *Size  `json:"size"`
	HorseBreed string `json:"horseBreed"`
}

func (Horse) IsMammalian()            {}
func (this Horse) GetSpecies() string { return this.Species }
func (this Horse) GetSize() *Size     { return this.Size }

func (Horse) IsAnimal() {}

type InnerDirectives struct {
	Message string `json:"message"`
}

type InnerInput struct {
	ID int `json:"id"`
}

type InnerObject struct {
	ID int `json:"id"`
}

type InputDirectives struct {
	Text          string           `json:"text"`
	NullableText  *string          `json:"nullableText,omitempty"`
	Inner         *InnerDirectives `json:"inner"`
	InnerNullable *InnerDirectives `json:"innerNullable,omitempty"`
	ThirdParty    *ThirdParty      `json:"thirdParty,omitempty"`
}

type InputWithEnumValue struct {
	Enum EnumTest `json:"enum"`
}

type LoopA struct {
	B *LoopB `json:"b"`
}

type LoopB struct {
	A *LoopA `json:"a"`
}

// Since gqlgen defines default implementation for a Map scalar, this tests that the builtin is _not_
// added to the TypeMap
type Map struct {
	ID string `json:"id"`
}

type Mutation struct {
}

type NestedInput struct {
	Field Email `json:"field"`
}

type NestedMapInput struct {
	Map map[string]interface{} `json:"map,omitempty"`
}

type ObjectDirectives struct {
	Text         string   `json:"text"`
	NullableText *string  `json:"nullableText,omitempty"`
	Order        []string `json:"order"`
}

type OmittableInput struct {
	ID     graphql.Omittable[*string]     `json:"id,omitempty"`
	Bool   graphql.Omittable[*bool]       `json:"bool,omitempty"`
	Str    graphql.Omittable[*string]     `json:"str,omitempty"`
	Int    graphql.Omittable[*int]        `json:"int,omitempty"`
	Time   graphql.Omittable[*time.Time]  `json:"time,omitempty"`
	Enum   graphql.Omittable[*Status]     `json:"enum,omitempty"`
	Scalar graphql.Omittable[*ThirdParty] `json:"scalar,omitempty"`
	Object graphql.Omittable[*OuterInput] `json:"object,omitempty"`
}

type OuterInput struct {
	Inner *InnerInput `json:"inner"`
}

type OuterObject struct {
	Inner *InnerObject `json:"inner"`
}

type Pet struct {
	ID      int    `json:"id"`
	Friends []*Pet `json:"friends,omitempty"`
}

type Query struct {
}

type Size struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type Slices struct {
	Test1 []*string `json:"test1,omitempty"`
	Test2 []string  `json:"test2,omitempty"`
	Test3 []*string `json:"test3"`
	Test4 []string  `json:"test4"`
}

type SpecialInput struct {
	Nesting *NestedInput `json:"nesting"`
}

type Subscription struct {
}

type User struct {
	ID      int        `json:"id"`
	Friends []*User    `json:"friends"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated,omitempty"`
	Pets    []*Pet     `json:"pets,omitempty"`
}

type ValidInput struct {
	Break       string `json:"break"`
	Default     string `json:"default"`
	Func        string `json:"func"`
	Interface   string `json:"interface"`
	Select      string `json:"select"`
	Case        string `json:"case"`
	Defer       string `json:"defer"`
	Go          string `json:"go"`
	Map         string `json:"map"`
	Struct      string `json:"struct"`
	Chan        string `json:"chan"`
	Else        string `json:"else"`
	Goto        string `json:"goto"`
	Package     string `json:"package"`
	Switch      string `json:"switch"`
	Const       string `json:"const"`
	Fallthrough string `json:"fallthrough"`
	If          string `json:"if"`
	Range       string `json:"range"`
	Type        string `json:"type"`
	Continue    string `json:"continue"`
	For         string `json:"for"`
	Import      string `json:"import"`
	Return      string `json:"return"`
	Var         string `json:"var"`
	Underscore  string `json:"_"`
}

// These things are all valid, but without care generate invalid go code
type ValidType struct {
	DifferentCase      string `json:"differentCase"`
	DifferentCaseOld   string `json:"different_case"`
	ValidInputKeywords bool   `json:"validInputKeywords"`
	ValidArgs          bool   `json:"validArgs"`
}

type XXIt struct {
	ID string `json:"id"`
}

type XxIt struct {
	ID string `json:"id"`
}

type AsdfIt struct {
	ID string `json:"id"`
}

type IIt struct {
	ID string `json:"id"`
}

type EnumTest string

const (
	EnumTestOk EnumTest = "OK"
	EnumTestNg EnumTest = "NG"
)

var AllEnumTest = []EnumTest{
	EnumTestOk,
	EnumTestNg,
}

func (e EnumTest) IsValid() bool {
	switch e {
	case EnumTestOk, EnumTestNg:
		return true
	}
	return false
}

func (e EnumTest) String() string {
	return string(e)
}

func (e *EnumTest) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnumTest(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnumTest", str)
	}
	return nil
}

func (e EnumTest) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *EnumTest) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e EnumTest) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}

type Status string

const (
	StatusOk    Status = "OK"
	StatusError Status = "ERROR"
)

var AllStatus = []Status{
	StatusOk,
	StatusError,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusOk, StatusError:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *Status) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e Status) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}
