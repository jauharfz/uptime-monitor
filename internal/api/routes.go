package api

import "net/http"

func Routes(app *Application) http.Handler {
	// Normal Routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.HealthTest)
	mux.HandleFunc("GET /openapi.yaml", app.GetYamlSpec)
	mux.HandleFunc("GET /docs", app.ShowSwaggerUI)
	mux.HandleFunc("POST /api/v1/users/register", app.PostUserRegister)
	mux.HandleFunc("POST /api/v1/users/login", app.PostUserLogin)

	// Public Routes
	mux.Handle("GET /api/v1/public/monitors/{id}/status", app.Cors(http.HandlerFunc(app.ShowPublicMonitor)))

	// Protected Routes
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("POST /api/v1/monitors", app.CreateMonitor)
	protectedMux.HandleFunc("GET /api/v1/monitors", app.ListMonitors)
	protectedMux.HandleFunc("GET /api/v1/monitors/{id}", app.ShowMonitor)
	protectedMux.HandleFunc("PATCH /api/v1/monitors/{id}", app.UpdateMonitor)
	protectedMux.HandleFunc("DELETE /api/v1/monitors/{id}", app.DeleteMonitor)
	protectedMux.HandleFunc("GET /api/v1/monitors/{id}/checks", app.CheckMonitor)
	protectedMux.HandleFunc("GET /api/v1/monitors/{id}/stats", app.ShowMonitorStats)

	mux.Handle("/", app.RequireAuth(protectedMux))

	return mux
}
