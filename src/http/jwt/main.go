// ref : https://covenant.tistory.com/203

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	// _ "go get github.com/go-redis/redis/v7"
)

var (
	router = gin.Default()
	redis  map[string]interface{} // redis 대용
)

func init() {
	// //Initializing redis
	// dsn := os.Getenv("REDIS_DSN")
	// if len(dsn) == 0 {
	// 	dsn = "localhost:6379"
	// }
	// client = redis.NewClient(&redis.Options{
	// 	Addr: dsn, //redis port
	// })
	// _, err := client.Ping().Result()
	// if err != nil {
	// 	panic(err)
	// }

	redis = make(map[string]interface{})
}

func main() {
	router.POST("/login", Login)
	router.POST("/logout", TokenAuthMiddleware(), Logout)
	router.POST("/token/refresh", Refresh)
	router.POST("/todo", TokenAuthMiddleware(), CreateTodo)
	// router.POST("/logout", Logout)
	// router.POST("/todo", CreateTodo)

	log.Fatal(router.Run(":2222"))
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	// verify the token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") // this should be in an env file
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	// if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	// is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	// Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) // the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) // convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}

		// Delete the previous Refresh Token
		delErr := DeleteAuth(refreshUuid)
		if delErr != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		// Create new pairs of refresh and access tokens
		ts, createErr := CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}

		// save the tokens metadata to redis
		saveErr := CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "123123",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	// compare the user from the request, with the one we defined
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	ts, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	saveErr := CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func Logout(c *gin.Context) {
	au, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	err = DeleteAuth(au.AccessUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userId uint64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	// td.AtExpires = time.Now().Add(time.Second * 5).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	// create access token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["autorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// create refresh token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

type RedisMapAccessTokenSaveFormat struct {
	UserId string
	ExTime time.Duration
}

func CreateAuth(userId uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) // converting Unix to UTC
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	// set uuid, userid, now
	redis[td.AccessUuid] = RedisMapAccessTokenSaveFormat{
		UserId: strconv.Itoa(int(userId)),
		ExTime: at.Sub(now),
	}
	redis[td.RefreshUuid] = RedisMapAccessTokenSaveFormat{
		UserId: strconv.Itoa(int(userId)),
		ExTime: rt.Sub(now),
	}
	return nil
}

func DeleteAuth(givenUuid string) error {
	// deleted, err := client.Del(givenUuid).Result()
	delete(redis, givenUuid)
	return nil
}

////////////////////////////////// Application - TODO
type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

// 요청 헤더에서 토큰 추출
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	// normally authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// 추출한 토큰 검증
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	// DEBUG DEBUG DEBUG - START
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("[Debug] debug fail, result: ", err)
	} else {
		fmt.Println(claims)
		fmt.Println("[Debug] debug ok, result: ", err)
	}
	// DEBUG DEBUG DEBUG - END
	if err != nil {
		return nil, err
	}

	return token, nil
}

// token 만료 검사
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

// token 추출 및 조회
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

// find token in database(=redis) via uuid
func FetchAuth(authD *AccessDetails) (uint64, error) {
	// userId, err := client.Get(authD.AccessUuid).Result()
	userId, exists := redis[authD.AccessUuid]
	if !exists {
		return 0, errors.New("not exists")
	}
	fmt.Println("[DEBUG] ", userId.(RedisMapAccessTokenSaveFormat).UserId)
	userID, _ := strconv.ParseUint(userId.(RedisMapAccessTokenSaveFormat).UserId, 10, 64)
	return userID, nil
}

func CreateTodo(c *gin.Context) {
	var td *Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	td.UserID = userId

	//you can proceed to save the Todo to a database
	//but we will just return it to the caller here:
	c.JSON(http.StatusCreated, td)
}
