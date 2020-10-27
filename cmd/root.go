package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	"github.com/rifqiakrm/golara-boilerplate/controllers"
	"github.com/rifqiakrm/golara-boilerplate/utils/cache"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cfgFile   string
	dbPool    *sql.DB
	cachePool *redis.Pool
	router    *gin.Engine
)

var rootCMD = &cobra.Command{
	Use:   "golara-boilerplate",
	Short: "Golang Laravel Microservice Boilerplate",
	Long:  "Golang Laravel Microservice Boilerplate",
	Run: func(cmd *cobra.Command, args []string) {
		//gin.SetMode(gin.ReleaseMode) //uncomment for production deployment
		router = gin.New()
		router.Use(cors.Default())
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		controllers.Init(
			dbPool,
		)
		cache.Init(cachePool)
		Routes()
		if err := router.Run(fmt.Sprintf(":%s", viper.GetString("app.port"))); err != nil {
			panic(err)
		}
	},
}

func init() {
	cobra.OnInitialize(splash, initconfig, initDB, initCache)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCMD.PersistentFlags().StringVar(&cfgFile, "configs", "configs/config.local.toml", "configs file (example is $HOME/configs.toml)")
}

// splash print plain text message to console
func splash() {
	fmt.Print(`
  ________ ________  .____       _____ __________    _____   
 /  _____/ \_____  \ |    |     /  _  \\______   \  /  _  \  
/   \  ___  /   |   \|    |    /  /_\  \|       _/ /  /_\  \ 
\    \_\  \/    |    \    |___/    |    \    |   \/    |    \
 \______  /\_______  /_______ \____|__  /____|_  /\____|__  /
        \/         \/        \/       \/       \/         \/ 
`)
}

func initconfig() {
	viper.SetConfigType("toml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// search configs in home directory with name "configs" (without extension)
		viper.AddConfigPath("./configs")
		viper.SetConfigName(os.Getenv("CONFIG_FILE"))
	}

	//read env
	viper.AutomaticEnv()

	// if a configs file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("config application:", err)
	}

	log.Println("using configs file:", viper.ConfigFileUsed())
}

func initDB() {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		"disable")

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbPool = db

	log.Println("database successfully connected!")
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initCache() {
	redisHost := fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port"))
	if redisHost == ":" {
		redisHost = "localhost:6379"
	}
	cachePool = newPool(redisHost)

	ctx := context.Background()
	_, err := cachePool.GetContext(ctx)

	if err != nil {
		panic("failed to connect to redis")
	}

	log.Println("redis successfully connected!")
	cleanupHook()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		cachePool.Close()
		os.Exit(0)
	}()
}
