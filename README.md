# GoLang_Basics

1. Go gives error if we import package and doesn't use it.
2. Variable starting with Capital letter means it is public else private 
3. no var style variable declaration

`<variable_name> := <value>`

4. **comma ok or err ok  syntax** - 
   1. In go there is no try catch bloack
   2. errors will be treated as variables
5. **go env** : tells about env variables of go

6. **Memory Management**
    |new| make|
    |---|-----|
    |Allocate memory but no init| Allocate memorty and INIT|
    |you will get a memory address| you will get a memory address|
    |zeroed storage| non-zerored storage|
7. Garbage happens automatically (Out of Scope or nil)
8. Pointers: Address of a variable
   1. var ptr *string //initialisation
   2. var ptr = &val. // pointer for already defined variable
9. Arrays: It's very basic, not much used in go lang
10. Slices(under the hood arrays)
11. Maps: key value pairs
12. Struct: equivalvent to classes(No inheritence, no super, no parent )
13. switch, if else - normal
   - switch have no break
   - It has fallthrough which can go to next case as well
14. Loops : only for loop is there
   - variation are more
   - goto(special thing - this acts as label)

15. Functions
16. Methods: When referencing methods by structs
17. Defer: 