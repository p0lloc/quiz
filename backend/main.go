package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Quizzes []Quiz = []Quiz{
	Quiz{
		Id:   "math",
		Name: "Math quiz",
		Questions: []QuizQuestion{
			{
				Id:   "question1",
				Name: "What is 2+2",
				Choices: []QuizChoice{
					{
						Id:      "choice1",
						Name:    "4",
						Correct: true,
					},
					{
						Id:      "choice2",
						Name:    "cat",
						Correct: false,
					},
					{
						Id:      "choice2",
						Name:    "42",
						Correct: false,
					},
					{
						Id:      "choice3",
						Name:    "pi",
						Correct: false,
					},
				},
			},
		},
	},
}

type CreateQuizRequest struct {
	Name      string         `json:"name"`
	Questions []QuizQuestion `json:"questions"`
}

var Games []Game = []Game{
	Game{
		Id:     "somerandomid",
		Code:   "123456",
		QuizId: "math",
	},
}

type Game struct {
	Id     string `json:"id"`
	Code   string `json:"code"`
	QuizId string `json:"quizId"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	app.Get("/funny", rootRoute)
	app.Post("/join", handleJoin)
	app.Get("/app/*", allApp)

	app.Get("/api/quizzes", getQuizzes)
	app.Post("/api/quizzes", createQuiz)
	app.Post("/api/quizzes/:quizId/host", hostQuiz)

	log.Fatal(app.Listen(":3000"))
}

func getQuizById(id string) *Quiz {
	for _, quiz := range Quizzes {
		if quiz.Id == id {
			return &quiz
		}
	}

	return nil
}

func getGameByCode(code string) *Game {
	for _, game := range Games {
		if game.Code == code {
			return &game
		}
	}

	return nil
}

func allApp(c *fiber.Ctx) error {
	return c.SendString("app!")
}

func hostQuiz(c *fiber.Ctx) error {

	quizId := c.Params("quizId")
	quiz := getQuizById(quizId)
	if quiz == nil {
		return c.Status(fiber.StatusNotFound).SendString("quiz not found")
	}

	game := Game{
		Id:     "somerandomid",
		Code:   "123456",
		QuizId: quizId,
	}

	Games = append(Games, game)

	return c.JSON(game)
}

func getQuizzes(c *fiber.Ctx) error {
	return c.JSON(Quizzes)
}

func createQuiz(c *fiber.Ctx) error {
	req := CreateQuizRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	Quizzes = append(Quizzes, Quiz{
		Id:        "id123",
		Name:      req.Name,
		Questions: req.Questions,
	})
	return getQuizzes(c)
}

func handleJoin(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	game := getGameByCode(code)
	if game == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(game)
}

func rootRoute(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
