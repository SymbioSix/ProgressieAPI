package main

import (
	"log"

	_ "github.com/SymbioSix/ProgressieAPI/docs"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	todo_r "github.com/SymbioSix/ProgressieAPI/routers/Todo"
	au_r "github.com/SymbioSix/ProgressieAPI/routers/auth"
	crs_r "github.com/SymbioSix/ProgressieAPI/routers/courses"
	dash_r "github.com/SymbioSix/ProgressieAPI/routers/dashboard"
	ln_r "github.com/SymbioSix/ProgressieAPI/routers/landing"
	lead_r "github.com/SymbioSix/ProgressieAPI/routers/leaderboard"
	quiz_r "github.com/SymbioSix/ProgressieAPI/routers/quiz"
	rank_r "github.com/SymbioSix/ProgressieAPI/routers/rank"
	todo_s "github.com/SymbioSix/ProgressieAPI/services/To_do_list"
	au_s "github.com/SymbioSix/ProgressieAPI/services/auth"
	crs_s "github.com/SymbioSix/ProgressieAPI/services/courses"
	dash_s "github.com/SymbioSix/ProgressieAPI/services/dashboard"
	ln_s "github.com/SymbioSix/ProgressieAPI/services/landing"
	lead_s "github.com/SymbioSix/ProgressieAPI/services/leaderboard"
	quiz_s "github.com/SymbioSix/ProgressieAPI/services/quiz"
	rank_s "github.com/SymbioSix/ProgressieAPI/services/rank"
	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/SymbioSix/ProgressieAPI/utils/swagger" // PROPS TO: github.com/gofiber/swagger (modified so it can runned in gofiber/fiber/v3)
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

var (
	app *fiber.App

	// TODO: Create New Service Controller and Router Variables
	AuthController au_s.AuthController
	AuthRouter     au_r.AuthRouter

	LandNavbarService ln_s.LandNavbarService
	LandNavbarRouter  ln_r.LandNavbarRouter

	LandHeroService ln_s.LandHeroService
	LandHeroRouter  ln_r.LandHeroRouter

	LandFaqService ln_s.LandFaqService
	LandFaqRouter  ln_r.LandFaqRouter

	LandFaqCategoryService ln_s.LandFaqCategoryService
	LandFaqCategoryRouter  ln_r.LandFaqCategoryRouter

	LandAboutUsService ln_s.AboutUsService
	LandAboutUsRouter  ln_r.LandAboutUsRouter

	LandFooterService ln_s.FooterService
	LandFooterRouter  ln_r.LandFooterRouter

	DashboardController dash_s.DashboardController
	DashboardRouter     dash_r.DashboardRouter

	CourseController crs_s.CourseController
	CourseRouter     crs_r.GetCourseRouter

	TodoController todo_s.ToDoListController
	TodoRouter     todo_r.SetupToDoListRoutes

	LeaderboardController lead_s.LeaderboardController
	LeaderboardRouter     lead_r.LeaderboardRouter

	RankController rank_s.RankController
	RankRouter     rank_r.RankRouter

	QuizController quiz_s.QuizService
	QuizRouter     quiz_r.GetQuizRouter

	QuizQuestionController quiz_s.QuizQuestionService
	QuizQuestionRouter     quiz_r.GetQuizQuestionRouter

	QuizAnswerMultipleChoiceController quiz_s.QuizAnswerMultipleChoiceService
	QuizAnswerMultipleChoiceRouter     quiz_r.GetQuizAMCRouter

	QuizResultController quiz_s.QuizResultService
	QuizResultRouter     quiz_r.GetQuizResultRouter
)

func init() {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Initialize Database and API Connectivity
	s.ConnectDatabase(&config)
	s.ConnectViaAPI(&config)

	// TODO: Initialize Routers and Controllers

	AuthController = au_s.NewAuthController(s.DB, s.Client)
	AuthRouter = au_r.NewRouteAuthController(AuthController)

	LandNavbarService = ln_s.NewLandNavbarService(s.DB)
	LandNavbarRouter = ln_r.NewLandNavbarRouter(LandNavbarService)

	LandHeroService = ln_s.NewLandHeroService(s.DB)
	LandHeroRouter = ln_r.NewLandHeroRouter(LandHeroService)

	LandFaqService = ln_s.NewLandFaqService(s.DB)
	LandFaqRouter = ln_r.NewLandFaqRouter(LandFaqService)

	LandFaqCategoryService = ln_s.NewLandFaqCategoryService(s.DB)
	LandFaqCategoryRouter = ln_r.NewLandFaqCategoryRouter(LandFaqCategoryService)

	LandAboutUsService = ln_s.NewAboutUsService(s.DB)
	LandAboutUsRouter = ln_r.NewLandAboutUsRouter(LandAboutUsService)

	LandFooterService = ln_s.NewFooterService(s.DB)
	LandFooterRouter = ln_r.NewLandFooterRouter(LandFooterService)

	DashboardController = dash_s.NewDashboardController(s.DB, s.Client)
	DashboardRouter = dash_r.NewRouteAuthController(DashboardController)

	CourseController = crs_s.NewCourseController(s.DB, s.Client)
	CourseRouter = crs_r.NewGetCourseRouter(CourseController)

	TodoController = todo_s.NewTodoController(s.DB, s.Client)
	TodoRouter = todo_r.NewSetupToDoListRoutes(TodoController)

	LeaderboardController = lead_s.NewLeaderboardController(s.DB, s.Client)
	LeaderboardRouter = lead_r.NewRouteLeaderboardController(LeaderboardController)

	RankController = rank_s.NewRankController(s.DB, s.Client)
	RankRouter = rank_r.NewRouteRankController(RankController)

	QuizController = quiz_s.NewQuizService(s.DB)
	QuizRouter = quiz_r.NewGetQuizRouter(QuizController)

	QuizQuestionController = quiz_s.NewQuizQuestionService(s.DB)
	QuizQuestionRouter = quiz_r.NewGetQuizQuestionRouter(QuizQuestionController)

	QuizAnswerMultipleChoiceController = quiz_s.NewQuizAnswerMultipleChoiceService(s.DB)
	QuizAnswerMultipleChoiceRouter = quiz_r.NewGetQuizAMCRouter(QuizAnswerMultipleChoiceController)

	QuizResultController = quiz_s.NewQuizResultService(s.DB)
	QuizResultRouter = quiz_r.NewGetQuizResultRouter(QuizResultController)

	app = fiber.New()
}

//	@title			Self-Ie API Services
//	@version		1.0
//	@description	RESTful Self-ie Academy API Services. Built to ensure Self-ie Services are good to be served!
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		selfieapi.up.railway.app
//	@BasePath	/v1

//	@accept		json
//	@produce	json

func main() {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	corsConfig := cors.Config{
		// Allow Origins Will Be Updated With Our Web Domain
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost:5175", "http://127.0.0.1:5175", "https://selfieprogressie.netlify.app"},
		AllowCredentials: true,
	}

	app.Use(cors.New(corsConfig))

	router := app.Group("/v1")

	router.Get("/liveness-check",
		healthcheck.NewHealthChecker(
			healthcheck.Config{
				Probe: func(c fiber.Ctx) bool { return true },
			},
		),
	)

	router.Get("/healthcheck",

		func(c fiber.Ctx) error {
			var database_status string = "ready"
			var supabase_api_status string = "ready"
			var overall_status string = "super healthy"
			healthmap := status.HealthMap{
				DatabaseStatus:    database_status,
				SupabaseAPIStatus: supabase_api_status,
				OverallStatus:     overall_status,
			}
			if s.DB.Error != nil && !s.Client.Rest.Ping() {
				database_status = "error"
				supabase_api_status = "error"
				overall_status = "having issue(s) : database and supabase"
				return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
			}
			if s.DB.Error != nil {
				database_status = "error"
				overall_status = "having issue(s) : database"
				return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
			}
			if !s.Client.Rest.Ping() {
				supabase_api_status = "error"
				overall_status = "having issue(s) : supabase"
				return c.Status(fiber.StatusInternalServerError).JSON(healthmap)
			}
			return c.Status(fiber.StatusOK).JSON(healthmap)
		})

	router.Get("/swagger/*", swagger.HandlerDefault)

	// Connect all the routes
	AuthRouter.AuthRoutes(router)
	LandNavbarRouter.LandNavbarRoutes(router)
	LandHeroRouter.LandHeroRoutes(router)
	LandFaqRouter.LandFaqRoutes(router)
	LandFaqCategoryRouter.LandFaqCategoryRoutes(router)
	LandAboutUsRouter.LandAboutUsRoutes(router)
	LandFooterRouter.LandFooterRoutes(router)
	DashboardRouter.DashboardRoutes(router)
	CourseRouter.GetCourseRoutes(router)
	TodoRouter.GetSetupToDoListRoutes(router)
	LeaderboardRouter.LeaderboardRoutes(router)
	RankRouter.RankRoutes(router)
	QuizRouter.GetQuizRouter(router)
	QuizAnswerMultipleChoiceRouter.GetQuizAMCRouter(router)
	QuizQuestionRouter.GetQuizQuestionRouter(router)
	QuizResultRouter.GetQuizResultRouter(router)

	// Serve The API
	s.StartServerWithGracefulShutdown(app, &config)
}

// HealthCheck godoc
//
//	@Summary		Get API Health Check Status
//	@Description	Get API Health Check Status
//	@Tags			CheckUp
//	@Produce		json
//	@Success		200	{object}	status.HealthMap
//	@Failure		500	{object}	status.HealthMap
//	@Router			/healthcheck [get]
func Unusable1() {}

// LivenessCheck godoc
//
//	@Summary		Get Liveness Check Status
//	@Description	Get Liveness Check Status
//	@Tags			CheckUp
//	@Produce		plain
//	@Router			/liveness-check [get]
func Unuable2() {}
