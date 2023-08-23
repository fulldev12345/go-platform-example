package placeholder

import (
	"platform/authorization"
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},
		//&SimpleMessageComponent{},
		authorization.NewAuthComponent(
			"protected",
			authorization.NewRoleCondition("Administrator"),
			CounterHandler{},
		),
		handling.NewRouter(
			handling.HandlerEntry{"", NameHandler{}},
			handling.HandlerEntry{"", DayHandler{}},
			//handling.HandlerEntry{ "",  CounterHandler{}},
			handling.HandlerEntry{"", AuthenticationHandler{}},
		).AddMethodAlias("/", NameHandler.GetNames),
	)
}
func Start() {
	sessions.RegisterSessionService()
	authorization.RegisterDefaultSignInService()
	authorization.RegisterDefaultUserService()
	RegisterPlaceholderUserStore()
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
