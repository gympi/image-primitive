package handlers

import (
    "html/template"
)

type Page struct {
 Title string
 Body  template.HTML
}
