package main

// les types
//
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias pour uint8
//
// rune // alias pour int32
//      // représente un "code point" Unicode
//
// float32 float64
//
// complex64 complex128


import (
	"io"
	"fmt"
	"math/cmplx"
	"math"
	"strings"
	"time"
	"github.com/user/stringutil"
)

type Vertex struct {
	X int
	Y int
}

// genre de methode de class
type otherType struct {
	X, Y float64
}

type otherType2 struct {
	X, Y float64
}

func (v otherType) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (vv *otherType) Scale(f float64) {
	vv.X = vv.X * f
	vv.Y = vv.Y * f
	fmt.Println(vv.X, vv.Y)
}

func (v *otherType2) Abss() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (vv *otherType2) Scalee(f float64) {
	vv.X = vv.X * f
	vv.Y = vv.Y * f
	fmt.Println(vv.X, vv.Y)
}

// meme chose mais sur une variable simple
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x + 5
	return
}

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println("Hello world of go!")
	fmt.Println(stringutil.Reverse("Hello world of go!"))
	fmt.Println(add(21, 21))

	// ':='  est la version courte de declaratione et d'assignation
	// uniquement disponible dans les fonctions
	// en dehors il faudra utiliser 'var'
	a, b := "hello", "world"
	fmt.Println(a, b)
	fmt.Println(swap(a, b))

	// le return de split peut assigner 2 valeurs
	// comme si dessous x et y
	x, y := split(40)
	fmt.Println(x, y)


	// pour declarer une variable utiliser 'var' et definir un type
	var ff int
	fmt.Println(ff)
	// si on ne precise pas le type il sera determiné a l'assignation de la
	// valeur
	var i, j int = 1, 2
	// 					  bool  bool  string
	var c, python, java = true, true, "no!"
	fmt.Println(i, j, c, python, java)

	// %T => type  %v => value (les arguments sont passés dans le Println)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var it int
	var fl float64
	var bl bool
	var str string
	fmt.Printf("%v %v %v %q\n", it, fl, bl, str)

	// cast en un autre type
	d, e := 3, 4
	g := float64(e + 1)
	var h = uint(d + e)
	fmt.Println(d, e, g, h)

	var k int
	// l est de type int
	l := k
	m := 3.14
	fmt.Printf("l is of type %T\n", l)
	fmt.Printf("k is of type %T\n", k)
	fmt.Printf("m is of type %T\n", m)


	const Truth = true
		fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// voir le if du pow
	// a la maniere du for on peux initialiser une variable
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// defer ne sera executé que lorsque la fonction return
	defer fmt.Printf(" world\n")
	defer fmt.Printf(" new")
	defer fmt.Printf(" brave")
	defer fmt.Printf("hello")

	// les pointeurs fonctionnent comme en C
	var pr *int

	ft := 42
	pr = &ft
	ft = 21
	fmt.Println("[", *pr, "]")

	ver := Vertex{1, 2}
	ver.X = 4
	fmt.Println(ver.X)

	fmt.Println(v1, p, v2, v3)

	var aa [2]string
	aa[0] = "Hello"
	aa[1] = "World"
	fmt.Println(aa[0], aa[1])
	fmt.Println(aa)

	sss := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("sss ==", sss)

	for i := 0; i < len(sss); i++ {
		fmt.Printf("sss[%d] == %d\n", i, sss[i])
	}

	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)

	sl := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("sl ==", sl)
	fmt.Println("sl[1:4] ==", sl[1:4])

	// missing low index implies 0
	fmt.Println("sl[:3] ==", sl[:3])

	// missing high index implies len(s)
	fmt.Println("sl[4:] ==", sl[4:])

	asl := make([]int, 5)
	printSlice("asl", asl)
	bsl := make([]int, 0, 5)
	// bsl = append(bsl, 2)
	// bsl = append(bsl, 3)
	// bsl = append(bsl, 4)
	printSlice("bsl", bsl)
	bsl = append(bsl, 1, 2, 3)
	printSlice("bsl", bsl)
	csl := bsl[:2]
	printSlice("csl", csl)
	dsl := csl[1:4]
	printSlice("dsl", dsl)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
