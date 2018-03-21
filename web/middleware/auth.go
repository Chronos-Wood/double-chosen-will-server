package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/double-chosen-will-server/web"
	log "github.com/sirupsen/logrus"
)

	type Session struct{
		UserName string `json:"userName"`
		SessionID string `json:"sessionID"`
		IP string `json:"ip"`
		jwt.StandardClaims
	}

	const(
		key="Q50MWQ!bcWDv.k:S,]0lk,\\qaet]a0m"
	)

	func SessionValidation(redisConnectionPool *redis.Pool) gin.HandlerFunc {
		return func(context *gin.Context){
			tokenString, err := context.Cookie("scookie")
			if err != nil {
				context.AbortWithStatusJSON(200, status.ResponseMessage{Status: status.UNAUTHORIZED, Message: "please login"})
				return
			}
			claims, err := jwt.ParseWithClaims(tokenString, &Session{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(key), nil
			})

			if err != nil {
				context.AbortWithStatusJSON(200, status.ResponseMessage{Status: status.UNAUTHORIZED, Message: err.Error()})
				return
			}

			session, ok := claims.Claims.(*Session)
			if !ok {
				context.AbortWithStatusJSON(200, status.ResponseMessage{Status: status.UNAUTHORIZED, Message: "malformatted token"})
				return
			}
			if session.hasExpired(redisConnectionPool.Get()) {
				context.AbortWithStatusJSON(200, status.ResponseMessage{Status: status.UNAUTHORIZED, Message: "login again"})
				return
			}

			if strings.HasPrefix(context.Request.RemoteAddr, session.IP) {
				context.AbortWithStatusJSON(200, status.ResponseMessage{Status: status.UNAUTHORIZED, Message: "invalid host"})
				return
			}

			context.Set("UserName", session.UserName)

		}
	}

	//do not valid session when beening parsing
	func (session Session) Valid() error {
		return nil;
	}

	func (session *Session) hasExpired(conn redis.Conn) bool {
		var resp interface{}
		var err error

		defer func() {
			if err != nil {
				log.Errorf("read session failed:%s", err.Error())
			}
			conn.Close()
		}()

		if conn.Err() != nil {
			log.Errorf("unable to connect to redis server:%s", conn.Err().Error())
		}

		if resp, err = conn.Do("SET", session.SessionID, 1, "EX", "300", "XX" ); resp == nil{
			return true;
		}
		return false;
	}



