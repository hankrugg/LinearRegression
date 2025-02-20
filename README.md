# Go Linear Regression

## Overview  
This project demonstrates a linear regression model implemented in Go. The model is trained on three different datasets, and the first 20 data points are visualized using a bar chart that compares actual vs. predicted values. The goal is to evaluate the model's accuracy across different datasets.

## Features  
- Implements linear regression from scratch in Go  
- Supports multiple datasets for comparison  
- Visualizes actual vs. predicted values using a bar chart  
- Outputs key regression metrics (e.g., mean squared error)  

## Installation  
Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).  

## Usage  
1. Clone the repository:  
   ```sh
   git clone https://github.com/hankrugg/LinearRegression.git
   cd LinearRegression
   ```
2. Run the program:  
   ```sh
   go run .
   ```
3. The program will:  
   - Load the datasets  
   - Train a linear regression model  
   - Display predicted vs. actual values in a bar chart  
   - Print regression metrics in the console  

## Future Enhancements  
- Implement multiple regression for more complex models  
- Optimize performance with Go concurrency  
- Add user-defined datasets and interactive visualization  
