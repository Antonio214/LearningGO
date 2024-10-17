# Temperature Converter CLI

## Description

The **Temperature Converter CLI** is a simple command-line application written in Go that converts temperatures between Celsius and Fahrenheit. This project is designed to help you get acquainted with Go's basic syntax, data types, control structures, functions, and tooling.

## Goals

- **Understand Go Environment Setup**: Ensure Go is correctly installed and your workspace is configured.
- **Learn Basic Syntax and Data Types**: Work with variables, constants, and basic data types like `int`, `float64`, and `string`.
- **Utilize Control Structures**: Implement `if` statements and basic user input validation.
- **Define and Call Functions**: Create functions for temperature conversion and input handling.
- **Use Standard Library Packages**: Employ packages like `fmt` for input/output and `strconv` for string conversion.
- **Practice Basic Tooling**: Compile and run the program using `go run` and `go build`.

## Features

- Convert temperatures from Celsius to Fahrenheit and vice versa.
- User-friendly prompts and outputs.
- Input validation to handle incorrect entries.

## Prerequisites

- Go installed on your system (Go 1.16 or later recommended).
- A command-line interface (Terminal, Command Prompt, etc.).
- Basic knowledge of navigating the command line.

## Getting Started

### 1. Clone or Download the Project

You can start by creating a new directory for your project or cloning a repository if one is provided.

```bash
# Create a new directory
mkdir temperature-converter-cli
cd temperature-converter-cli

# Initialize a new Go module
go mod init temperature-converter-cli
```

### 2. Write the Code

Create a `main.go` file and start coding your temperature converter application. Implement the following:

- Prompt the user to select the conversion type.
- Accept a temperature value as input.
- Perform the conversion using appropriate formulas.
- Display the result to the user.

### 3. Run the Application

You can run the program directly without compiling using:

```bash
go run main.go
```

### 3.1 Test the application

You can test the application by running the following command:

```bash
go test
```

### 4. Build the Application

To compile the program into an executable, use:

```bash
go build
```

This command generates an executable file (`temperature-converter-cli` or `temperature-converter-cli.exe` on Windows) in your current directory.

### 5. Execute the Compiled Program

Run the compiled executable:

```bash
# On Linux or macOS
./temperature-converter-cli

# On Windows
temperature-converter-cli.exe
```

## Usage Instructions

1. **Select Conversion Type**

   When you run the program, you'll be prompted to choose the type of conversion:

   ```
   Temperature Converter
   ---------------------
   Select conversion type:
   1. Celsius to Fahrenheit
   2. Fahrenheit to Celsius
   Enter your choice (1 or 2):
   ```

2. **Enter Temperature Value**

   After selecting the conversion type, input the temperature value you wish to convert:

   ```
   Enter the temperature value to convert:
   ```

3. **View the Result**

   The program will display the converted temperature:

   ```
   Result: XX.X degrees Fahrenheit
   ```

4. **Sample Run**

   ```
   Temperature Converter
   ---------------------
   Select conversion type:
   1. Celsius to Fahrenheit
   2. Fahrenheit to Celsius
   Enter your choice (1 or 2): 1
   Enter the temperature value to convert: 100
   Result: 212 degrees Fahrenheit
   ```

## Project Structure

```
temperature-converter-cli/
│
├── go.mod              // Go module file
├── converter.go        // Temperature conversion functions
├── converter_test.go   // Test cases for temperature conversion
├── main.go             // Main application file
└── README.md           // Project documentation
```

## Important Notes

- **Input Validation**: Ensure your program can handle invalid inputs gracefully (e.g., non-numeric temperature values or invalid menu choices).
- **Functions**: Create separate functions for:
  - Displaying the menu and getting the user's choice.
  - Reading and validating the temperature input.
  - Performing the temperature conversion.
- **Constants**: Use constants where appropriate, such as for menu options or conversion formulas.

## Learning Outcomes

By completing this project, you will:

- Gain hands-on experience with Go's basic syntax and program structure.
- Learn how to read user input and display output.
- Understand how to work with control flow statements and functions.
- Get familiar with compiling and running Go programs.
- Use standard library packages effectively.

## Additional Resources

- [The Go Programming Language Documentation](https://golang.org/doc/)

## License

This project is open-source and available under the [MIT License](LICENSE).

## Conclusion

Enjoy building your Temperature Converter CLI! This foundational project will set the stage for more complex applications as you continue to learn Go.