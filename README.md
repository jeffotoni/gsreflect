# gsreflect
Taking the tag of a struct and allowing it to be default 

## Install

```
go get github.com/jeffotoni/gsreflect
```

## Example

```go

// struct config
type Config struct {
	Domain  string  `default:"s3go.gov"`
	Schema  string  `default:"https"`
	Port    int     `default:"9002"`
	Port2   int64   `default:"324564566675"`
	Db      string  `default:"postgresql"`
	Cluster int32   `default:"10"`
	Cost    float32 `default:"98.56"`
	Unit    bool    `default:"false"`
	Passwd  string  `default:"x37c$%2"`
	Login   string  `default:"postgres"`
}

// example
func main() {

	// declare confg
	c := Config{}

	// load value default
	// on struct
	err := gsr.Default(&c)

	// treat or error
	if err != nil {

		fmt.Println("error: ", err)
		return
	}

	// print screen
	fmt.Println("Domain: ", c.Domain)
	fmt.Println("Schema: ", c.Schema)
	fmt.Println("Port: ", c.Port)
	fmt.Println("Port2: ", c.Port2)
	fmt.Println("Db: ", c.Db)
	fmt.Println("Cluster: ", c.Cluster)
	fmt.Println("Cost: ", c.Cost)
	fmt.Println("Unit: ", c.Unit)
	fmt.Println("Passwd: ", c.Passwd)
	fmt.Println("Login: ", c.Login)
}
```

## Contributing

- Fork the repo on GitHub
- Clone the project to your own machine
- Create a *branch* with your modifications `git checkout -b gsr-cool`.
- Then _commit_ your changes `git commit -m 'contributing to a nice improvement'`
- Make a _push_ to your _branch_ `git push origin gsr-cool`.
- Submit a **Pull Request** so that we can review your changes and update them
