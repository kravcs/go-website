// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1569083795, 0),

		Content: string("document.getElementById(\"text\").innerHTML = \"Hello World<br/><p>This website is served by Go!</p>\""),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1569083798, 0),

		Content: string("<html>\n    <head>\n        <title>Wow cool!</title>\n        <link rel=\"stylesheet\" href=\"styles.css\" />\n    </head>\n    <body>\n        <div class=\"container\">\n            <p id=\"text\" class=\"text\"></p>\n        </div>\n        <script src=\"app.js\"></script>\n    </body>\n</html>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "styles.css",
		FileModTime: time.Unix(1569082748, 0),

		Content: string(""),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1569082748, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.js"
			file3, // "index.html"
			file4, // "styles.css"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`public`, &embedded.EmbeddedBox{
		Name: `public`,
		Time: time.Unix(1569082748, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
			"styles.css": file4,
		},
	})
}
