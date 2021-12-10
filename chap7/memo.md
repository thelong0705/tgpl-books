# Interfaces as contracts

- concrete types
    - specifes the exact representation
    - show the essential operations for that representation
        - arithmetic for number
        - indexing append range for slices
    - a concrete type provide additional behaviors through methods
    - you know what it is and what you can do with a concrete types

- interface type
    - abstract
    - does not expose the representation or internal structure
    - reveals only some of their methods
    - know nothing about what it is only know about what it can do

- fmt.Printf vs fmt.Sprintf
    - fmt.Printf print to the standard output
    - fmt.Spritf print to a string
    - share the same format coe
    - both are wrappers around a third function fmt.Fprintf

- fmt.Fprintf
    - func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
    - F stands for file and indicates formatted output should be written to the file
        - printf case the os.Stdout is an osFile
        - sprintf case the argument is not a file but it resembles one:
          pointer to a memory which bytes can be written
    - first params of FprintF is not a file either. It 's an io.Writer
        - has Write(p []byte) (n int, err error) method
            - write len(p) bytes from p to the underlying data stream
            - returns the number of bytes written from p

- io.Writer
    - define the contract between Fprintf and its callers
    - requires the caller provide a value that has method write
    - guarantees Fprintf will do its job give any values satisfies io.writer

- substitutability
    - freedom to subsitute one type for another that satisfies the same interface

# interface types

- an interface types specifies a set of methods that a concrete type must possess to be considered an instance of that
  interface
- some used interfaces
    - io.writer
    - io.reader
    - io.closer
- declaration of new interface types as combinations of existing ones

```go
type ReadWriter interface {
    Reader
    Writer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

- like struct embedding let us name another interface as a short hand for writing out all of its methods
- it s called embedding an interface

# interface satisfaction
- a type statisfies an interface if it possess all the methods the interface requires
- a concrete type is a interface meaning that it satisfies the the interface
- assignability rule for interfaces is very simple: an expression maybe assigned to an interface only if its type satisfies the interface
- concrete type T
  - has some method with type T
  - has some method with type *T
  - it s legal to call a *T method on variable type T
    - syntactic sugar: value of type T does not have all the method that a *T poiter does
    - might satisfy fewer interfaces
```go
type IntSet struct { /* ... */ }
func (*IntSet) String() string
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

var s IntSet
var _ = s.String() // OK: s is a variable and &s has a String method

```
- however since only *Intset has a string method only *Intset satisfied the fmt.Stringer interface
- `godoc -analysis=type` displays the method of each type and the relationship between interfaces and concrete types
- an intercace wraps and conceals the concrete type and value that it holds
  - only the methods revealed by the interface type may be called even if the concrete type has others
- an interface with more methods demand more on the type that implement it
  - empty interface {} has no methods dont demands on the types that satisfy it
    - we can assign any value to the empty interface
    - interface{} contain any type we can do nothing to the value it holds
- non empty interface types such as io.Writer are most often satisfied by a pointer type particurlaly when one of or more of the interface methods implies some kind of mutation to the receivers as the write method does
- but not only pointer type can satisfy interface
- an concrete type may satisfy many unrelated interfaces
- interfaces are but one useful way to group related concrete types together and express the facets they share in common.
- each grouping of concrete types based on their shared behaviors can be expressed as an interface type. Unlike class based languages in which of the set of interfaces satisfied by is explicit in Go we can define new abstractions or grouping of interest when we need them with out modify the declartion of the concrete type. This is particularly useful when the concrete type comes from a package written by a different author. Of course, there do need to be underlying commonalities in the concrete types.

