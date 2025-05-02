# goPassM
Password manager written in Go.

![logo goPassM](logo.png)
## Description

I made this little project in order to discover the Go language concepts. This is a tool in CLI for now.

You can add and store a combo of service url/login/password in a encrypted file.

## Getting Started

### Dependencies

- Install Go see -> [Go site](https://go.dev/doc/install)

### Installing

```bash
git clone https://github.com/Vak00/goPassM
```

### Executing program

* How to run the program
* Step-by-step bullets

```bash
go run cmd/run.go
```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```

## Authors

[Vak00](https://github.com/Vak00)

## Roadmap

- [] Do the tests
- [] Put the vault file and .master file somewhere else + hide them
- [] Harmonize logs + better managment of the logs
- [] Use Interface instead of Struct for the Command
- [] Purpose random generated password for entries
- [] Add check for the master password (lenght + complexity)

### Minor
- [] Create an UI
- [] Create session manageemnt to avoid the user put password each time
- [] Try using biometric system
- [] Use Cobra (cli manager)
## Version History

// TODO
<!-- * 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release -->

## License

This project is licensed under the MIT License - see the LICENSE file for details

## Acknowledgments

Inspiration, code snippets, etc.
* [effective-go](https://go.dev/doc/effective_go)
