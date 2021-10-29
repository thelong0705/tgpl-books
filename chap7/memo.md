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
  -