package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Mirinjamamul/go-poc-api/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	// Create the users table if it doesn't exist
	createUsersTable()
}

func createUsersTable() {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        poc_lead VARCHAR(100),
        point_of_contact VARCHAR(100),
        employee_id VARCHAR(100),
        full_name VARCHAR(200),
        company VARCHAR(100),
        unit VARCHAR(100),
        department VARCHAR(100),
        section VARCHAR(100),
        job_category VARCHAR(100),
        designation VARCHAR(100),
        employee_type VARCHAR(100),
        work_location VARCHAR(100),
        reporting_employee_id VARCHAR(100),
        reporting_employee_name VARCHAR(200),
        reporting_designation VARCHAR(100),
        gender VARCHAR(10),
        religion VARCHAR(50),
        national_id VARCHAR(100),
        name VARCHAR(100),
        date_of_birth VARCHAR(100),
        father_name VARCHAR(200),
        mother_name VARCHAR(200),
        blood VARCHAR(10),
        personal_phone VARCHAR(15),
        present_address TEXT,
        permanent_address TEXT,
        marital_status VARCHAR(20),
        emergency_contact_name VARCHAR(200),
        emergency_contact_relation VARCHAR(50),
        emergency_contact_address TEXT,
        emergency_contact_phone VARCHAR(15),
        personal_email VARCHAR(100),
        tin_no VARCHAR(100),
        tin_circle VARCHAR(50),
        tin_zone VARCHAR(50),
        employee_office_mobile VARCHAR(15),
        employee_office_email VARCHAR(100),
        education_secondary_institute VARCHAR(200),
        education_secondary_major VARCHAR(100),
        education_secondary_institution VARCHAR(200),
        education_secondary_result VARCHAR(50),
        education_secondary_passing_year VARCHAR(100),
        education_higher_secondary_institute VARCHAR(200),
        education_higher_secondary_major VARCHAR(100),
        education_higher_secondary_institution VARCHAR(200),
        education_higher_secondary_result VARCHAR(50),
        education_higher_secondary_passing_year VARCHAR(100),
        education_graduation_institute VARCHAR(200),
        education_graduation_major VARCHAR(100),
        education_graduation_institution VARCHAR(200),
        education_graduation_result VARCHAR(50),
        education_graduation_passing_year VARCHAR(100),
        education_post_graduation_institute VARCHAR(200),
        education_post_graduation_major VARCHAR(100),
        education_post_graduation_institution VARCHAR(200),
        education_post_graduation_result VARCHAR(50),
        education_post_graduation_passing_year VARCHAR(100),
        work_experience_first_name VARCHAR(200),
        work_experience_first_designation VARCHAR(100),
        work_experience_first_date_to VARCHAR(100),
        work_experience_first_date_from VARCHAR(100),
        work_experience_second_name VARCHAR(200),
        work_experience_second_designation VARCHAR(100),
        work_experience_second_date_to VARCHAR(100),
        work_experience_second_date_from VARCHAR(100),
        work_experience_third_name VARCHAR(200),
        work_experience_third_designation VARCHAR(100),
        work_experience_third_date_to VARCHAR(100),
        work_experience_third_date_from VARCHAR(100),
        work_experience_fourth_name VARCHAR(200),
        work_experience_fourth_designation VARCHAR(100),
        work_experience_fourth_date_to VARCHAR(100),
        work_experience_fourth_date_from VARCHAR(100),
        work_experience_fifth_name VARCHAR(200),
        work_experience_fifth_designation VARCHAR(100),
        work_experience_fifth_date_to VARCHAR(100),
        work_experience_fifth_date_from VARCHAR(100),
        work_experience_sixth_name VARCHAR(200),
        work_experience_sixth_designation VARCHAR(100),
        work_experience_sixth_date_to VARCHAR(100),
        work_experience_sixth_date_from VARCHAR(100)
    );
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating users table: ", err)
	}
	fmt.Println("Users table is ensured to exist")
}

func Close() {
	db.Close()
}

