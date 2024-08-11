package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mirinjamamul/go-poc-api/database"
	"github.com/Mirinjamamul/go-poc-api/models"
	"github.com/xuri/excelize/v2"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Parse the incoming file
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Error parsing file:", err)
		http.Error(w, "Unable to parse file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	log.Printf("Uploaded file: %s\n", header.Filename)

	// Open the Excel file
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Println("Error opening Excel file:", err)
		http.Error(w, "Unable to open Excel file", http.StatusInternalServerError)
		return
	}

	// Get the first sheet name
	sheet := f.GetSheetName(4)
	if sheet == "" {
		log.Println("Sheet not found")
		http.Error(w, "Sheet not found", http.StatusInternalServerError)
		return
	}

	// Read the rows from the sheet
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Println("Error reading rows from Excel sheet:", err)
		http.Error(w, "Unable to read rows from Excel sheet", http.StatusInternalServerError)
		return
	}

	// Skip header row
	for i, row := range rows {
		if i == 0 {
			log.Println("Skipping header row")
			continue
		}

		// Log the current row number

		// Create user from row data
		user := models.User{
			Poc_Lead:                                getValue(row, 0),
			Point_of_contact:                        getValue(row, 1),
			Employee_Id:                             getValue(row, 2),
			Full_Name:                               getValue(row, 3),
			Company:                                 getValue(row, 4),
			Unit:                                    getValue(row, 5),
			Department:                              getValue(row, 6),
			Section:                                 getValue(row, 7),
			Job_Category:                            getValue(row, 8),
			Designation:                             getValue(row, 9),
			Employee_Type:                           getValue(row, 10),
			Work_Location:                           getValue(row, 11),
			Reporting_Employee_ID:                   getValue(row, 12),
			Reporting_Employee_Name:                 getValue(row, 13),
			Reporting_Designation:                   getValue(row, 14),
			Gender:                                  getValue(row, 15),
			Religion:                                getValue(row, 16),
			National_id:                             getValue(row, 17),
			Name:                                    getValue(row, 18),
			Date_of_Birth:                           getValue(row, 19),
			Father_name:                             getValue(row, 20),
			Mother_Name:                             getValue(row, 21),
			Blood:                                   getValue(row, 22),
			Personal_Phone:                          getValue(row, 23),
			Present_Address:                         getValue(row, 24),
			Permanent_Address:                       getValue(row, 25),
			Marital_Status:                          getValue(row, 26),
			Emergency_Contact_name:                  getValue(row, 27),
			Emergency_Contact_Relation:              getValue(row, 28),
			Emergency_Contact_Address:               getValue(row, 29),
			Emergency_Contact_Phone:                 getValue(row, 30),
			Personal_Email:                          getValue(row, 31),
			Tin_No:                                  getValue(row, 32),
			Tin_Circle:                              getValue(row, 33),
			Tin_Zone:                                getValue(row, 34),
			Employee_Office_Mobile:                  getValue(row, 35),
			Employee_Office_Email:                   getValue(row, 36),
			Education_Secondary_Institute:           getValue(row, 37),
			Education_Secondary_Major:               getValue(row, 38),
			Education_Secondary_Institution:         getValue(row, 39),
			Education_Secondary_Result:              getValue(row, 40),
			Education_Secondary_Passing_Year:        getValue(row, 41),
			Education_Higher_Secondary_Institute:    getValue(row, 42),
			Education_Higher_Secondary_Major:        getValue(row, 43),
			Education_Higher_Secondary_Institution:  getValue(row, 44),
			Education_Higher_Secondary_Result:       getValue(row, 45),
			Education_Higher_Secondary_Passing_Year: getValue(row, 46),
			Education_Graduation_Institute:          getValue(row, 47),
			Education_Graduation_Major:              getValue(row, 48),
			Education_Graduation_Institution:        getValue(row, 49),
			Education_Graduation_Result:             getValue(row, 50),
			Education_Graduation_Passing_Year:       getValue(row, 51),
			Education_POST_Graduation_Institute:     getValue(row, 52),
			Education_POST_Graduation_Major:         getValue(row, 53),
			Education_POST_Graduation_Institution:   getValue(row, 54),
			Education_POST_Graduation_Result:        getValue(row, 55),
			Education_POST_Graduation_Passing_Year:  getValue(row, 56),
			Work_Experience_First_Name:              getValue(row, 57),
			Work_Experience_First_Designation:       getValue(row, 58),
			Work_Experience_First_Date_To:           getValue(row, 59),
			Work_Experience_First_Date_From:         getValue(row, 60),
			Work_Experience_Second_Name:             getValue(row, 61),
			Work_Experience_Second_Designation:      getValue(row, 62),
			Work_Experience_Second_Date_To:          getValue(row, 63),
			Work_Experience_Second_Date_From:        getValue(row, 64),
			Work_Experience_Third_Name:              getValue(row, 65),
			Work_Experience_Third_Designation:       getValue(row, 66),
			Work_Experience_Third_Date_To:           getValue(row, 67),
			Work_Experience_Third_Date_From:         getValue(row, 68),
			Work_Experience_Fourth_Name:             getValue(row, 69),
			Work_Experience_Fourth_Designation:      getValue(row, 70),
			Work_Experience_Fourth_Date_To:          getValue(row, 71),
			Work_Experience_Fourth_Date_From:        getValue(row, 72),
			Work_Experience_Fifth_Name:              getValue(row, 73),
			Work_Experience_Fifth_Designation:       getValue(row, 74),
			Work_Experience_Fifth_Date_To:           getValue(row, 75),
			Work_Experience_Fifth_Date_From:         getValue(row, 76),
			Work_Experience_Sixth_Name:              getValue(row, 77),
			Work_Experience_Sixth_Designation:       getValue(row, 78),
			Work_Experience_Sixth_Date_To:           getValue(row, 79),
			Work_Experience_Sixth_Date_From:         getValue(row, 80),
		}

		// Log the user details being inserted
		log.Printf("Inserting user: %+v\n", user)

		// Insert user into the database
		err = database.CreateUser(user)
		if err != nil {
			log.Println("Error inserting user into database:", err)
		} else {
			log.Printf("User inserted successfully: %s\n", user.Full_Name)
		}
	}

	fmt.Fprintln(w, "File uploaded and processed successfully")
	log.Println("UploadFile finished")
}

func getValue(row []string, index int) string {
	if index >= len(row) {
		return "" // or any default value you'd like to use
	}
	return row[index]
}
