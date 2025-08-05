package pkg

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func sanitizeSensitiveData(bytes []byte) string {
	var dataParse = make(map[string]interface{})
	_ = json.Unmarshal(bytes, &dataParse)

	// if dataParse["email"] != nil {
	// 	dataParse["email"] = "*****"
	// }

	if dataParse["password"] != nil {
		dataParse["password"] = "*****"
	}

	if dataParse["repeat_password"] != nil {
		dataParse["repeat_password"] = "*****"
	}

	if dataParse["phone"] != nil {
		dataParse["phone"] = "*****"
	}

	// if dataParse["email"] != nil {
	// 	dataParse["email"] = "*****"
	// }

	if dataParse["token"] != nil {
		dataParse["token"] = "*****"
	}

	if dataParse["key"] != nil {
		dataParse["key"] = "*****"
	}

	if dataParse["secret_key"] != nil {
		dataParse["secret_key"] = "*****"
	}

	res, _ := json.Marshal(dataParse)
	return string(res)
}

func getIp(c *fiber.Ctx) string {
	ip := c.Get("X-Forwarded-For")

	if ip != "" {
		ips := strings.Split(ip, ",")
		ip = strings.TrimSpace(ips[0])
	} else {
		ip = c.IP()
	}

	return ip
}

func CreateAccessLog(ctx *fiber.Ctx, ptr string, statusCode int, resp any) {

	if Cfg.Application.EnableLog {
		logFormat := ptr +
			" " +
			time.Now().Format("2006/01/02 15:04:05") +
			" " +
			getIp(ctx) +
			" " +
			ctx.Method() +
			" " +
			strconv.Itoa(statusCode) +
			" " +
			"CONTENT_TYPE=" + string(ctx.Request().Header.ContentType()) +
			" " +
			"ROUTE=" + ctx.Route().Path

		if ctx.Request().URI().QueryString() != nil {
			logFormat = logFormat + " QUERY_URL=" + string(ctx.Request().URI().QueryString())
		}

		if ctx.Body() != nil {
			body := string(ctx.Request().Body())

			helper := make(map[string]interface{})

			err := json.Unmarshal([]byte(body), &helper)
			if err == nil {
				bytes, err := json.Marshal(helper)
				if err == nil {
					// Sanitize some input body
					var dataSanitize = sanitizeSensitiveData(bytes)
					logFormat = logFormat + " PAYLOAD=" + dataSanitize
				}
			}
		}

		bytes, err := json.Marshal(resp)
		if err == nil {
			var dataSanitize = sanitizeSensitiveData(bytes)
			if dataSanitize != "null" {
				logFormat = logFormat + " RESPONSE=" + dataSanitize
			}
		}

		// Kalo misalnya ada cookie tambahin
		executor := ctx.Cookies("cms_email", "")
		if executor != "" {
			logFormat = logFormat + " EXECUTOR=" + executor
		}
		// if ctx.Cookie() {
		// extract jwt
		// logFormat = logFormat + " EXECUTOR=" + email_executor
		// }

		if Cfg.Application.EnableLogToFile {
			// Write log to file
			go func() {
				f, err := os.OpenFile(Cfg.Application.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}

				defer f.Close()

				if _, err := f.WriteString(logFormat + "\n"); err != nil {
					log.Println(err)
				}
			}()
		}

		color.Magenta(logFormat)
	}
}
