
# Sample project on how to debug Go

## Delve

### Debug application

1. Compile and run the application in "debug mode"

		```
		dlv debug
		```

2. Follow instructions to navigate

### Debug test

1. Compile and run the test in "debug mode"

		```
		dlv test ./domain/
		```

2. Follow instructions to navigate

---

## VS Code

1. Open the project into VS Code using cli (or manually)

		```
		cd go-debug/
		code .
		```

2. Open a file and place a breakpoint

3. Click on the left icon "Debug"

4. Choose the mode "Debug" from the top dropdown

---

## Links

* https://github.com/go-delve/delve
* https://github.com/go-delve/delve/blob/master/Documentation/cli/getting_started.md
* https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code
* https://scotch.io/tutorials/debugging-go-code-with-visual-studio-code
