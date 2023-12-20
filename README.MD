## Linear Regression and Pearson Correlation Coefficient

This Go program calculates the linear regression line and Pearson correlation coefficient for a given dataset.

### Usage

1. Make sure you have Go installed on your system.
2. Compile the program by running the following command:

   ````shell
   go build main.go
   ```

3. Run the program with a file path as an argument:

   ````shell
   ./main <file_path>
   ```

   Replace `<file_path>` with the path to your data file.

### Input

The program expects a file containing a list of floating-point numbers, with each number on a separate line.

### Output

The program calculates the linear regression line and Pearson correlation coefficient based on the input data and prints the results to the console.

Example output:

```
Linear Regression Line: y = 0.123456x + 0.654321
Pearson Correlation Coefficient: 0.9876543210
```

### Algorithm

The program uses the following algorithms to perform the calculations:

- **Linear Regression**: The program calculates the slope and intercept of the linear regression line using the least squares method.
- **Pearson Correlation Coefficient**: The program calculates the Pearson correlation coefficient, which measures the linear relationship between two variables.

### File Reading

The program reads the data file specified as a command-line argument and processes the data accordingly. The file should contain one floating-point number per line.

### Dependencies

The program does not have any external dependencies.

### License

This project is licensed under the [MIT License](LICENSE).

### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

### Author

The code in this repository was written by [Hussain Alboori](https://github.com/hussainalboori).

Feel free to customize the content based on your preferences and add your own information.