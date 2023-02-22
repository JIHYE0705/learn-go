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