func CreateUser(user models.User) error {
	// Construct the SQL INSERT statement based on the User struct fields
	sql := `INSERT INTO users (
		Poc_Lead, Point_of_contact, Employee_Id, Full_Name, Company, Unit, Department, Section,
		Job_Category, Designation, Employee_Type, Work_Location, Reporting_Employee_ID, 
		Reporting_Employee_Name, Reporting_Designation, Gender, Religion, National_id, Name, 
		Date_of_Birth, Father_name, Mother_Name, Blood, Personal_Phone, Present_Address, 
		Permanent_Address, Marital_Status, Emergency_Contact_name, Emergency_Contact_Relation, 
		Emergency_Contact_Address, Emergency_Contact_Phone, Personal_Email, Tin_No, Tin_Circle, 
		Tin_Zone, Employee_Office_Mobile, Employee_Office_Email, Education_Secondary_Institute, 
		Education_Secondary_Major, Education_Secondary_Institution, Education_Secondary_Result, 
		Education_Secondary_Passing_Year, Education_Higher_Secondary_Institute, 
		Education_Higher_Secondary_Major, Education_Higher_Secondary_Institution, 
		Education_Higher_Secondary_Result, Education_Higher_Secondary_Passing_Year, 
		Education_Graduation_Institute, Education_Graduation_Major, Education_Graduation_Institution, 
		Education_Graduation_Result, Education_Graduation_Passing_Year, Education_POST_Graduation_Institute, 
		Education_POST_Graduation_Major, Education_POST_Graduation_Institution, 
		Education_POST_Graduation_Result, Education_POST_Graduation_Passing_Year, 
		Work_Experience_First_Name, Work_Experience_First_Designation, Work_Experience_First_Date_To, 
		Work_Experience_First_Date_From, Work_Experience_Second_Name, Work_Experience_Second_Designation, 
		Work_Experience_Second_Date_To, Work_Experience_Second_Date_From, Work_Experience_Third_Name, 
		Work_Experience_Third_Designation, Work_Experience_Third_Date_To, Work_Experience_Third_Date_From, 
		Work_Experience_Fourth_Name, Work_Experience_Fourth_Designation, Work_Experience_Fourth_Date_To, 
		Work_Experience_Fourth_Date_From, Work_Experience_Fifth_Name, Work_Experience_Fifth_Designation, 
		Work_Experience_Fifth_Date_To, Work_Experience_Fifth_Date_From, Work_Experience_Sixth_Name, 
		Work_Experience_Sixth_Designation, Work_Experience_Sixth_Date_To, Work_Experience_Sixth_Date_From
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, 
		$23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, 
		$43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, 
		$63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81
	)`

	// Prepare the statement
	stmt, err := db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	// Log the parameters
	params := []interface{}{
		user.Poc_Lead, user.Point_of_contact, user.Employee_Id, user.Full_Name, user.Company, user.Unit, user.Department,
		user.Section, user.Job_Category, user.Designation, user.Employee_Type, user.Work_Location, user.Reporting_Employee_ID,
		user.Reporting_Employee_Name, user.Reporting_Designation, user.Gender, user.Religion, user.National_id, user.Name,
		user.Date_of_Birth, user.Father_name, user.Mother_Name, user.Blood, user.Personal_Phone, user.Present_Address,
		user.Permanent_Address, user.Marital_Status, user.Emergency_Contact_name, user.Emergency_Contact_Relation,
		user.Emergency_Contact_Address, user.Emergency_Contact_Phone, user.Personal_Email, user.Tin_No, user.Tin_Circle,
		user.Tin_Zone, user.Employee_Office_Mobile, user.Employee_Office_Email, user.Education_Secondary_Institute,
		user.Education_Secondary_Major, user.Education_Secondary_Institution, user.Education_Secondary_Result,
		user.Education_Secondary_Passing_Year, user.Education_Higher_Secondary_Institute, user.Education_Higher_Secondary_Major,
		user.Education_Higher_Secondary_Institution, user.Education_Higher_Secondary_Result,
		user.Education_Higher_Secondary_Passing_Year, user.Education_Graduation_Institute, user.Education_Graduation_Major,
		user.Education_Graduation_Institution, user.Education_Graduation_Result, user.Education_Graduation_Passing_Year,
		user.Education_POST_Graduation_Institute, user.Education_POST_Graduation_Major,
		user.Education_POST_Graduation_Institution, user.Education_POST_Graduation_Result,
		user.Education_POST_Graduation_Passing_Year, user.Work_Experience_First_Name, user.Work_Experience_First_Designation,
		user.Work_Experience_First_Date_To, user.Work_Experience_First_Date_From, user.Work_Experience_Second_Name,
		user.Work_Experience_Second_Designation, user.Work_Experience_Second_Date_To, user.Work_Experience_Second_Date_From,
		user.Work_Experience_Third_Name, user.Work_Experience_Third_Designation, user.Work_Experience_Third_Date_To,
		user.Work_Experience_Third_Date_From, user.Work_Experience_Fourth_Name, user.Work_Experience_Fourth_Designation,
		user.Work_Experience_Fourth_Date_To, user.Work_Experience_Fourth_Date_From, user.Work_Experience_Fifth_Name,
		user.Work_Experience_Fifth_Designation, user.Work_Experience_Fifth_Date_To, user.Work_Experience_Fifth_Date_From,
		user.Work_Experience_Sixth_Name, user.Work_Experience_Sixth_Designation, user.Work_Experience_Sixth_Date_To,
		user.Work_Experience_Sixth_Date_From,
	}

	// Execute the statement with user data
	_, err = stmt.Exec(params...)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func GetUsers() []models.User {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Poc_Lead, &user.Point_of_contact, &user.Employee_Id, &user.Full_Name, &user.Company, &user.Unit, &user.Department, &user.Section, &user.Job_Category, &user.Designation, &user.Employee_Type, &user.Work_Location, &user.Reporting_Employee_ID, &user.Reporting_Employee_Name, &user.Reporting_Designation, &user.Gender, &user.Religion, &user.National_id, &user.Name, &user.Date_of_Birth, &user.Father_name, &user.Mother_Name, &user.Blood, &user.Personal_Phone, &user.Present_Address, &user.Permanent_Address, &user.Marital_Status, &user.Emergency_Contact_name, &user.Emergency_Contact_Relation, &user.Emergency_Contact_Address, &user.Emergency_Contact_Phone, &user.Personal_Email, &user.Tin_No, &user.Tin_Circle, &user.Tin_Zone, &user.Employee_Office_Mobile, &user.Employee_Office_Email, &user.Education_Secondary_Institute, &user.Education_Secondary_Major, &user.Education_Secondary_Institution, &user.Education_Secondary_Result, &user.Education_Secondary_Passing_Year, &user.Education_Higher_Secondary_Institute, &user.Education_Higher_Secondary_Major, &user.Education_Higher_Secondary_Institution, &user.Education_Higher_Secondary_Result, &user.Education_Higher_Secondary_Passing_Year, &user.Education_Graduation_Institute, &user.Education_Graduation_Major, &user.Education_Graduation_Institution, &user.Education_Graduation_Result, &user.Education_Graduation_Passing_Year, &user.Education_POST_Graduation_Institute, &user.Education_POST_Graduation_Major, &user.Education_POST_Graduation_Institution, &user.Education_POST_Graduation_Result, &user.Education_POST_Graduation_Passing_Year, &user.Work_Experience_First_Name, &user.Work_Experience_First_Designation, &user.Work_Experience_First_Date_To, &user.Work_Experience_First_Date_From, &user.Work_Experience_Second_Name, &user.Work_Experience_Second_Designation, &user.Work_Experience_Second_Date_To, &user.Work_Experience_Second_Date_From, &user.Work_Experience_Third_Name, &user.Work_Experience_Third_Designation, &user.Work_Experience_Third_Date_To, &user.Work_Experience_Third_Date_From, &user.Work_Experience_Fourth_Name, &user.Work_Experience_Fourth_Designation, &user.Work_Experience_Fourth_Date_To, &user.Work_Experience_Fourth_Date_From, &user.Work_Experience_Fifth_Name, &user.Work_Experience_Fifth_Designation, &user.Work_Experience_Fifth_Date_To, &user.Work_Experience_Fifth_Date_From, &user.Work_Experience_Sixth_Name, &user.Work_Experience_Sixth_Designation, &user.Work_Experience_Sixth_Date_To, &user.Work_Experience_Sixth_Date_From)
		users = append(users, user)
	}
	return users
}

