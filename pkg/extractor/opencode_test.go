package extractor

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenCodeIncrementalChildChangesAndOldSchemas(t *testing.T) {
	for _, updated := range []bool{true, false} {
		name := "timestamps"
		if !updated {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "opencode space #.db")
			db, err := sql.Open("sqlite", path)
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			column := ""
			if updated {
				column = ", time_updated INTEGER DEFAULT 1"
			}
			_, err = db.Exec(`CREATE TABLE session(id TEXT PRIMARY KEY,title TEXT,time_created INTEGER,directory TEXT` + column + `);
				CREATE TABLE message(id TEXT PRIMARY KEY,session_id TEXT,time_created INTEGER,data TEXT` + column + `);
				CREATE TABLE part(id TEXT PRIMARY KEY,message_id TEXT,time_created INTEGER,data TEXT` + column + `);
				INSERT INTO session(id,title,time_created,directory) VALUES('s','Session expiration',1,'D:\work\example-project');
				INSERT INTO message(id,session_id,time_created,data) VALUES('m1','s',1,'{"role":"user"}'),('m2','s',2,'{"role":"assistant"}');
				INSERT INTO part(id,message_id,time_created,data) VALUES('p1','m1',1,'{"type":"text","text":"Explain session expiration"}'),('p2','m2',2,'{"type":"text","text":"Refresh tokens control session expiration."}');`)
			if err != nil {
				t.Fatal(err)
			}
			e := NewOpenCodeExtractor()
			var firstRevision string
			convs, err := e.ExtractIncremental(path, func(_, revision string) bool { firstRevision = revision; return true })
			if err != nil || len(convs) != 1 || len(convs[0].Messages) != 2 || convs[0].Project != "example-project" {
				t.Fatalf("bad extraction: %+v %v", convs, err)
			}
			convs, err = e.ExtractIncremental(path, func(_, revision string) bool { return revision == "" || revision != firstRevision })
			if err != nil {
				t.Fatal(err)
			}
			if updated && len(convs) != 0 {
				t.Fatal("unchanged session extracted")
			}
			if !updated && (firstRevision != "" || len(convs) != 1) {
				t.Fatal("legacy schema must not cache")
			}
			if updated {
				if _, err := db.Exec(`UPDATE part SET data='{"type":"text","text":"Changed child only"}',time_updated=2 WHERE id='p2'`); err != nil {
					t.Fatal(err)
				}
				convs, err = e.ExtractIncremental(path, func(_, revision string) bool { return revision != firstRevision })
				if err != nil || len(convs) != 1 || convs[0].Messages[1].Content != "Changed child only" {
					t.Fatalf("child update missed: %+v %v", convs, err)
				}
			}
		})
	}
}

func TestProjectNameRejectsUserHomeDirectories(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{`C:\Users\alice`, `/home/alice`, `/Users/alice`, home} {
		if got := projectName(path); got != "" {
			t.Fatalf("exposed home directory name %q from %s", got, path)
		}
	}
	for _, path := range []string{`C:\Users\alice\work\payments`, `/home/alice/work/payments`} {
		if got := projectName(path); got != "payments" {
			t.Fatalf("projectName(%s)=%q", path, got)
		}
	}
}
