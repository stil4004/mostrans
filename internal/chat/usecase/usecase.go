package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"service/internal/chat"
	"service/pkg/ai_client"
	"sync"
)

type ChatUseCase struct {
	repo        chat.Repository
	aiApiClient ai_client.UseCase

	responseMapMtx *sync.Mutex
	responseMap    map[int]string

	errResponseMapMtx *sync.Mutex
	errResponseMap    map[int]string
}

func NewChatUseCase(repo chat.Repository, aiApiClient ai_client.UseCase) chat.UseCase {
	return &ChatUseCase{
		repo:        repo,
		aiApiClient: aiApiClient,

		responseMapMtx: new(sync.Mutex),
		responseMap:    make(map[int]string),

		errResponseMapMtx: new(sync.Mutex),
		errResponseMap:    make(map[int]string),
	}
}

func (c *ChatUseCase) ProcessMessage(ctx context.Context, req chat.ProcessMessageRequest) (chat.ProcessMessageResponse, error) {
	var (
		resp ai_client.GetBrigV1Response
		err  error
	)
	if req.MessageText == "" {
		rnd := rand.Intn(5)
		c.errResponseMapMtx.Lock()
		msg := c.errResponseMap[rnd]
		c.errResponseMapMtx.Unlock()

		return chat.ProcessMessageResponse{
			ResponseMessage: msg,
		}, errors.New("empty data sent")
	}
	a := req.AIType
	switch a {
	case 1:
		resp, err = c.aiApiClient.GetBrigV1(ctx, req.MessageText)
		if err != nil {
			rnd := rand.Intn(5)
			c.errResponseMapMtx.Lock()
			msg := c.errResponseMap[rnd]
			c.errResponseMapMtx.Unlock()
			log.Println("Err on first", err)
			return chat.ProcessMessageResponse{
				ResponseMessage: msg,
			}, errors.New("empty type sent")
		}
		log.Println(resp)

	case 2:
		rnd := rand.Intn(5)
		c.errResponseMapMtx.Lock()
		msg := c.errResponseMap[rnd]
		c.errResponseMapMtx.Unlock()
		log.Println("Err on first", err)
		return chat.ProcessMessageResponse{
			ResponseMessage: msg,
		}, errors.New("empty type sent")

	default:
		resp, err = c.aiApiClient.GetBrigV1(ctx, req.MessageText)
		if err != nil {
			rnd := rand.Intn(5)
			c.errResponseMapMtx.Lock()
			msg := c.errResponseMap[rnd]
			c.errResponseMapMtx.Unlock()

			return chat.ProcessMessageResponse{
				ResponseMessage: msg,
			}, errors.New("empty type sent default")
		}
	}
	responseDB, err := c.repo.GetInfoFromBatch(ctx, chat.GetInfoFromBatchRequest{
		Periods:  resp.Period,
		Stations: resp.Stations,
	})
	if err != nil {
		rnd := rand.Intn(5)
		c.errResponseMapMtx.Lock()
		msg := c.errResponseMap[rnd]
		c.errResponseMapMtx.Unlock()
		log.Println("Error from db", err)
		return chat.ProcessMessageResponse{
			ResponseMessage: msg,
		}, errors.New("error from db")
	}

	rnd := rand.Intn(5)
	c.responseMapMtx.Lock()
	msg := fmt.Sprintf(c.responseMap[rnd], responseDB.Stations[0], responseDB.PeopleFlow)
	c.responseMapMtx.Unlock()

	return chat.ProcessMessageResponse{
		ResponseMessage: msg,
	}, nil
}

// func (a *AuthUseCase) LogIn(ctx context.Context, req auth.LogInRequest) (auth.LogInResponse, error) {

// 	resp, err := a.repo.CheckLogIn(ctx, auth.CheckLogInRequest{
// 		NickName: req.NickName,
// 		Password: req.Password,
// 	})
// 	if err != nil {
// 		return auth.LogInResponse{
// 			Authenticated: false,
// 		}, err
// 	}

// 	if !resp.Authenticated {
// 		return auth.LogInResponse{
// 			Authenticated: false,
// 		}, errors.New("wrong password")
// 	}

// 	// TODO
// 	// token := jwt.NewWithClaims(
// 	// 	jwt.SigningMethodHS256,
// 	// 	&jwt.StandardClaims{
// 	// 		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
// 	// 		IssuedAt:  time.Now().Unix(),

// 	// 	},
// 	// )
// 	respUsr, err := a.repo.GetUserByLogin(ctx, auth.GetUserByLoginRequest{
// 		NickName: req.NickName,
// 	})
// 	if err != nil {
// 		return auth.LogInResponse{
// 			Authenticated: false,
// 		}, err
// 	}

// 	return auth.LogInResponse{
// 		User:          respUsr.UserResp,
// 		Authenticated: true,
// 		Access:        "efwefwe.sefsefesf.wadawwd",
// 	}, nil
// }

func generatePasswordHashJWT(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte("231231232")))
}

func (c *ChatUseCase) Cache() {
	c.errResponseMapMtx.Lock()
	c.errResponseMap = map[int]string{
		0: "К сожалению, не получилось обработать ваш запрос",
		1: "Извинте я очень старалась, но не смогла обработать ваш запрос...",
		2: "Видимо произошла какая-то ошибка, попробуйте снова, или измените запрос",
		3: "Что-то пошло не так, к сожалению не смогла обработать ваш запрос",
		4: "Такой запрос не получилось обработать, возможно, вина даже на моей стороне...",
		5: "Мне очень жаль, но запрос не получилось обработать",
	}
	c.errResponseMapMtx.Unlock()

	c.responseMapMtx.Lock()
	c.responseMap = map[int]string{
		0: "Похоже вы искали сколько было человек на станции %[1]s - по моим данным это %[2]s  за день",
		1: "За запрошенное время на станции %[1]s прошло %[2]s человек",
		2: "Нашла данные по вашему запросу - в это время на станции %[1]s пассажиропоток составил %[2]s человек",
		3: "Удалось найти информацию по вашему запросу: на станции %[1]s - поток составил %[2]s",
		4: "Исходя из запрошенных данных поток составил %[2]s на станции  %[1]s",
		5: "Судя по моим данным, поток на этой станции составил %[2]s",
	}
	c.responseMapMtx.Unlock()
}
