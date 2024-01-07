package main

import (
	// "fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"context"
)

type User struct{
	Name  				 string	   `json:"name"`
	PhoneNumber  		 string    `json:"phone_number"`
	OTP   				 string    `json:"otp"`
	// OTPExpirationTime    time.Time     `json:"otp_expration_time"`
}

var dbURL = "postgres://postgres:S1122001:)@localhost:5432/postgres"

// var db *sql.DB

func setUserInfo(userName, phoneNumber string)User{
	var user User
	user.Name = userName
	user.PhoneNumber = phoneNumber

	return user
}

func checkPhoneNumberDuplicate(phoneNumber string, phoneNumbers pgx.Rows) bool{
	for _, phoneNum := range phoneNumbers.RawValues(){
		if phoneNumber == string(phoneNum){
			return true
		}
	}
	return false
}


func createNewUser(c * gin.Context){
	var user User

	name, _ := c.GetQuery("name")
	phoneNumber, _:= c.GetQuery("phone_number")
	
	user = setUserInfo(name, phoneNumber)

	// Open a connection with the database
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic("Can't connect to the database")
	}
	defer conn.Close(context.Background())  // Close a connection in the end of the program
	
	
	rows, err := conn.Query(context.Background(), "SELECT phone_number FROM users")
	
	flag := checkPhoneNumberDuplicate(phoneNumber, rows)
	if !flag{
		// Database insert commend
		var insertedID int
		conn.QueryRow(context.Background(), "INSERT INTO users(name, phone_number) VALUES ($1, $2)", user.Name , user.PhoneNumber).Scan(&insertedID)

		c.IndentedJSON(http.StatusOK, "Added Successfully")
	}else{
		c.IndentedJSON(http.StatusBadRequest, "Bad Request")
	}
}



func generateOTP()string {
	var otp []int
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < 4; i++{
		otp = append(otp, rand.Intn(10))
	}

	str := ""
	for i:=0; i<len(otp); i++{
		str += strconv.Itoa(otp[i])
	}
	
	return str
}

func insertOTP(c * gin.Context){
	phoneNumber, _ := c.GetQuery("phone_number")
	otp := generateOTP()

	// Open a connection with the database
	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil{
		panic("Can't connect to the database")
	}
	defer conn.Close(context.Background()) // Close a connection in the end of the program

	// Database update commend
	var insertedID int
	conn.QueryRow(context.Background(), "UPDATE users SET otp=$1 otp_expiration_time=$2 WHERE phone_number=$3", otp, time.Now().Minute(), phoneNumber).Scan(&insertedID)

	c.IndentedJSON(http.StatusOK, "Updated Successfully, your OTP is "+ otp +", and it will expire after 1 min")
}


func verifyOTP(c * gin.Context){
	phoneNumber, _ := c.GetQuery("phone_number")
	otp, _ := c.GetQuery("otp")

	// Open a connection with the database
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil{
		panic("Can't connect to the database")
	}
	defer conn.Close(context.Background())  // Close a connection in the end of the program

	// // Database select commend
	rows, _ := conn.Query(context.Background(), "SELECT otp, otp_expiration_time FROM users WHERE phone_number=$1", phoneNumber)
	
	if err != nil{
		panic(err)
	}
	// fmt.Println(rows)

	for rows.Next(){
		var val_otp string
		var otpExpirationTime int

		err1 := rows.Scan(&val_otp, &otpExpirationTime)
		if err1 != nil{
			panic(err1)
		}

		if time.Now().Minute() - otpExpirationTime > 1{
			c.IndentedJSON(http.StatusForbidden, "Expired")
		} else{
			if otp == val_otp{
				c.IndentedJSON(http.StatusOK, "Verifed")
			} else{
				c.IndentedJSON(http.StatusNotFound, "Not Found")
			}
		}
	}
}



func main(){
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST("/api/users", createNewUser)
	router.POST("/api/users/generateotp", insertOTP)
	router.POST("/api/users/verifyotp", verifyOTP)
	router.Run("localhost:5000")	
}