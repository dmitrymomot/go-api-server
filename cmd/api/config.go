package main

import (
	"time"

	"github.com/dmitrymomot/go-env"
	_ "github.com/joho/godotenv/autoload" // Load .env file automatically
)

var (
	// Application
	appName     = env.GetString("APP_NAME", "api")
	appDebug    = env.GetBool("APP_DEBUG", false)
	appLogLevel = env.GetString("APP_LOG_LEVEL", "info")

	// HTTP Router
	httpPort                  = env.GetInt("HTTP_PORT", 8080)
	httpRequestTimeout        = env.GetDuration("HTTP_REQUEST_TIMEOUT", time.Second*10)
	httpServerShutdownTimeout = env.GetDuration("HTTP_SERVER_SHUTDOWN_TIMEOUT", time.Second*5)

	// Cors
	corsAllowedOrigins     = env.GetStrings("CORS_ALLOWED_ORIGINS", ",", []string{"*"})
	corsAllowedMethods     = env.GetStrings("CORS_ALLOWED_METHODS", ",", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"})
	corsAllowedHeaders     = env.GetStrings("CORS_ALLOWED_HEADERS", ",", []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-ID", "X-Request-Id", "Origin", "User-Agent", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Pragma", "Referer"})
	corsAllowedCredentials = env.GetBool("CORS_ALLOWED_CREDENTIALS", true)
	corsMaxAge             = env.GetInt("CORS_MAX_AGE", 300)

	// Build tag is set up while deployment
	buildTag        = "undefined"
	buildTagRuntime = env.GetString("COMMIT_HASH", buildTag)

	// DB
	dbConnString   = env.MustString("DATABASE_URL")
	dbMaxOpenConns = env.GetInt("DATABASE_MAX_OPEN_CONNS", 20)
	dbMaxIdleConns = env.GetInt("DATABASE_IDLE_CONNS", 2)

	// Redis
	redisConnString = env.GetString("REDIS_URL", "redis://localhost:6379/0")

	// Worker
	workerConcurrency     = env.GetInt("WORKER_CONCURRENCY", 10)
	workerShutdownTimeout = env.GetDuration("WORKER_SHUTDOWN_TIMEOUT", time.Second*5)
	workerLogLevel        = env.GetString("WORKER_LOG_LEVEL", "info")
	queueName             = env.GetString("QUEUE_NAME", "default")
	queueTaskDeadline     = env.GetDuration("QUEUE_TASK_DEADLINE", time.Minute)
	queueMaxRetry         = env.GetInt("QUEUE_TASK_RETRY_LIMIT", 3)
)