func GetUser(id string) (models.User, bool) {
	var user models.User
	err := db.QueryRow(`
		SELECT Poc_Lead, Point_of_contact, Employee_Id, Full_Name, Company, Unit, Department, Section, Job_Category, Designation,
		       Employee_Type, Work_Location, Reporting_Employee_ID, Reporting_Employee_Name, Reporting_Designation, Gender, Religion, National_id,
		       Name, Date_of_Birth, Father_name, Mother_Name, Blood, Personal_Phone, Present_Address, Permanent_Address, Marital_Status,
		       Emergency_Contact_name, Emergency_Contact_Relation, Emergency_Contact_Address, Emergency_Contact_Phone, Personal_Email, Tin_No,
		       Tin_Circle, Tin_Zone, Employee_Office_Mobile, Employee_Office_Email, Education_Secondary_Institute, Education_Secondary_Major,
		       Education_Secondary_Institution, Education_Secondary_Result, Education_Secondary_Passing_Year, Education_Higher_Secondary_Institute,
		       Education_Higher_Secondary_Major, Education_Higher_Secondary_Institution, Education_Higher_Secondary_Result, Education_Higher_Secondary_Passing_Year,
		       Education_Graduation_Institute, Education_Graduation_Major, Education_Graduation_Institution, Education_Graduation_Result,
		       Education_Graduation_Passing_Year, Education_POST_Graduation_Institute, Education_POST_Graduation_Major, Education_POST_Graduation_Institution,
		       Education_POST_Graduation_Result, Education_POST_Graduation_Passing_Year, Work_Experience_First_Name, Work_Experience_First_Designation,
		       Work_Experience_First_Date_To, Work_Experience_First_Date_From, Work_Experience_Second_Name, Work_Experience_Second_Designation,
		       Work_Experience_Second_Date_To, Work_Experience_Second_Date_From, Work_Experience_Third_Name, Work_Experience_Third_Designation,
		       Work_Experience_Third_Date_To, Work_Experience_Third_Date_From, Work_Experience_Fourth_Name, Work_Experience_Fourth_Designation,
		       Work_Experience_Fourth_Date_To, Work_Experience_Fourth_Date_From, Work_Experience_Fifth_Name, Work_Experience_Fifth_Designation,
		       Work_Experience_Fifth_Date_To, Work_Experience_Fifth_Date_From, Work_Experience_Sixth_Name, Work_Experience_Sixth_Designation,
		       Work_Experience_Sixth_Date_To, Work_Experience_Sixth_Date_From
		FROM users WHERE Employee_Id = $1`, id).Scan(
		&user.Poc_Lead, &user.Point_of_contact, &user.Employee_Id, &user.Full_Name, &user.Company, &user.Unit, &user.Department, &user.Section,
		&user.Job_Category, &user.Designation, &user.Employee_Type, &user.Work_Location, &user.Reporting_Employee_ID, &user.Reporting_Employee_Name,
		&user.Reporting_Designation, &user.Gender, &user.Religion, &user.National_id, &user.Name, &user.Date_of_Birth, &user.Father_name, &user.Mother_Name,
		&user.Blood, &user.Personal_Phone, &user.Present_Address, &user.Permanent_Address, &user.Marital_Status, &user.Emergency_Contact_name,
		&user.Emergency_Contact_Relation, &user.Emergency_Contact_Address, &user.Emergency_Contact_Phone, &user.Personal_Email, &user.Tin_No, &user.Tin_Circle,
		&user.Tin_Zone, &user.Employee_Office_Mobile, &user.Employee_Office_Email, &user.Education_Secondary_Institute, &user.Education_Secondary_Major,
		&user.Education_Secondary_Institution, &user.Education_Secondary_Result, &user.Education_Secondary_Passing_Year, &user.Education_Higher_Secondary_Institute,
		&user.Education_Higher_Secondary_Major, &user.Education_Higher_Secondary_Institution, &user.Education_Higher_Secondary_Result,
		&user.Education_Higher_Secondary_Passing_Year, &user.Education_Graduation_Institute, &user.Education_Graduation_Major, &user.Education_Graduation_Institution,
		&user.Education_Graduation_Result, &user.Education_Graduation_Passing_Year, &user.Education_POST_Graduation_Institute, &user.Education_POST_Graduation_Major,
		&user.Education_POST_Graduation_Institution, &user.Education_POST_Graduation_Result, &user.Education_POST_Graduation_Passing_Year,
		&user.Work_Experience_First_Name, &user.Work_Experience_First_Designation, &user.Work_Experience_First_Date_To, &user.Work_Experience_First_Date_From,
		&user.Work_Experience_Second_Name, &user.Work_Experience_Second_Designation, &user.Work_Experience_Second_Date_To, &user.Work_Experience_Second_Date_From,
		&user.Work_Experience_Third_Name, &user.Work_Experience_Third_Designation, &user.Work_Experience_Third_Date_To, &user.Work_Experience_Third_Date_From,
		&user.Work_Experience_Fourth_Name, &user.Work_Experience_Fourth_Designation, &user.Work_Experience_Fourth_Date_To, &user.Work_Experience_Fourth_Date_From,
		&user.Work_Experience_Fifth_Name, &user.Work_Experience_Fifth_Designation, &user.Work_Experience_Fifth_Date_To, &user.Work_Experience_Fifth_Date_From,
		&user.Work_Experience_Sixth_Name, &user.Work_Experience_Sixth_Designation, &user.Work_Experience_Sixth_Date_To, &user.Work_Experience_Sixth_Date_From,
	)
	if err != nil {
		return user, false
	}
	return user, true
}

