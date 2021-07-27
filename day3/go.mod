module hello

go 1.16

//replace api/Config => ./Config

//replace api/Routes => ./Routes

//replace api/Models => ./Models

//replace api/Controllers => ./Controllers

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
