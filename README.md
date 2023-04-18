## Resrouces

* How to build blockchain from scratch (POW): https://builtin.com/blockchain/create-your-own-blockchain
* How to build blockchain from scratch (POS): https://levelup.gitconnected.com/build-a-proof-of-stake-blockchain-in-go-a765cb217d35

## Troubleshoot

* How to build/run Go outside of GOPATH: https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire

## Private notes

*Don't try to think hardly these notes. It's just my though to make it easy to understand code*

* In Golang, we don't have class but we can use struct as a class.
* And we use Pointer receiver and Value receiver as an alias to simulates a class by struct.
* See blow example:

```
type Person struct {
    Name string
    Age  int
}

func (p *Person) String() string {
    return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

//Pointer receiver: can change the value of object
func (p *Person) IncreaseAge() {
    p.Age += 1
}

//Value receiver: can not change the value of object
func (p Person) NotReallyIncreaseAge() {
    p.Age += 1 //warning: ineffective assignment to field Person.Age
}
```

* See more here (Vietnamese version): https://techmaster.vn/posts/36540/pointer-va-receiver-trong-golang

* Functions these start with uppercase is public access, start with lowercase is private access