package main

import (
	"net/http"

	"github.com/jrovieri/bookings/internal/config"
	"github.com/jrovieri/bookings/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJson)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	// mux.With(Auth).Get("/admin/dashboard", handlers.Repo.AdminDashboard)
	// mux.With(Auth).Get("/admin/reservations-new", handlers.Repo.AdminNewReservations)
	// mux.With(Auth).Get("/admin/reservations-all", handlers.Repo.AdminAllReservations)
	// mux.With(Auth).Get("/admin/reservations-calendar", handlers.Repo.AdminReservationsCalendar)

	mux.Get("/admin/dashboard", handlers.Repo.AdminDashboard)
	mux.Get("/admin/reservations-new", handlers.Repo.AdminNewReservations)
	mux.Get("/admin/reservations-all", handlers.Repo.AdminAllReservations)
	mux.Get("/admin/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
	mux.Post("/admin/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)

	mux.Get("/admin/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservations)
	mux.Post("/admin/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservations)

	mux.Get("/admin/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservation)
	mux.Get("/admin/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservation)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
