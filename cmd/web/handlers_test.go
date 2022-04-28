package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {

	app := newApplication(t)
	ts := newTestServer(t, app.routes())

	sd, _, body := ts.get(t, "/ping")

	if sd != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, sd)
	}

	if string(body) != "ok" {
		t.Errorf("want body to equal %q", "ok")
	}
	// app := &application{
	// 	errorLog: log.New(io.Discard, "", 0),
	// 	infoLog:  log.New(io.Discard, "", 0),
	// }

	// ts := httptest.NewTLSServer(app.routes())

	// defer ts.Close()

	// rs, err := ts.Client().Get(ts.URL + "/ping")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if rs.StatusCode != http.StatusOK {
	// 	t.Errorf("want %d;got %d", http.StatusOK, rs.StatusCode)
	// }

	// defer rs.Body.Close()

	// body, err := io.ReadAll(rs.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if string(body) != "OK" {
	// 	t.Errorf("want body to equal %q", "OK")
	// }

	//rr := httptest.NewRecorder()

	// r, err := http.NewRequest(http.MethodGet, "/", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// ping(rr, r)

	// rs := rr.Result()

	// if rs.StatusCode != http.StatusOK {
	// 	t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	// }

	// defer rs.Body.Close()

	// body, err := io.ReadAll(rs.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if string(body) != "OK" {
	// 	t.Errorf("want body to equal %q", "OK")
	// }
}
