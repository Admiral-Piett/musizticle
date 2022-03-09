package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<div class="tenor-gif-embed" data-postid="17424068" data-share-method="host" data-aspect-ratio="2.19178" data-width="80%"><a href="https://tenor.com/view/star-wars-obi-wan-kenobi-hello-there-hi-there-greeting-gif-17424068">Star Wars Obi Wan Kenobi GIF</a>from <a href="https://tenor.com/search/star+wars-gifs">Star Wars GIFs</a></div> <script type="text/javascript" async src="https://tenor.com/embed.js"></script>`)
	}
}
