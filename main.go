package main

import "fmt"
import "os"
import "encoding/csv"
import "io"

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


func workingWithCSV(openFile bool,fileName string, csvData [][]string) {
	// AI CODE

// 1. Change type to io.ReadWriteCloser so .Close() works on both types of files
    var file io.ReadWriteCloser
    var err error // 2. Declare err here at the top level

    if openFile {
        // 3. Use '=' instead of ':=' so we don't create temporary block variables
        file, err = os.Open(fileName)
        if err != nil {
            fmt.Println("Error opening file: ", err)
            return
        }
        defer file.Close()

        // 4. Move the reading logic inside here! We only read if the file already exists.
        reader := csv.NewReader(file)
        csvData, err = reader.ReadAll()
        if err != nil {
            fmt.Println("Error reading file: ", err)
            return
        }
    } else {
        // 5. Use '=' here too so it updates the outer 'file' variable
        file, err = os.Create(fileName)
        if err != nil {
            fmt.Println("Error creating file: ", err)
            return
        }
        defer file.Close()
        
        // If we just created the file, it's empty, so we skip reading.
        // Your code will just use the empty grid you made in makeCSV()!
    }

	// END AI CODE
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
		workingWithCSV(false, fileName, csvData);
	} else {
		workingWithCSV(true, fileName, csvData);
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