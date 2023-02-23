# Go-Lang 연습 repo

## 1.1 Packages and Imports
### method 이름이 대문자로 시작하는 경우와 소문자로 시작하는 경우가 존재
  - 대문자로 시작하는 method: Java 에서 public method 와 유사 -> export 가능한 method(다른 파일에서 사용 가능)
  - 소문자로 시작하는 method: Java 에서 private method 와 유사 -> export 불가능한 method(다른 파일에서 사용 불가능)

## 1.2 Variables and Constants
### Const(상수)
- 변하지 않는 값 -> Java의 final
### Variable(변수)
- Golang 은 변수의 타입을 지정해주어야 한다


### 선언방법
#### 타입을 직접 지정해주는 방법
```go 
    var name string = "jihye" 
```
#### 축약시킨 코드
```go 
    name := "jihye"
```
- 축약된 방식으로 변수를 선언하면 go 가 자동으로 type 을 찾아주지만, 정해진 type 을 임의로 변경할 수 없다

## 1.3-4 Functions
> Go 는 생성된 변수 또는 함수가 사용되지 않으면 컴파일 에러를 발생시킨다

### Function
- multiple value 를 반환하는 function  에서 하나의 값만 return 받으면 에러  
    -> 하나를 생략하고 싶다면 '_'(underscore) 를 넣어준다

### Argument
- 반복하는 Function
- 여러개의 인자값을 한번에 받을 때 사용
- 인자의 타입 앞에 '...' 추가하면 배열처럼 사용할 수 있음

### NAKED return
- return 할 변수를 명시하지 않고 return 하는 것

### defer
- function 이 끝났을 때 추가적으로 다른 동작을 더 할 수 있게 해줌  
  ex) 함수 실행 후 API 호출, 파일 생성 후 열기 등등

## 1.4 for, range, ...args
- Go 에서 loop 는 for 만 사용이 가능하다
- `.forEach`, `.map`, `for in` 모두 사용 불가
- `range` 는 오로지 `for` 안에서만 동작
    > #### range
    > - `range` 는 배열에 loop 를 적용할 수 있도록 한다
    > - `range` 는 값을 주는것이 아닌 배열의 index 를 반환 한다

## 1.5 If with a Twist
### if, else if, else
- Go 에서는 소괄호가 없다
```go
if 조건 {
	소스코드
}
```
- 조건에 변수선언이 가능하다
```go
if 변수선언; 조건 {
	소스코드
}
```

## 1.6 Switch
- 소괄호 없이 사용한다
```go
switch 변수 {
case 조건1:
	return 
case 조건2:
	return
}
```
- `if-else` 나 `if-else if` 가 난무하는 경우를 피할 수 있다

## 1.7 Pointers
- 변수 앞에 `&` 를 붙여주면 해당 변수의 메모리값을 알 수 있다(메모리 주소 확인 가능)
- 메모리 주소가 저장된 변수 앞에 `*` 을 붙여주면 메모리에 저장된 값을 볼 수 있다
- 메모리에 저장된 object 를 복사하고 싶다면 메모리 주소를 저장해주어야 한다

```go
a := 2
b := &a // 메모리 주소를 저장하려면 &
fmt.Println(*b) // 메모리에 저장된 값을 보고 싶으면 *
*b = 10 // 메모리에 저장된 값을 변경하고 싶다면 * 로 값을 불러옴 (*b == a)
```

## 1.8 Arrays and Slices
### Array
- Go 에서 배열을 생성하려면 크기와 타입을 지정해주어야한다
```go
변수이름 := [배열 크기] 배열 타입 {구성 요소}
```
- 구성요소를 추가할 때 배열 크기를 넘어설 수 없다
### Slice
- 크기를 지정하지 않는 배열
```go
변수이름 := [] 배열타입 {구성 요소}
```
- 배열크기와 상관없이 구성요소 추가 가능
```go
slice변수 = append(slice변수, 추가하고 싶은 값)
```
- append 함수는 slice 와 값을 받아서 새로운 slice 변수를 반환한다