// range permet de recuperer l'index et la valeur
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	//  _ permet de sauter cet element
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	for index, _:= range pow {
		fmt.Printf("%d\n", index)
	}

	type Vertex struct {
		Lat, Long float64
	}

	var mk map[string]Vertex

// pour faire un tableau associatif (hash table)
	mk = make(map[string]Vertex)
	mk["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(mk["Bell Labs"])

	var mk2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
		// on peux ne pas preciser le type (vertex dans ce cas)
		"Apple": {40.68433, -74.39967},
		"kfc":	{37.42202, -122.08408},
	}
	fmt.Println(mk2)
	fmt.Println(mk2["Google"])
	fmt.Println(mk2["Google"].Lat)

	om := make(map[string]int)

	om["Answer"] = 42
	fmt.Println("The value:", om["Answer"])

	om["Answer"] = 48
	fmt.Println("The value:", om["Answer"])

	delete(om, "Answer")
	fmt.Println("The value:", om["Answer"])

// on peux verifier la presence d'une valeur en recuperant 2 valeurs
// ov sera la valeur et ok sera un bool renvoyant faux si vide
	ov, ok := om["Answer"]
	fmt.Println("The value:", ov, "Present?", ok)


	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(
				pos(i),
				neg(-2*i),
			)
		}

// faire appel a une sorte de methode associé a un type
// (voir lignes 36 a 42)
	vver := otherType{3, 4}
		fmt.Println(vver.Abs())
	// poitneur sur struct dans la methode (ligne 44)
	vver.Scale(10)
		fmt.Println("methode avec poiteur sur struct", vver.Abs())

	mf := MyFloat(-42.42)
	fmt.Println(mf.Abs())

	vv := &otherType2{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", vv, vv.Abss())
	vv.Scalee(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", vv, vv.Abss())

	// interface
	var itf I
	itf = &T{"hello printed by the interface"}
	describe(itf)
	// ce M prend un type T en argument
	itf.M()

	itf = F(math.Pi)
	describe(itf)
	// ce M prend un type F en argument
	// c'est le meme principe que le polymorphisme en C++
	itf.M()

	var itf_vide interface{}
	// ce describe prend un type interface vide
	describe(itf_vide)

	itf_vide = 42
	describe(itf_vide)

	itf_vide = "hello"
	describe(itf_vide)

	do(21)
	do("hello")
	do(true)

	apers := Person{"Arthur Dent", 42}
	zpers := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(apers, zpers)

	if err := run(); err != nil {
		fmt.Println(err)
	}

// strings reader exemple
	reader_func()

// channel system
	fmt.Println("basic channel")
	channel_sum()
	fmt.Println("channel with buffer")
	channel_buffered()
	fmt.Println("fibo by channel")
	make_fibo()
	fmt.Println("fibo by channel with select")
	fibo_with_select()
    fmt.Println("select works in asynchrone way")
    bomb_with_select()

    gorout()
}

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func gorout() {
    f("direct")
    go f("goroutine")
    go func(msg string) {
        fmt.Println(msg)
    }("going")
    fmt.Scanln()
    fmt.Println("done")
}

func bomb_with_select() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibo_with_select() {
	c := make(chan int)
	quit := make(chan int)
    // cette finction est utilisée de maniere asynchrone, le channel c
    // dans le go func sera initialisé dans le fibonacci_select.
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(":", <-c)
		}
		quit <- 0
	}()
	fibonacci_select(c, quit)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// ce close est due a l'utilisation d'un range sur le channel
	close(c)
}

func make_fibo() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// le range d'un channel demande forcement un close
	for i := range c {
		fmt.Println(i)
	}
}

func channel_buffered() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func channel_sum() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	sum := 0
	fmt.Println(a)
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func reader_func() {
	rdr := strings.NewReader("Hello, Reader!")

	brd := make([]byte, 8)
	for {
		n, err := rdr.Read(brd)
		fmt.Printf("%v bytes read,  err = %v b = %v\n", n, err, brd)
		fmt.Printf("%v b[:n] = %q\n",n,  brd[:n])
		if err == io.EOF {
			break
		}
	}
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printBoard(ss [][]string) {
	for i := 0; i < len(ss); i++ {
		fmt.Printf("%s\n", strings.Join(ss[i], " "))
	}
}
