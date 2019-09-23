module flamingo

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
)

require (
	github.com/boltdb/bolt v1.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.10
	github.com/joho/godotenv v1.3.0
	github.com/pquerna/ffjson v0.0.0-20190918152532-477b94a9a7ec
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
)
