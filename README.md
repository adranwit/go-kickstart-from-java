# go-kickstart-from-java



- [Motivation](#motivation)
- [DataType](#datatype)
- [Code organization](#organization)
- [Pointer](#pointer)
- [Function](#function)
- [Struct](#struct)
- [Collection](#struct)



	
<a name="motivation"></a>
## Motivation

This project was developed to provide kick-start with go-lang for java developer.

[Go Playground](https://play.golang.org/)



<a name="datatype"></a>
## Data Type
This section provides the closest golang and java type, classes mapping.


### Basic data types    
 
 | Golang      | Java           |
 |-------------|----------------|
 |   int,  uint64, int64     | long           | 
 |   float64     | double           | 
 |   float32     | float           | 
 |   int,  uint32, int32, rune     | int           | 
 |   int,  uint16, int16     | short           | 
 |   uint8, int8,  byte       | byte          |  
 |   bool      | bool         |  
 |   string     | java.lang.String   | 
 
 
### Object type
  |   interface{} | java.lang.Object |



### Time type 
 | Golang      | Java           |
  |-------------|----------------|
 |   time.Time     | java.util.Date   | 
  
### Collection types    


 | Golang      | Java           |
 |-------------|----------------|
 |   map[k]v     | java.util.Map   | 
 |   []T  (slice)       | java.util.List   | 
 |   chan T (channel)    | java.util.concurrent.BlockingQueue<T>  | 
 |   [3]T (array)   | T[] | 
    
  
 ### Concurrent type
 
  | Golang      | Java           |
  |-------------|----------------|
  |   sync.Mutex,sync.RWMutex    | java.util.concurrent.locks.Lock, synchronized  | 
  |   sync.WaitGroup       | java.util.concurrent.CountDownLatch  | 
  | go func() {} (go routine)  |java.lang.Thread|





<a name="organization"></a>
## Code organization

```text
$GOPATH/
    src/
        github.com/viant/toolbox/
            toolbox/
                collection.go
                collection_test.go
                file_logger.go
            dsc/
                manager.go
    pkg/
           darwin_amd64
    bin/
```

### Packages

Go use directory structure to uniquely map corresponding pacakge. 
Within a single folder, the package name has to be same for the all source files 
which belong to that directory or for test file it can be postfixed with _test package.
We develop our Go programs in the $GOPATH directory, 
where we organize source code files into directories as packages. 


### Code scope

  |     Java  | Golang    |
  |-------------|----------------|
  | public    | global (exported) : upper case identifers | 
  |  package (default) | local (unexported): any lower case identiers | 
  | private/protected  | n/a                     | 
      
In Go packages, all identifiers will be exported to other packages if they tart with upper case identifer.


### Importing Third-Party Packages

#### Instalation   
        go get github.com/viant/toolbox

#### Import

```go

    import  "github.com/viant/toolbox"
    
```
`
### Standard library
   Go’s standard library comes with lot of useful packages.
    The packages from the standard library are available at the “pkg” subdirectory of the GOROOT directory.
       
   [Standard Library Package](https://golang.org/pkg/)

```go

    package main;
      
    import  "fmt" 
    import "net/http"
 

```

### Main Package

 The package “main” is used to direct the Go compiler that the package should compile as an executable program. 
 The main function in the package “main” will be the entry point of the program. 


[Run the code](https://play.golang.org/p/1Afks35zfX)

```go
    
    package main
    
    import ("fmt")
    
    
    func main(){
     fmt.Println("Hello, World!")
    }

```



### Declaration

### Assignment with initializers
```go
    
   var i, j int64 = 0, 1
   var k = "abc"

```


### Short assignment with implicit type.
```go


    
    i := 0 
    text := "abc";
    x1, x2, x3 := true, 0, "xyz"
    

```

### Zero values

Any variable of a basic data types without explicit values are 0 or false for boolean, or "" for string


```go

	var i int
	var b bool
	var s string
	fmt.Printf("Zero values %v %v %v %q\n", i,  b, s)

```


<a name="pointes"></a>
### Pointer
Go has pointers. A pointer holds the memory address of a value. 

**ampersand** operator generates a pointer to its operand operator generates a pointer to its operand
**star** operator dereferences the pointer to the underlying value. 
```go
i := 42
p := &i
*p++
```


<a name="function"></a>
### Function


In **Java** all primitive values are passed by value (a new copy of the variable is created and passed to called function or method)
, while all object variables are pased by references.

In golang all values are passed by value.


    
[Run the code](https://play.golang.org/p/wbWwAcEn3X)

```go

func incrementNumber(i int) {
	i++;
	fmt.Printf("Increment: %v\n", i);	
}


func main() {
	i := 0;
	fmt.Printf("Before Increment: %v\n", i);	
	incrementNumber(i);
	fmt.Printf("After Increment: %v\n", i);	
	
}

```

#### paassing pointer
In case a variable is passed by pointer, a new copy of pointer to the same memory address is created.


[Run the code](https://play.golang.org/p/s7eZq0g7eS)

```go


func incrementNumber(i *int) {
	*i++;
	fmt.Printf("Increment: %v\n", *i);	
}



func main() {
	i := 0;
	fmt.Printf("Before Increment: %v\n", i);	
	incrementNumber(&i);
	fmt.Printf("After Increment: %v\n", i);	
	
}

```



### deferred function

A defer statement defers the execution of a function until the surrounding function returns. 

In java the functionality is achieve by try/finally blocks

```java
    
    int i = 0;
    try {
           i++;
    
    } finally { //executes no matther if try throws exception or not
        System.out.println("i: " + i);
    }

```



[Run the code](https://play.golang.org/p/2C5VntoNFk)

```go

func incrementNumber(i *int) {
	*i++;
	defer func() {*i++}()
	fmt.Printf("Increment: %v\n", *i);	

}



func main() {
	i := 0;
	fmt.Printf("Before Increment: %v\n", i);	
	incrementNumber(&i);
	fmt.Printf("After Increment: %v\n", i);	
	
}

```


### init function

When go initialize packages it will look for function called "init" to call it at the beginning of the execution time.

https://play.golang.org/p/gPat9Ph0lh

```go

var i int

func init(){
	i++
}

func main() {
	i++;
	fmt.Printf("Number: %v\n", i)
}

```


### go routines

A goroutine is a lightweight thread of execution.


[Run the code](https://play.golang.org/p/xc27V8Kjed)

```go


func printNumber(i int) {
	fmt.Printf("Number: %v\n", i)
}

func main() {
	
	
	for i := 0; i<10;i++ {
		go printNumber(i);
	}
	fmt.Printf("AFTER LOOP\n");
	time.Sleep(1);
	
}


```
   
   
[Run the code](https://play.golang.org/p/xAsixxqEWE)
    
```go


func printNumber(i int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done();
	fmt.Printf("Number: %v\n", i)
	
}

func main() {
	
	waitGroup := &sync.WaitGroup{};
	waitGroup.Add(10);
	
	for i := 0; i<10;i++ {
		go printNumber(i, waitGroup);
	}
	fmt.Printf("AFTER LOOP\n");
	waitGroup.Wait();
	
}

```    

## Struct

A struct is a collection of fields and/or methods.

In java the nearest abstraction is a class.


### Java Class
```java

public class Student {

    private int id;
    private String name;
    private Date enrolmentDate;


    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Date getEnrolmentDate() {
        return enrolmentDate;
    }

    public void setEnrolmentDate(Date enrolmentDate) {
        this.enrolmentDate = enrolmentDate;
    }
}


```


### Go Struct

[Run the code](https://play.golang.org/p/XLORar-CBo)


```go

type Student struct {
	id int
	name string
	registrationDate *time.Time
}

func(s *Student) Id() int{
    return s.id;
}

func(s *Student) Name() string{
    return s.name;
}

func(s *Student) SetName(name string) {
    s.name = name;
}


func NewStudent(id int,name string, registrationDate *time.Time) *Student {
	return &Student{
		id:id,
		name:name,
		registrationDate:registrationDate,
	}
}


```

### Methods
In Go, a method is a function that is declared with a receiver. A receiver is a value or a pointer of a named or struct type. 
All the methods for a given type belong to the type’s method set.


## Inheritance

Go provides simple inheritence methanism called embedding. 

Embedding allows methods of one struct to get "promoted" to appear to exist on another struct.


[Run the code](https://play.golang.org/p/769gVrSy1x)

```go


type Person struct {
	id int
	name string
}

func(p *Person) Id() int{
    return p.id;
}

func(p *Person) Name() string{
    return p.name;
}

func(p *Person) SetName(name string) {
    p.name = name;
}

type Student struct {
	Person
	registrationDate *time.Time
}



func(s *Student) RegistrationDate() *time.Time {
    return s.registrationDate;
}

func NewStudent(id int,name string, registrationDate *time.Time) *Student {
	return &Student{
		Person: Person{
		id:id,
		name:name,
		},
		registrationDate:registrationDate,
	}
}



func main() {
	var now = time.Now();
	var student = NewStudent(1, "John Smith", &now);
	fmt.Printf("Student[id: %v, name:%v]\n", student.Id(), student.Name());

}

```

### Interfaces

Interfaces in Go are special types type that specifies a one or more methods. In order to a struct to be consider a specific interface, 
it has to provide all the method defined by this interface.
In Java, besides providing all interface methods, the class has to also use "implements" keyword to be consider a specific interface.


[Run the code](https://play.golang.org/p/dhf5hZk6hj)

```go


type Animal interface {
	Greeting() string
}

type Cat struct {}

func (c Cat) Greeting() string {
	return "meow"
}


type Dog struct {Animal}

func (d Dog) Greeting() string {
	return "bark"
}


func SayHello(animal Animal) {
	fmt.Printf("%T said %v\n", animal, animal.Greeting())
}

func main() {
     cat := &Cat{}
     dog := &Dog{}
     SayHello(dog)
     SayHello(cat);

}


```

#### Abstract clases

[abstract-classes-in-golang](http://blog.bingecoder.net/index.php/2016/06/20/abstract-classes-in-golang/)


## Collections
        
### Slices
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.


[Run the code](https://play.golang.org/p/1jzhVOweWm)
```go
    
       random := rand.New(rand.NewSource(time.Now().UnixNano()))
   
       var aSlice = make([]int, 1);
       
       for i:=0;i<10;i++ {
           aSlice = append(aSlice, random.Int())
       }
   
   
       fmt.Printf("First iteration\n")
       for index, item := range aSlice {
           fmt.Printf("[%v]: %v\n", index, item)
       }
   	
       fmt.Printf("Second iteration\n")
       for i:=0;i<10;i++ {
             fmt.Printf("[%v]: %v\n", i, aSlice[i])
     
       }

```

### Map
In Go a map is an unordered collection of key-value pairs. 
Maps are **not** thread safe.

[Run the code](https://play.golang.org/p/V92pz6AW87)


```go
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
        var aMap =  make(map[int]int)
    
       for i:=0;i<10;i+=2 {
           aMap[i]=random.Int()
       }
   
   
       for key, value := range aMap {
           fmt.Printf("[%v]: %v\n", key, value)
       }
   	
   	
   	
       elem, hasElem:= aMap[2]
       fmt.Printf("has: %v, elem: %v\n", elem, hasElem)
       
       delete(aMap, 2)
   
       elem, hasElem= aMap[2]
       fmt.Printf("has: %v, elem: %v\n", elem, hasElem)
      
	
	
```

#### Thread safe map

Secure read and write operation with mutex.

[Run the code](https://play.golang.org/p/p0VhTexYFo)
```go


type K string
type V string

type ConcurrentMap struct {
	*sync.RWMutex
	aMap map[K]V
}

func (m *ConcurrentMap) Put(key K, value V) {
	m.Lock()
	defer m.Unlock()
	m.aMap[key] = value
}


func (m *ConcurrentMap) Get(key K) V {
	m.RLock()
	defer m.RUnlock()
	return m.aMap[key]
}



func main() {
	aMap := &ConcurrentMap{
		&sync.RWMutex{},
		make(map[K]V),
	}
	
	aMap.Put("k1", "v1")
	aMap.Put("k2", "v2")
	
	fmt.Printf("k1: %v\n", aMap.Get("k1"))
	fmt.Printf("k2: %v\n", aMap.Get("k2"))

}

```

### Channel

Channels are special type, similarly to a blocking queues, it allows to send and receive data.


[Run the code]()https://play.golang.org/p/wPl0U3cHTM)

```go
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ch = make(chan int, 1)
	ch <- random.Int()    // Send v to channel ch.
	v := <-ch  // Receive from ch, and assign it to v
	fmt.Printf("Rand: %v\n", v)
```
          
          
#### select

The select statement allow go to wait on data appearing on multiple channels.


[Run the code](https://play.golang.org/p/I_aOaQjYY8)
```go

func consume(channel1 chan int, channel2 chan int, quit chan bool) {

for {
        var data int
        select {

        case data = <-channel1:
            fmt.Printf("Consuming data from channel1: %v\n", data)

        case data = <-channel2:
            fmt.Printf("Consuming data from channel2: %v\n", data)


        case <-quit:
            fmt.Println("quit")
            return
        }
    }

}



func main() {
	
	var channel1 = make(chan int)
	var channel2 = make(chan int)
	var quitNotification = make(chan bool)
	
	go consume(channel1, channel2, quitNotification)
	
	channel1 <- 100;
	channel1 <- 200;
	quitNotification <- true
	
	
} 


```
           
#### timeouts


[Run the code](https://play.golang.org/p/S9oC-3FsIb)

```go


func consume(channel1 chan int, channel2 chan int, quit chan bool) {

for {
        var data int
        select {

        case data = <-channel1:
            fmt.Printf("Consuming data from channel1: %v\n", data)

        case data = <-channel2:
            fmt.Printf("Consuming data from channel2: %v\n", data)

       case <-time.After(200 * time.Millisecond):
            fmt.Println("timeout")
      



        case <-quit:
            fmt.Println("quit")
            return
        }
    }

}



func main() {
	
	var channel1 = make(chan int)
	var channel2 = make(chan int)
	var quitNotification = make(chan bool)
	
	go consume(channel1, channel2, quitNotification)
	
	channel1 <- 100;
	channel1 <- 200;
	
	time.Sleep(600 * time.Millisecond);
	
	quitNotification <- true
	
	time.Sleep(100 * time.Millisecond);
	
	
} 


```


### Http



First HTTP server
```go

package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

```



## Web service


  https://github.com/viant/toolbox/blob/master/service_router.go  
  https://github.com/viant/toolbox/blob/master/service_router_test.go  
   
  Check app2

    
## Database 

    Check app3


### Vertica setup

1. Install ODBC manager http://www.odbcmanager.net/
2. Install Vertica driver https://my.vertica.com/client_drivers/7.1.x/vertica-odbc-7.1.2-0.mac.pkg
3. Change  set DriverManagerEncoding=UTF-16 in /Library/Vertica/ODBC/lib/vertica.ini 
4. Install go driver go get github.com/alexbrainman/odbc
5. Create/modify Vertica driver details /etc/odbcinst.ini   

```text
[ODBC Drivers]
Vertica=Installed
UsageCount=1

[Vertica]
Description=HP Vertica ODBC Driver
Driver=/Library/Vertica/ODBC/lib/libverticaodbc.dylib
Setup=/Library/Vertica/ODBC/lib/libverticaodbc.dylib
ErrorMessagesPath=/Library/Vertica/ODBC/messages/
DriverManagerEncoding=UTF-16
UsageCount=1
```

6. Install Vertica driver odbcinst -i -d -f /etc/odbcinst.ini   
7. Create dcc.Config configuration 

```text
{
  "DriverName": "odbc",
  "Descriptor": "driver=Vertica;Database=[database];ServerName=[server];port=[port];user=[user];password=[password]",
  "Parameters": {
    "user": "user",
    "password":"****",
    "database":"dw",
    "server":"10.55.1.181",
    "port":"5433"
  }
}

```

8. Add vertica  import	_ "github.com/alexbrainman/odbc"










