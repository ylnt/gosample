package hello

import (
	"context"
	"expvar"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"gopkg.in/tokopedia/logging.v1"
)

type ServerConfig struct {
	Name string
}

type Config struct {
	Server ServerConfig
}

type HelloWorldModule struct {
	cfg       *Config
	something string
	stats     *expvar.Int
}

func NewHelloWorldModule() *HelloWorldModule {

	var cfg Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	// this message only shows up if app is run with -debug option, so its great for debugging
	logging.Debug.Println("hello init called", cfg.Server.Name)

	return &HelloWorldModule{
		cfg:       &cfg,
		something: "John Doe",
		stats:     expvar.NewInt("rpsStats"),
	}

}

func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), r.URL.Path)
	defer span.Finish()

	hlm.stats.Add(1)
	hlm.someSlowFuncWeWantToTrace(ctx, w)
}

<<<<<<< HEAD
func (hlm *HelloWorldModule) SayIntroToGo(w http.ResponseWriter, r *http.Request) {
	type student struct {
		no    int
		name  string
		age   int
		major string
	}

	type students struct {
		person []student
	}

	stu := students{}
	stu.person = append(stu.person, student{1, "Andi", 12, "Mathematics"})
	stu.person = append(stu.person, student{2, "Anton", 13, "Computer Science"})
	stu.person = append(stu.person, student{3, "Budi", 17, "Information System"})
	stu.person = append(stu.person, student{4, "Calvin", 23, "Computer Science"})
	stu.person = append(stu.person, student{5, "Dennis", 12, "Mathematics"})

	for i := 0; i < len(stu.person); i++ {
		if (stu.person[i].no % 2 == 1) {
			w.Write([]byte("Odd:"))
			w.Write([]byte(stu.person[i].name))
		}

		w.Write([]byte("\n"))
		w.Write([]byte("\n"))

		if (stu.person[i].age < 14) && (stu.person[i].major == "Computer Science") {
			w.Write([]byte("Under 14 and Computer Science:"))
			w.Write([]byte(stu.person[i].name))
		}

		w.Write([]byte("\n"))
		w.Write([]byte("\n"))
	}
}

=======
>>>>>>> 00f41174e6e7cd1c36a3d3ac10eea8aaea40a57c
func (hlm *HelloWorldModule) someSlowFuncWeWantToTrace(ctx context.Context, w http.ResponseWriter) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "someSlowFuncWeWantToTrace")
	defer span.Finish()

	w.Write([]byte("Hello " + hlm.something))
}