func UpdateUser(id string, user models.User) {
	_, err := db.Exec(`
		UPDATE users SET 
			Poc_Lead = $2, Point_of_contact = $3, Employee_Id = $4, Full_Name = $5, Company = $6, Unit = $7, Department = $8, Section = $9,
			Job_Category = $10, Designation = $11, Employee_Type = $12, Work_Location = $13, Reporting_Employee_ID = $14, Reporting_Employee_Name = $15,
			Reporting_Designation = $16, Gender = $17, Religion = $18, National_id = $19, Name = $20, Date_of_Birth = $21, Father_name = $22, 
			Mother_Name = $23, Blood = $24, Personal_Phone = $25, Present_Address = $26, Permanent_Address = $27, Marital_Status = $28, 
			Emergency_Contact_name = $29, Emergency_Contact_Relation = $30, Emergency_Contact_Address = $31, Emergency_Contact_Phone = $32, 
			Personal_Email = $33, Tin_No = $34, Tin_Circle = $35, Tin_Zone = $36, Employee_Office_Mobile = $37, Employee_Office_Email = $38, 
			Education_Secondary_Institute = $39, Education_Secondary_Major = $40, Education_Secondary_Institution = $41, Education_Secondary_Result = $42, 
			Education_Secondary_Passing_Year = $43, Education_Higher_Secondary_Institute = $44, Education_Higher_Secondary_Major = $45, 
			Education_Higher_Secondary_Institution = $46, Education_Higher_Secondary_Result = $47, Education_Higher_Secondary_Passing_Year = $48, 
			Education_Graduation_Institute = $49, Education_Graduation_Major = $50, Education_Graduation_Institution = $51, Education_Graduation_Result = $52, 
			Education_Graduation_Passing_Year = $53, Education_POST_Graduation_Institute = $54, Education_POST_Graduation_Major = $55, 
			Education_POST_Graduation_Institution = $56, Education_POST_Graduation_Result = $57, Education_POST_Graduation_Passing_Year = $58, 
			Work_Experience_First_Name = $59, Work_Experience_First_Designation = $60, Work_Experience_First_Date_To = $61, Work_Experience_First_Date_From = $62, 
			Work_Experience_Second_Name = $63, Work_Experience_Second_Designation = $64, Work_Experience_Second_Date_To = $65, Work_Experience_Second_Date_From = $66, 
			Work_Experience_Third_Name = $67, Work_Experience_Third_Designation = $68, Work_Experience_Third_Date_To = $69, Work_Experience_Third_Date_From = $70, 
			Work_Experience_Fourth_Name = $71, Work_Experience_Fourth_Designation = $72, Work_Experience_Fourth_Date_To = $73, Work_Experience_Fourth_Date_From = $74, 
			Work_Experience_Fifth_Name = $75, Work_Experience_Fifth_Designation = $76, Work_Experience_Fifth_Date_To = $77, Work_Experience_Fifth_Date_From = $78, 
			Work_Experience_Sixth_Name = $79, Work_Experience_Sixth_Designation = $80, Work_Experience_Sixth_Date_To = $81, Work_Experience_Sixth_Date_From = $82 
		WHERE Employee_Id = $1`,
		id, user.Poc_Lead, user.Point_of_contact, user.Employee_Id, user.Full_Name, user.Company, user.Unit, user.Department, user.Section,
		user.Job_Category, user.Designation, user.Employee_Type, user.Work_Location, user.Reporting_Employee_ID, user.Reporting_Employee_Name,
		user.Reporting_Designation, user.Gender, user.Religion, user.National_id, user.Name, user.Date_of_Birth, user.Father_name, user.Mother_Name,
		user.Blood, user.Personal_Phone, user.Present_Address, user.Permanent_Address, user.Marital_Status, user.Emergency_Contact_name,
		user.Emergency_Contact_Relation, user.Emergency_Contact_Address, user.Emergency_Contact_Phone, user.Personal_Email, user.Tin_No, user.Tin_Circle,
		user.Tin_Zone, user.Employee_Office_Mobile, user.Employee_Office_Email, user.Education_Secondary_Institute, user.Education_Secondary_Major,
		user.Education_Secondary_Institution, user.Education_Secondary_Result, user.Education_Secondary_Passing_Year, user.Education_Higher_Secondary_Institute,
		user.Education_Higher_Secondary_Major, user.Education_Higher_Secondary_Institution, user.Education_Higher_Secondary_Result,
		user.Education_Higher_Secondary_Passing_Year, user.Education_Graduation_Institute, user.Education_Graduation_Major, user.Education_Graduation_Institution,
		user.Education_Graduation_Result, user.Education_Graduation_Passing_Year, user.Education_POST_Graduation_Institute, user.Education_POST_Graduation_Major,
		user.Education_POST_Graduation_Institution, user.Education_POST_Graduation_Result, user.Education_POST_Graduation_Passing_Year,
		user.Work_Experience_First_Name, user.Work_Experience_First_Designation, user.Work_Experience_First_Date_To, user.Work_Experience_First_Date_From,
		user.Work_Experience_Second_Name, user.Work_Experience_Second_Designation, user.Work_Experience_Second_Date_To, user.Work_Experience_Second_Date_From,
		user.Work_Experience_Third_Name, user.Work_Experience_Third_Designation, user.Work_Experience_Third_Date_To, user.Work_Experience_Third_Date_From,
		user.Work_Experience_Fourth_Name, user.Work_Experience_Fourth_Designation, user.Work_Experience_Fourth_Date_To, user.Work_Experience_Fourth_Date_From,
		user.Work_Experience_Fifth_Name, user.Work_Experience_Fifth_Designation, user.Work_Experience_Fifth_Date_To, user.Work_Experience_Fifth_Date_From,
		user.Work_Experience_Sixth_Name, user.Work_Experience_Sixth_Designation, user.Work_Experience_Sixth_Date_To, user.Work_Experience_Sixth_Date_From,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(id string) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
