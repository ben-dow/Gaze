package socket

import (
	"context"
	"fmt"
	"io"
	"net/http"
)
import "nhooyr.io/websocket"

func WebSocketHttp(w http.ResponseWriter, r *http.Request) {
	c, _ := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"logging"},
	})

	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	if c.Subprotocol() != "logging" {
		c.Close(websocket.StatusPolicyViolation, "client must speak echo subprotocol")
		return
	}

	for {
		err := util(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			return
		}
	}

}

func util(ctx context.Context, c *websocket.Conn) error {
	fmt.Println("starting at top")

	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	err = w.Close()
	return err
}
