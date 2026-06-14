package main

import "fmt"
import "os"
import "encoding/csv"

func makeCSV(row int, column int) [][]string {
	var csv [][]string;
	for i := 0; i < row; i++ {
		csv = append(csv, make([]string, column));

	}

	return csv;
}

func printCSV(csv [][]string) {
	for i := 0; i < len(csv); i++ {
		for j := 0; j < len(csv[i]); j++ {
			fmt.Print(csv[i][j] + " ");
		}
		fmt.Print("\n");
	}
}

func openCSV(fileName string, create bool) {
	var csvData [][]string;
	if create {
		fmt.Print("rows: ");
		var row int;
		fmt.Scanln(&row);
		fmt.Print("columns: ");
		var column int;
		fmt.Scanln(&column);
		csvData = makeCSV(row, column);
	} else {
		file, err := os.Open(fileName);
		defer file.Close();
		if err != nil {
			fmt.Println("Error opening file: ", err);
			return;
		}	

		reader := csv.NewReader(file);
		csvData, err = reader.ReadAll();
		if err != nil {
			fmt.Println("Error reading file: ", err);
			return;
		}

		var running = true;
		for running {
			printCSV(csvData);
			fmt.Print("command(exit,edit,save): ");
			var command string;
			fmt.Scanln(&command);

			switch command {
			case "exit":
				running = false;
			case "edit":
				fmt.Print("row: ");
				var row int;
				fmt.Scanln(&row);
				fmt.Print("column: ");
				var column int;
				fmt.Scanln(&column);
				fmt.Print("value: ");
				var value string;
				fmt.Scanln(&value);
				csvData[row][column] = value;
			case "save":
				file, err := os.Create(fileName);
				if err != nil {
					fmt.Println("Error creating file: ", err);
					return;
				}

				writer := csv.NewWriter(file);
				err = writer.WriteAll(csvData);
				if err != nil {
					fmt.Println("Error writing file: ", err);
					return;
				}
				writer.Flush();
				fmt.Println("File saved successfully.");
			}
		}

	}

	}


func main() {
	fmt.Println("csvEditor");
	var running bool = true;
	for running {
		fmt.Print("command(exit,create,open): ");
		var command string;
		fmt.Scanln(&command);

		switch command {
		case "exit":
			running = false
		case "create":
			
			openCSV("new.csv", true);
		case "open":
			fmt.Print("file name: ");
			var fileName string;
			fmt.Scanln(&fileName);
			openCSV(fileName, false);
		}
	}
	
}