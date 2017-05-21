package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"
)

const (
	FILE   = "file"
	EDIT   = "edit"
	METHOD = "method"
	SECRET = "secret"
)

type (
	Output struct {
		Head   string
		Body   string
		Verify string
	}

	Input struct {
		Test string `json:"test"`
		jwt.StandardClaims
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "JWT-CLI"
	app.Usage = "generate jwt"
	app.Description = "JWT Sentence generator for cli"
	app.UsageText = "jet-cli [Global Option] argument"
	app.Version = "0.1.0"
	app.Author = "sys-cat"
	app.Email = "systemcat01@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  FILE + ", f",
			Usage: "json file path",
		},
		cli.StringFlag{
			Name:  EDIT + ", e",
			Usage: "direct input json data",
		},
		cli.StringFlag{
			Name:  METHOD + ", m",
			Usage: "convert method",
		},
		cli.StringFlag{
			Name:  SECRET + ", s",
			Usage: "secret key",
		},
	}

	app.Action = Cli

	app.Run(os.Args)
}

func Cli(c *cli.Context) error {
	if err := validate(c); err != nil {
		fmt.Println("---------validate")
		fmt.Println(err)
		return err
	}
	claim := Input{}
	if c.String(EDIT) != "" {
		claim1, err := InputStdin(c.String(EDIT))
		if err != nil {
			fmt.Println("---------InputStdin")
			fmt.Println(err)
			return err
		} else {
			claim = claim1
		}
	}
	if claim.Test == "" {
		claim2, err := InputFile(c.String(FILE))
		if err != nil {
			fmt.Println("---------InputFile")
			fmt.Println(err)
			return err
		} else {
			claim = claim2
		}
	}
	var method string
	if c.String(METHOD) == "" {
		method = "HS256"
	} else {
		method = c.String(METHOD)
	}
	result, err := Generate(claim, method, c.String(SECRET))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Header  : %s\n", result.Head)
	fmt.Printf("Body    : %s\n", result.Body)
	fmt.Printf("Verify  : %s\n", result.Verify)
	return nil
}

func validate(c *cli.Context) error {
	err := errors.New("")
	if !FileExists(c.String(FILE)) && !EditExists(c.String(EDIT)) {
		err = errors.New("File and Edit is empty")
	}
	if err.Error() != "" {
		log.Fatal(err)
		return err
	}
	return nil
}

func FileExists(file string) bool {
	if file == "" {
		return false
	}
	return true
}

func EditExists(edit string) bool {
	if edit == "" {
		return false
	}
	return true
}

func InputFile(path string) (Input, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return Input{}, err
	}
	defer f.Close()
	var j string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		j += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return Input{}, err
	}
	result := Input{}
	if err := json.Unmarshal([]byte(j), &result); err != nil {
		log.Fatal(err)
		return Input{}, err
	}
	return result, nil
}

func InputStdin(data string) (Input, error) {
	result := Input{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return Input{}, err
	}
	return result, nil
}

func Generate(source Input, method string, secret string) (Output, error) {
	if method == "" {
		return Output{}, errors.New("Invalid method")
	}
	if secret == "" {
		return Output{}, errors.New("Invalid Secret Key")
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(method), source)
	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
		return Output{}, err
	}
	fmt.Printf("Original: %s\n", signedString)
	parsed := strings.Split(signedString, ".")
	if len(parsed) == 3 {
		res := Output{
			Head:   parsed[0],
			Body:   parsed[1],
			Verify: parsed[2],
		}
		return res, nil
	}
	return Output{}, errors.New("Parse missing")
}
